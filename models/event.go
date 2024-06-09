package models

import (
	"time"

	"dunky.com/eventbooking/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"` // For required field
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

//var events = []Event{}

// Method
func (e Event) Save() error {
	// later: add it to the database
	query := `INSERT INTO events (name, description, location, date_time, user_id) 
	VALUES(?, ?, ?, ?, ?);`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

// Normal function
func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	var dateTimeStr string

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &dateTimeStr, &event.UserID)
		if err != nil {
			return nil, err
		}
		event.DateTime, err = time.Parse("2006-01-02 15:04:05", dateTimeStr)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	//fmt.Println("Executing query:", query, "with id:", id)
	row := db.DB.QueryRow(query, id)
	var event Event
	var dateTimeStr string
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &dateTimeStr, &event.UserID)
	if err != nil {
		return nil, err
	}
	event.DateTime, err = time.Parse("2006-01-02 15:04:05", dateTimeStr)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e Event) Update() error {
	query := `
		UPDATE events
		SET name = ?, description = ?, location = ?, date_time = ? 
		WHERE id = ?;
		`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	return err

}

func (e Event) Delete() error {
	query := `DELETE FROM event WHERE ID = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID)
	return err
}
