package server

import (
	"GIM/apps/cloud_msg/internal/server/cloud_msg"
	"GIM/pkg/commands"
)

type server struct {
	cloudMessageServer cloud_msg.CloudMessageServer
}

func NewServer(cloudMessageServer cloud_msg.CloudMessageServer) commands.MainInstance {
	return &server{cloudMessageServer: cloudMessageServer}
}

func (s *server) Initialize() (err error) {
	return
}

func (s *server) RunLoop() {
	s.cloudMessageServer.Run()
}

func (s *server) Destroy() {

}
