package service

import (
	"errors"

	"github.com/ilhm-rai/mygram/entity"
	"github.com/ilhm-rai/mygram/model"
	"github.com/ilhm-rai/mygram/repository"
	"gorm.io/gorm"
)

type commentServiceImpl struct {
	CommentRepository repository.CommentRepository
}

func (service *commentServiceImpl) DeleteById(id int) (err error) {
	err = service.CommentRepository.DeleteById(uint(id))
	return
}

func (service *commentServiceImpl) FindById(id int) (response model.CommentResponse, err error) {
	comment, err := service.CommentRepository.FindById(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("comment_not_found")
		}
		return
	}
	response = service.EntityToResponse(comment)
	return
}

func (service *commentServiceImpl) FindAll() (response []model.CommentResponse, err error) {
	comments, err := service.CommentRepository.FindAll()
	for _, comment := range comments {
		response = append(response, service.EntityToResponse(comment))
	}
	return
}

func (service *commentServiceImpl) Save(request model.SaveCommentRequest) (err error) {
	err = service.CommentRepository.Save(service.RequestToEntity(request))
	return
}

func (service *commentServiceImpl) RequestToEntity(data model.SaveCommentRequest) entity.Comment {
	return entity.Comment{
		ID:      data.ID,
		UserId:  data.UserId,
		PhotoId: data.PhotoId,
		Message: data.Message,
	}
}

func (service *commentServiceImpl) EntityToResponse(data entity.Comment) model.CommentResponse {
	return model.CommentResponse{
		ID:      data.ID,
		UserId:  data.UserId,
		PhotoId: data.PhotoId,
		Message: data.Message,
	}
}

func NewCommentService(commentRepository *repository.CommentRepository) CommentService {
	return &commentServiceImpl{
		CommentRepository: *commentRepository,
	}
}
