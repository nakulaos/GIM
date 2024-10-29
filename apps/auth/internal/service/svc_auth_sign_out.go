package service

import (
	"GIM/pkg/common/xlog"
	"GIM/pkg/proto/pb_auth"
	"context"
)

func (s *authService) SignOut(ctx context.Context, req *pb_auth.SignOutReq) (resp *pb_auth.SignOutResp, _ error) {
	resp = new(pb_auth.SignOutResp)
	var (
		err error
	)
	// 1. 移除ws对应platform的serverId
	// 2. 移除会话session
	_, _, err = s.online.UserOnline(req.Uid, 0, req.Platform)
	if err != nil {
		resp.Set(ERROR_CODE_AUTH_GRPC_SERVICE_FAILURE, ERROR_AUTH_GRPC_SERVICE_FAILURE)
		xlog.Warn(resp.Code, resp.Msg, err.Error())
		return
	}
	err = s.userCache.SignOut(req.Uid, req.Platform)
	if err != nil {
		resp.Set(ERROR_CODE_AUTH_LOGOUT_FAILED, ERROR_AUTH_LOGOUT_FAILED)
		xlog.Warn(resp.Code, resp.Msg, err.Error())
		return
	}
	return
}
