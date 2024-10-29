package server

import (
	"GIM/apps/chat_msg/internal/server/chat_msg"
	"GIM/pkg/commands"
)

type server struct {
	chatMessageServer chat_msg.ChatMessageServer
}

func NewServer(chatMessageServer chat_msg.ChatMessageServer) commands.MainInstance {
	return &server{chatMessageServer: chatMessageServer}
}

func (s *server) Initialize() (err error) {
	return
}

func (s *server) RunLoop() {
	s.chatMessageServer.Run()
}

func (s *server) Destroy() {

}
