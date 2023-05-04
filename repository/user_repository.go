package repository

import "github.com/ilhm-rai/mygram/entity"

type UserRepository interface {
	Save(user entity.User) (err error)
	FindByEmailOrUsername(email string, username string) (user entity.User, err error)
}
