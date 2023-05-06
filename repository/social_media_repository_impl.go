package repository

import (
	"github.com/ilhm-rai/mygram/entity"
	"gorm.io/gorm"
)

type socialMediaRepositoryImpl struct {
	DB *gorm.DB
}

func (repository *socialMediaRepositoryImpl) DeleteById(id uint) (err error) {
	err = repository.DB.Where("id = ?", id).Delete(&entity.SocialMedia{}).Error
	return
}

func (repository *socialMediaRepositoryImpl) FindById(id uint) (socialMedia entity.SocialMedia, err error) {
	err = repository.DB.Where("id = ?", id).First(&socialMedia).Error
	return
}

func (repository *socialMediaRepositoryImpl) FindAll() (socialMedia []entity.SocialMedia, err error) {
	err = repository.DB.Find(&socialMedia).Error
	return
}

func (controller *socialMediaRepositoryImpl) Save(socialMedia entity.SocialMedia) (err error) {
	return controller.DB.Save(&socialMedia).Error
}

func NewSocialMediaRepository(database *gorm.DB) SocialMediaRepository {
	return &socialMediaRepositoryImpl{
		DB: database,
	}
}
