package entity

import (
	"time"
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
