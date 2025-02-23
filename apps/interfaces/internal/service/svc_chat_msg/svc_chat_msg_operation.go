package svc_chat_msg

import (
	"GIM/apps/interfaces/internal/dto/dto_chat_msg"
	"GIM/pkg/common/xlog"
	"GIM/pkg/proto/pb_chat_msg"
	"GIM/pkg/proto/pb_enum"
	"GIM/pkg/xhttp"
	"github.com/jinzhu/copier"
)

func (s *chatMessageService) MessageOperation(req *dto_chat_msg.MessageOperationReq, uid int64, platform int32) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		reqArgs = new(pb_chat_msg.MessageOperationReq)
		reply   *pb_chat_msg.MessageOperationResp
	)
	copier.Copy(reqArgs, req)
	reqArgs.SenderId = uid
	reqArgs.Platform = pb_enum.PLATFORM_TYPE(platform)
	reply = s.chatMessageClient.MessageOperation(reqArgs)
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
