package svc_red_env_receive

import (
	"GIM/apps/interfaces/internal/config"
	"GIM/apps/interfaces/internal/dto/dto_red_env_receive"
	red_env_receive_client "GIM/apps/red_env_receive/client"
	"GIM/pkg/xhttp"
)

type RedEnvReceiveService interface {
	GrabRedEnvelope(params *dto_red_env_receive.GrabRedEnvelopeReq, uid int64) (resp *xhttp.Resp)
	OpenRedEnvelope(params *dto_red_env_receive.OpenRedEnvelopeReq, uid int64) (resp *xhttp.Resp)
}

type redEnvReceiveService struct {
	redEnvReceiveClient red_env_receive_client.RedEnvReceiveClient
}

func NewRedEnvReceiveService(conf *config.Config) RedEnvReceiveService {
	redEnvReceiveClient := red_env_receive_client.NewRedEnvReceiveClient(conf.Etcd, conf.RedEnvReceiveServer, conf.Jaeger, conf.Name)
	return &redEnvReceiveService{redEnvReceiveClient: redEnvReceiveClient}
}
