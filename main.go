package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/zngue/go_helper/pkg"
	"github.com/zngue/go_helper/pkg/sign_chan"
	"github.com/zngue/go_user_login/app/model"
	"github.com/zngue/go_user_login/app/router"
)

/*
*@Author Administrator
*@Date 21/5/2021 12:24
*@desc
 */
func main() {
	if conErr := pkg.NewConfig(); conErr != nil {
		logrus.Fatal(conErr)
		return
	}
	mysql, err := pkg.NewMysql()
	if err != nil {
		logrus.Fatal(err)
		return
	}
	if mysql != nil {

		mysql.AutoMigrate(new(model.Account), new(model.User), new(model.UserAction))
		//auto.Auto(mysql)
	}
	port := viper.GetString("AppPort")
	run, errs := pkg.GinRun(port, func(engine *gin.Engine) {
		group := engine.Group("point")
		router.Router(group)
	})
	if errs != nil {
		sign_chan.SignLog(errs)
	}
	go func() {
		err = run.ListenAndServe()
		if err != nil {
			sign_chan.SignLog(err)
		}
	}()
	sign_chan.SignChalNotify()
	sign_chan.ListClose(func(ctx context.Context) error {
		return run.Shutdown(ctx)
	})
}
