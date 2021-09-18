package service

import (
	"github.com/zngue/go_helper/pkg"
	"github.com/zngue/go_user_login/app/model"
	"github.com/zngue/go_user_login/app/request"
)

type UserAction interface {
	Detail(request *request.UserActionRequest) (*model.UserAction, error)
	List(request request.UserActionRequest) (*[]model.UserAction, error)
}

type userAction struct {
}

func (u *userAction) List(request request.UserActionRequest) (*[]model.UserAction, error) {
	panic("implement me")
}

func (u *userAction) Detail(request *request.UserActionRequest) (*model.UserAction, error) {
	var one model.UserAction
	request.ReturnType = 3
	err := request.Common(pkg.MysqlConn.Model(&model.UserAction{})).First(&one).Error
	return &one, err
}

func NewUserAction() UserAction {

	return new(userAction)
}
