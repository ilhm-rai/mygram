package repository

import "github.com/ilhm-rai/mygram/entity"

type SocialMediaRepository interface {
	FindAll() (socialMedia []entity.SocialMedia, err error)
	FindById(id uint) (socialMedia entity.SocialMedia, err error)
	Save(socialMedia entity.SocialMedia) (err error)
	DeleteById(id uint) (err error)
}
