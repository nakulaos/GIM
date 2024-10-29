package biz_chat_invite

import (
	"GIM/domain/cache"
	"GIM/domain/cr/cr_chat"
	"GIM/domain/cr/cr_user"
	"GIM/domain/po"
	"GIM/domain/repo"
	"GIM/pkg/common/xlog"
	"GIM/pkg/proto/pb_chat"
	"GIM/pkg/proto/pb_dist"
	"GIM/pkg/proto/pb_enum"
	"GIM/pkg/proto/pb_invite"
	"GIM/pkg/proto/pb_msg"
	"GIM/pkg/proto/pb_user"
	"github.com/jinzhu/copier"
)

func ConstructChatInviteNotificationMessage(
	inviteReq *pb_invite.InitiateChatInviteReq,
	invitees []*po.ChatInvite,
	chatCache cache.ChatCache,
	userCache cache.UserCache,
	chatRepo repo.ChatRepository,
	userRepo repo.UserRepository) (req *pb_dist.ChatInviteNotificationReq, err error) {
	/*
		1. 获取邀请人信息
		2. 获取被邀请人serverId
		3. 如果是群聊则获取群信息
		4. 构建请求参数


	*/
	var (
		userInfo    *pb_user.BasicUserInfo
		chatInfo    *pb_chat.ChatInfo
		userSrvMaps map[int64]int64
	)
	// 1、获取邀请人信息
	userInfo, err = cr_user.GetBasicUserInfo(userCache, userRepo, inviteReq.InitiatorUid)
	if err != nil {
		xlog.Warn(ERROR_CODE_CHAT_INVITE_GET_INVITER_INFO_FAILED, ERROR_CHAT_INVITE_GET_INVITER_INFO_FAILED, err.Error())
		return
	}
	if userInfo.Uid == 0 {
		xlog.Warn(ERROR_CODE_CHAT_INVITE_GET_INVITER_INFO_FAILED, ERROR_CHAT_INVITE_GET_INVITER_INFO_FAILED)
		return
	}
	// 2、获取被邀请人serverId
	userSrvMaps, err = cr_user.GetUserServerList(userCache, userRepo, inviteReq.InviteeUids)
	if err != nil {
		xlog.Warn(ERROR_CODE_CHAT_INVITE_GET_INVITEE_INFO_FAILED, ERROR_CHAT_INVITE_GET_INVITEE_INFO_FAILED, err.Error())
		return
	}
	if len(userSrvMaps) == 0 {
		xlog.Warn(ERROR_CODE_CHAT_INVITE_GET_INVITEE_INFO_FAILED, ERROR_CHAT_INVITE_GET_INVITEE_INFO_FAILED)
		return
	}
	if inviteReq.ChatType == pb_enum.CHAT_TYPE_GROUP {
		// 3、获取群信息
		chatInfo, err = cr_chat.GetGroupChatInfo(chatCache, chatRepo, inviteReq.ChatId)
		if err != nil {
			xlog.Warn(ERROR_CODE_CHAT_INVITE_GET_CHAT_INFO_FAILED, ERROR_CHAT_INVITE_GET_CHAT_INFO_FAILED, err.Error())
			return
		}
		if chatInfo.ChatId == 0 {
			xlog.Warn(ERROR_CODE_CHAT_INVITE_GET_CHAT_INFO_FAILED, ERROR_CHAT_INVITE_GET_CHAT_INFO_FAILED)
			return
		}
	}
	// 4、构建请求参数
	var (
		invite        *po.ChatInvite
		ok            bool
		initiatorInfo = new(pb_invite.InitiatorInfo)
		serverId      int64
	)
	req = &pb_dist.ChatInviteNotificationReq{
		SenderId:       inviteReq.InitiatorUid,
		SenderPlatform: inviteReq.Platform,
		Notifications:  make([]*pb_dist.ChatInviteNotification, 0),
	}
	copier.Copy(initiatorInfo, userInfo)
	for _, invite = range invitees {
		ci := &pb_msg.ChatInvite{
			InviteId:      invite.InviteId,
			CreatedTs:     invite.CreatedTs,
			InviteeUid:    invite.InviteeUid,
			ChatType:      pb_enum.CHAT_TYPE(invite.ChatType),
			InvitationMsg: invite.InvitationMsg,
			InitiatorInfo: initiatorInfo,
			ChatInfo:      nil,
		}
		if serverId, ok = userSrvMaps[invite.InviteeUid]; ok == false {
			continue
		}
		if pb_enum.CHAT_TYPE(invite.ChatType) == pb_enum.CHAT_TYPE_GROUP {
			ci.ChatInfo = chatInfo
		}
		notification := &pb_dist.ChatInviteNotification{
			ReceiverId:       invite.InviteeUid,
			ReceiverServerId: serverId,
			Invite:           ci,
		}
		req.Notifications = append(req.Notifications, notification)
	}
	return
}
