package user

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (svc Service) signup(user User) (User, error) {
	return svc.repo.save(user)
}
