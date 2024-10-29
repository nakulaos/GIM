package svc_chat

import (
	"GIM/apps/interfaces/internal/dto/dto_chat"
	"GIM/pkg/common/xlog"
	"GIM/pkg/proto/pb_chat"
	"GIM/pkg/xhttp"
	"github.com/jinzhu/copier"
)

func (s *chatService) RemoveGroupChatMember(params *dto_chat.RemoveGroupChatMemberReq, uid int64) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		reqArgs = new(pb_chat.RemoveGroupChatMemberReq)
		reply   *pb_chat.RemoveGroupChatMemberResp
	)
	copier.Copy(reqArgs, params)
	reqArgs.Uid = uid
	reply = s.chatClient.RemoveGroupChatMember(reqArgs)
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
