package ctrl_lbs

import (
	"GIM/apps/interfaces/internal/service/svc_lbs"
)

type LbsCtrl struct {
	lbsService svc_lbs.LbsService
}

func NewLbsCtrl(lbsService svc_lbs.LbsService) *LbsCtrl {
	return &LbsCtrl{lbsService: lbsService}
}
