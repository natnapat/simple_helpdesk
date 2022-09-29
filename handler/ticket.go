package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/natnapat/simple_helpdesk/errs"
	"github.com/natnapat/simple_helpdesk/service"
)

type ticketHandler struct {
	ticketSrv service.TicketService
}

func NewTicketHandler(ticketSrv service.TicketService) ticketHandler {
	return ticketHandler{ticketSrv: ticketSrv}
}

func (h ticketHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		handleError(w, errs.NewValidationError("incorrect request format"))
		return
	}

	req := service.TicketRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handleError(w, errs.NewValidationError("incorrect request format"))
		return
	}

	res, err := h.ticketSrv.CreateTicket(req)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h ticketHandler) UpdateTicket(w http.ResponseWriter, r *http.Request) {
	ticketID, _ := strconv.Atoi(mux.Vars(r)["id"])

	if r.Header.Get("Content-Type") != "application/json" {
		handleError(w, errs.NewValidationError("incorrect request format"))
		return
	}

	req := service.TicketRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handleError(w, errs.NewValidationError("incorrect request format"))
		return
	}

	res, err := h.ticketSrv.UpdateTicket(ticketID, req)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h ticketHandler) GetTickets(w http.ResponseWriter, r *http.Request) {
	tickets, err := h.ticketSrv.GetTickets()
	if err != nil {
		handleError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tickets)
}

func (h ticketHandler) GetTicketsByStatus(w http.ResponseWriter, r *http.Request) {
	status := mux.Vars(r)["status"]
	tickets, err := h.ticketSrv.GetTicketsByStatus(status)
	if err != nil {
		handleError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tickets)
}
