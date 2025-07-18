package event

import (
	"github.com/kxw07/REST-API/storage"
	"log/slog"
)

type Repository struct {
}

func (rep Repository) getAllEvents() ([]Event, error) {
	getEvents := `
	SELECT id, user_id, name, description, location, date_time
	FROM events
	`

	rows, err := storage.DB.Query(getEvents)
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
	getEventById := `
	SELECT id, user_id, name, description, location, date_time
	FROM events
	WHERE id = ?
	`

	row := storage.DB.QueryRow(getEventById, eventId)
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
	insertEvent := `
	INSERT INTO events (user_id, name, description, location, date_time)
	VALUES (?, ?, ?, ?, ?)
	RETURNING id, user_id, name, description, location, date_time
	`

	var result Event
	err := storage.DB.
		QueryRow(insertEvent, event.UserID, event.Name, event.Description, event.Location, event.DateTime).
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
	SET user_id = ?, name = ?, description = ?, location = ?, date_time = ?
	WHERE id = ?
	RETURNING id, user_id, name, description, location, date_time
	`

	var result Event
	err := storage.DB.
		QueryRow(updateEventSql, event.UserID, event.Name, event.Description, event.Location, event.DateTime, event.ID).
		Scan(&result.ID, &result.UserID, &result.Name, &result.Description, &result.Location, &result.DateTime)

	if err != nil {
		slog.Error("error when updating event", "error", err)
		return Event{}, err
	}

	return result, nil
}
