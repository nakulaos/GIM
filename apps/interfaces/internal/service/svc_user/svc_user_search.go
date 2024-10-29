package svc_user

import (
	"GIM/apps/interfaces/internal/dto/dto_user"
	"GIM/pkg/common/xlog"
	"GIM/pkg/proto/pb_user"
	"GIM/pkg/xhttp"
	"github.com/jinzhu/copier"
)

func (s *userService) Search(params *dto_user.SearchUserReq, uid int64) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		reqArgs = &pb_user.SearchUserReq{}
		reply   *pb_user.SearchUserResp
		data    = dto_user.SearchUserReqResp{List: make([]*dto_user.UserSummary, 0)}
	)
	copier.Copy(reqArgs, params)
	reqArgs.Uid = uid
	reply = s.userClient.SearchUser(reqArgs)
	if reply == nil {
		resp.SetResult(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		return
	}
	if reply.Code > 0 {
		resp.SetResult(reply.Code, reply.Msg)
		return
	}
	copier.Copy(&data.List, reply.List)
	data.Total = reply.Total
	resp.Data = data
	return
}
