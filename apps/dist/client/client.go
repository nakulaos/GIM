package dist_client

import (
	"GIM/pkg/common/xgrpc"
	"GIM/pkg/common/xlog"
	"GIM/pkg/conf"
	"GIM/pkg/proto/pb_dist"
	"context"
	"google.golang.org/grpc"
)

type DistClient interface {
	ChatInviteNotification(req *pb_dist.ChatInviteNotificationReq) (resp *pb_dist.ChatInviteNotificationResp)
}

type distClient struct {
	opt *xgrpc.ClientDialOption
}

func NewDistClient(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) DistClient {
	return &distClient{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *distClient) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}

func (c *distClient) ChatInviteNotification(req *pb_dist.ChatInviteNotificationReq) (resp *pb_dist.ChatInviteNotificationResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_dist.NewDistClient(conn)
	var err error
	resp, err = client.ChatInviteNotification(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}
