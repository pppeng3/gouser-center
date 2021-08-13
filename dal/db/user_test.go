package db

import (
	"context"
	"fmt"
	"gouser_center/dal/db/model"
	"testing"

	"github.com/sirupsen/logrus"
)

var userCtx = context.Background()

func TestGetUser(t *testing.T) {
	user, err := GetUser(userCtx, model.User{Username: "pp"})
	if err != nil {
		logrus.Error(err)
		t.FailNow()
	}
	logrus.Infof("%+v", user)
}

func TestRegister(t *testing.T) {
	user := model.User{
		Username: "ppp",
		Password: "123456",
	}
	extra := model.UserExtra{
		Nickname:    "ppp",
		PhoneNumber: "123jdsakAS_",
	}
	err := CreateUserWithExtra(userCtx, user, extra)
	if err != nil {
		logrus.Error(err)
		t.FailNow()
	}
}

func TestGetExtra(t *testing.T) {
	user := model.User{}
	user.ID = 1
	extra, err := GetUserExtra(userCtx, user)
	if err != nil {
		logrus.Error(err)
		t.FailNow()
	}
	fmt.Printf("%+v", extra)
}
