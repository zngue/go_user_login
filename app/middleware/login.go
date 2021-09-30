package middleware

import (
	"github.com/spf13/cast"
	"github.com/zngue/go_user_login/app/model"
	"github.com/zngue/go_user_login/app/request"
	"github.com/zngue/go_user_login/app/service"

	"github.com/gin-gonic/gin"
	"github.com/zngue/go_helper/pkg/api"
	"github.com/zngue/go_helper/pkg/jwt"
	"github.com/zngue/go_helper/pkg/rep"
)

var WechatUserInfo *model.User

func TokenCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if len(token) == 0 {
			rep.Error(ctx, rep.Msg("用户信息错误"), rep.Code(422))
		}
		parse, err := new(jwt.AuthJwt).Parse(token)
		if err != nil {
			api.Error(ctx, api.Err(err))
			return
		}
		userInfo := parse.UserInfo
		user := userInfo.(map[string]interface{})
		if id, ok := user["ID"]; ok {
			newID := cast.ToInt(id)
			if newID > 0 {
				ctx.Set("userid", id)
				var req request.UserRequest
				req.ID = newID
				detail, errs := service.NewUser().Detail(&req)
				if errs != nil {
					api.Error(ctx, api.Err(errs))
					return
				}
				WechatUserInfo = detail
			} else {
				api.Error(ctx, api.Msg("无效token"))
				return
			}
		} else {
			api.Error(ctx, api.Msg("无效用户信息"))
			return
		}

	}
}
