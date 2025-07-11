package event

import (
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
