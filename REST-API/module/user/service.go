package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (svc Service) signup(user User) (User, error) {
	hashValue, err := svc.hashPassword(user.Password)
	if err != nil {
		return User{}, err
	}

	user.Password = hashValue

	return svc.repo.save(user)
}

func (svc Service) login(user User) error {
	storedValue := svc.repo.get(user.Email)
	verifiedResult := svc.verifyPassword(user.Password, storedValue)

	if !verifiedResult {
		return errors.New("verified failed")
	}

	return nil
}

func (svc Service) verifyPassword(inputValue, storedValue string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedValue), []byte(inputValue))
	return err == nil
}

func (svc Service) hashPassword(password string) (string, error) {
	hashValue, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hashValue), err
}
