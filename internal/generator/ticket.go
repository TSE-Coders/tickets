package generator

const (
	MaxDifficulty = 10
)

func init() {
	Regions.addRegion(&Regions.NYC, "NYC")
	Regions.addRegion(&Regions.Boston, "Boston")
	Regions.addRegion(&Regions.Denver, "Denver")

	Products.addProduct(&Products.APM, "APM")
	Products.addProduct(&Products.DBM, "DBM")
	Products.addProduct(&Products.Monitors, "Monitors")
	Products.addProduct(&Products.SynRUM, "Synthetics/RUM")
}

func NewTicket() Ticket {
	return Ticket{
		TicketID:   "12345678",
		Region:     Regions.NYC,
		Difficulty: 5,
		Product:    Products.SynRUM,
	}
}
func (t Ticket) WithTicketID(id string) Ticket {
	t.TicketID = id
	return t
}
func (t Ticket) WithRegion(region string) Ticket {
	t.Region = region
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
