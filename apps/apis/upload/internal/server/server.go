package server

import (
	"GIM/apps/apis/upload/internal/config"
	"GIM/apps/apis/upload/internal/router"
	"GIM/pkg/commands"
	"GIM/pkg/common/xetcd"
	"GIM/pkg/common/xgin"
	"GIM/pkg/common/xlog"
	"GIM/pkg/utils"
	"flag"
)

type server struct {
	ginServer *xgin.GinServer
	cfg       *config.Config
}

func NewServer() commands.MainInstance {
	return &server{cfg: config.NewConfig()}
}

func (s *server) Initialize() (err error) {
	flag.Parse()
	s.ginServer = xgin.NewGinServer()
	router.Register(s.ginServer.Engine)
	return
}

func (s *server) RunLoop() {
	err := xetcd.RegisterEtcd(s.cfg.Etcd.Schema, s.cfg.Etcd.Endpoints, utils.GetServerIP(), s.cfg.Port, s.cfg.Name, xetcd.TIME_TO_LIVE)
	if err != nil {
		xlog.Error(err.Error())
		return
	}
	s.ginServer.Run(s.cfg.Port)
}

func (s *server) Destroy() {

}
