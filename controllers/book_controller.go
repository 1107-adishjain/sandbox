package controllers

import (
	"github.com/1107-adishjain/sandbox/app"
	"github.com/1107-adishjain/sandbox/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// this bookInput struct contains the json request body fields
type BookInput struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
	Year      int    `json:"year"`
}

func CreateBook(app *app.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input BookInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}
		// Check if a book with the same title exists
		var existing models.Books
		if err := app.DB.Where("title = ?", input.Title).First(&existing).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "book with the same title already exists"})
			return
		}

		book := models.Books{
			Author:    input.Author,
			Title:     input.Title,
			Publisher: input.Publisher,
			Year:      input.Year,
		}

		if err := app.DB.Create(&book).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create book"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "book created successfully"})
	}
}

func DeleteBook(app *app.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookID := c.Param("id")

		if bookID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "book id is required"})
			return
		}
		// Try to delete the book directly and check RowsAffected in gorm if the book id does not exist then the erorr is nil but RowsAffected will be 0
		result := app.DB.Delete(&models.Books{}, bookID)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete book"})
			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "book deleted successfully"})
	}
}

func GetBooks(app *app.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		var books []models.Books
		if err := app.DB.Find(&books).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch books"})
			return
		}
		c.JSON(http.StatusOK, books)
	}
}

func GetBooksbyID(app *app.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookID := c.Param("id")

		if bookID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "book id is required"})
			return
		}

		var book models.Books
		if err := app.DB.First(&book, bookID).Error; err != nil {
			if err.Error() == "record not found" {
				c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch book"})
			}
			return
		}

		c.JSON(http.StatusOK, book)
	}
}
