package event

import (
	"time"
)

type Event struct {
	ID          int64     `json:"id,omitempty"`
	UserID      int64     `json:"userId,omitempty"`
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
}
