package service

import (
	"GIM/apps/lbs/internal/config"
	"GIM/domain/cache"
	"GIM/domain/mrepo"
	"GIM/domain/repo"
	"GIM/pkg/proto/pb_lbs"
	"context"
)

type LbsService interface {
	ReportLngLat(ctx context.Context, req *pb_lbs.ReportLngLatReq) (resp *pb_lbs.ReportLngLatResp, err error)
	PeopleNearby(ctx context.Context, req *pb_lbs.PeopleNearbyReq) (resp *pb_lbs.PeopleNearbyResp, err error)
}

type lbsService struct {
	cfg       *config.Config
	lbsRepo   mrepo.LbsRepository
	userRepo  repo.UserRepository
	locRepo   repo.UserLocationRepository
	userCache cache.UserCache
}

func NewLbsService(cfg *config.Config, lbsRepo mrepo.LbsRepository, userRepo repo.UserRepository, locRepo repo.UserLocationRepository, userCache cache.UserCache) LbsService {
	return &lbsService{cfg: cfg, lbsRepo: lbsRepo, userRepo: userRepo, locRepo: locRepo, userCache: userCache}
}
