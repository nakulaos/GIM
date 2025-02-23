package service

import (
	"GIM/domain/po"
	"GIM/pkg/common/xants"
	"GIM/pkg/common/xlog"
	"GIM/pkg/entity"
	"GIM/pkg/proto/pb_chat"
	"GIM/pkg/proto/pb_chat_member"
	"GIM/pkg/proto/pb_enum"
	"GIM/pkg/utils"
	"context"
)

func (s *chatService) getChat(chatId int64) (chat *po.Chat, err error) {
	var (
		w = entity.NewMysqlQuery()
	)
	w.SetFilter("chat_id=?", chatId)
	chat, err = s.chatRepo.Chat(w)
	return
}

func (s *chatService) getChatMember(chatId int64, uid int64) (member *pb_chat_member.ChatMemberInfo, err error) {
	var (
		w = entity.NewMysqlQuery()
	)
	w.SetFilter("chat_id=?", chatId)
	w.SetFilter("uid=?", uid)
	member, err = s.chatMemberRepo.ChatMember(w)
	return
}

func (s *chatService) RemoveGroupChatMember(ctx context.Context, req *pb_chat.RemoveGroupChatMemberReq) (resp *pb_chat.RemoveGroupChatMemberResp, _ error) {
	/*
		1. 判断uid 是否在chat中
		2. 校验权限移除能移除的
		3. 发送移除成员消息
	*/

	resp = new(pb_chat.RemoveGroupChatMemberResp)
	var (
		member *pb_chat_member.ChatMemberInfo
		u      = entity.NewMysqlUpdate()
		err    error
	)
	if len(req.MemberList) == 0 {
		return
	}
	member, err = s.getChatMember(req.ChatId, req.Uid)
	if err != nil {
		resp.Set(ERROR_CODE_CHAT_QUERY_DB_FAILED, ERROR_CHAT_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_QUERY_DB_FAILED, ERROR_CHAT_QUERY_DB_FAILED, err.Error())
		return
	}
	if member.Uid == 0 {
		resp.Set(ERROR_CODE_CHAT_QUERY_DB_FAILED, ERROR_CHAT_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_QUERY_DB_FAILED, ERROR_CHAT_QUERY_DB_FAILED)
		return
	}

	u.SetFilter("chat_id=?", req.ChatId)
	u.SetFilter("uid IN(?)", req.MemberList)
	u.SetFilter("role_id < ?", member.RoleId)
	u.SetFilter("deleted_ts=?", 0)
	u.Set("status", int32(pb_enum.CHAT_STATUS_REMOVED))
	u.Set("deleted_ts", utils.NowUnix())

	_, err = s.removeChatMember(u, req.ChatId, req.MemberList, pb_enum.CHAT_TYPE_GROUP)
	if err != nil {
		resp.Set(ERROR_CODE_CHAT_UPDATE_VALUE_FAILED, ERROR_CHAT_UPDATE_VALUE_FAILED)
		xlog.Warn(resp.Code, resp.Msg, err.Error())
		return
	}

	// 推送移除成员
	xants.Submit(func() {
		s.quitGroupChatMessage(
			req.ChatId,
			req.MemberList,
			pb_enum.SUB_TOPIC_CHAT_REMOVE_CHAT_MEMBER,
			pb_enum.MSG_TYPE_REMOVE_CHAT_MEMBER)
	})
	return
}
