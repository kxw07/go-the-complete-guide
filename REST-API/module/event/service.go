package event

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (svc Service) getEvents(userId int64) ([]Event, error) {
	return svc.repo.getAllEvents(userId)
}

func (svc Service) getEvent(eventId int64, userId int64) (Event, error) {
	return svc.repo.getEvent(eventId, userId)
}

func (svc Service) createEvent(event Event, userId int64) (Event, error) {
	return svc.repo.createEvent(event, userId)
}

func (svc Service) updateEvent(event Event, userId int64) (Event, error) {
	return svc.repo.updateEvent(event, userId)
}

func (svc Service) deleteEvent(eventId int64, userId int64) error {
	return svc.repo.deleteEvent(eventId, userId)
}

func (svc Service) registerEvent(eventId int64, userId int64) error {
	return svc.repo.registerEvent(eventId, userId)
}

func (svc Service) unregisterEvent(eventId int64, userId int64) error {
	// return svc.repo.unregisterEvent(eventId, userId)
	return svc.repo.runUnregisterEventWithInterfaceExcer(eventId, userId)
}
