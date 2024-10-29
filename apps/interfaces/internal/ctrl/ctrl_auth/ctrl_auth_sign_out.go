package ctrl_auth

import (
	"GIM/apps/interfaces/internal/dto/dto_auth"
	"GIM/pkg/common/xgin"
	"GIM/pkg/common/xlog"
	"GIM/pkg/xhttp"
	"github.com/gin-gonic/gin"
)

func (ctrl *AuthCtrl) SignOut(ctx *gin.Context) {
	var (
		params   = new(dto_auth.SignOutReq)
		uid      int64
		platform int32
		resp     *xhttp.Resp
	)
	uid = xgin.GetUid(ctx)
	if uid == 0 {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_USER_ID_DOESNOT_EXIST, xhttp.ERROR_HTTP_USER_ID_DOESNOT_EXIST)
		return
	}
	platform = xgin.GetPlatform(ctx)
	if platform == 0 {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_PLATFORM_DOESNOT_EXIST, xhttp.ERROR_HTTP_PLATFORM_DOESNOT_EXIST)
		return
	}
	params.Uid = uid
	params.Platform = platform
	resp = ctrl.authService.SignOut(params)
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}
