package models

import (
	// "fmt"
	"restapi/db"
	"time"

	
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

var events = []Event{}



func (e *Event) Save() error {

	query := `
	INSERT INTO events(name,description,location,dateTime,user_id)
	values(?,?,?,?,?)`
	stmt, err := db.DB.Prepare(query) // Prepare creates a prepared statement for later queries or executions
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}
	// fmt.Println(result)
	id, err := result.LastInsertId()
	e.Id = id
	return err
}

func GetAllEvent() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query) // Query executes a query that returns rows, typically a SELECT
	// instead of exec() we use Query for select statemnts
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event,error){
	query:="SELECT * FROM events where id = ?"
	row:=db.DB.QueryRow(query,id) // QueryRow executes a query that is expected to return at most one row.
	var event Event
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err!=nil{
		return nil,err
	}
	return &event,nil
}

func (event Event) Update() error{
	query:=	`UPDATE events SET  name=?,description=?,location=?,dateTime=?
	WHERE id= ?`
	stmt,err:=db.DB.Prepare(query)
	if err!=nil{
		return err
	}
	defer stmt.Close()
	_,err=stmt.Exec(event.Name,event.Description,event.Location,event.DateTime,event.Id)
	return err 
}

func (event Event) Delete() error{
	query:=	`DELETE FROM events WHERE id=?`
	stmt,err:=db.DB.Prepare(query)
	if err!=nil{
		return err
	}
	defer stmt.Close()
	_,err=stmt.Exec(event.Id)
	return err 
}

func (e Event) Register(userId int64) error{
	query:=`
	INSERT INTO registrations (event_id,user_id) VALUES(?,?)
	`
	stmt,err:=db.DB.Prepare(query)
	if err!=nil{
		return err
	}
	defer stmt.Close()

	_,err=stmt.Exec(e.Id,userId)
	return err 

}
func (e Event) CancelRegistration(userId int64) error{
	query:=`
	DELETE FROM registrations WHERE event_id=? AND user_id=?
	`
	stmt,err:=db.DB.Prepare(query)
	if err!=nil{
		return err
	}
	defer stmt.Close()

	_,err=stmt.Exec(e.Id,userId)
	return err 
}