package db

import (
	"library/utils"
	"log"

	// "gorm.io/driver/sqlite"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func InitDB(dbPath string) *Repository {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// GORM automatically creates/updates the table
	db.AutoMigrate(&utils.Book{}, &utils.User{})

	return &Repository{DB: db}
}