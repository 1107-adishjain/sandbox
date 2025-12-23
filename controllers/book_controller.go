package controllers

import(
	"github.com/gin-gonic/gin"
)

func CreateBook()  gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implementation for creating a book
	}
}

func DeleteBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implementation for deleting a book
	}	
}

func GetBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implementation for getting books
	}		
}

func GetBooksbyID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implementation for getting a book by ID
	}				
}	