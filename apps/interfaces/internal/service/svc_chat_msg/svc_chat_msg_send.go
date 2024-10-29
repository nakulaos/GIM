package svc_chat_msg

import (
	"GIM/apps/interfaces/internal/dto/dto_chat_msg"
	"GIM/pkg/common/xlog"
	"GIM/pkg/common/xsnowflake"
	"GIM/pkg/proto/pb_enum"
	"GIM/pkg/proto/pb_msg"
	"GIM/pkg/utils"
	"GIM/pkg/xhttp"
	"github.com/jinzhu/copier"
)

func (s *chatMessageService) SendChatMessage(req *dto_chat_msg.SendChatMessageReq, uid int64, platform int32) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		reqArgs = &pb_msg.SendChatMessageReq{Msg: new(pb_msg.CliChatMessage)}
		reply   *pb_msg.SendChatMessageResp
	)
	copier.Copy(reqArgs.Msg, req)
	reqArgs.Topic = pb_enum.TOPIC_CHAT
	reqArgs.SubTopic = pb_enum.SUB_TOPIC_CHAT_MSG
	reqArgs.Msg.CliMsgId = xsnowflake.NewSnowflakeID()
	reqArgs.Msg.SentTs = utils.NowUnix()
	reqArgs.Msg.SenderId = uid
	reqArgs.Msg.SenderPlatform = pb_enum.PLATFORM_TYPE(platform)
	reply = s.msgClient.SendChatMessage(reqArgs)
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
