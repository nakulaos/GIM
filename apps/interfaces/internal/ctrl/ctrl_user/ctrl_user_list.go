package ctrl_user

import (
	"GIM/apps/interfaces/internal/dto/dto_user"
	"GIM/pkg/common/xgin"
	"GIM/pkg/common/xlog"
	"GIM/pkg/xhttp"
	"github.com/gin-gonic/gin"
)

func (ctrl *UserCtrl) UserList(ctx *gin.Context) {
	var (
		params = new(dto_user.UserListReq)
		resp   *xhttp.Resp
		err    error
	)
	if err = xgin.ShouldBindQuery(ctx, params); err != nil {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR, err.Error())
		return
	}
	resp = ctrl.userService.UserList(params)
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}
