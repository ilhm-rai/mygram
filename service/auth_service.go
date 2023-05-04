package service

import "github.com/ilhm-rai/mygram/model"

type AuthService interface {
	Register(request model.RegisterUserRequest) (err error)
	Login(request model.LoginUserRequest) (token string, err error)
}
