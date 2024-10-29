package service

import (
	"GIM/pkg/common/xlog"
	"GIM/pkg/entity"
	"GIM/pkg/proto/pb_user"
	"context"
)

func (s *userService) GetUserList(ctx context.Context, req *pb_user.GetUserListReq) (resp *pb_user.GetUserListResp, _ error) {
	/*
		1. 批量获取用户信息，只查数据库
	*/
	resp = &pb_user.GetUserListResp{List: make([]*pb_user.UserInfo, 0)}
	var (
		w   = entity.NewMysqlQuery()
		err error
	)
	w.SetFilter("uid IN(?)", req.UidList)
	resp.List, err = s.userRepo.UserList(w)
	if err != nil {
		resp.Set(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED)
		xlog.Warn(resp.Code, resp.Msg, err.Error())
		return
	}
	return
}
