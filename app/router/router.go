package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_user_login/app/api/account"
	"github.com/zngue/go_user_login/app/api/message"
	"github.com/zngue/go_user_login/app/api/verification"
)

// Router /*
func Router(group *gin.RouterGroup) {
	accountGroup := group.Group("account")
	{
		accountGroup.POST("edit", account.Edit)
		accountGroup.GET("list", account.List)
		accountGroup.GET("detail", account.Detail)
	}
	messageGroup := group.Group("message")
	{
		messageGroup.Any("token", message.Token)
		messageGroup.GET("qrcodeCreate", message.QrcodeCreate)
	}
	verificationGroup := group.Group("verification")
	{
		verificationGroup.GET("verify", verification.Verification)
	}

}
