package svc_auth

import (
	"GIM/apps/interfaces/internal/dto/dto_auth"
	"GIM/pkg/common/xlog"
	"GIM/pkg/proto/pb_auth"
	"GIM/pkg/xhttp"
	"github.com/jinzhu/copier"
)

func (s *authService) SignUp(params *dto_auth.SignUpReq) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		req      = new(pb_auth.SignUpReq)
		reply    *pb_auth.SignUpResp
		authResp = new(dto_auth.AuthResp)
	)
	copier.Copy(req, params)

	reply = s.authClient.SignUp(req)
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

	copier.Copy(authResp, reply)
	resp.Data = authResp
	return
}
