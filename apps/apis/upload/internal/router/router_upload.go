package router

import (
	"GIM/apps/apis/upload/dig"
	"GIM/apps/apis/upload/internal/ctrl"
	"GIM/apps/apis/upload/internal/service"
	"github.com/gin-gonic/gin"
)

func registerUploadRouter(group *gin.RouterGroup) {
	var svc service.UploadService
	dig.Invoke(func(s service.UploadService) {
		svc = s
	})
	ctrl := ctrl.NewUploadCtrl(svc)
	router := group.Group("upload")
	router.POST("avatar", ctrl.UploadAvatar)
	router.GET("presigned", ctrl.Presigned)
}
