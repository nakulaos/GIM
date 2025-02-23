package server

import (
	"GIM/apps/chat/internal/server/chat"
	"GIM/pkg/commands"
)

type server struct {
	chatServer chat.ChatServer
}

func NewServer(chatServer chat.ChatServer) commands.MainInstance {
	return &server{chatServer: chatServer}
}

func (s *server) Initialize() (err error) {

	return
}

func (s *server) RunLoop() {
	s.chatServer.Run()
}

func (s *server) Destroy() {

}
