package entity

import (
	"time"

	"github.com/ilhm-rai/mygram/helper"
	"gorm.io/gorm"
)

type User struct {
	ID          uint   `gorm:"primarykey"`
	Username    string `gorm:"uniqueIndex;not null"`
	Email       string `gorm:"uniqueIndex;not null"`
	Password    string `gorm:"not null"`
	Age         int    `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Photos      []Photo
	SocialMedia []SocialMedia
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = helper.GeneratePass(u.Password)
	return
}
