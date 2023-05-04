package repository

import (
	"github.com/ilhm-rai/mygram/entity"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	DB *gorm.DB
}

func (repository *userRepositoryImpl) FindByEmailOrUsername(email string, username string) (user entity.User, err error) {
	err = repository.DB.Where("username = ?", username).First(&user).Error
	return
}

func (repository *userRepositoryImpl) Save(user entity.User) (err error) {
	err = repository.DB.Create(&user).Error
	return
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		DB: database,
	}
}
