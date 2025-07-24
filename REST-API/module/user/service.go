package user

import (
	"errors"
	"fmt"
	"github.com/kxw07/REST-API/utils"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (svc Service) signup(user User) (User, error) {
	hashValue, err := utils.HashPassword(user.Password)
	if err != nil {
		return User{}, errors.New(fmt.Sprintf("failed to hash password: %v", err))
	}

	user.Password = hashValue

	return svc.repo.save(user)
}

func (svc Service) login(user User) (string, error) {
	storedUser, _ := svc.repo.get(user)
	verifiedResult := utils.VerifyPassword(storedUser.Password, user.Password)
	if !verifiedResult {
		return "", errors.New("verified failed")
	}

	token, err := utils.GenerateToken(storedUser.Email, storedUser.ID)
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to generate token: %v", err))
	}

	return token, nil
}
