package event

import (
	"time"
)

type Event struct {
	ID          int
	UserID      int       `binding:"required"`
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
}
