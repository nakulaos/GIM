package service

import (
	chat_client "GIM/apps/chat/client"
	"GIM/apps/chat_invite/internal/config"
	dist_client "GIM/apps/dist/client"
	user_client "GIM/apps/user/client"
	"GIM/domain/cache"
	"GIM/domain/repo"
	"GIM/pkg/common/xkafka"
	"GIM/pkg/proto/pb_invite"
	"context"
)

type ChatInviteService interface {
	InitiateChatInvite(ctx context.Context, req *pb_invite.InitiateChatInviteReq) (resp *pb_invite.InitiateChatInviteResp, err error)
	ChatInviteHandle(ctx context.Context, req *pb_invite.ChatInviteHandleReq) (resp *pb_invite.ChatInviteHandleResp, err error)
	ChatInviteList(ctx context.Context, req *pb_invite.ChatInviteListReq) (resp *pb_invite.ChatInviteListResp, err error)
}

type chatInviteService struct {
	cfg              *config.Config
	chatInviteRepo   repo.ChatInviteRepository
	userRepo         repo.UserRepository
	avatarRepo       repo.AvatarRepository
	chatMemberRepo   repo.ChatMemberRepository
	chatRepo         repo.ChatRepository
	chatCache        cache.ChatCache
	chatMessageCache cache.ChatMessageCache
	chatMemberCache  cache.ChatMemberCache
	userCache        cache.UserCache
	chatClient       chat_client.ChatClient
	userClient       user_client.UserClient
	distClient       dist_client.DistClient
	producer         *xkafka.Producer
	cacheProducer    *xkafka.Producer
}

func NewChatInviteService(
	cfg *config.Config,
	chatInviteRepo repo.ChatInviteRepository,
	userRepo repo.UserRepository,
	avatarRepo repo.AvatarRepository,
	chatMemberRepo repo.ChatMemberRepository,
	chatRepo repo.ChatRepository,
	chatCache cache.ChatCache,
	chatMessageCache cache.ChatMessageCache,
	chatMemberCache cache.ChatMemberCache,
	userCache cache.UserCache,
) ChatInviteService {
	svc := &chatInviteService{
		cfg:              cfg,
		chatInviteRepo:   chatInviteRepo,
		userRepo:         userRepo,
		avatarRepo:       avatarRepo,
		chatMemberRepo:   chatMemberRepo,
		chatRepo:         chatRepo,
		chatCache:        chatCache,
		chatMessageCache: chatMessageCache,
		chatMemberCache:  chatMemberCache,
		userCache:        userCache}
	svc.chatClient = chat_client.NewChatClient(cfg.Etcd, cfg.ChatServer, cfg.Jaeger, cfg.Name)
	svc.userClient = user_client.NewUserClient(cfg.Etcd, cfg.UserServer, cfg.Jaeger, cfg.Name)
	svc.distClient = dist_client.NewDistClient(cfg.Etcd, cfg.DistServer, cfg.Jaeger, cfg.Name)
	svc.producer = xkafka.NewKafkaProducer(cfg.MsgProducer.Address, cfg.MsgProducer.Topic)
	svc.cacheProducer = xkafka.NewKafkaProducer(cfg.CacheProducer.Address, cfg.CacheProducer.Topic)
	return svc
}
