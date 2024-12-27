package service

import (
	"GIM/pkg/common/xlog"
	"GIM/pkg/entity"
	"GIM/pkg/proto/pb_chat_member"
	"GIM/pkg/proto/pb_enum"
	"context"
)

func (s *chatMemberService) GetContactList(ctx context.Context, req *pb_chat_member.GetContactListReq) (resp *pb_chat_member.GetContactListResp, _ error) {
	/*
		1. 获取联系人列表
	*/
	resp = new(pb_chat_member.GetContactListResp)
	var (
		w   = entity.NewMysqlQuery()
		err error
	)
	w.SetFilter("owner_id=?", req.Uid)
	w.SetFilter("chat_type=?", pb_enum.CHAT_TYPE_PRIVATE)
	w.SetFilter("chat_id>?", req.LastChatId)
	w.SetLimit(req.Limit)
	w.SetSort("chat_id ASC")
	resp.List, err = s.chatMemberRepo.ChatMemberBasicInfoList(w)
	if err != nil {
		resp.Set(ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED, err.Error())
		return
	}
	return
}
