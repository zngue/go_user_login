package request

import (
	"github.com/zngue/go_helper/pkg"
	"gorm.io/gorm"
)

type UserRequest struct {
	pkg.CommonRequest
	ID      int    `form:"id" field:"id" where:"eq" default:"0"`
	OpenID  string `form:"OpenID" field:"open_id" where:"eq" default:""`
	UserStr string `form:"userStr" field:"user_str" where:"eq" default:""`
}

func (a *UserRequest) Common(db *gorm.DB) *gorm.DB {
	return a.Init(db, *a)
}
