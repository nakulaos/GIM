package router

import (
	"GIM/apps/interfaces/dig"
	"GIM/apps/interfaces/internal/ctrl/ctrl_chat"
	"GIM/apps/interfaces/internal/service/svc_chat"
	"github.com/gin-gonic/gin"
)

func registerChatRouter(group *gin.RouterGroup) {
	var svc svc_chat.ChatService
	dig.Invoke(func(s svc_chat.ChatService) {
		svc = s
	})
	ctrl := ctrl_chat.NewChatCtrl(svc)
	router := group.Group("chat")
	router.POST("create_group_chat", ctrl.CreateGroupChat)
	router.POST("edit_group_chat", ctrl.EditGroupChat)
	router.POST("delete_contact", ctrl.DeleteContact)
	router.POST("quit_group_chat", ctrl.QuitGroupChat)
	router.POST("remove_group_member", ctrl.RemoveGroupChatMember)
	router.GET("group_chat_details", ctrl.GroupChatDetails)
}
