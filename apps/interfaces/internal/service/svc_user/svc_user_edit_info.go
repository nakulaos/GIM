package svc_user

import (
	"GIM/apps/interfaces/internal/dto/dto_user"
	"GIM/pkg/common/xlog"
	"GIM/pkg/proto/pb_user"
	"GIM/pkg/xhttp"
	"github.com/jinzhu/copier"
)

func (s *userService) EditUserInfo(params *dto_user.EditUserInfoReq, uid int64) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		req   = new(pb_user.EditUserInfoReq)
		reply *pb_user.EditUserInfoResp
	)
	copier.Copy(req, params)
	req.Uid = uid
	reply = s.userClient.EditUserInfo(req)
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
	return
}
