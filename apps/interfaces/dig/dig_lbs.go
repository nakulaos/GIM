package dig

import (
	"GIM/apps/interfaces/internal/service/svc_lbs"
)

func init() {
	Provide(svc_lbs.NewLbsService)
}
