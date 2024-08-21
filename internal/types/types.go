package types

type Office struct {
	Name string `db:"name"`
}

type Product struct {
	Name string `db:"name"`
}

const (
	MaxTicketDifficulty = 10
)

type Ticket struct {
	Id         string `db:"id" json:"id"`
	Office     string `db:"office" json:"office"`
	Difficulty uint8  `db:"difficulty" json:"difficulty"`
	Product    string `db:"product" json:"product"`
}

func NewTicket() Ticket {
	return Ticket{
		Id:         "12345678",
		Office:     "NYC",
		Difficulty: 5,
		Product:    "APM",
	}
}
func (t Ticket) WithId(id string) Ticket {
	t.Id = id
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
