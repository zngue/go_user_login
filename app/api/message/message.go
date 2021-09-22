package message

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_helper/pkg/rep"
	"github.com/zngue/go_user_login/app/service"
)

func Token(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "1")
	if len(id) == 0 {
		return
	}
	service.NewMessage().SetMessage(ctx, id)
}
func QrcodeCreate(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "")
	action := ctx.DefaultQuery("action", "login")
	if len(id) == 0 || len(action) == 0 {
		rep.ParameterError(ctx, rep.Msg("id action 参数必填"))
		return
	}
	loginData, err := service.NewMessage().QrcodeCreate(id, action)
	rep.DataWithErr(ctx, err, loginData)
	return
}
