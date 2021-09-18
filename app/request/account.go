package request

import (
	"github.com/zngue/go_helper/pkg"
	"gorm.io/gorm"
)

type AccountRequest struct {
	pkg.CommonRequest
	ID            int    `form:"id" field:"id" where:"eq" default:"0"`
	Name          string `form:"name" field:"name" where:"like" default:""`
	Appid         string `form:"app_id" field:"app_id" where:"like" default:""`
	AccountOrigin string `form:"account_origin" field:"account_origin" where:"like" default:""`
	Account       string `form:"account" field:"account" where:"like" default:""`
	Token         string `form:"token" field:"token" where:"eq" default:""`
}

func (a *AccountRequest) Common(db *gorm.DB) *gorm.DB {
	return a.Init(db, *a)
}
