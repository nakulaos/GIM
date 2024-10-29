package service

import (
	"GIM/domain/po"
	"GIM/pkg/common/xants"
	"GIM/pkg/common/xlog"
	"GIM/pkg/entity"
	"GIM/pkg/proto/pb_chat"
	"context"
	"github.com/jinzhu/copier"
)

func (s *chatService) GetChatInfo(ctx context.Context, req *pb_chat.GetChatInfoReq) (resp *pb_chat.GetChatInfoResp, _ error) {
	resp = &pb_chat.GetChatInfoResp{ChatInfo: &pb_chat.ChatInfo{}}
	var (
		w    = entity.NewMysqlQuery()
		chat *po.Chat
		err  error
	)
	w.SetFilter("chat_id = ?", req.ChatId)
	chat, err = s.chatRepo.Chat(w)
	if err != nil {
		resp.Set(ERROR_CODE_CHAT_QUERY_DB_FAILED, ERROR_CHAT_QUERY_DB_FAILED)
		xlog.Warn(resp.Code, resp.Msg, err.Error())
		return
	}
	copier.Copy(resp.ChatInfo, chat)
	xants.Submit(func() {
		s.cacheChatInfo(chat)
	})
	return
}

func (s *chatService) cacheChatInfo(chat *po.Chat) {
	var (
		chatInfo = new(pb_chat.ChatInfo)
	)
	copier.Copy(chatInfo, chat)
	s.chatCache.SetGroupChatInfo(chatInfo)
}
