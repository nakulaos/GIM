package service

import (
	"GIM/apps/convo/internal/config"
	"GIM/domain/cache"
	"GIM/domain/repo"
	"GIM/pkg/proto/pb_convo"
	"context"
)

type ConvoService interface {
	ConvoList(ctx context.Context, req *pb_convo.ConvoListReq) (resp *pb_convo.ConvoListResp, err error)
	ConvoChatSeqList(ctx context.Context, req *pb_convo.ConvoChatSeqListReq) (resp *pb_convo.ConvoChatSeqListResp, err error)
}

// conversation
type convoService struct {
	cfg              *config.Config
	chatMemberRepo   repo.ChatMemberRepository
	convoCache       cache.ConvoCache
	chatMessageCache cache.ChatMessageCache
}

func NewConvoService(cfg *config.Config, chatMemberRepo repo.ChatMemberRepository, convoCache cache.ConvoCache, chatMessageCache cache.ChatMessageCache) ConvoService {
	svc := &convoService{
		cfg:              cfg,
		chatMemberRepo:   chatMemberRepo,
		convoCache:       convoCache,
		chatMessageCache: chatMessageCache}
	return svc
}
