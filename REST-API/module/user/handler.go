package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	svc Service
}

func (handler Handler) RegisterRoutes(server *gin.Engine) {
	server.POST("/user", handler.signup)
}

func (handler Handler) signup(context *gin.Context) {
	var user User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input."})
		return
	}

	user, err := handler.svc.signup(user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully.", "user": user})
}
