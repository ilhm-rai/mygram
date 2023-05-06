package repository

import (
	"github.com/ilhm-rai/mygram/entity"
	"gorm.io/gorm"
)

type commentRepositoryImpl struct {
	DB *gorm.DB
}

func (repository *commentRepositoryImpl) DeleteById(id uint) (err error) {
	err = repository.DB.Where("id = ?", id).Delete(&entity.Comment{}).Error
	return
}

func (repository *commentRepositoryImpl) FindById(id uint) (comment entity.Comment, err error) {
	err = repository.DB.Where("id = ?", id).First(&comment).Error
	return
}

func (repository *commentRepositoryImpl) FindAll() (comments []entity.Comment, err error) {
	err = repository.DB.Find(&comments).Error
	return
}

func (controller *commentRepositoryImpl) Save(comment entity.Comment) (err error) {
	return controller.DB.Save(&comment).Error
}

func NewCommentRepository(database *gorm.DB) CommentRepository {
	return &commentRepositoryImpl{
		DB: database,
	}
}
