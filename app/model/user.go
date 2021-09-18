package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint           `gorm:"primarykey"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at" form:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updated_at" form:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at" form:"deleted_at"`
	Subscribe     int32          `json:"subscribe"`
	OpenID        string         `json:"openid"`
	Nickname      string         `json:"nickname"`
	Sex           int32          `json:"sex"`
	City          string         `json:"city"`
	Country       string         `json:"country"`
	Province      string         `json:"province"`
	Language      string         `json:"language"`
	Headimgurl    string         `json:"headimgurl"`
	SubscribeTime int32          `json:"subscribe_time"`
	UnionID       string         `json:"unionid"`
	Action        string         `json:"action"`
	Type          int            `json:"type"`
	UserStr       string         `json:"user_str"`
}

func (User) TableName() string {
	return "wechat_user"
}
