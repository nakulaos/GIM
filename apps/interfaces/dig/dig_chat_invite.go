package dig

import (
	"GIM/apps/interfaces/internal/service/svc_chat_invite"
)

func init() {
	Provide(svc_chat_invite.NewChatInviteService)
}
