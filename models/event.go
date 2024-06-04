package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required"` // For required field
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

var events = []Event{}

// Method
func (e Event) Save() {
	// later: add it to the database
	events = append(events, e)
}

// Normal function
func GetAllEvents() []Event {
	return events
}
