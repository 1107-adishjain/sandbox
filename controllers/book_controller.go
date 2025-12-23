package controllers

import(
	"github.com/gin-gonic/gin"
	"github.com/1107-adishjain/sandbox/app"
	"github.com/1107-adishjain/sandbox/models"
)
// this book struct contains the json request body fields
type Book struct{
	Author string `json:"author"`
	Title string `json:"title"`
	Publisher string `json:"publisher"`
	Year int `json:"year"`
}
func CreateBook(app *app.Application)  gin.HandlerFunc {
	return func(c *gin.Context) {
		var book Book
		if err:=c.ShouldBindJSON(&book);err!=nil{
			c.JSON(400,gin.H{"error":"invalid request body"})
			return
		}
		err:= app.DB.Create(&book).Error
		if err!=nil{
			c.JSON(500,gin.H{"error":"failed to create book"})
			return
		}
		c.JSON(201,gin.H{"message":"book created successfully"})
	}
}

func DeleteBook(app *app.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implementation for deleting a book
	}	
}

func GetBooks(app *app.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		var books []models.Book
		err:= app.DB.Find(&books).Error
		if err!=nil{
			c.JSON(500,gin.H{"error":"failed to fetch books"})
			return
		}
		c.JSON(200,books)
	}		
}

func GetBooksbyID(app *app.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implementation for getting a book by ID
	}				
}	