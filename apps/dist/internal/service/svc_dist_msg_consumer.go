package service

import (
	"github.com/Shopify/sarama"
	"lark/apps/dist/internal/logic"
	gw_client "lark/apps/msg_gateway/client"
	"lark/pkg/common/xkafka"
	"lark/pkg/common/xlog"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_obj"
)

func (s *distService) Setup(_ sarama.ConsumerGroupSession) error {
	close(s.consumerGroup.Ready)
	return nil
}
func (s *distService) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (s *distService) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	var (
		msg *sarama.ConsumerMessage
		err error
	)
	for {
		select {
		case msg = <-claim.Messages():
			if msg == nil {
				continue
			}
			if err = s.msgHandle[msg.Topic](msg.Value, string(msg.Key)); err != nil {
				continue
			}
			session.MarkMessage(msg, "")
		case <-session.Context().Done():
			return nil
		}
	}
	return nil
}

func (s *distService) MessageHandler(msg []byte, msgKey string) (err error) {
	switch msgKey {
	case constant.CONST_MSG_KEY_MSG:
		s.sendChatMessage(msg)
	case constant.CONST_MSG_KEY_OPERATION:
		s.messageOperation(msg)
	default:
	}
	return
}

func (s *distService) getChatMembers(chatId int64) (serverMembers map[int64][]*pb_obj.Int64Array) {
	var (
		hashmap map[string]string
	)
	// 1万人占用 753KB Redis Memory
	hashmap = s.chatMemberCache.GetAllDistChatMembers(chatId)
	if len(hashmap) > 0 {
		serverMembers = logic.GetMembersFromHash(hashmap)
	} else {
		hashmap = s.getMemberList(chatId)
		serverMembers = logic.GetMembersFromHash(hashmap)
	}
	return
}

func (s *distService) getMemberList(chatId int64) (members map[string]string) {
	var (
		userListReq = &pb_chat_member.GetDistMemberListReq{ChatId: chatId}
		resp        *pb_chat_member.GetDistMemberListResp
	)
	resp = s.chatMemberClient.GetDistMemberList(userListReq)
	if resp == nil {
		xlog.Warn(ERROR_CODE_DIST_GRPC_SERVICE_FAILURE, ERROR_DIST_GRPC_SERVICE_FAILURE)
		return
	}
	if resp.Code > 0 {
		xlog.Warn(resp.Code, resp.Msg)
		return
	}
	return resp.Members
}

func (s *distService) getClient(serverId int64) (client gw_client.MessageGatewayClient) {
	s.mutex.RLock()
	client, _ = s.clients[serverId]
	s.mutex.RUnlock()
	return
}

func (s *distService) getProducer(serverId int64) (producer *xkafka.Producer) {
	s.mutex.RLock()
	producer, _ = s.producers[serverId]
	s.mutex.RUnlock()
	return
}
