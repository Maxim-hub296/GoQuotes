package models

import "gorm.io/gorm"

type Quote struct {
	gorm.Model
	Author   string `gorm:"not null"`
	Text     string `gorm:"not null"`
	Favorite bool   `gorm:"not null"`

	UserID uint
	User   User
}

func (q *Quote) Create(author string, text string, userId uint) {
	q.Author = author
	q.Text = text
	q.Favorite = false
	q.UserID = userId
}
