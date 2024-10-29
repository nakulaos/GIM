package svc_chat_invite

import (
	chat_invite_client "GIM/apps/chat_invite/client"
	"GIM/apps/interfaces/internal/config"
	"GIM/apps/interfaces/internal/dto/dto_chat_invite"
	"GIM/pkg/xhttp"
)

type ChatInviteService interface {
	InitiateChatInvite(params *dto_chat_invite.InitiateChatInviteReq, uid int64) (resp *xhttp.Resp)
	ChatInviteList(params *dto_chat_invite.ChatInviteListReq, uid int64) (resp *xhttp.Resp)
	ChatInviteHandle(params *dto_chat_invite.ChatInviteHandleReq, uid int64) (resp *xhttp.Resp)
}

type chatInviteService struct {
	chatInviteClient chat_invite_client.ChatInviteClient
}

func NewChatInviteService() ChatInviteService {
	conf := config.GetConfig()
	chatInviteClient := chat_invite_client.NewChatInviteClient(conf.Etcd, conf.ChatInviteServer, conf.Jaeger, conf.Name)
	return &chatInviteService{chatInviteClient: chatInviteClient}
}
