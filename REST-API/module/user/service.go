package user

import (
	"errors"
	"fmt"
	"github.com/kxw07/REST-API/utils/hash"
	"github.com/kxw07/REST-API/utils/jwt"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (svc Service) signup(user User) (User, error) {
	hashValue, err := hash.Do(user.Password)
	if err != nil {
		return User{}, errors.New(fmt.Sprintf("failed to hash password: %v", err))
	}

	user.Password = hashValue

	return svc.repo.save(user)
}

func (svc Service) login(user User) (string, error) {
	storedUser, _ := svc.repo.get(user)
	verifiedResult := hash.Compare(storedUser.Password, user.Password)
	if !verifiedResult {
		return "", errors.New("verified failed")
	}

	token, err := jwt.GenerateToken(storedUser.Email, storedUser.ID)
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to generate token: %v", err))
	}

	return token, nil
}
