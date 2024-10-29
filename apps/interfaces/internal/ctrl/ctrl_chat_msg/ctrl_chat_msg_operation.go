package ctrl_chat_msg

import (
	"GIM/apps/interfaces/internal/dto/dto_chat_msg"
	"GIM/pkg/common/xgin"
	"GIM/pkg/common/xlog"
	"GIM/pkg/xhttp"
	"github.com/gin-gonic/gin"
)

func (ctrl *ChatMessageCtrl) MessageOperation(ctx *gin.Context) {
	var (
		params   = new(dto_chat_msg.MessageOperationReq)
		uid      int64
		platform int32
		resp     *xhttp.Resp
		err      error
	)
	if err = xgin.BindJSON(ctx, params); err != nil {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR, err.Error())
		return
	}
	uid = xgin.GetUid(ctx)
	if uid == 0 {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_USER_ID_DOESNOT_EXIST, xhttp.ERROR_HTTP_USER_ID_DOESNOT_EXIST)
		return
	}
	platform = xgin.GetPlatform(ctx)
	if platform == 0 {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_PLATFORM_DOESNOT_EXIST, xhttp.ERROR_HTTP_PLATFORM_DOESNOT_EXIST)
		return
	}
	resp = ctrl.chatMessageService.MessageOperation(params, uid, platform)
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}
