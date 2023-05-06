package repository

import (
	"github.com/ilhm-rai/mygram/entity"
	"gorm.io/gorm"
)

type photoRepositoryImpl struct {
	DB *gorm.DB
}

func (repository *photoRepositoryImpl) DeleteById(id uint) (err error) {
	err = repository.DB.Where("id = ?", id).Delete(&entity.Photo{}).Error
	return
}

func (repository *photoRepositoryImpl) FindById(id uint) (photo entity.Photo, err error) {
	err = repository.DB.Where("id = ?", id).First(&photo).Error
	return
}

func (repository *photoRepositoryImpl) FindAll() (photos []entity.Photo, err error) {
	err = repository.DB.Find(&photos).Error
	return
}

func (controller *photoRepositoryImpl) Save(photo entity.Photo) (err error) {
	return controller.DB.Save(&photo).Error
}

func NewPhotoRepository(database *gorm.DB) PhotoRepository {
	return &photoRepositoryImpl{
		DB: database,
	}
}
