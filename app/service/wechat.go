package service

import (
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/config"
	accountWechat "github.com/zngue/go_helper/pkg/wechat"
	"github.com/zngue/go_user_login/app/model"
)

type Wechat interface {
	AccountOffice(account2 *model.Account) *officialaccount.OfficialAccount
}
type wechat struct {
}

func (w *wechat) AccountOffice(account2 *model.Account) *officialaccount.OfficialAccount {
	commonWechat := accountWechat.NewWechatClient().CommonWechat()
	memory := cache.NewMemory()
	c := &config.Config{
		AppID:     account2.AppID,
		AppSecret: account2.AppKey,
		Token:     account2.Token,
		Cache:     memory,
	}
	return commonWechat.GetOfficialAccount(c)
}

func NewWechat() Wechat {
	return new(wechat)
}
