package dig

import (
	"GIM/apps/interfaces/internal/service/svc_order"
)

func init() {
	Provide(svc_order.NewOrderService)
}
