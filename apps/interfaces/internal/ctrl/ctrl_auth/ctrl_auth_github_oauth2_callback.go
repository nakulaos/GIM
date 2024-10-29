package ctrl_auth

import (
	"GIM/apps/interfaces/internal/dto/dto_auth"
	"GIM/pkg/proto/pb_enum"
	"GIM/pkg/xhttp"
	"github.com/gin-gonic/gin"
)

// 测试示例 lark/examples/github_oauth2/main.go

func (ctrl *AuthCtrl) GithubOAuth2Callback(ctx *gin.Context) {
	var (
		params = new(dto_auth.GithubOauthCallbackReq)
		resp   *xhttp.Resp
	)
	params.Code = ctx.Query("code")
	params.Platform = pb_enum.PLATFORM_TYPE_WEB
	resp = ctrl.authService.GithubOAuth2Callback(params)
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	//ctx.Redirect(http.StatusMovedPermanently, "http://localhost:9080")
	xhttp.Success(ctx, resp.Data)
}
