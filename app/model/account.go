package model

import (
	"github.com/zngue/go_helper/pkg/model"
)

type Account struct {
	model.ZngModel
	ID            uint   `gorm:"primarykey"`
	Name          string `json:"name" form:"name"`
	Token         string `json:"token" form:"token"`
	AppKey        string `json:"appkey" form:"appkey"`
	AppID         string `json:"appid" form:"appid"`
	AccountOrigin string `json:"account_origin" form:"account_origin"`
	Account       string `json:"account" form:"account"`
	Desc          string `json:"desc" form:"desc"`
}

func (Account) TableName() string {
	return "wechat_account"
}
