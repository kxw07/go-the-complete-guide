package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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

func VerifyPassword(storedValue, inputValue string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedValue), []byte(inputValue))
	return err == nil
}

func HashPassword(password string) (string, error) {
	hashValue, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hashValue), err
}

func VerifyToken(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	return err
}
