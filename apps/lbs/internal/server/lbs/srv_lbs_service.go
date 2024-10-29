package lbs

import (
	"GIM/pkg/proto/pb_lbs"
	"context"
)

func (s *lbsServer) ReportLngLat(ctx context.Context, req *pb_lbs.ReportLngLatReq) (resp *pb_lbs.ReportLngLatResp, err error) {
	return s.lbsService.ReportLngLat(ctx, req)
}

func (s *lbsServer) PeopleNearby(ctx context.Context, req *pb_lbs.PeopleNearbyReq) (resp *pb_lbs.PeopleNearbyResp, err error) {
	return s.lbsService.PeopleNearby(ctx, req)
}
