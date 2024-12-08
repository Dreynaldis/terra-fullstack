package route

import (
	"backend/internal/delivery/handler"
	"backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, userUsecase *usecase.UserUsecase) {
	router.POST("/users/login", func(c *gin.Context) {
		handler.Login(c, userUsecase)
	})
	router.POST("/users/register", func(c *gin.Context) {
		handler.Register(c, userUsecase)
	})
}
