package svc_order

import (
	"GIM/apps/interfaces/internal/dto/dto_order"
	"GIM/pkg/common/xlog"
	"GIM/pkg/proto/pb_order"
	"GIM/pkg/xhttp"
	"github.com/jinzhu/copier"
)

func (s *orderService) CreateRedEnvelopeOrder(params *dto_order.CreateRedEnvelopeOrderReq, uid int64) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		req   = new(pb_order.CreateRedEnvelopeOrderReq)
		reply *pb_order.CreateRedEnvelopeOrderResp
	)
	_ = copier.Copy(req, params)
	req.Uid = uid
	reply = s.orderClient.CreateRedEnvelopeOrder(req)
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
	resp.Data = reply.PayUrl
	return
}
