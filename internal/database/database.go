package database

import (
	"GoQuotes/internal/models"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	var err error
	db, err = gorm.Open(sqlite.Open("goquotes.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("не удалось подключиться к базе:", err)
	}

	// миграции
	err = db.AutoMigrate(&models.User{}, &models.Quote{})
	if err != nil {
		log.Fatal("не удалось мигрировать:", err)
	}

	fmt.Println("База готова")
	return db

}
