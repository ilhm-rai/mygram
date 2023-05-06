package service

import "github.com/ilhm-rai/mygram/model"

type SocialMediaService interface {
	FindAll() (response []model.SocialMediaResponse, err error)
	FindById(id int) (response model.SocialMediaResponse, err error)
	Save(request model.SaveSocialMediaRequest) (err error)
	DeleteById(id int) (err error)
}
