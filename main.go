package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_helper/pkg/common_run"
	_ "github.com/zngue/go_helper/pkg/jwt"
	"github.com/zngue/go_user_login/app/router"
)

/*
*@Author Administrator
*@Date 21/5/2021 12:24
*@desc
 */
func main() {
	common_run.CommonGinRun(
		common_run.FnRouter(func(engine *gin.Engine) {
			group := engine.Group("point")
			router.Router(group)
		}),
		common_run.IsRegisterCenter(true),
	)

}
