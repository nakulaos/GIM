package service

import (
	"GIM/apps/apis/upload/internal/config"
	"GIM/apps/apis/upload/internal/dto"
	chat_client "GIM/apps/chat/client"
	user_client "GIM/apps/user/client"
	"GIM/pkg/xhttp"
	"github.com/gin-gonic/gin"
)

/*
https://www.cnblogs.com/peteremperor/p/16301336.html
https://www.cnblogs.com/liuqingzheng/p/16124105.html

https://www.lanol.cn/post/599.html
*/

type UploadService interface {
	UploadAvatar(ctx *gin.Context, req *dto.UploadAvatarReq) (resp *xhttp.Resp)
	Presigned(ctx *gin.Context, req *dto.PresignedReq) (resp *xhttp.Resp)
}

type uploadService struct {
	cfg        *config.Config
	userClient user_client.UserClient
	chatClient chat_client.ChatClient
}

func NewUploadService(conf *config.Config) UploadService {
	chatClient := chat_client.NewChatClient(conf.Etcd, conf.ChatServer, conf.Jaeger, conf.Name)
	userClient := user_client.NewUserClient(conf.Etcd, conf.UserServer, conf.Jaeger, conf.Name)
	return &uploadService{cfg: conf, chatClient: chatClient, userClient: userClient}
}
