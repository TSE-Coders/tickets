package generator

import "github.com/TSE-Coders/tickets/internal/store"

type Generator struct {
	ticketCount uint
	storeConfig store.DBConnectionConfig
	store       *store.DB
}

type Ticket struct {
	TicketID   string `json:"ticket_id"`
	Region     string `json:"region"`
	Difficulty uint8  `json:"difficulty"`
	Product    string `json:"product"`
}
