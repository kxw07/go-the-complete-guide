package event

import "math/rand"

type Service struct {
	repo Repository
}

func (svc Service) getEvents() []Event {
	return svc.repo.getAllEvents()
}

func (svc Service) createEvent(event Event) (Event, error) {
	event.UserID = rand.Int()

	return svc.repo.createEvent(event)
}
