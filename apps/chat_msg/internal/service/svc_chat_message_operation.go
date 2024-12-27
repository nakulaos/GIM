package service

import (
	"GIM/domain/po"
	"GIM/pkg/common/xlog"
	"GIM/pkg/constant"
	"GIM/pkg/entity"
	"GIM/pkg/proto/pb_chat_msg"
	"GIM/pkg/proto/pb_enum"
	"GIM/pkg/proto/pb_mq"
	"GIM/pkg/proto/pb_msg"
	"GIM/pkg/utils"
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/proto"
)

func (s *chatMessageService) MessageOperation(ctx context.Context, req *pb_chat_msg.MessageOperationReq) (resp *pb_chat_msg.MessageOperationResp, err error) {
	/*
		主要作用是处理消息的前置操作
		1. 判断是否超过10分钟
		2. 判断是否有权限操作
		3. 消息推送
	*/

	resp = new(pb_chat_msg.MessageOperationResp)
	var (
		message *po.Message
		nowTs   = utils.NowUnix()
		inbox   = &pb_mq.InboxMessage{
			Topic:    pb_enum.TOPIC_MSG_OPR,
			SubTopic: pb_enum.SUB_TOPIC_CHAT_OPERATION,
		}
		operation = new(pb_msg.MessageOperation)
	)
	message, err = s.GetMessage(req.ChatId, req.SeqId)
	if err != nil {
		resp.Set(ERROR_CODE_CHAT_MSG_QUERY_DB_FAILED, ERROR_CHAT_MSG_QUERY_DB_FAILED)
		return
	}
	// 1、超过10分钟无法测回
	if nowTs-message.SrvTs > constant.CONST_SECOND_10_MINUTES {
		resp.Set(ERROR_CODE_CHAT_MSG_BEYOND_OPERABLE_TIME, ERROR_CHAT_MSG_BEYOND_OPERABLE_TIME)
		xlog.Warn(resp.Code, resp.Msg)
		return
	}
	switch pb_enum.MSG_OPERATION(message.Status) {
	case pb_enum.MSG_OPERATION_RECALL:
		// 重复操作
		return
	}
	// 2、成员身份判断
	if message.SenderId != req.SenderId {
		// 无权操作
		return
	}

	// 3、消息推送
	copier.Copy(operation, req)
	operation.SrvMsgId = message.SrvMsgId
	operation.ChatType = pb_enum.CHAT_TYPE(message.ChatType)
	inbox.Body, _ = proto.Marshal(operation)
	_, _, err = s.producer.EnQueue(inbox, constant.CONST_MSG_KEY_OPERATION+utils.GetChatPartition(req.ChatId))
	if err != nil {
		resp.Set(ERROR_CODE_CHAT_MSG_ENQUEUE_FAILED, ERROR_CHAT_MSG_ENQUEUE_FAILED)
		xlog.Warn(resp.Code, resp.Msg, err.Error())
		return
	}
	return
}

func (s *chatMessageService) GetMessage(chatId int64, seqId int64) (message *po.Message, err error) {
	var (
		w = entity.NewNormalQuery()
	)
	message, _ = s.chatMessageCache.GetChatMessage(chatId, seqId)
	if message.SrvMsgId > 0 {
		return
	}
	w.SetFilter("chat_id=?", chatId)
	w.SetFilter("seq_id=?", seqId)
	message, err = s.chatMessageRepo.Message(w)
	if err != nil {
		xlog.Warn(ERROR_CODE_CHAT_MSG_QUERY_DB_FAILED, ERROR_CHAT_MSG_QUERY_DB_FAILED, err.Error())
		return
	}
	return
}
