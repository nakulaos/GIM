package service

import (
	"GIM/apps/dist/internal/logic"
	"GIM/pkg/common/xlog"
	"GIM/pkg/constant"
	"GIM/pkg/proto/pb_dist"
	"GIM/pkg/proto/pb_enum"
	"GIM/pkg/proto/pb_gw"
	"GIM/pkg/proto/pb_obj"
	"GIM/pkg/utils"
	"context"
	"google.golang.org/protobuf/proto"
	"sync"
)

func (s *distService) ChatInviteNotification(ctx context.Context, req *pb_dist.ChatInviteNotificationReq) (resp *pb_dist.ChatInviteNotificationResp, err error) {
	resp = new(pb_dist.ChatInviteNotificationResp)
	var (
		index        int
		notification *pb_dist.ChatInviteNotification
		serverId     int64
		wg           = new(sync.WaitGroup)
	)
	for index, _ = range req.Notifications {
		notification = req.Notifications[index]
		var (
			distMembers map[int64][]*pb_obj.Int64Array
			body        []byte
			chatId      int64
		)
		if notification.Invite.ChatInfo == nil {
			chatId = notification.Invite.InviteId
		} else {
			chatId = notification.Invite.ChatInfo.ChatId
		}
		body, err = proto.Marshal(notification.Invite)
		if err != nil {
			xlog.Warn(ERROR_CODE_DIST_PROTOCOL_MARSHAL_FAILED, ERROR_DIST_PROTOCOL_MARSHAL_FAILED, err.Error())
			return
		}
		distMembers = logic.GetDistMembers(notification.ReceiverServerId, notification.ReceiverId, 0)
		for serverId, _ = range distMembers {
			msgReq := &pb_gw.SendTopicMessageReq{
				Topic:          pb_enum.TOPIC_CHAT_INVITE,
				SubTopic:       pb_enum.SUB_TOPIC_CHAT_INVITE_REQUEST,
				Members:        distMembers[serverId],
				SenderId:       notification.Invite.InitiatorInfo.Uid,
				SenderPlatform: 0,
				Body:           body,
			}
			s.asyncSendMessage(wg, msgReq, serverId, constant.CONST_MSG_KEY_CHAT_INVITE+utils.GetChatPartition(chatId))
		}
	}
	wg.Wait()
	return
}
