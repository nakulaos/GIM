package dig

import (
	"GIM/apps/interfaces/internal/service/svc_red_env_receive"
)

func init() {
	Provide(svc_red_env_receive.NewRedEnvReceiveService)
}
