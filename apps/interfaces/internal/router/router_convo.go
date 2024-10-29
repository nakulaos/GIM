package router

import (
	"GIM/apps/interfaces/dig"
	"GIM/apps/interfaces/internal/ctrl/ctrl_convo"
	"GIM/apps/interfaces/internal/service/svc_convo"
	"github.com/gin-gonic/gin"
)

func registerConvoRouter(group *gin.RouterGroup) {
	var svc svc_convo.ConvoService
	dig.Invoke(func(s svc_convo.ConvoService) {
		svc = s
	})
	ctrl := ctrl_convo.NewConvoCtrl(svc)
	router := group.Group("convo")
	router.POST("list", ctrl.ConvoList)
	router.GET("chat_seq_list", ctrl.ConvoChatSeqList)
}
