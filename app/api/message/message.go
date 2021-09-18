package message

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_helper/pkg/rep"
	"github.com/zngue/go_user_login/app/service"
)

type Data struct {
	ImgUrl  string `json:"imgUrl"`
	RandStr string `json:"randStr"`
}

func Token(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "1")
	if len(id) == 0 {
		return
	}
	service.NewMessage().SetMessage(ctx, id)
}
func QrcodeCreate(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "")
	login := ctx.DefaultQuery("action", "login")
	if len(id) == 0 || len(login) == 0 {
		rep.ParameterError(ctx, rep.Msg("id action 参数必填"))
		return
	}
	code, randStr := service.NewMessage().QrcodeCreate(id, login)
	rep.Success(ctx, rep.Data(Data{
		ImgUrl:  code,
		RandStr: randStr,
	}))
}
