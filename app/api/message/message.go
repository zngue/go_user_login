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
	token := ctx.DefaultQuery("id", "1")
	if len(token) == 0 {
		return
	}
	service.NewMessage().SetMessage(ctx, token)
}
func QrcodeCreate(ctx *gin.Context) {
	token := ctx.DefaultQuery("token", "znhjaldakljdfsdad")
	login := ctx.DefaultQuery("action", "login")
	code, randStr := service.NewMessage().QrcodeCreate(token, login)
	rep.Success(ctx, rep.Data(Data{
		ImgUrl:  code,
		RandStr: randStr,
	}))
}
