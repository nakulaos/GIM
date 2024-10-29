package dig

import (
	"GIM/apps/interfaces/internal/service/svc_red_env"
)

func init() {
	Provide(svc_red_env.NewRedEnvService)
}
