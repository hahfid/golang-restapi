package routes

import (
	controllers "golang-restapi/controller"
	"golang-restapi/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	secured := r.Group("/api")
	secured.Use(middleware.AuthMiddleware())
	{
		secured.GET("/books", controllers.GetBooks)
		secured.POST("/books", controllers.CreateBook)
		secured.GET("/books/:id", controllers.GetBookByID)
		secured.PUT("/books/:id", controllers.UpdateBook)
		secured.DELETE("/books/:id", controllers.DeleteBook)

	}
}
