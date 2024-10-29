package router

import (
	"GIM/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	group := engine.Group("api")
	group.Use(middleware.JwtAuth())
	registerUploadRouter(group)
}
