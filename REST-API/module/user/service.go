package user

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var secretKey = "ThisIsSecretKey"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (svc Service) signup(user User) (User, error) {
	hashValue, err := svc.hashPassword(user.Password)
	if err != nil {
		return User{}, errors.New(fmt.Sprintf("failed to hash password: %v", err))
	}

	user.Password = hashValue

	return svc.repo.save(user)
}

func (svc Service) login(user User) (string, error) {
	storedUser, _ := svc.repo.get(user.Email)
	verifiedResult := svc.verifyPassword(storedUser.Password, user.Password)
	if !verifiedResult {
		return "", errors.New("verified failed")
	}

	token, err := svc.generateToken(storedUser)
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to generate token: %v", err))
	}

	return token, nil
}

func (svc Service) generateToken(user User) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  user.Email,
		"userId": user.ID,
		"expire": time.Now().Add(time.Hour * 1).Unix(),
	}).SignedString([]byte(secretKey))
}

func (svc Service) verifyPassword(storedValue, inputValue string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedValue), []byte(inputValue))
	return err == nil
}

func (svc Service) hashPassword(password string) (string, error) {
	hashValue, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hashValue), err
}
