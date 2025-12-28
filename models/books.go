package models

import (
	"time"
	"gorm.io/gorm"
)

// this book struct represents the books table in the database
type Books struct {
	ID        uint      `json:"id" gorm:"primaryKey, autoIncrement"`
	Author    string    `json:"author" gorm:"not null" required:"true"`
	Title     string    `json:"title" gorm:"not null" reqquired:"true"`
	Publisher string    `json:"publisher" gorm:"not null" required:"true"`
	Year      int       `json:"year" gorm:"not null" required:"true"`
	CreatedAt time.Time `json:"created_at"`
}

func MigrateBooks(db *gorm.DB) error {
	return db.AutoMigrate(&Books{})
}
