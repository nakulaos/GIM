package service

import (
	"GIM/pkg/common/xlog"
	"GIM/pkg/constant"
	"GIM/pkg/entity"
	"GIM/pkg/proto/pb_enum"
	"GIM/pkg/proto/pb_msg"
	"github.com/IBM/sarama"
	"google.golang.org/protobuf/proto"
)

// Setup is run at the beginning of a new session, before ConsumeClaim
func (s *chatService) Setup(_ sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(s.consumerGroup.Ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (s *chatService) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (s *chatService) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case msg, ok := <-claim.Messages():
			if ok == false {
				xlog.Info("message channel was closed")
				return nil
			}
			s.msgHandle[msg.Topic](msg.Value, string(msg.Key))
			session.MarkMessage(msg, "")
		case <-session.Context().Done():
			return nil
		}
	}
	return nil
}

func (s *chatService) MessageHandler(msg []byte, msgKey string) (err error) {
	switch msgKey {
	case constant.CONST_MSG_KEY_READ_RECEIPT:
		// 客户端更新已读消息
		s.updateReadReceiptSeq(msg)
		return
	case constant.CONST_MSG_KEY_CHAT_SEQ:
		s.updateChatSeq(msg)
		return
	default:
		return
	}
}

func (s *chatService) updateReadReceiptSeq(msg []byte) (err error) {
	var (
		read = new(pb_msg.ReadReceipt)
	)
	if err = proto.Unmarshal(msg, read); err != nil {
		err = nil
		return
	}
	s.updateReadSeq(read)
	return
}

func (s *chatService) updateReadSeq(read *pb_msg.ReadReceipt) {
	var (
		u = entity.NewMysqlUpdate()
	)
	u.SetFilter("chat_id=?", read.ChatId)
	u.SetFilter("uid=?", read.Uid)
	u.SetFilter("read_seq<?", read.SeqId)
	u.Set("read_seq", read.SeqId)
	s.chatMemberRepo.UpdateChatMember(u)
}

func (s *chatService) updateChatSeq(msg []byte) (err error) {
	var (
		read = new(pb_msg.ChatSeq)
		u    = entity.NewMysqlUpdate()
	)
	if err = proto.Unmarshal(msg, read); err != nil {
		err = nil
		return
	}
	u.SetFilter("chat_id=?", read.ChatId)
	u.SetFilter("seq_id<?", read.SeqId)
	u.Set("seq_id", read.SeqId)
	u.Set("srv_ts", read.SrvTs)
	s.chatRepo.UpdateChat(u)
	if read.MsgFrom != pb_enum.MSG_FROM_USER {
		return
	}
	r := &pb_msg.ReadReceipt{
		Uid:    read.SenderId,
		ChatId: read.ChatId,
		SeqId:  read.SeqId,
	}
	s.updateReadSeq(r)
	return
}
