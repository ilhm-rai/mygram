package service

import "github.com/ilhm-rai/mygram/model"

type PhotoService interface {
	FindAll() (response []model.PhotoResponse, err error)
	FindById(id int) (response model.PhotoResponse, err error)
	Save(request model.SavePhotoRequest) (err error)
	DeleteById(id int) (err error)
}
