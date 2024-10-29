package ctrl_chat_msg

import (
	"GIM/apps/interfaces/internal/dto/dto_chat_msg"
	"GIM/pkg/common/xgin"
	"GIM/pkg/common/xlog"
	"GIM/pkg/xhttp"
	"github.com/gin-gonic/gin"
)

func (ctrl *ChatMessageCtrl) GetChatMessageList(ctx *gin.Context) {
	var (
		params = new(dto_chat_msg.GetChatMessageListReq)
		resp   *xhttp.Resp
		err    error
	)
	if err = xgin.ShouldBindQuery(ctx, params); err != nil {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR, err.Error())
		return
	}
	resp = ctrl.chatMessageService.GetChatMessageList(params)
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}

// 弃用
//func (ctrl *ChatMessageCtrl) GetChatMessages(ctx *gin.Context) {
//	var (
//		params = new(dto_chat_msg.GetChatMessagesReq)
//		resp   *xhttp.Resp
//		err    error
//	)
//	if err = ctx.ShouldBindQuery(params); err != nil {
//		xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_REQ_DESERIALIZE_FAILED, xhttp.ERROR_HTTP_REQ_DESERIALIZE_FAILED)
//		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_DESERIALIZE_FAILED, xhttp.ERROR_HTTP_REQ_DESERIALIZE_FAILED, err.Error())
//		return
//	}
//	resp = ctrl.chatMessageService.GetChatMessages(params)
//	if resp.Code > 0 {
//		xhttp.Error(ctx, resp.Code, resp.Msg)
//		return
//	}
//	xhttp.Success(ctx, resp.Data)
//}
