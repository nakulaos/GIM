package chat_member

import (
	"GIM/apps/chat_member/internal/config"
	"GIM/apps/chat_member/internal/service"
	"GIM/pkg/common/xgrpc"
	"GIM/pkg/proto/pb_chat_member"
	"google.golang.org/grpc"
	"io"
)

type ChatMemberServer interface {
	Run()
}

type chatMemberServer struct {
	pb_chat_member.UnimplementedChatMemberServer
	cfg               *config.Config
	grpcServer        *xgrpc.GrpcServer
	chatMemberService service.ChatMemberService
}

func NewChatMemberServer(cfg *config.Config, chatMemberService service.ChatMemberService) ChatMemberServer {
	srv := &chatMemberServer{cfg: cfg, chatMemberService: chatMemberService}
	return srv
}

func (s *chatMemberServer) Run() {
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

	pb_chat_member.RegisterChatMemberServer(srv, s)
	s.grpcServer = xgrpc.NewGrpcServer(s.cfg.GrpcServer, s.cfg.Etcd)
	s.grpcServer.RunServer(srv)
}
