package route

import (
	mw "github.com/1107-adishjain/sandbox/middleware"
	"github.com/1107-adishjain/sandbox/controllers"
	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/1107-adishjain/sandbox/app"
	"net/http"
)

func Routes(app *app.Application) *gin.Engine {

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowHeaders:  []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        3600,
	}))
	router.Use(mw.SecurityHeaders())
	limiter := tollbooth.NewLimiter(100, nil)
	router.Use(tollbooth_gin.LimitHandler(limiter))
	router.Use(func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 20<<20) // 20 MB limit
		c.Next()
	})
	r1 := router.Group("/api/v1")
	{
		r1.GET("/healthcheck", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"server working": true})
		})
		r1.POST("/create_books",controllers.CreateBook(app))
		r1.DELETE("/delete_book/:id",controllers.DeleteBook(app))
		r1.GET("/books",controllers.GetBooks(app))
		r1.GET("/books/:id",controllers.GetBooksbyID(app))
	}
	return router

}
