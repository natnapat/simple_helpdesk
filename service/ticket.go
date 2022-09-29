package service

type TicketResponse struct {
	TicketID    int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Contact     string `json:"contact"`
	Status      string `json:"status"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}

type TicketRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Contact     string `json:"contact"`
	Status      int    `json:"status"`
}

type TicketService interface {
	CreateTicket(TicketRequest) (*TicketResponse, error)
	UpdateTicket(int, TicketRequest) (*TicketResponse, error)
	GetTickets() ([]TicketResponse, error)
	GetTicketsByStatus(string) ([]TicketResponse, error)
}
