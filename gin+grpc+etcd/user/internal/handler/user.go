package handler

import (
	"context"
	"user/internal/repository"
	"user/internal/service/service"
	"user/pkg/e"
)

type UserService struct {
}

// 实例化UserService
func NewUserService() *UserService {
	return &UserService{}
}

// UserLogin，用户登陆
func (*UserService) UserLogin(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailRespose, err error) {
	var user repository.User
	resp = new(service.UserDetailRespose)
	resp.Code = e.Success
	err = user.ShowUserInfo(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}
	resp.UserDetail = repository.BuildUser(user)
	return resp, nil
}

// UserRegister 用户注册
func (*UserService) Usergister(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailRespose, err error) {
	var user repository.User
	resp = new(service.UserDetailRespose)
	resp.Code = e.Success
	user, err = user.UserCreate(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}
	resp.UserDetail = repository.BuildUser(user)
	return resp, nil
}
