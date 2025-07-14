package event

import (
	"github.com/kxw07/REST-API/storage"
	"log/slog"
)

type Repository struct {
}

func (rep Repository) getAllEvents() []Event {
	getEvents := `
	SELECT id, user_id, name, description, location, date_time
	FROM events
	`

	rows, err := storage.DB.Query(getEvents)
	if err != nil {
		slog.Error("error when querying events", "error", err)
		return nil
	}

	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.UserID, &event.Name, &event.Description, &event.Location, &event.DateTime)
		if err != nil {
			slog.Error("error when scanning row", "error", err)
			continue
		}
		events = append(events, event)
	}

	return events
}

func (rep Repository) getEvent(eventId int64) Event {
	getEventById := `
	SELECT id, user_id, name, description, location, date_time
	FROM events
	WHERE id = ?
	`

	row := storage.DB.QueryRow(getEventById, eventId)
	if row.Err() != nil {
		slog.Error("error when querying event by id", "error", row.Err())
		return Event{}
	}
	var event Event
	err := row.Scan(&event.ID, &event.UserID, &event.Name, &event.Description, &event.Location, &event.DateTime)
	if err != nil {
		slog.Error("error when scanning row", "error", err)
		return Event{}
	}

	return event
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
		slog.Error("error when create event", "error", err)
		return Event{}, err
	}

	return result, nil
}
