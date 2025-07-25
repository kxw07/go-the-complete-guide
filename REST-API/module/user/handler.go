package user

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (handler Handler) RegisterRoutes(server *gin.Engine) {
	server.POST("/user/signup", handler.signUp)
	server.POST("/user/login", handler.login)
}

func (handler Handler) signUp(context *gin.Context) {
	var user User
	if err := context.ShouldBindJSON(&user); err != nil {
		slog.Info("singUp: bind json failed", "error", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input."})
		return
	}

	user, err := handler.svc.signup(user)
	if err != nil {
		slog.Info("signUp: do signUp failed", "error", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully.", "user": user})
}

func (handler Handler) login(context *gin.Context) {
	var user User
	if err := context.ShouldBindJSON(&user); err != nil {
		slog.Info("login: bind json failed", "error", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input."})
		return
	}

	token, err := handler.svc.login(user)
	if err != nil {
		slog.Info("login: do login failed", "error", err)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful.", "token": token})
}
