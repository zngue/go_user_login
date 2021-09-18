package service

import (
	"github.com/xinliangnote/go-util/md5"
	"github.com/zngue/go_helper/pkg"
	"github.com/zngue/go_user_login/app/model"
	"github.com/zngue/go_user_login/app/request"
)

type Account interface {
	Edit(account2 *model.Account) error
	List(request *request.AccountRequest) (*[]model.Account, error)
	Detail(request *request.AccountRequest) (*model.Account, error)
}
type account struct {
}

func (a *account) Detail(request *request.AccountRequest) (*model.Account, error) {
	accountModel := pkg.MysqlConn.Model(&model.Account{})
	request.ReturnType = 3
	var accountOne model.Account
	err := request.Common(accountModel).First(&accountOne).Error
	return &accountOne, err
}

func (a *account) List(request *request.AccountRequest) (*[]model.Account, error) {
	accountModel := pkg.MysqlConn.Model(&model.Account{})
	var accountList []model.Account
	err := request.Common(accountModel).Find(&accountList).Error
	return &accountList, err
}

func (a *account) Edit(account2 *model.Account) error {
	var err error
	account2.Token = md5.MD5(account2.AccountOrigin)
	if account2.ID == 0 {
		err = pkg.MysqlConn.Model(&model.Account{}).Create(&account2).Error
	} else {
		err = pkg.MysqlConn.Model(&model.Account{}).Where("id = ?", account2.ID).Updates(map[string]interface{}{
			"name":           account2.Name,
			"token":          account2.Token,
			"desc":           account2.Desc,
			"app_key":        account2.AppKey,
			"app_id":         account2.AppID,
			"account_origin": account2.AccountOrigin,
			"account":        account2.Account,
		}).Error
	}
	return err
}

func NewAccount() Account {
	return new(account)
}
