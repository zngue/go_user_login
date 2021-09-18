package account

import (
	"github.com/zngue/go_helper/pkg/rep"
	"github.com/zngue/go_user_login/app/request"

	"github.com/gin-gonic/gin"
	"github.com/zngue/go_user_login/app/model"
	"github.com/zngue/go_user_login/app/service"
)

func Edit(ctx *gin.Context) {
	var data model.Account
	if err := ctx.ShouldBind(&data); err != nil {
		rep.Error(ctx, rep.Err(err))
		return
	}
	if err := service.NewAccount().Edit(&data); err != nil {
		rep.Error(ctx, rep.Err(err))
		return
	}
	rep.Success(ctx, rep.Msg("操作成功"))
	return
}

func List(ctx *gin.Context) {
	var re request.AccountRequest
	if err := ctx.ShouldBind(&re); err != nil {
		rep.Error(ctx, rep.Err(err))
		return
	}
	if list, err := service.NewAccount().List(&re); err != nil {
		rep.Error(ctx, rep.Err(err))
		return
	} else {
		rep.Success(ctx, rep.Data(list), rep.Msg("操作成功"))
		return
	}
}

func Detail(ctx *gin.Context) {

	var re request.AccountRequest
	if err := ctx.ShouldBind(&re); err != nil {
		rep.Error(ctx, rep.Err(err))
		return
	}
	if one, err := service.NewAccount().Detail(&re); err != nil {
		rep.Error(ctx, rep.Err(err))
		return
	} else {
		rep.Success(ctx, rep.Data(one), rep.Msg("操作成功"))
		return
	}

}
