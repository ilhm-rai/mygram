package service

import "github.com/ilhm-rai/mygram/model"

type CommentService interface {
	FindAll() (response []model.CommentResponse, err error)
	FindById(id int) (response model.CommentResponse, err error)
	Save(request model.SaveCommentRequest) (err error)
	DeleteById(id int) (err error)
}
