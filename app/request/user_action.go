package request

import (
	"github.com/zngue/go_helper/pkg"
	"gorm.io/gorm"
)

type UserActionRequest struct {
	pkg.CommonRequest
	UserStr string `form:"userStr" field:"user_str" where:"eq" default:""`
	VCode   int    `form:"vCode" field:"code" where:"eq" default:"0"`
	VAction string `form:"vAction" field:"action" where:"eq" default:""`
}

func (u *UserActionRequest) Common(db *gorm.DB) *gorm.DB {
	return u.Init(db, *u)
}
