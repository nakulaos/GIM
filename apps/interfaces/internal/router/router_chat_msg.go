package router

import (
	"GIM/apps/interfaces/dig"
	"GIM/apps/interfaces/internal/ctrl/ctrl_chat_msg"
	"GIM/apps/interfaces/internal/service/svc_chat_msg"
	"github.com/gin-gonic/gin"
)

func registerChatMessageRouter(group *gin.RouterGroup) {
	var svc svc_chat_msg.ChatMessageService
	dig.Invoke(func(s svc_chat_msg.ChatMessageService) {
		svc = s
	})
	ctrl := ctrl_chat_msg.NewChatMessageCtrl(svc)
	router := group.Group("chat_msg")
	router.GET("list", ctrl.GetChatMessageList)
	router.GET("search", ctrl.Search)
	router.POST("operation", ctrl.MessageOperation)
	router.POST("send_msg", ctrl.SendChatMessage)
}
