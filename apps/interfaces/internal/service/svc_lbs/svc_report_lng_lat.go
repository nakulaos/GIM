package svc_lbs

import (
	"GIM/apps/interfaces/internal/dto/dto_lbs"
	"GIM/pkg/common/xlog"
	"GIM/pkg/proto/pb_lbs"
	"GIM/pkg/xhttp"
	"github.com/jinzhu/copier"
)

func (s *lbsService) ReportLngLat(params *dto_lbs.ReportLngLatReq, uid int64) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		req   = new(pb_lbs.ReportLngLatReq)
		reply *pb_lbs.ReportLngLatResp
	)
	copier.Copy(req, params)
	req.Uid = uid
	reply = s.lbsClient.ReportLngLat(req)
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
