package ctrl_auth

import (
	"GIM/apps/interfaces/internal/dto/dto_auth"
	"GIM/pkg/common/xgin"
	"GIM/pkg/common/xlog"
	"GIM/pkg/proto/pb_enum"
	"GIM/pkg/xhttp"
	"github.com/gin-gonic/gin"
)

func (ctrl *AuthCtrl) GoogleOAuth2Callback(ctx *gin.Context) {
	var (
		params = new(dto_auth.GoogleOauthCallbackReq)
		resp   *xhttp.Resp
		err    error
	)
	if err = xgin.ShouldBindQuery(ctx, params); err != nil {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR, err.Error())
		return
	}
	params.Platform = pb_enum.PLATFORM_TYPE_WEB
	resp = ctrl.authService.GoogleOAuth2Callback(params)
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}
