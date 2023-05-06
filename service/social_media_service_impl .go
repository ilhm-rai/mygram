package service

import (
	"errors"

	"github.com/ilhm-rai/mygram/entity"
	"github.com/ilhm-rai/mygram/model"
	"github.com/ilhm-rai/mygram/repository"
	"gorm.io/gorm"
)

type socialMediaServiceImpl struct {
	SocialMediaRepository repository.SocialMediaRepository
}

func (service *socialMediaServiceImpl) DeleteById(id int) (err error) {
	err = service.SocialMediaRepository.DeleteById(uint(id))
	return
}

func (service *socialMediaServiceImpl) FindById(id int) (response model.SocialMediaResponse, err error) {
	socialMedia, err := service.SocialMediaRepository.FindById(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("social_media_not_found")
		}
		return
	}
	response = service.EntityToResponse(socialMedia)
	return
}

func (service *socialMediaServiceImpl) FindAll() (response []model.SocialMediaResponse, err error) {
	socialMedia, err := service.SocialMediaRepository.FindAll()
	for _, media := range socialMedia {
		response = append(response, service.EntityToResponse(media))
	}
	return
}

func (service *socialMediaServiceImpl) Save(request model.SaveSocialMediaRequest) (err error) {
	err = service.SocialMediaRepository.Save(service.RequestToEntity(request))
	return
}

func (service *socialMediaServiceImpl) RequestToEntity(data model.SaveSocialMediaRequest) entity.SocialMedia {
	return entity.SocialMedia{
		ID:             data.ID,
		UserId:         data.UserId,
		Name:           data.Name,
		SocialMediaURL: data.SocialMediaURL,
	}
}

func (service *socialMediaServiceImpl) EntityToResponse(data entity.SocialMedia) model.SocialMediaResponse {
	return model.SocialMediaResponse{
		ID:             data.ID,
		UserId:         data.UserId,
		Name:           data.Name,
		SocialMediaURL: data.SocialMediaURL,
	}
}

func NewSocialMediaService(socialMediaRepository *repository.SocialMediaRepository) SocialMediaService {
	return &socialMediaServiceImpl{
		SocialMediaRepository: *socialMediaRepository,
	}
}
