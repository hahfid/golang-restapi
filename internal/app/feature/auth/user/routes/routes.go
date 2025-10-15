package routes

import (
	"github.com/gin-gonic/gin"
	"golang-restapi/internal/app/feature/auth/user/handler"
)

func RegisterRoutes(r *gin.Engine, handler *handler.AuthHandler) {
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
}
