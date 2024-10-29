package dig

import (
	"GIM/apps/interfaces/internal/service/svc_user"
)

func init() {
	Provide(svc_user.NewUserService)
}
