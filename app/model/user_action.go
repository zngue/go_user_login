package model

import (
	"time"

	"gorm.io/gorm"
)

type UserAction struct {
	ID         uint           `gorm:"primarykey"`
	CreatedAt  time.Time      `gorm:"column:created_at" json:"created_at" form:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at" json:"updated_at" form:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at" form:"deleted_at"`
	Action     string         `json:"action"`
	Type       int            `json:"type"`
	UserStr    string         `json:"user_str"`    //随机字符串
	OpenID     string         `json:"openid"`      //OpenID
	AddTime    int            `json:"add_time"`    //添加时间
	Code       string         `json:"code"`        //如果是action  code 那么存验证码
	ExpireTime int            `json:"expire_time"` //过期时间
}

func (UserAction) TableName() string {
	return "wechat_user_action"
}
