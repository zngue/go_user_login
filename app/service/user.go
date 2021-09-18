package service

import (
	"time"

	"github.com/zngue/go_helper/pkg"
	"github.com/zngue/go_user_login/app/model"
	"github.com/zngue/go_user_login/app/request"
)

type User interface {
	Edit(user2 *model.User) error
	EditUserAndAction(user2 *model.User, action *ScanResp) error
	Detail(request *request.UserRequest) (*model.User, error)
}
type user struct {
}

func (u *user) EditUserAndAction(user2 *model.User, action *ScanResp) error {
	begin := pkg.MysqlConn.Begin()
	defer begin.Rollback()
	if user2.ID == 0 {
		begin.Model(&model.User{}).Create(&user2)
	}
	unix := time.Now().Unix()
	userAction := model.UserAction{
		Action:     action.Action,
		Type:       action.Type,
		UserStr:    action.UserStr,
		OpenID:     user2.OpenID,
		AddTime:    int(unix),
		ExpireTime: int(unix + 60*5),
		Code:       action.Code,
	}
	begin.Model(&model.UserAction{}).Create(&userAction)
	if err := begin.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (u *user) Detail(request *request.UserRequest) (*model.User, error) {
	tx := pkg.MysqlConn.Model(&model.User{})
	var one model.User
	err := request.Common(tx).First(&one).Error
	return &one, err
}

func (u *user) Edit(user2 *model.User) error {
	tx := pkg.MysqlConn.Model(&model.User{})
	if user2.ID > 0 {
		return tx.Where("id = ? ", user2.ID).Updates(Struct2Map(*user2)).Error
	} else {
		return tx.Create(user2).Error
	}
}
func NewUser() User {
	return new(user)
}
