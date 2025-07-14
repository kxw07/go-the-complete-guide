package event

import "math/rand"

type Service struct {
	repo Repository
}

func (svc Service) getEvents() ([]Event, error) {
	return svc.repo.getAllEvents()
}

func (svc Service) getEvent(id int64) (Event, error) {
	return svc.repo.getEvent(id)
}

func (svc Service) createEvent(event Event) (Event, error) {
	event.UserID = rand.Int()

	return svc.repo.createEvent(event)
}
