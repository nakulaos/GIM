package svc_order

import (
	"GIM/apps/interfaces/internal/config"
	"GIM/apps/interfaces/internal/dto/dto_order"
	order_client "GIM/apps/order/client"
	"GIM/pkg/xhttp"
)

type OrderService interface {
	CreateRedEnvelopeOrder(params *dto_order.CreateRedEnvelopeOrderReq, uid int64) (resp *xhttp.Resp)
}

type orderService struct {
	orderClient order_client.OrderClient
}

func NewOrderService(conf *config.Config) OrderService {
	orderClient := order_client.NewOrderClient(conf.Etcd, conf.OrderServer, conf.Jaeger, conf.Name)
	return &orderService{orderClient: orderClient}
}
