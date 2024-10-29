package svc_auth

import (
	"GIM/apps/interfaces/internal/dto/dto_auth"
	"GIM/pkg/common/xlog"
	"GIM/pkg/proto/pb_auth"
	"GIM/pkg/xhttp"
	"github.com/jinzhu/copier"
)

func (s *authService) GoogleOAuth2Callback(params *dto_auth.GoogleOauthCallbackReq) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		req   = new(pb_auth.GoogleOAuth2CallbackReq)
		reply *pb_auth.GoogleOAuth2CallbackResp
	)
	_ = copier.Copy(req, params)
	reply = s.authClient.GoogleOAuth2Callback(req)
	if reply == nil {
		resp.SetResult(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		return
	}
	if reply.Code > 0 {
		resp.SetResult(reply.Code, reply.Msg)
		xlog.Warn(reply.Code, reply.Msg)
		return
	}
	resp.Data = reply.AuthUserInfo
	return
}
