package svc_auth

import (
	"GIM/apps/interfaces/internal/dto/dto_auth"
	"GIM/pkg/common/xlog"
	"GIM/pkg/proto/pb_auth"
	"GIM/pkg/xhttp"
	"github.com/jinzhu/copier"
)

func (s *authService) GithubOAuth2Callback(params *dto_auth.GithubOauthCallbackReq) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		req   = new(pb_auth.GithubOAuth2CallbackReq)
		reply *pb_auth.GithubOAuth2CallbackResp
	)
	_ = copier.Copy(req, params)
	reply = s.authClient.GithubOAuth2Callback(req)
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
