package generator

type Generator struct {
	ticketCount uint
}

type Ticket struct {
	TicketID   string `json:"ticket_id"`
	Region     string `json:"region"`
	Difficulty uint8  `json:"difficulty"`
	Product    string `json:"product"`
}
