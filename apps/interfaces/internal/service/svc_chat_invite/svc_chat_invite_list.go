package svc_chat_invite

import (
	"GIM/apps/interfaces/internal/dto/dto_chat_invite"
	"GIM/pkg/common/xlog"
	"GIM/pkg/proto/pb_invite"
	"GIM/pkg/xhttp"
	"github.com/jinzhu/copier"
)

func (s *chatInviteService) ChatInviteList(params *dto_chat_invite.ChatInviteListReq, uid int64) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		reqArgs = new(pb_invite.ChatInviteListReq)
		reply   *pb_invite.ChatInviteListResp
	)
	copier.Copy(reqArgs, params)
	reqArgs.Uid = uid
	reply = s.chatInviteClient.ChatInviteList(reqArgs)
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
	resp.Data = reply.List
	return
}
