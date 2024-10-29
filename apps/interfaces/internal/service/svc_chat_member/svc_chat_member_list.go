package svc_chat_member

import (
	"GIM/apps/interfaces/internal/dto/dto_chat_member"
	"GIM/pkg/common/xlog"
	"GIM/pkg/proto/pb_chat_member"
	"GIM/pkg/xhttp"
	"github.com/jinzhu/copier"
)

func (s *chatMemberService) ChatMemberList(params *dto_chat_member.ChatMemberListReq) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		req   = &pb_chat_member.GetChatMemberListReq{}
		reply *pb_chat_member.GetChatMemberListResp
	)
	copier.Copy(req, params)

	reply = s.chatMemberClient.GetChatMemberList(req)
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
