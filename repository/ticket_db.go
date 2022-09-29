package repository

import (
	"github.com/jmoiron/sqlx"
)

type ticketRepositoryDB struct {
	db *sqlx.DB
}

func NewTicketRepositoryDB(db *sqlx.DB) ticketRepositoryDB {
	return ticketRepositoryDB{db: db}
}

func (r ticketRepositoryDB) Create(ticket Ticket) (*Ticket, error) {
	query := `insert into tickets (title, description, contact, status)
	values (:title,:description,:contact,:status) returning *`
	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}

	var returning Ticket
	err = stmt.Get(&returning, ticket)
	if err != nil {
		return nil, err
	}

	ticket.TicketID = returning.TicketID
	ticket.Created_at = returning.Created_at
	return &ticket, nil
}

func (r ticketRepositoryDB) Update(id int, ticket Ticket) (*Ticket, error) {
	query := `UPDATE tickets SET title = :title, description = :description, contact = :contact, status = :status WHERE id = :id RETURNING *`
	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}

	var returning Ticket
	err = stmt.Get(&returning, ticket)
	if err != nil {
		return nil, err
	}

	ticket = returning
	return &ticket, nil
}

func (r ticketRepositoryDB) GetAll() ([]Ticket, error) {
	tickets := []Ticket{}
	query := "SELECT * FROM tickets ORDER BY status ASC, updated_at DESC"
	err := r.db.Select(&tickets, query)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r ticketRepositoryDB) GetbyStatus(status int) ([]Ticket, error) {
	tickets := []Ticket{}
	query := `SELECT * FROM tickets WHERE status = $1 ORDER BY updated_at DESC`
	err := r.db.Select(&tickets, query, status)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}
