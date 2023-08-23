package ctrl_red_env

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/internal/dto/dto_red_env"
	"lark/pkg/common/xgin"
	"lark/pkg/common/xlog"
	"lark/pkg/xhttp"
)

func (ctrl *RedEnvCtrl) GiveRedEnvelope(ctx *gin.Context) {
	var (
		params = new(dto_red_env.GiveRedEnvelopeReq)
		resp   *xhttp.Resp
		uid    int64
		err    error
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
	params.SenderPlatform = xgin.GetPlatform(ctx)
	resp = ctrl.redEnvService.GiveRedEnvelope(params, uid)
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}
