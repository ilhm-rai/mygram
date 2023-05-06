package repository

import "github.com/ilhm-rai/mygram/entity"

type PhotoRepository interface {
	FindAll() (photos []entity.Photo, err error)
	FindById(id uint) (photo entity.Photo, err error)
	Save(photo entity.Photo) (err error)
	DeleteById(id uint) (err error)
}
