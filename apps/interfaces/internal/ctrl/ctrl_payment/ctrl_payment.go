package ctrl_payment

import (
	"GIM/apps/interfaces/internal/service/svc_payment"
	"GIM/pkg/xhttp"
	"github.com/gin-gonic/gin"
)

type PaymentCtrl struct {
	paymentService svc_payment.PaymentService
}

func NewPaymentCtrl(paymentService svc_payment.PaymentService) *PaymentCtrl {
	return &PaymentCtrl{paymentService: paymentService}
}

func (ctrl *PaymentCtrl) Edit(ctx *gin.Context) {
	xhttp.Success(ctx, nil)
}

func (ctrl *PaymentCtrl) Info(ctx *gin.Context) {
	xhttp.Success(ctx, nil)
}
