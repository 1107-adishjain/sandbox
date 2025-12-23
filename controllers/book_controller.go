package controllers

import(
	"github.com/gin-gonic/gin"
	"github.com/1107-adishjain/sandbox/app"
)

func CreateBook(app *app.Application)  gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implementation for creating a book
	}
}

func DeleteBook(app *app.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implementation for deleting a book
	}	
}

func GetBooks(app *app.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implementation for getting books
	}		
}

func GetBooksbyID(app *app.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implementation for getting a book by ID
	}				
}	