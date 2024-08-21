package generator

const (
	MaxDifficulty = 10
)

type Ticket struct {
	TicketID   string `json:"ticket_id"`
	Office     string `json:"office"`
	Difficulty uint8  `json:"difficulty"`
	Product    string `json:"product"`
}

func NewTicket() Ticket {
	return Ticket{
		TicketID:   "12345678",
		Office:     "NYC",
		Difficulty: 5,
		Product:    "APM",
	}
}
func (t Ticket) WithTicketID(id string) Ticket {
	t.TicketID = id
	return t
}
func (t Ticket) WithOffice(office string) Ticket {
	t.Office = office
	return t
}
func (t Ticket) WithDifficulty(difficulty uint8) Ticket {
	t.Difficulty = difficulty
	return t
}
func (t Ticket) WithProduct(product string) Ticket {
	t.Product = product
	return t
}
