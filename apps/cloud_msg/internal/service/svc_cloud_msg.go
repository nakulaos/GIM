package service

import (
	"GIM/apps/cloud_msg/internal/config"
	"GIM/pkg/common/xkafka"
	"GIM/pkg/obj"
	"github.com/IBM/sarama"
)

type CloudMessageService interface {
}

type cloudMessageService struct {
	cfg           *config.Config
	consumerGroup *xkafka.MConsumerGroup
	msgHandle     map[string]obj.KafkaMessageHandler
	msgCount      int64
}

func NewCloudMessageService(cfg *config.Config) CloudMessageService {
	svc := &cloudMessageService{cfg: cfg}
	svc.msgHandle = make(map[string]obj.KafkaMessageHandler)
	svc.msgHandle[cfg.MsgConsumer.Topic[0]] = svc.MessageHandler

	svc.consumerGroup = xkafka.NewMConsumerGroup(&xkafka.MConsumerGroupConfig{KafkaVersion: sarama.V3_2_1_0, OffsetsInitial: sarama.OffsetNewest, IsReturnErr: false},
		cfg.MsgConsumer.Topic,
		cfg.MsgConsumer.Address,
		cfg.MsgConsumer.GroupID)
	svc.consumerGroup.RegisterHandler(svc)

	return svc
}
