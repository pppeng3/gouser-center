package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;column:username;type:varchar(16);not null"`
	Password string `gorm:"column:password;not null;"`
}

type UserExtra struct {
	gorm.Model
	UserID      uint32 `gorm:"unique;column:user_id;type:int(11) unsigned;not null"` // user表的id
	Nickname    string `gorm:"column:nickname;not null;varchar(16)"`
	Email       string `gorm:"column:email;type:varchar(32)"`
	PhoneNumber string `gorm:"column:phone_number;type:varchar(16)"`
	AvatarUrl   string `gorm:"column:avatar_url;type:varchar(255)"`
}
