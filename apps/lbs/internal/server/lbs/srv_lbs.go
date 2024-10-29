package lbs

import (
	"GIM/apps/lbs/internal/config"
	"GIM/apps/lbs/internal/service"
	"GIM/pkg/common/xgrpc"
	"GIM/pkg/proto/pb_lbs"
	"google.golang.org/grpc"
	"io"
)

type LbsServer interface {
	Run()
}

type lbsServer struct {
	pb_lbs.UnimplementedLbsServer
	cfg        *config.Config
	lbsService service.LbsService
	grpcServer *xgrpc.GrpcServer
}

func NewLbsServer(cfg *config.Config, lbsService service.LbsService) LbsServer {
	return &lbsServer{cfg: cfg, lbsService: lbsService}
}

func (s *lbsServer) Run() {
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

	pb_lbs.RegisterLbsServer(srv, s)
	s.grpcServer = xgrpc.NewGrpcServer(s.cfg.GrpcServer, s.cfg.Etcd)
	s.grpcServer.RunServer(srv)
}
