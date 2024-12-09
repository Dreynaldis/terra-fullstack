package handler

import (
	"backend/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CallbackGoogleHandler(c *gin.Context, usecase *usecase.UserUsecase) {
	var body struct {
		Token string `json:"token"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, username, err := usecase.HandleGoogleCallback(c.Request.Context(), body.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"token": token,
		"user":  username,
	})
}

func Check(c *gin.Context) {
	user, _ := c.Get("user")
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "authorized"})
}
