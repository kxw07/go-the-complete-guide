package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log/slog"
	"time"
)

var secretKey = "ThisIsSecretKey"

func GenerateToken(email string, userId int64) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"expire": time.Now().Add(time.Hour * 1).Unix(),
	}).SignedString([]byte(secretKey))
}

func VerifyToken(tokenString string) (int64, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, fmt.Errorf("token parsing error: %w", err)
	}

	if !parsedToken.Valid {
		return 0, fmt.Errorf("token invalid: %v", tokenString)
	}

	claims := parsedToken.Claims.(jwt.MapClaims)
	slog.Info(fmt.Sprintf("claims: %v", claims))

	return int64(claims["userId"].(float64)), nil
}
