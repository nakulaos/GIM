package server

import (
	"GIM/apps/lbs/internal/server/lbs"
	"GIM/pkg/commands"
)

type server struct {
	lbsServer lbs.LbsServer
}

func NewServer(lbsServer lbs.LbsServer) commands.MainInstance {
	return &server{lbsServer}
}

func (s *server) Initialize() (err error) {
	return
}

func (s *server) RunLoop() {
	s.lbsServer.Run()
}

func (s *server) Destroy() {

}
