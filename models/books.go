package models

import (
	"gorm.io/gorm"
)

// this book struct represents the books table in the database
type Book struct {
	ID        uint   `json:"id" gorm:"primaryKey, autoIncrement"`
	Author    string `json:"author"`
	Title     string `json:"title" gorm:"not null"`
	Publisher string `json:"publisher"`
	Year      int    `json:"year"`
}

func MigrateBooks(db *gorm.DB) error {
	return db.AutoMigrate(&Book{})
}
