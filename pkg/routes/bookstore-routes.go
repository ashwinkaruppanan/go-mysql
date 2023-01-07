package routes

import (
	"github.com/ashwin/go-mysql/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var RegisterBookStoreRoutes = func(router *gin.Engine) {
	router.POST("/book/", controllers.CreateBook)
	router.GET("/book/", controllers.GetBook)
	router.GET("/book/:id", controllers.GetBookById)
	router.PUT("/book/:id", controllers.UpdateBook)
	router.DELETE("/book/:id", controllers.DeleteBook)

}
