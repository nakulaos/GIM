package server

import (
	"GIM/apps/interfaces/internal/config"
	"GIM/apps/interfaces/internal/router"
	"GIM/pkg/commands"
	"GIM/pkg/common/xgin"
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
	s.ginServer.Run(s.cfg.Port)
}

func (s *server) Destroy() {

}
