package dig

import (
	"GIM/apps/interfaces/internal/service/svc_convo"
)

func init() {
	Provide(svc_convo.NewConvoService)
}
