package db

import (
	"context"
	"errors"
	"gouser_center/dal/db/model"

	ex_errors "github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func GetUser(ctx context.Context, user model.User) (res model.User, err error) {
	logrus.Info(user.Username)
	db := db.WithContext(ctx)
	err = db.Model(&user).Where(&user).Take(&res).Error
	return
}

func CreateUserWithExtra(ctx context.Context, user model.User, extra model.UserExtra) (err error) {
	err = db.WithContext(ctx).Transaction(func(tx *gorm.DB) (err error) {
		result := tx.Model(&user).FirstOrCreate(&user, user)
		if err = result.Error; err != nil {
			err = ex_errors.Errorf("Create user error: %v", err)
			return
		}
		if result.RowsAffected == 0 { //插入条数为0,用户名冲突
			err = ex_errors.New("用户名已存在")
			return
		}
		if user.ID == 0 {
			return errors.New("user_id is 0")
		}
		logrus.Infof("Create User: %+v", user)
		extra.UserID = uint32(user.ID)
		if err = tx.Model(&extra).Create(&extra).Error; err != nil {
			err = ex_errors.Errorf("Create userExtra error: %v", err)
			return
		}
		logrus.Infof("Create UserExtra: %+v", extra)
		if extra.ID <= 0 {
			return errors.New("user_extra_id is 0")
		}

		return
	})
	if err != nil {
		logrus.Warnf("Register error: %+v", err)
	}
	return err
}

func GetUserExtra(ctx context.Context, user model.User) (extra model.UserExtra, err error) {
	if user.ID <= 0 {
		return extra, errors.New("user_id is empty")
	}
	err = db.WithContext(ctx).Model(&extra).Where("user_id = ?", user.ID).First(&extra).Error
	return
}

func GetUserInfo(ctx context.Context, userID uint32, username string) (userInfo *model.UserExtra, err error) {
	userInfo = &model.UserExtra{}
	db := db.WithContext(ctx)
	if userID > 0 {
		err = db.Model(&model.UserExtra{}).Where("user_id = ?", userID).First(&userInfo).Error
	} else {
		var user model.User
		err = db.Model(&model.User{}).Select("id").Where("username = ?", username).First(&user).Error
		if err != nil {
			return
		}

		err = db.Model(&model.UserExtra{}).Where("user_id = ?", user.ID).First(&userInfo).Error
	}

	if err != nil {
		err = ex_errors.WithStack(err)
	}
	return
}
