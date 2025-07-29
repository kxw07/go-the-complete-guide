package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/kxw07/REST-API/utils"
	"log/slog"
	"net/http"
)

func Authenticate(context *gin.Context) {
	token := context.GetHeader("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No Authorization"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		slog.Info("getEvents: valid token failed", "error", err)
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No Authorization"})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
