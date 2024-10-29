package service

import (
	"GIM/pkg/common/xants"
	"GIM/pkg/common/xlog"
	"GIM/pkg/entity"
	"GIM/pkg/proto/pb_user"
	"context"
)

func (s *userService) GetBasicUserInfoList(ctx context.Context, req *pb_user.GetBasicUserInfoListReq) (resp *pb_user.GetBasicUserInfoListResp, err error) {
	/*
		1. 查询用户基本信息列表
		2. 缓存基本信息列表
	*/

	resp = &pb_user.GetBasicUserInfoListResp{List: make([]*pb_user.BasicUserInfo, 0)}
	var (
		w = entity.NewMysqlQuery()
	)
	w.SetFilter("uid IN(?)", req.Uids)
	resp.List, err = s.userRepo.BasicUserInfoList(w)
	if err != nil {
		resp.Set(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED)
		xlog.Warn(resp.Code, resp.Msg, err.Error())
		return
	}
	xants.Submit(func() {
		s.cacheBasicUsers(resp.List)
	})
	return
}
