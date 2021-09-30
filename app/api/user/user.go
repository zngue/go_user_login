package user

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/zngue/go_helper/pkg/api"
	"github.com/zngue/go_user_login/app/request"
	"github.com/zngue/go_user_login/app/service"
)

func Detail(ctx *gin.Context) {

}

func GetUserInfo(ctx *gin.Context) {
	userid, exists := ctx.Get("userid")
	if !exists {
		api.Error(ctx, api.Msg("无效用户参数"))
		return
	}
	var userReq request.UserRequest
	userReq.ID = cast.ToInt(userid)
	detail, err := service.NewUser().Detail(&userReq)
	api.DataWithErr(ctx, err, detail)
	return

}
