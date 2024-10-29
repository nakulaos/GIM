package chat

import (
	"GIM/apps/chat/internal/config"
	"GIM/apps/chat/internal/service"
	"GIM/pkg/common/xgrpc"
	"GIM/pkg/proto/pb_chat"
	"google.golang.org/grpc"
	"io"
)

type ChatServer interface {
	Run()
}

type chatServer struct {
	pb_chat.UnimplementedChatServer
	cfg         *config.Config
	grpcServer  *xgrpc.GrpcServer
	chatService service.ChatService
}

func NewChatServer(cfg *config.Config, chatService service.ChatService) ChatServer {
	return &chatServer{cfg: cfg, chatService: chatService}
}

func (s *chatServer) Run() {
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

	pb_chat.RegisterChatServer(srv, s)
	s.grpcServer = xgrpc.NewGrpcServer(s.cfg.GrpcServer, s.cfg.Etcd)
	s.grpcServer.RunServer(srv)
}
