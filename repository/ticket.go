package repository

import "time"

type Ticket struct {
	TicketID    int        `db:"id"`
	Title       string     `db:"title"`
	Description string     `db:"description"`
	Contact     string     `db:"contact"`
	Status      int        `db:"status"`
	Created_at  time.Time  `db:"created_at"`
	Updated_at  *time.Time `db:"updated_at"`
}

type TicketRepository interface {
	Create(Ticket) (*Ticket, error)
	Update(int, Ticket) (*Ticket, error)
	GetAll() ([]Ticket, error)
	GetbyStatus(int) ([]Ticket, error)
}
