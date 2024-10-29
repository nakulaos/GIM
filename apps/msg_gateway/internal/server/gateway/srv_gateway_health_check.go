package gateway

import (
	"GIM/pkg/proto/pb_gw"
	"context"
)

func (s *gatewayServer) HealthCheck(ctx context.Context, req *pb_gw.HealthCheckReq) (resp *pb_gw.HealthCheckResp, err error) {
	resp = new(pb_gw.HealthCheckResp)
	return
}
