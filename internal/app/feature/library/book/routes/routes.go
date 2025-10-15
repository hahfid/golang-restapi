package routes

import (
	"github.com/gin-gonic/gin"
	"golang-restapi/internal/app/feature/library/book/handler"
)

func RegisterRoutes(r *gin.RouterGroup, handler *handler.BookHandler) {
	r.GET("/books", handler.GetBooks)
	r.POST("/books", handler.CreateBook)
	r.GET("/books/:id", handler.GetBook)
	r.PUT("/books/:id", handler.UpdateBook)
	r.DELETE("/books/:id", handler.DeleteBook)
}
