package red_env_client

import (
	"GIM/pkg/common/xgrpc"
	"GIM/pkg/common/xlog"
	"GIM/pkg/conf"
	"GIM/pkg/proto/pb_red_env"
	"context"
	"google.golang.org/grpc"
)

// 暂不开放
type RedEnvClient interface {
	GiveRedEnvelope(req *pb_red_env.GiveRedEnvelopeReq) (resp *pb_red_env.GiveRedEnvelopeResp)
}

type redEnvClient struct {
	opt *xgrpc.ClientDialOption
}

func NewRedEnvClient(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) RedEnvClient {
	return &redEnvClient{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *redEnvClient) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}

func (c redEnvClient) GiveRedEnvelope(req *pb_red_env.GiveRedEnvelopeReq) (resp *pb_red_env.GiveRedEnvelopeResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_red_env.NewRedEnvClient(conn)
	var err error
	resp, err = client.GiveRedEnvelope(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}
