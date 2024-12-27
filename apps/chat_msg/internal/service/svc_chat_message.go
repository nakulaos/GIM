package service

import (
	chat_member_client "GIM/apps/chat_member/client"
	"GIM/apps/chat_msg/internal/config"
	"GIM/domain/cache"
	"GIM/domain/mrepo"
	"GIM/domain/repo"
	"GIM/pkg/common/xkafka"
	"GIM/pkg/proto/pb_chat_msg"
	"context"
)

type ChatMessageService interface {
	GetChatMessageList(ctx context.Context, req *pb_chat_msg.GetChatMessageListReq) (resp *pb_chat_msg.GetChatMessageListResp, err error)
	SearchMessage(ctx context.Context, req *pb_chat_msg.SearchMessageReq) (resp *pb_chat_msg.SearchMessageResp, err error)
	MessageOperation(ctx context.Context, req *pb_chat_msg.MessageOperationReq) (resp *pb_chat_msg.MessageOperationResp, err error)
}

type chatMessageService struct {
	cfg              *config.Config
	chatMessageRepo  repo.ChatMessageRepository
	messageHotRepo   mrepo.MessageHotRepository
	chatMessageCache cache.ChatMessageCache
	chatMemberCache  cache.ChatMemberCache
	chatMemberClient chat_member_client.ChatMemberClient
	producer         *xkafka.Producer
}

func NewChatMessageService(
	cfg *config.Config,
	chatMessageRepo repo.ChatMessageRepository,
	messageHotRepo mrepo.MessageHotRepository,
	chatMessageCache cache.ChatMessageCache,
	chatMemberCache cache.ChatMemberCache) ChatMessageService {
	svc := &chatMessageService{
		cfg:              cfg,
		chatMessageRepo:  chatMessageRepo,
		messageHotRepo:   messageHotRepo,
		chatMessageCache: chatMessageCache,
		chatMemberCache:  chatMemberCache}
	svc.chatMemberClient = chat_member_client.NewChatMemberClient(cfg.Etcd, cfg.ChatMemberServer, cfg.GrpcServer.Jaeger, cfg.Name)
	svc.producer = xkafka.NewKafkaProducer(cfg.MsgProducer.Address, cfg.MsgProducer.Topic)
	return svc
}
