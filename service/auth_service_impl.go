package service

import (
	"errors"

	"github.com/ilhm-rai/mygram/entity"
	"github.com/ilhm-rai/mygram/helper"
	"github.com/ilhm-rai/mygram/model"
	"github.com/ilhm-rai/mygram/repository"
	"gorm.io/gorm"
)

type authServiceImpl struct {
	userRepository repository.UserRepository
}

func (service *authServiceImpl) Login(request model.LoginUserRequest) (token string, err error) {
	user, err := service.userRepository.FindByEmailOrUsername(request.EmailOrUsername, request.EmailOrUsername)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("account_not_found")
		}
		return
	}

	valid := helper.ComparePass(user.Password, request.Password)

	if !valid {
		err = errors.New("invalid_password")
		return
	}

	token = helper.GenerateToken(user.ID, user.Email, user.Username)
	return
}

func (service *authServiceImpl) Register(request model.RegisterUserRequest) (err error) {
	user := entity.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
		Age:      request.Age,
	}
	err = service.userRepository.Save(user)
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		err = errors.New("email_or_username_exist")
	}
	return
}

func NewAuthService(userRepository *repository.UserRepository) AuthService {
	return &authServiceImpl{
		userRepository: *userRepository,
	}
}
