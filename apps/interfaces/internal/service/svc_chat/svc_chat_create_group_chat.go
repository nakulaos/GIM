package svc_chat

import (
	"GIM/apps/interfaces/internal/dto/dto_chat"
	"GIM/pkg/common/xlog"
	"GIM/pkg/proto/pb_chat"
	"GIM/pkg/xhttp"
	"github.com/jinzhu/copier"
)

func (s *chatService) CreateGroupChat(params *dto_chat.CreateGroupChatReq, uid int64) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		reqArgs = new(pb_chat.CreateGroupChatReq)
		reply   *pb_chat.CreateGroupChatResp
	)
	copier.Copy(reqArgs, params)
	reqArgs.CreatorUid = uid
	reply = s.chatClient.CreateGroupChat(reqArgs)
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
