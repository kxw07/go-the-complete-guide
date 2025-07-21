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
	event.UserID = rand.Int63()

	return svc.repo.createEvent(event)
}

func (svc Service) updateEvent(event Event) (Event, error) {
	return svc.repo.updateEvent(event)
}

func (svc Service) deleteEvent(id int64) error {
	return svc.repo.deleteEvent(id)
}
