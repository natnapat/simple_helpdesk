package service

import (
	"strings"
	"time"

	"github.com/natnapat/simple_helpdesk/errs"
	"github.com/natnapat/simple_helpdesk/logs"
	"github.com/natnapat/simple_helpdesk/repository"
)

type ticketService struct {
	ticketRepo repository.TicketRepository
}

func NewTicketService(ticketRepo repository.TicketRepository) ticketService {
	return ticketService{ticketRepo: ticketRepo}
}

func (s ticketService) CreateTicket(req TicketRequest) (*TicketResponse, error) {
	//validation
	if req.Title == "" || req.Description == "" || req.Contact == "" {
		logs.Error("incorrect request")
		return nil, errs.NewValidationError("incorrect request")
	}
	if req.Status < 0 || req.Status > 3 {
		logs.Error("incorrect request")
		return nil, errs.NewValidationError("incorrect request")
	}

	//setup repo
	ticket := repository.Ticket{
		TicketID:    0,
		Title:       req.Title,
		Description: req.Description,
		Contact:     req.Contact,
		Status:      req.Status,
		Created_at:  time.Time{},
		Updated_at:  &time.Time{},
	}
	newTicket, err := s.ticketRepo.Create(ticket)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	statusRes := checkStatus(newTicket.Status)

	//setup response
	res := TicketResponse{
		TicketID:    newTicket.TicketID,
		Title:       newTicket.Title,
		Description: newTicket.Description,
		Contact:     newTicket.Contact,
		Status:      statusRes,
		Created_at:  newTicket.Created_at.Format("2006-01-02 15:04:05"),
		Updated_at:  newTicket.Updated_at.Format("2006-01-02 15:04:05"),
	}
	return &res, nil
}

func (s ticketService) UpdateTicket(id int, req TicketRequest) (*TicketResponse, error) {
	//validation
	if req.Title == "" || req.Description == "" || req.Contact == "" {
		logs.Error("incorrect request")
		return nil, errs.NewValidationError("incorrect request")
	}
	if req.Status < 0 || req.Status > 3 {
		logs.Error("incorrect request")
		return nil, errs.NewValidationError("incorrect request")
	}

	//setup repo
	ticket := repository.Ticket{
		TicketID:    id,
		Title:       req.Title,
		Description: req.Description,
		Contact:     req.Contact,
		Status:      req.Status,
		Created_at:  time.Time{},
		Updated_at:  &time.Time{},
	}
	newTicket, err := s.ticketRepo.Update(id, ticket)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	//setup response
	statusRes := checkStatus(newTicket.Status)

	res := TicketResponse{
		TicketID:    newTicket.TicketID,
		Title:       newTicket.Title,
		Description: newTicket.Description,
		Contact:     newTicket.Contact,
		Status:      statusRes,
		Created_at:  newTicket.Created_at.Format("2006-01-02 15:04:05"),
		Updated_at:  newTicket.Updated_at.Format("2006-01-02 15:04:05"),
	}
	return &res, nil
}

func (s ticketService) GetTickets() ([]TicketResponse, error) {
	tickets, err := s.ticketRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	//setup response
	ticketResponses := []TicketResponse{}
	var updatedAt string
	for _, ticket := range tickets {
		if ticket.Updated_at == nil {
			updatedAt = ""
		} else {
			updatedAt = ticket.Updated_at.Format("2006-01-02 15:04:05")
		}

		status := checkStatus(ticket.Status)

		ticketResponse := TicketResponse{
			TicketID:    ticket.TicketID,
			Title:       ticket.Title,
			Description: ticket.Description,
			Contact:     ticket.Contact,
			Status:      status,
			Created_at:  ticket.Created_at.Format("2006-01-02 15:04:05"),
			Updated_at:  updatedAt,
		}
		ticketResponses = append(ticketResponses, ticketResponse)
	}
	return ticketResponses, nil
}

func (s ticketService) GetTicketsByStatus(status string) ([]TicketResponse, error) {
	//change ticket status format
	lower := strings.ToLower(status)
	var statusID int
	switch lower {
	case "pending":
		statusID = 0
	case "accepted":
		statusID = 1
	case "resolved":
		statusID = 2
	case "rejected":
		statusID = 3
	default:
		logs.Error("incorrect status")
		return nil, errs.NewValidationError("incorrect status")
	}

	//repo
	tickets, err := s.ticketRepo.GetbyStatus(statusID)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	//response
	ticketResponses := []TicketResponse{}
	var updatedAt string
	for _, ticket := range tickets {
		if ticket.Updated_at == nil {
			updatedAt = ""
		} else {
			updatedAt = ticket.Updated_at.Format("2006-01-02 15:04:05")
		}

		statusRes := checkStatus(ticket.Status)

		ticketResponse := TicketResponse{
			TicketID:    ticket.TicketID,
			Title:       ticket.Title,
			Description: ticket.Description,
			Contact:     ticket.Contact,
			Status:      statusRes,
			Created_at:  ticket.Created_at.Format("2006-01-02 15:04:05"),
			Updated_at:  updatedAt,
		}
		ticketResponses = append(ticketResponses, ticketResponse)
	}
	return ticketResponses, nil
}

//utilities
func checkStatus(status int) string {
	var result string
	switch status {
	case 0:
		result = "pending"
	case 1:
		result = "accepted"
	case 2:
		result = "resolved"
	case 3:
		result = "rejected"
	}
	return result
}
