package tickets

import (
	"database/sql"
	"errors"
	"strings"

	_ "github.com/go-sql-driver/mysql" //this is needed for the mysql driver
	"github.com/mickelsonm/demo-api/helpers/database"
)

//Tickets is a collection of tickets
type Tickets []Ticket

//Ticket is a structured ticket
type Ticket struct {
	ID        int    `json:"id"`
	ShortDesc string `json:"short_desc"`
	LongDesc  string `json:"long_desc"`
}

var (
	getAllTickets = `select ID, ShortDesc, LongDesc from Tickets;`
	getTicket     = `select ID, ShortDesc, LongDesc from Tickets where ID = ?;`
	addTicket     = `insert into Tickets(ShortDesc,LongDesc,Created,LastUpdated) values (?, ?, UTC_TIMESTAMP(), null);`
	updateTicket  = `update Tickets set ShortDesc = ?, LongDesc = ?, LastUpdated = UTC_TIMESTAMP() where ID = ?;`
	deleteTicket  = `delete from Tickets where ID = ?;`
)

//GetAll - Gets all tickets
func GetAll() (Tickets, error) {
	var tickets Tickets
	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return tickets, err
	}
	defer db.Close()

	stmt, err := db.Prepare(getAllTickets)
	if err != nil {
		return tickets, err
	}
	defer stmt.Close()

	res, err := stmt.Query()
	if err != nil {
		return tickets, err
	}

	for res.Next() {
		var t Ticket

		err = res.Scan(&t.ID, &t.ShortDesc, &t.LongDesc)
		if err != nil {
			return tickets, err
		}

		tickets = append(tickets, t)
	}
	defer res.Close()

	return tickets, nil
}

//Get - gets a ticket by an ID
func (t *Ticket) Get() error {
	if t.ID == 0 {
		return errors.New("Invalid ticket ID")
	}

	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(getTicket)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var tk Ticket

	row := stmt.QueryRow(t.ID)

	err = row.Scan(&tk.ID, &tk.ShortDesc, &tk.LongDesc)
	if err != nil {
		return err
	}

	t.ID = tk.ID
	t.ShortDesc = tk.ShortDesc
	t.LongDesc = tk.LongDesc

	return nil
}

//Add - adds a ticket
func (t *Ticket) Add() error {
	if len(strings.TrimSpace(t.ShortDesc)) == 0 {
		return errors.New("Ticket must have a short description")
	}

	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(addTicket)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(t.ShortDesc, t.LongDesc)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	t.ID = int(id)

	return nil
}

//Update - updates a ticket
func (t *Ticket) Update() error {
	if len(strings.TrimSpace(t.ShortDesc)) == 0 {
		return errors.New("Ticket must have a short description")
	}

	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(updateTicket)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(t.ShortDesc, t.LongDesc, t.ID); err != nil {
		return err
	}

	return nil
}

//Delete - deletes a ticket by its ID
func (t *Ticket) Delete() error {
	if t.ID == 0 {
		return errors.New("Invalid ticket ID")
	}

	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(deleteTicket)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(t.ID); err != nil {
		return err
	}

	return nil
}
