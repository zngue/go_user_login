package verification

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zngue/go_helper/pkg/jwt"
	"github.com/zngue/go_helper/pkg/rep"
	"github.com/zngue/go_user_login/app/request"
	"github.com/zngue/go_user_login/app/service"
)

func Verification(ctx *gin.Context) {
	var req request.UserActionRequest
	if err := ctx.ShouldBind(&req); err != nil {
		rep.Error(ctx, rep.Err(err))
		return
	}
	if req.UserStr == "" {
		rep.Error(ctx, rep.Msg("无效参数"), rep.Code(422))
		return
	}
	detail, err := service.NewUserAction().Detail(&req)
	if err != nil {
		rep.Error(ctx, rep.Err(err))
		return
	}
	if detail.Action == "code" && detail.ExpireTime < int(time.Now().Unix()) {
		if detail.ExpireTime < int(time.Now().Unix()) {
			rep.Error(ctx, rep.Msg("验证码失效请重新获取"))
			return
		}
	}
	var userReq request.UserRequest
	userReq.OpenID = detail.OpenID
	user, errs := service.NewUser().Detail(&userReq)
	if errs != nil {
		rep.Error(ctx, rep.Err(errs))
		return
	}
	token, errToken := new(jwt.AuthJwt).CreateToken(&user)
	if errToken != nil {
		rep.Error(ctx, rep.Err(errToken))
		return
	}
	rep.Success(ctx, rep.Data(token))
}
