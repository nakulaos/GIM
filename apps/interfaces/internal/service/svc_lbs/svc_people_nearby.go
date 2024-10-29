package svc_lbs

import (
	"GIM/apps/interfaces/internal/dto/dto_lbs"
	"GIM/pkg/common/xlog"
	"GIM/pkg/proto/pb_lbs"
	"GIM/pkg/xhttp"
	"github.com/jinzhu/copier"
)

func (s *lbsService) PeopleNearby(params *dto_lbs.PeopleNearbyReq, uid int64) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		req   = new(pb_lbs.PeopleNearbyReq)
		reply *pb_lbs.PeopleNearbyResp
	)
	copier.Copy(req, params)
	req.Uid = uid
	reply = s.lbsClient.PeopleNearby(req)
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
	resp.Data = reply
	return
}
