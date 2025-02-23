package xgin

import (
	"GIM/pkg/constant"
	"GIM/pkg/utils"
	"GIM/pkg/xhttp"
	"github.com/gin-gonic/gin"
)

func GetPlatform(ctx *gin.Context) (platform int32) {
	var (
		value  any
		exists bool
	)
	value, exists = ctx.Get(constant.USER_PLATFORM)
	if exists == false {
		xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_PLATFORM_DOESNOT_EXIST, xhttp.ERROR_HTTP_PLATFORM_DOESNOT_EXIST)
		return
	}
	platform, _ = utils.ToInt32(value)
	if platform == 0 {
		xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_PLATFORM_DOESNOT_EXIST, xhttp.ERROR_HTTP_PLATFORM_DOESNOT_EXIST)
		return
	}
	return
}
