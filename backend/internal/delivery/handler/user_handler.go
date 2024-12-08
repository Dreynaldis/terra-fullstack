package handler

import (
	"backend/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context, usecase *usecase.UserUsecase) {

	var loginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := usecase.Login(c, loginReq.Email, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Register(c *gin.Context, usecase *usecase.UserUsecase) {
	var createUserReq struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&createUserReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := usecase.Register(c, createUserReq.Username, createUserReq.Email, createUserReq.Password)
	if err != nil {
		if err.Error() == "email already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
			return
		}
		if err.Error() == "username already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": "username already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created successfully"})
}
