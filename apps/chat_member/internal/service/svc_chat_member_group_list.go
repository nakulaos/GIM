package service

import (
	"GIM/pkg/common/xlog"
	"GIM/pkg/entity"
	"GIM/pkg/proto/pb_chat_member"
	"GIM/pkg/proto/pb_enum"
	"context"
)

func (s *chatMemberService) GetGroupChatList(ctx context.Context, req *pb_chat_member.GetGroupChatListReq) (resp *pb_chat_member.GetGroupChatListResp, _ error) {

	/*
		1. 获取群聊列表
	*/
	resp = new(pb_chat_member.GetGroupChatListResp)
	var (
		w   = entity.NewMysqlQuery()
		err error
	)
	w.SetFilter("uid=?", req.Uid)
	w.SetFilter("chat_type=?", pb_enum.CHAT_TYPE_GROUP)
	w.SetFilter("chat_id>?", req.LastChatId)
	w.SetLimit(req.Limit)
	w.SetSort("chat_id ASC")
	resp.List, err = s.chatMemberRepo.GroupChatBasicInfoList(w)
	if err != nil {
		resp.Set(ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED)
		xlog.Warn(resp.Code, resp.Msg, err.Error())
		return
	}
	return
}
