package ctrl_order

import (
	"GIM/apps/interfaces/internal/service/svc_order"
	"GIM/pkg/xhttp"
	"github.com/gin-gonic/gin"
)

type OrderCtrl struct {
	orderService svc_order.OrderService
}

func NewOrderCtrl(orderService svc_order.OrderService) *OrderCtrl {
	return &OrderCtrl{orderService: orderService}
}

func (ctrl *OrderCtrl) Edit(ctx *gin.Context) {
	xhttp.Success(ctx, nil)
}

func (ctrl *OrderCtrl) Info(ctx *gin.Context) {
	xhttp.Success(ctx, nil)
}
