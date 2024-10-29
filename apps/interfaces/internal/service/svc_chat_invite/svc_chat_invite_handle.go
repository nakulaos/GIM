package svc_chat_invite

import (
	"GIM/apps/interfaces/internal/dto/dto_chat_invite"
	"GIM/pkg/common/xlog"
	"GIM/pkg/proto/pb_invite"
	"GIM/pkg/xhttp"
	"github.com/jinzhu/copier"
)

func (s *chatInviteService) ChatInviteHandle(params *dto_chat_invite.ChatInviteHandleReq, uid int64) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		reqArgs = new(pb_invite.ChatInviteHandleReq)
		reply   *pb_invite.ChatInviteHandleResp
	)
	copier.Copy(reqArgs, params)
	reqArgs.HandlerUid = uid
	reply = s.chatInviteClient.ChatInviteHandle(reqArgs)
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
