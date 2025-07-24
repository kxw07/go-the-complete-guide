package user

import "golang.org/x/crypto/bcrypt"

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

	user.Password = string(hashValue)

	return svc.repo.save(user)
}

func (svc Service) hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
}
