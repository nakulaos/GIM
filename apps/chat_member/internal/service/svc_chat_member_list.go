package service

import (
	"GIM/pkg/common/xlog"
	"GIM/pkg/entity"
	"GIM/pkg/proto/pb_chat_member"
	"context"
)

func (s *chatMemberService) GetChatMemberList(ctx context.Context, req *pb_chat_member.GetChatMemberListReq) (resp *pb_chat_member.GetChatMemberListResp, _ error) {
	// 获取用户群聊用户列表的信息
	resp = &pb_chat_member.GetChatMemberListResp{}
	var (
		w   = entity.NewMysqlQuery()
		err error
	)
	w.SetFilter("chat_id=?", req.ChatId)
	w.SetFilter("uid>?", req.LastUid)
	w.SetSort("uid ASC")
	w.SetLimit(req.Limit)
	resp.List, err = s.chatMemberRepo.GroupChatMemberInfoList(w)
	if err != nil {
		resp.Set(ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED)
		xlog.Warn(resp.Code, resp.Msg, err.Error())
		return
	}
	return
}
