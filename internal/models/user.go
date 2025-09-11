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

func (u *User) SetData(username string, password string) {
	u.Username = username
	hash := sha256.Sum256([]byte(password))
	u.Password = hex.EncodeToString(hash[:])
}

func (u *User) CheckPassword(password string) bool {
	hash := sha256.Sum256([]byte(password))
	return u.Password == hex.EncodeToString(hash[:])
}

func (u *User) Delete(db *gorm.DB) error {
	return db.Unscoped().Delete(&u).Error
}
