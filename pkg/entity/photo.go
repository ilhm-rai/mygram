package entity

import "time"

type Photo struct {
	ID        uint `gorm:"primarykey"`
	UserId    uint
	Title     string `gorm:"not null"`
	Caption   string
	PhotoURL  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User
	Comments  []Comment
}
