package router

import (
	"GIM/apps/interfaces/dig"
	"GIM/apps/interfaces/internal/ctrl/ctrl_chat_invite"
	"GIM/apps/interfaces/internal/service/svc_chat_invite"
	"github.com/gin-gonic/gin"
)

func registerChatInviteRouter(group *gin.RouterGroup) {
	var svc svc_chat_invite.ChatInviteService
	dig.Invoke(func(s svc_chat_invite.ChatInviteService) {
		svc = s
	})
	ctrl := ctrl_chat_invite.NewChatInviteCtrl(svc)
	router := group.Group("chat_invite")
	router.POST("initiate", ctrl.InitiateChatInvite)
	router.POST("handle", ctrl.ChatInviteHandle)
	router.GET("list", ctrl.ChatInviteList)
}
