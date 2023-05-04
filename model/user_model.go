package model

type RegisterUserRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Age      int    `json:"age" validate:"required,gte=8"`
}

type LoginUserRequest struct {
	EmailOrUsername string `json:"emailOrUsername" validate:"required"`
	Password        string `json:"password" validate:"required"`
}
