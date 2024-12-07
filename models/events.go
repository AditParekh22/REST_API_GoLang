package models

import (
	"fmt"
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

// var events []Event = []Event{}

func (e Event) Save() error {
	// Later save to database
	query := `
	INSERT INTO events (name, description, location, dataTime, user_id)
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return err
}

/*
This is a function not a method as
we do not call it on a particular
event but for all events.
*/
func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var e Event

		if err := rows.Scan(&e.ID, &e.Name, &e.Description,
			&e.Location, &e.DateTime, &e.UserId); err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}

func GetEventID(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE ID = ?` // ? add for protection against SQL injection
	row := db.DB.QueryRow(query, id)

	// defer rows.Close()
	var e Event
	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (e Event) Update() error {
	query := `
	UPDATE events 
	SET name = ?, description = ?, location = ?, dataTime = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	fmt.Println(result)

	return err

}

func (e Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID)
	return err
}
