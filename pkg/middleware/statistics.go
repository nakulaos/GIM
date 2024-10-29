package middleware

import (
	"GIM/pkg/common/xredis"
	"GIM/pkg/constant"
	"github.com/gin-gonic/gin"
)

func Statistics() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.FullPath()
		xredis.ZIncrBy(constant.RK_SYNC_API_REQUESTS, 1, path)
		ctx.Next()
	}
}
