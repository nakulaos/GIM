package service

import (
	"GIM/pkg/common/xlog"
	"GIM/pkg/entity"
	"GIM/pkg/proto/pb_user"
	"context"
)

func (s *userService) GetServerIdList(ctx context.Context, req *pb_user.GetServerIdListReq) (resp *pb_user.GetServerIdListResp, _ error) {
	resp = new(pb_user.GetServerIdListResp)
	var (
		w   = entity.NewMysqlQuery()
		err error
	)
	w.SetFilter("uid IN(?)", req.Uids)
	resp.List, err = s.userRepo.UserServerList(w)
	if err != nil {
		resp.Set(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED)
		xlog.Warn(resp.Code, resp.Msg, err.Error())
		return
	}
	return
}
