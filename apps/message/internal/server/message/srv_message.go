package message

import (
	"GIM/apps/message/internal/config"
	"GIM/apps/message/internal/service"
	"GIM/pkg/common/xgrpc"
	"GIM/pkg/common/xkafka"
	"GIM/pkg/proto/pb_msg"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"io"
)

type MessageServer interface {
	Run()
}

type messageServer struct {
	pb_msg.UnimplementedMessageServer
	cfg            *config.Config
	grpcServer     *xgrpc.GrpcServer
	messageService service.MessageService

	validate *validator.Validate
	producer *xkafka.Producer
}

func NewMessageServer(cfg *config.Config, messageService service.MessageService) MessageServer {
	srv := &messageServer{cfg: cfg, validate: validator.New(), messageService: messageService}
	srv.producer = xkafka.NewKafkaProducer(cfg.MsgProducer.Address, cfg.MsgProducer.Topic)
	return srv
}

func (s *messageServer) Run() {
	var (
		srv    *grpc.Server
		closer io.Closer
	)
	srv, closer = xgrpc.NewServer(s.cfg.GrpcServer)
	defer func() {
		if closer != nil {
			closer.Close()
		}
	}()
	pb_msg.RegisterMessageServer(srv, s)
	s.grpcServer = xgrpc.NewGrpcServer(s.cfg.GrpcServer, s.cfg.Etcd)
	s.grpcServer.RunServer(srv)
}
