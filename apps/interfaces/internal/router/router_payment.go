package router

import (
	"GIM/apps/interfaces/dig"
	"GIM/apps/interfaces/internal/ctrl/ctrl_payment"
	"GIM/apps/interfaces/internal/service/svc_payment"
	"github.com/gin-gonic/gin"
)

func registerPaymentRouter(group *gin.RouterGroup) {
	var svc svc_payment.PaymentService
	dig.Invoke(func(s svc_payment.PaymentService) {
		svc = s
	})
	ctrl := ctrl_payment.NewPaymentCtrl(svc)
	router := group.Group("payment")
	router.POST("edit", ctrl.Edit)
	router.GET("info", ctrl.Info)
}
