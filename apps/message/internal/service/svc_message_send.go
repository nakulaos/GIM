package service

import (
	"GIM/pkg/common/xlog"
	"GIM/pkg/common/xsnowflake"
	"GIM/pkg/constant"
	"GIM/pkg/proto/pb_chat_member"
	"GIM/pkg/proto/pb_enum"
	"GIM/pkg/proto/pb_mq"
	"GIM/pkg/proto/pb_msg"
	"GIM/pkg/utils"
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/proto"
)

func (s *messageService) SendChatMessage(ctx context.Context, req *pb_msg.SendChatMessageReq) (resp *pb_msg.SendChatMessageResp, _ error) {
	/*
		1. 参数校验
		2. 重复消息校验
		3. 获取发送者信息
		4. 补充消息内容
		5. 将消息推送到kafka消息队列
	*/
	resp = new(pb_msg.SendChatMessageResp)
	var (
		inbox = &pb_mq.InboxMessage{
			Topic:    req.Topic,
			SubTopic: req.SubTopic,
		}
		chatMsg    = new(pb_msg.SrvChatMessage)
		senderInfo *pb_chat_member.ChatMemberInfo
		seqId      int64
		result     string
		ok         bool
		err        error
	)
	// 1、参数校验
	if err = s.validate.Struct(req.Msg); err != nil {
		resp.Set(ERROR_CODE_MESSAGE_VALIDATOR_FAILED, ERROR_MESSAGE_VALIDATOR_FAILED)
		xlog.Warn(resp.Code, resp.Msg, err.Error())
		return
	}
	if chatMsg.AssocId, err = s.verifyMessage(req); err != nil {
		resp.Set(ERROR_CODE_MESSAGE_VALIDATOR_FAILED, ERROR_MESSAGE_VALIDATOR_FAILED)
		xlog.Warn(resp.Code, resp.Msg, err.Error())
		return
	}
	// 2、重复消息校验
	result, ok = s.chatMessageCache.RepeatMessageVerify(s.cfg.Redis.Prefix, req.Msg.ChatId, req.Msg.CliMsgId)
	if ok == false {
		resp.Set(ERROR_CODE_MESSAGE_VALIDATOR_FAILED, result)
		xlog.Warn(resp.Code, resp.Msg)
		return
	}

	// 3、获取发送者信息
	senderInfo, err = s.getSenderInfo(req.Msg.ChatId, req.Msg.SenderId, resp)
	if err != nil {
		resp.Set(ERROR_CODE_MESSAGE_GET_SENDER_INFO_FAILED, err.Error())
		xlog.Warn(resp.Code, resp.Msg, err.Error())
		return
	}
	if senderInfo.Uid == 0 {
		if resp.Code > 0 {
			xlog.Warn(resp.Code, resp.Msg)
		} else {
			xlog.Warn(ERROR_CODE_MESSAGE_GET_SENDER_INFO_FAILED, ERROR_MESSAGE_GET_SENDER_INFO_FAILED)
		}
		return
	}

	// 4、补充消息内容
	if seqId, err = s.chatMessageCache.IncrSeqID(req.Msg.ChatId); err != nil {
		resp.Set(ERROR_CODE_MESSAGE_INCR_SEQ_ID_FAILED, ERROR_MESSAGE_INCR_SEQ_ID_FAILED)
		xlog.Warn(resp.Code, resp.Msg, err.Error())
		return
	}
	copier.Copy(chatMsg, req.Msg)
	chatMsg.SrvMsgId = xsnowflake.NewSnowflakeID()
	chatMsg.ChatType = senderInfo.ChatType
	chatMsg.SeqId = seqId
	chatMsg.SrvTs = utils.NowUnix()
	chatMsg.MsgFrom = pb_enum.MSG_FROM_USER
	chatMsg.SenderName = senderInfo.Alias
	chatMsg.SenderAvatar = senderInfo.MemberAvatar
	inbox.Body, _ = proto.Marshal(chatMsg)
	// 5、将消息推送到kafka消息队列
	if s.producer == nil {
		resp.Set(ERROR_CODE_MESSAGE_PRODUCER_IS_NULL, ERROR_MESSAGE_PRODUCER_IS_NULL)
		xlog.Warn(resp.Code, resp.Msg)
		return
	}
	_, _, err = s.producer.EnQueue(inbox, constant.CONST_MSG_KEY_MSG+utils.GetChatPartition(chatMsg.ChatId))
	if err != nil {
		resp.Set(ERROR_CODE_MESSAGE_ENQUEUE_FAILED, ERROR_MESSAGE_ENQUEUE_FAILED)
		xlog.Warn(resp.Code, resp.Msg, err.Error())
		return
	}
	chatSeq := &pb_msg.ChatSeq{
		ChatId:   chatMsg.ChatId,
		SeqId:    seqId,
		SrvTs:    chatMsg.SrvTs,
		SenderId: chatMsg.SenderId,
		MsgFrom:  chatMsg.MsgFrom,
	}

	// 通过消息队列异步更新user的read_seq，以及chat的seq_id
	_, _, err = s.seqProducer.EnQueue(chatSeq, constant.CONST_MSG_KEY_CHAT_SEQ+utils.GetChatPartition(chatMsg.ChatId))
	if err != nil {
		xlog.Warn(ERROR_CODE_MESSAGE_ENQUEUE_FAILED, ERROR_MESSAGE_ENQUEUE_FAILED, err.Error())
	}
	return
}

func (s *messageService) getSenderInfo(chatId int64, uid int64, resp *pb_msg.SendChatMessageResp) (info *pb_chat_member.ChatMemberInfo, err error) {
	var (
		req   *pb_chat_member.GetChatMemberInfoReq
		reply *pb_chat_member.GetChatMemberInfoResp
	)
	info, err = s.chatMemberCache.GetChatMemberInfo(chatId, uid)
	if err != nil {
		xlog.Warn(ERROR_CODE_MESSAGE_REDIS_GET_FAILED, ERROR_MESSAGE_REDIS_GET_FAILED, err.Error())
	}
	if info.Uid > 0 {
		err = s.authentication(info)
		if err != nil {
			return
		}
		return
	}
	req = &pb_chat_member.GetChatMemberInfoReq{
		ChatId: chatId,
		Uid:    uid,
	}
	reply = s.chatMemberClient.GetChatMemberInfo(req)
	if reply == nil {
		resp.Set(ERROR_CODE_MESSAGE_GRPC_SERVICE_FAILURE, ERROR_MESSAGE_GRPC_SERVICE_FAILURE)
		xlog.Warn(resp.Code, resp.Msg)
		return
	}
	if reply.Code > 0 {
		resp.Set(reply.Code, reply.Msg)
		xlog.Warn(resp.Code, resp.Msg)
		return
	}
	err = s.authentication(reply.Info)
	if err != nil {
		return
	}
	info = reply.Info
	return
}

func (s *messageService) authentication(info *pb_chat_member.ChatMemberInfo) (err error) {
	switch pb_enum.CHAT_STATUS(info.Status) {
	case pb_enum.CHAT_STATUS_QUITTED, pb_enum.CHAT_STATUS_DELETED, pb_enum.CHAT_STATUS_REMOVED:
		// 非联 Chat 成员
		err = ERROR_MESSAGE_NON_CHAT_MEMBER_ERR
	case pb_enum.CHAT_STATUS_NON_CONTACT:
		// 非联系人
		err = ERROR_MESSAGE_NON_CONTACT_ERR
	}
	return
}
