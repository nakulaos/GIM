package router

import (
	"GIM/apps/interfaces/dig"
	"GIM/apps/interfaces/internal/ctrl/ctrl_chat_member"
	"GIM/apps/interfaces/internal/service/svc_chat_member"
	"github.com/gin-gonic/gin"
)

func registerChatMemberRouter(group *gin.RouterGroup) {
	var svc svc_chat_member.ChatMemberService
	dig.Invoke(func(s svc_chat_member.ChatMemberService) {
		svc = s
	})
	ctrl := ctrl_chat_member.NewChatMemberCtrl(svc)
	router := group.Group("chat_member")
	router.GET("list", ctrl.ChatMemberList)
	router.GET("contact_list", ctrl.ContactList)
	router.GET("group_chat_list", ctrl.GroupChatList)
}
