package service

import (
	"GIM/domain/do"
	"GIM/pkg/common/xlog"
	"GIM/pkg/constant"
	"GIM/pkg/proto/pb_enum"
	"GIM/pkg/utils"
	"github.com/IBM/sarama"
	"github.com/spf13/cast"
)

func (s *cacheService) Setup(_ sarama.ConsumerGroupSession) error {
	close(s.consumerGroup.Ready)
	return nil
}
func (s *cacheService) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (s *cacheService) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
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

func (s *cacheService) MessageHandler(msg []byte, key string) (err error) {
	// 这一部分的缓存主要用来兜底的
	if len(msg) == 0 {
		return
	}
	switch key {
	case constant.CONST_MSG_KEY_CACHE_AGREE_INVITATION:
		var (
			obj    = new(do.KeyMaps)
			chatId int64
			slot   int
		)
		utils.ByteToObj(msg, obj)
		chatId = cast.ToInt64(obj.Key)
		slot = cast.ToInt(obj.Ex)
		err = s.chatMemberCache.HMSetChatMembers(chatId, slot, obj.Maps)
	case constant.CONST_MSG_KEY_CACHE_REMOVE_CHAT_MEMBER:
		var (
			obj = map[string][]string{}
		)
		utils.ByteToObj(msg, obj)
		err = s.chatMemberCache.HDelChatMembers(obj)
	case constant.CONST_MSG_KEY_CACHE_CREATE_GROUP_CHAT:
		var (
			obj    = new(do.KeyFieldValue)
			chatId int64
			uid    int64
			value  string
		)
		utils.ByteToObj(msg, obj)
		chatId, _ = utils.ToInt64(obj.Key)
		uid, _ = utils.ToInt64(obj.Field)
		value = utils.ToString(obj.Value)
		err = s.chatMemberCache.HSetNXChatMember(chatId, pb_enum.CHAT_TYPE_GROUP, uid, value)
	}
	return
}
