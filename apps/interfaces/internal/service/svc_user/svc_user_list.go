package svc_user

import (
	"GIM/apps/interfaces/internal/dto/dto_user"
	"GIM/pkg/common/xlog"
	"GIM/pkg/proto/pb_user"
	"GIM/pkg/utils"
	"GIM/pkg/xhttp"
)

func (s *userService) UserList(params *dto_user.UserListReq) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		reqArgs = &pb_user.GetBasicUserInfoListReq{Uids: utils.SplitToInt64List(params.Uids, ",")}
		reply   *pb_user.GetBasicUserInfoListResp
	)
	reply = s.userClient.GetBasicUserInfoList(reqArgs)
	if reply == nil {
		resp.SetResult(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		return
	}
	if reply.Code > 0 {
		resp.SetResult(reply.Code, reply.Msg)
		return
	}
	resp.Data = reply.List
	return
}
