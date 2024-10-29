package dig

import (
	"GIM/apps/interfaces/internal/service/svc_payment"
)

func init() {
	Provide(svc_payment.NewPaymentService)
}
