package model

import (
	"github.com/kxw07/REST-API/storage"
	"log/slog"
	"time"
)

type Event struct {
	ID          int
	UserID      int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
}

func GetAll() []Event {
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

func AddEvent(event Event) {
	insertEvents := `
	INSERT INTO events (id, user_id, name, description, location, date_time)
	VALUES (?, ?, ?, ?, ?, ?)
	`

	stmt, err := storage.DB.Prepare(insertEvents)
	if err != nil {
		slog.Error("error when preparing statement", "error", err)
	}

	_, err = stmt.Exec(event.ID, event.UserID, event.Name, event.Description, event.Location, event.DateTime)
	if err != nil {
		slog.Error("error when executing statement", "error", err)
	}
}
