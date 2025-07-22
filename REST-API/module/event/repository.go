package event

import (
	"github.com/kxw07/REST-API/storage"
	"log/slog"
)

type Repository struct {
	sto *storage.Storage
}

func NewRepository(sto *storage.Storage) *Repository {
	return &Repository{sto: sto}
}

func (rep Repository) getAllEvents() ([]Event, error) {
	getEventsSql := `
	SELECT id, user_id, name, description, location, date_time
	FROM events
	`

	rows, err := rep.sto.GetDB().Query(getEventsSql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.UserID, &event.Name, &event.Description, &event.Location, &event.DateTime)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func (rep Repository) getEvent(eventId int64) (Event, error) {
	getEventByIdSql := `
	SELECT id, user_id, name, description, location, date_time
	FROM events
	WHERE id = ?
	`

	row := rep.sto.GetDB().QueryRow(getEventByIdSql, eventId)
	if row.Err() != nil {
		return Event{}, row.Err()
	}
	var event Event
	err := row.Scan(&event.ID, &event.UserID, &event.Name, &event.Description, &event.Location, &event.DateTime)
	if err != nil {
		return Event{}, err
	}

	return event, nil
}

func (rep Repository) createEvent(event Event) (Event, error) {
	insertEventSql := `
	INSERT INTO events (user_id, name, description, location, date_time)
	VALUES (?, ?, ?, ?, ?)
	RETURNING id, user_id, name, description, location, date_time
	`

	var result Event
	err := rep.sto.GetDB().
		QueryRow(insertEventSql, event.UserID, event.Name, event.Description, event.Location, event.DateTime).
		Scan(&result.ID, &result.UserID, &result.Name, &result.Description, &result.Location, &result.DateTime)

	if err != nil {
		slog.Error("error when creating event", "error", err)
		return Event{}, err
	}

	return result, nil
}

func (rep Repository) updateEvent(event Event) (Event, error) {
	updateEventSql := `
	UPDATE events
	SET name = ?, description = ?, location = ?, date_time = ?
	WHERE id = ?
	RETURNING name, description, location, date_time
	`

	var result Event
	err := rep.sto.GetDB().
		QueryRow(updateEventSql, event.Name, event.Description, event.Location, event.DateTime, event.ID).
		Scan(&result.Name, &result.Description, &result.Location, &result.DateTime)

	if err != nil {
		slog.Error("error when updating event", "error", err)
		return Event{}, err
	}

	return result, nil
}

func (rep Repository) deleteEvent(id int64) error {
	deleteEventSql := `
	DELETE FROM events WHERE id = ?
	`
	_, err := rep.sto.GetDB().Exec(deleteEventSql, id)

	if err != nil {
		slog.Error("error when deleting event", "error", err)
		return err
	}

	return nil
}
