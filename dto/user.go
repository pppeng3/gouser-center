package dto

import (
	"gouser_center/dal/db/model"
	user_center "gouser_center/proto/user-center"

	"gorm.io/gorm"
)

func RPCUser2ModelUser(user *user_center.User) *model.User {
	if user == nil {
		return &model.User{}
	}
	return &model.User{
		Username: user.Username,
		Password: user.Password,
		Model: gorm.Model{
			ID: uint(user.Id),
		},
	}
}

func ModelUser2RPCUser(user *model.User) *user_center.User {
	if user == nil {
		return &user_center.User{}
	}
	return &user_center.User{
		Id:       int32(user.ID),
		Username: user.Username,
		Password: user.Password,
	}
}

func RPCUserExtra2ModelUserExtra(extra *user_center.UserExtra) *model.UserExtra {
	if extra == nil {
		return &model.UserExtra{}
	}
	return &model.UserExtra{
		Model:       gorm.Model{ID: uint(extra.Id)},
		UserID:      uint32(extra.UserId),
		Nickname:    extra.Nickname,
		PhoneNumber: extra.PhoneNumber,
		Email:       extra.Email,
		AvatarUrl:   extra.AvatarUrl,
	}
}

func ModelUserExtra2RPCUserExtra(extra *model.UserExtra) *user_center.UserExtra {
	if extra == nil {
		return &user_center.UserExtra{}
	}
	return &user_center.UserExtra{
		Id:          int32(extra.ID),
		UserId:      int32(extra.UserID),
		Nickname:    extra.Nickname,
		PhoneNumber: extra.PhoneNumber,
		Email:       extra.Email,
		AvatarUrl:   extra.AvatarUrl,
	}
}
