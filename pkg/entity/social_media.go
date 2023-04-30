package entity

import "time"

type SocialMedia struct {
	ID             uint `gorm:"primarykey"`
	UserId         uint
	Name           string `gorm:"not null"`
	SocialMediaURL string `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	User           User
}
