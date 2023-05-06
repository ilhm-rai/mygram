package service

import (
	"errors"

	"github.com/ilhm-rai/mygram/entity"
	"github.com/ilhm-rai/mygram/model"
	"github.com/ilhm-rai/mygram/repository"
	"gorm.io/gorm"
)

type photoServiceImpl struct {
	PhotoRepository repository.PhotoRepository
}

func (service *photoServiceImpl) DeleteById(id int) (err error) {
	err = service.PhotoRepository.DeleteById(uint(id))
	return
}

func (service *photoServiceImpl) FindById(id int) (response model.PhotoResponse, err error) {
	photo, err := service.PhotoRepository.FindById(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("photo_not_found")
		}
		return
	}
	response = service.EntityToResponse(photo)
	return
}

func (service *photoServiceImpl) FindAll() (response []model.PhotoResponse, err error) {
	photos, err := service.PhotoRepository.FindAll()
	for _, photo := range photos {
		response = append(response, service.EntityToResponse(photo))
	}
	return
}

func (service *photoServiceImpl) Save(request model.SavePhotoRequest) (err error) {
	err = service.PhotoRepository.Save(service.RequestToEntity(request))
	return
}

func (service *photoServiceImpl) RequestToEntity(data model.SavePhotoRequest) entity.Photo {
	return entity.Photo{
		ID:       data.ID,
		UserId:   data.UserId,
		Title:    data.Title,
		Caption:  data.Caption,
		PhotoURL: data.PhotoURL,
	}
}

func (service *photoServiceImpl) EntityToResponse(data entity.Photo) model.PhotoResponse {
	return model.PhotoResponse{
		ID:       data.ID,
		UserId:   data.UserId,
		Title:    data.Title,
		Caption:  data.Caption,
		PhotoURL: data.PhotoURL,
	}
}

func NewPhotoService(photoRepository *repository.PhotoRepository) PhotoService {
	return &photoServiceImpl{
		PhotoRepository: *photoRepository,
	}
}
