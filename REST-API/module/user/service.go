package user

type Service struct {
	repo Repository
}

func (svc Service) signup(user User) (User, error) {
	return svc.repo.save(user)
}
