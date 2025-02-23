package service

import (
	msg_client "GIM/apps/message/client"
	"GIM/apps/msg_gateway/internal/config"
	"GIM/apps/msg_gateway/internal/server/websocket/ws"
	"GIM/pkg/common/xkafka"
	"github.com/go-playground/validator/v10"
)

type WsService interface {
	Run()
	MessageCallback(msg *ws.Message)
}

type wsService struct {
	conf      *config.Config
	validate  *validator.Validate
	msgClient msg_client.MsgClient
	producer  *xkafka.Producer
}

func NewWsService(conf *config.Config) WsService {
	msgClient := msg_client.NewMsgClient(conf.Etcd, conf.MessageServer, conf.GrpcServer.Jaeger, conf.Name)
	svc := &wsService{conf: conf, validate: validator.New(), msgClient: msgClient}
	svc.producer = xkafka.NewKafkaProducer(conf.MsgProducer.Address, conf.MsgReadProducer.Topic)
	return svc
}

func (s *wsService) Run() {

}
