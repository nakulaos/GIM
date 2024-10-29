package chat_invite

import (
	"GIM/apps/chat_invite/internal/config"
	"GIM/apps/chat_invite/internal/service"
	"GIM/pkg/common/xgrpc"
	"GIM/pkg/proto/pb_invite"
	"google.golang.org/grpc"
	"io"
)

type ChatInviteServer interface {
	Run()
}

type chatInviteServer struct {
	pb_invite.UnimplementedInviteServer
	cfg            *config.Config
	grpcServer     *xgrpc.GrpcServer
	requestService service.ChatInviteService
}

func NewChatInviteServer(cfg *config.Config, requestService service.ChatInviteService) ChatInviteServer {
	return &chatInviteServer{cfg: cfg, requestService: requestService}
}

func (s *chatInviteServer) Run() {
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

	pb_invite.RegisterInviteServer(srv, s)
	s.grpcServer = xgrpc.NewGrpcServer(s.cfg.GrpcServer, s.cfg.Etcd)
	s.grpcServer.RunServer(srv)
}
