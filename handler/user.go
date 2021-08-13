package handler

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"gouser_center/dal/db"
	"gouser_center/dal/db/model"
	"gouser_center/dto"
	"gouser_center/proto/base"
	user_center "gouser_center/proto/user-center"
	"gouser_center/service"
)

func Register(ctx context.Context, req *user_center.RegisterRequest) (resp *user_center.RegisterResponse, err error) {
	err = db.CreateUserWithExtra(ctx, *dto.RPCUser2ModelUser(req.User), *dto.RPCUserExtra2ModelUserExtra(req.UserExtra))
	if err != nil {
		return
	}
	resp = &user_center.RegisterResponse{}
	return resp, nil
}

func Login(ctx context.Context, req *user_center.LoginRequest) (resp *user_center.LoginResponse, err error) {
	user, err := db.GetUser(ctx, model.User{Username: req.Username})
	if err != nil {
		return
	}

	if user.Password != req.Password {
		return nil, errors.New("用户名或密码错误")
	}
	extra, err := db.GetUserExtra(ctx, user)
	if err != nil {
		return
	}

	accessToken, accessExp, err := service.CreateAccessToken(user.ID, extra)
	if err != nil {
		return nil, errors.Errorf("create accessToken error, err: %+v", err)
	}
	refreshToken, refreshExp, err := service.CreateRefreshToken(user.ID, extra)
	if err != nil {
		return nil, errors.Errorf("create ")
	}

	resp = &user_center.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		AccessExp:    accessExp,
		RefreshExp:   refreshExp,
		UserId:       uint32(user.ID),
	}
	return
}

func Delete(ctx context.Context, req *user_center.DeleteRequest) (resp *user_center.DeleteResponse, err error) {
	return
}

func CheckToken(ctx context.Context, req *user_center.CheckTokenRequest) (resp *user_center.CheckTokenResponse, err error) {
	logrus.Infoln(req.Token)
	logrus.Infoln(req.GetIsRefresh())

	claims, err := service.ParseToken(req.Token, req.GetIsRefresh())
	if err != nil {
		logrus.Errorln(err)
		return
	}
	logrus.Infof("%s access", claims.Extra.Nickname)
	resp = &user_center.CheckTokenResponse{
		BaseResp: &base.BaseResp{StatusCode: 0},
	}
	err = claims.Valid()
	if err != nil {
		logrus.Errorln(err)
		return
	} else {
		logrus.Infoln("CheckToken Successed")
	}
	return
}

func Refresh(ctx context.Context, req *user_center.RefreshRequest) (resp *user_center.RefreshResponse, err error) {
	claims, err := service.ParseToken(req.RefreshToken, true)
	if err != nil {
		return
	}
	logrus.Infof("%s refresh", claims.Extra.Nickname)
	accessToken, accessExp, err := service.CreateAccessToken(claims.UserID, claims.Extra)
	if err != nil {
		return nil, errors.Errorf("create accessToken error, err: %+v", err)
	}
	refreshToken, refreshExp, err := service.CreateRefreshToken(claims.UserID, claims.Extra)
	if err != nil {
		return nil, errors.Errorf("create refreshToken error, err: %+v", err)
	}
	resp = &user_center.RefreshResponse{AccessToken: accessToken, RefreshToken: refreshToken, AccessExp: accessExp, RefreshExp: refreshExp}
	return
}

func GetUserInfo(ctx context.Context, req *user_center.GetUserInfoRequest) (resp *user_center.GetUserInfoResponse, err error) {
	resp = &user_center.GetUserInfoResponse{}
	userInfo, err := db.GetUserInfo(ctx, uint32(req.UserId), req.Username)
	if err != nil {
		return
	}
	resp.UserInfo = dto.ModelUserExtra2RPCUserExtra(userInfo)
	if req.Username != "" {
		resp.Username = req.Username
	} else {
		user, err := db.GetUser(ctx, model.User{Model: gorm.Model{ID: uint(req.UserId)}})
		if err != nil {
			return nil, errors.WithStack(err)
		}
		resp.Username = user.Username
	}
	return
}

func AddUser(ctx context.Context, req *user_center.AddUserRequest) (resp *user_center.AddUserResponse, err error) {
	return
}
