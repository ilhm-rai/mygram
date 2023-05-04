package entity

import "time"

type Comment struct {
	ID        uint `gorm:"primarykey"`
	UserId    uint
	PhotoId   uint
	Message   string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User
	Photo     Photo
}
