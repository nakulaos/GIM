package svc_lbs

import (
	"GIM/apps/interfaces/internal/config"
	"GIM/apps/interfaces/internal/dto/dto_lbs"
	lbs_client "GIM/apps/lbs/client"
	"GIM/pkg/xhttp"
)

type LbsService interface {
	ReportLngLat(params *dto_lbs.ReportLngLatReq, uid int64) (resp *xhttp.Resp)
	PeopleNearby(params *dto_lbs.PeopleNearbyReq, uid int64) (resp *xhttp.Resp)
}

type lbsService struct {
	cfg       *config.Config
	lbsClient lbs_client.LbsClient
}

func NewLbsService(conf *config.Config) LbsService {
	lbsClient := lbs_client.NewLbsClient(conf.Etcd, conf.LbsServer, conf.Jaeger, conf.Name)
	return &lbsService{cfg: conf, lbsClient: lbsClient}
}
