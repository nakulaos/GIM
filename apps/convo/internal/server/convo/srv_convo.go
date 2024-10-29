package convo

import (
	"GIM/apps/convo/internal/config"
	"GIM/apps/convo/internal/service"
	"GIM/pkg/common/xgrpc"
	"GIM/pkg/common/xmonitor"
	"GIM/pkg/proto/pb_convo"
	"google.golang.org/grpc"
	"io"
)

type ConvoServer interface {
	Run()
}

// conversation
type convoServer struct {
	pb_convo.UnimplementedConvoServer
	cfg          *config.Config
	convoService service.ConvoService
	grpcServer   *xgrpc.GrpcServer
}

func NewConvoServer(cfg *config.Config, convoService service.ConvoService) ConvoServer {
	srv := &convoServer{cfg: cfg, convoService: convoService}
	return srv
}

func (s *convoServer) Run() {
	var (
		srv    *grpc.Server
		closer io.Closer
	)
	xmonitor.RunMonitor(s.cfg.Monitor.Port)

	srv, closer = xgrpc.NewServer(s.cfg.GrpcServer)
	defer func() {
		if closer != nil {
			closer.Close()
		}
	}()

	pb_convo.RegisterConvoServer(srv, s)
	s.grpcServer = xgrpc.NewGrpcServer(s.cfg.GrpcServer, s.cfg.Etcd)
	s.grpcServer.RunServer(srv)
}
