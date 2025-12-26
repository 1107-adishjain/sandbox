package controllers

import (
	"github.com/1107-adishjain/sandbox/app"
	"github.com/1107-adishjain/sandbox/models"
	"github.com/gin-gonic/gin"
)

// this book struct contains the json request body fields
type Book struct{
	Author string `json:"author"`
	Title string `json:"title"`
	Publisher string `json:"publisher"`
	Year int `json:"year"`
}
func CreateBook(app *app.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Books
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": "invalid request body"})
			return
		}
		// Check if a book with the same title exists
		var existing models.Books
		if err := app.DB.Where("title = ?", input.Title).First(&existing).Error; err == nil {
			c.JSON(400, gin.H{"error": "book with the same title already exists"})
			return
		}
		if err := app.DB.Create(&input).Error; err != nil {
			c.JSON(500, gin.H{"error": "failed to create book"})
			return
		}
		c.JSON(201, gin.H{"message": "book created successfully"})
	}
}

func DeleteBook(app *app.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookID := c.Param("id")
		if bookID == "" {
			c.JSON(400, gin.H{"error": "book id is required"})
			return
		}
		// Try to delete the book directly and check RowsAffected
		result := app.DB.Delete(&models.Books{}, bookID)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": "failed to delete book"})
			return
		}
		if result.RowsAffected == 0 {
			c.JSON(404, gin.H{"error": "book not found"})
			return
		}
		c.JSON(200, gin.H{"message": "book deleted successfully"})
	}
}

func GetBooks(app *app.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		var books []models.Books
		if err := app.DB.Find(&books).Error; err != nil {
			c.JSON(500, gin.H{"error": "failed to fetch books"})
			return
		}
		c.JSON(200, books)
	}
}

func GetBooksbyID(app *app.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookID := c.Param("id")
		if bookID == "" {
			c.JSON(400, gin.H{"error": "book id is required"})
			return
		}
		var book models.Books
		if err := app.DB.First(&book, bookID).Error; err != nil {
			if err.Error() == "record not found" {
				c.JSON(404, gin.H{"error": "book not found"})
			} else {
				c.JSON(500, gin.H{"error": "failed to fetch book"})
			}
			return
		}
		c.JSON(200, book)
	}
}