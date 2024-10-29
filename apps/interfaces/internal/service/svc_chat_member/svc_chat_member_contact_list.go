package svc_chat_member

import (
	"GIM/apps/interfaces/internal/dto/dto_chat_member"
	"GIM/pkg/common/xlog"
	"GIM/pkg/proto/pb_chat_member"
	"GIM/pkg/xhttp"
	"github.com/jinzhu/copier"
)

func (s *chatMemberService) ContactList(params *dto_chat_member.ContactListReq, uid int64) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		reqArgs = &pb_chat_member.GetContactListReq{}
		reply   *pb_chat_member.GetContactListResp
	)
	copier.Copy(reqArgs, params)
	reqArgs.Uid = uid

	reply = s.chatMemberClient.GetContactList(reqArgs)
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
