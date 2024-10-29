package dig

import (
	"GIM/apps/interfaces/internal/service/svc_chat_msg"
)

func init() {
	Provide(svc_chat_msg.NewChatMessageService)
}
