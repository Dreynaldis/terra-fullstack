package route

import (
	"backend/internal/delivery/handler"
	"backend/internal/usecase"
	"backend/internal/utils"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, userUsecase *usecase.UserUsecase) {
	router.POST("/users/login", func(c *gin.Context) {
		handler.Login(c, userUsecase)
	})
	router.POST("/users/register", func(c *gin.Context) {
		handler.Register(c, userUsecase)
	})
	router.POST("/auth/google/callback", func(c *gin.Context) {
		handler.CallbackGoogleHandler(c, userUsecase)
	})
	router.GET("/users/:username", func(c *gin.Context) {
		handler.GetUserProfile(c, userUsecase)
	})
	router.GET("/auth/check", utils.AuthMiddleware(), handler.Check)
}
