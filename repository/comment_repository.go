package repository

import "github.com/ilhm-rai/mygram/entity"

type CommentRepository interface {
	FindAll() (comments []entity.Comment, err error)
	FindById(id uint) (comment entity.Comment, err error)
	Save(comment entity.Comment) (err error)
	DeleteById(id uint) (err error)
}
