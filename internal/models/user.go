package models

import (
	"crypto/sha256"
	"encoding/hex"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func (u *User) SetPassword(password string) {
	hash := sha256.Sum256([]byte(password))
	u.Password = hex.EncodeToString(hash[:])
}

func (u *User) CheckPassword(password string) bool {
	hash := sha256.Sum256([]byte(password))
	return u.Password == hex.EncodeToString(hash[:])
}
