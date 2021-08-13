package main

import (
	"context"
	"fmt"
	"gouser_center/handler"
	user_center "gouser_center/proto/user-center"
)

type UserCenterServer struct {
}

func (*UserCenterServer) Register(ctx context.Context, req *user_center.RegisterRequest) (*user_center.RegisterResponse, error) {
	fmt.Println("Register called")
	return handler.Register(ctx, req)
}
func (*UserCenterServer) Login(ctx context.Context, req *user_center.LoginRequest) (*user_center.LoginResponse, error) {
	fmt.Println("Login called")
	return handler.Login(ctx, req)
}
func (*UserCenterServer) Delete(ctx context.Context, req *user_center.DeleteRequest) (*user_center.DeleteResponse, error) {
	fmt.Println("Delete called")
	return handler.Delete(ctx, req)
}
func (*UserCenterServer) CheckToken(ctx context.Context, req *user_center.CheckTokenRequest) (*user_center.CheckTokenResponse, error) {
	fmt.Println("CheckToken called")
	return handler.CheckToken(ctx, req)
}
func (*UserCenterServer) Refresh(ctx context.Context, req *user_center.RefreshRequest) (*user_center.RefreshResponse, error) {
	fmt.Println("Refresh called")
	return handler.Refresh(ctx, req)
}
func (*UserCenterServer) GetUserInfo(ctx context.Context, req *user_center.GetUserInfoRequest) (*user_center.GetUserInfoResponse, error) {
	fmt.Println("GetUserInfo called")
	return handler.GetUserInfo(ctx, req)
}
func (*UserCenterServer) AddUser(ctx context.Context, req *user_center.AddUserRequest) (*user_center.AddUserResponse, error) {
	fmt.Println("AddUser called")
	return handler.AddUser(ctx, req)
}
