package svc_payment

import (
	"GIM/apps/interfaces/internal/config"
	payment_client "GIM/apps/payment/client"
)

type PaymentService interface {
}

type paymentService struct {
	paymentClient payment_client.PaymentClient
}

func NewPaymentService(conf *config.Config) PaymentService {
	paymentClient := payment_client.NewPaymentClient(conf.Etcd, conf.PaymentServer, conf.Jaeger, conf.Name)
	return &paymentService{paymentClient: paymentClient}
}
