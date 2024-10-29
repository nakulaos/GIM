package service

import (
	"GIM/pkg/common/xlog"
	"GIM/pkg/entity"
	"GIM/pkg/proto/pb_chat"
	"GIM/pkg/proto/pb_enum"
	"GIM/pkg/utils"
	"context"
)

func (s *chatService) DeleteContact(ctx context.Context, req *pb_chat.DeleteContactReq) (resp *pb_chat.DeleteContactResp, _ error) {
	resp = new(pb_chat.DeleteContactResp)
	var (
		u   = entity.NewMysqlUpdate()
		err error
	)
	u.SetFilter("chat_id=?", req.ChatId)
	u.SetFilter("owner_id=?", req.Uid)
	u.SetFilter("deleted_ts=?", 0)
	u.Set("status", int32(pb_enum.CHAT_STATUS_DELETED))
	u.Set("deleted_ts", utils.NowUnix())
	/*
		1. 单聊也一样，它们同时属于同一个chat_id，只是单聊只有两个成员，群聊有多个成员
	*/
	_, err = s.removeChatMember(u, req.ChatId, []int64{req.Uid, req.ContactId}, pb_enum.CHAT_TYPE_PRIVATE)
	if err != nil {
		resp.Set(ERROR_CODE_CHAT_UPDATE_VALUE_FAILED, ERROR_CHAT_UPDATE_VALUE_FAILED)
		xlog.Warn(resp.Code, resp.Msg, err.Error())
		return
	}
	return
}
