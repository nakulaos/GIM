package router

import (
	"GIM/apps/interfaces/dig"
	"GIM/apps/interfaces/internal/ctrl/ctrl_lbs"
	"GIM/apps/interfaces/internal/service/svc_lbs"
	"github.com/gin-gonic/gin"
)

func registerLbsRouter(group *gin.RouterGroup) {
	var svc svc_lbs.LbsService
	dig.Invoke(func(s svc_lbs.LbsService) {
		svc = s
	})
	ctrl := ctrl_lbs.NewLbsCtrl(svc)
	router := group.Group("lbs")
	router.POST("report_lng_lat", ctrl.ReportLngLat)
	router.GET("people_nearby", ctrl.PeopleNearby)
}
