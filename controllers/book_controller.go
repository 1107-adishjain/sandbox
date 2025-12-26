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
func CreateBook(app *app.Application)  gin.HandlerFunc {
	return func(c *gin.Context) {
		var book Book
		if err:=c.ShouldBindJSON(&book);err!=nil{
			c.JSON(400,gin.H{"error":"invalid request body"})
			return
		}
		// this creates a book record in the database. it runs a query like "INSERT INTO books (author, title, publisher, year) VALUES (...)"
		err:= app.DB.Create(&book).Error
		if err!=nil{
			c.JSON(500,gin.H{"error":"failed to create book"})
			return
		}
		// return the created book as response(status-201-created)
		c.JSON(201,gin.H{"message":"book created successfully"})
	}
}

func DeleteBook(app *app.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		book_id := c.Param("id")
		if book_id==""{
			c.JSON(400,gin.H{"error":"book id is required"})
			return
		}
		// this deletes the book from the database using its ID. it runs a query like "DELETE FROM books WHERE id=book_id"
		err:= app.DB.Delete(&models.Books{}, book_id).Error
		if err!=nil{
			c.JSON(500,gin.H{"error":"failed to delete book"})
			return
		}
		c.JSON(200,gin.H{"message":"book deleted successfully"})
	}	
}

func GetBooks(app *app.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		var books []models.Books
		// this fetches all the books from the database. it runs a query like "SELECT * FROM books"
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
		book_id:= c.Param("id")
		if book_id==""{
			c.JSON(400,gin.H{"error":"book id is required"})
			return
		}
		var book models.Books
		// This gets the book by using its ID. It runs a query like "SELECT * FROM books WHERE id=book_id"
		err:= app.DB.First(&book,book_id)
		if err!=nil{
			c.JSON(500,gin.H{"error":"failed to fetch book or record not found"})
			return
		}
		c.JSON(200,book)
	}				
}	