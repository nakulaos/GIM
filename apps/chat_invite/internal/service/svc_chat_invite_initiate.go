package service

import (
	"GIM/business/biz_chat_invite"
	"GIM/domain/po"
	"GIM/pkg/common/xants"
	"GIM/pkg/common/xlog"
	"GIM/pkg/common/xsnowflake"
	"GIM/pkg/entity"
	"GIM/pkg/proto/pb_dist"
	"GIM/pkg/proto/pb_enum"
	"GIM/pkg/proto/pb_invite"
	"context"
	"github.com/jinzhu/copier"
)

func (s *chatInviteService) InitiateChatInvite(_ context.Context, req *pb_invite.InitiateChatInviteReq) (resp *pb_invite.InitiateChatInviteResp, _ error) {
	/*
		聊天邀请发起
		1. 条件校验
			-- 私聊邀请只能邀请一个
			-- 群聊校验邀请人是否已经在群聊中
		2. 入库，邀请

	*/

	resp = new(pb_invite.InitiateChatInviteResp)
	var (
		invite = new(po.ChatInvite)
		w      = entity.NewMysqlQuery()
		err    error
	)

	// 参数校验
	switch req.ChatType {
	case pb_enum.CHAT_TYPE_PRIVATE:
		var (
			inviteeUid int64
		)
		if len(req.InviteeUids) != 1 {
			resp.Set(ERROR_CODE_CHAT_INVITE_PARAMETER_ERROR, ERROR_CHAT_INVITE_PARAMETER_ERROR)
			xlog.Warn(resp.Code, resp.Msg)
			return
		}
		inviteeUid = req.InviteeUids[0]
		if req.InitiatorUid == inviteeUid {
			resp.Set(ERROR_CODE_CHAT_INVITE_INITIATOR_INVITEE_SAME, ERROR_CHAT_INVITE_INITIATOR_INVITEE_SAME)
			xlog.Warn(resp.Code, resp.Msg)
			return
		}
		req.ChatId = xsnowflake.NewSnowflakeID()
	case pb_enum.CHAT_TYPE_GROUP:
		var (
			memberCount int64
		)
		if req.ChatId <= 0 {
			resp.Set(ERROR_CODE_CHAT_INVITE_PARAM_ERR, ERROR_CHAT_INVITE_PARAM_ERR)
			xlog.Warn(resp.Code, resp.Msg, req.String())
			return
		}
		w.SetFilter("chat_id=?", req.ChatId)
		w.SetFilter("uid IN(?)", req.InviteeUids)
		memberCount, err = s.chatMemberRepo.ChatMemberCount(w)
		if err != nil {
			resp.Set(ERROR_CODE_CHAT_INVITE_QUERY_DB_FAILED, ERROR_CHAT_INVITE_QUERY_DB_FAILED)
			xlog.Warn(resp.Code, resp.Msg, err.Error())
			return
		}
		if memberCount > 0 {
			resp.Set(ERROR_CODE_CHAT_INVITE_REPEAT_INVITATION, ERROR_CHAT_INVITE_REPEAT_INVITATION)
			xlog.Warn(resp.Code, resp.Msg, req.String())
			return
		}
	}
	copier.Copy(invite, req)
	var (
		invites = make([]*po.ChatInvite, len(req.InviteeUids))
		index   int
		uid     int64
	)
	for index, uid = range req.InviteeUids {
		ci := *invite
		ci.InviteId = xsnowflake.NewSnowflakeID()
		ci.InviteeUid = uid
		invites[index] = &ci
	}
	err = s.chatInviteRepo.CreateChatInvites(invites)
	if err != nil {
		resp.Set(ERROR_CODE_CHAT_INVITE_INSERT_VALUE_FAILED, ERROR_CHAT_INVITE_INSERT_VALUE_FAILED)
		xlog.Warn(resp.Code, resp.Msg, err)
		return
	}
	xants.Submit(func() {
		s.sendChatInviteNotificationMessage(req, invites)
	})
	return
}

func (s *chatInviteService) sendChatInviteNotificationMessage(inviteReq *pb_invite.InitiateChatInviteReq, invitees []*po.ChatInvite) {
	var (
		req  *pb_dist.ChatInviteNotificationReq
		resp *pb_dist.ChatInviteNotificationResp
		err  error
	)
	req, err = biz_chat_invite.ConstructChatInviteNotificationMessage(
		inviteReq,
		invitees,
		s.chatCache,
		s.userCache,
		s.chatRepo,
		s.userRepo)
	if err != nil {
		return
	}
	if req == nil {
		return
	}
	resp = s.distClient.ChatInviteNotification(req)
	if resp == nil {
		xlog.Warn(ERROR_CODE_CHAT_INVITE_GRPC_SERVICE_FAILURE, ERROR_CHAT_INVITE_GRPC_SERVICE_FAILURE)
		return
	}
	if resp.Code > 0 {
		xlog.Warn(resp.Code, resp.Msg)
	}
}
