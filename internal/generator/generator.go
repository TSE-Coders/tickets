package generator

import (
	"math/rand"
	"strconv"
)

func NewGenerator() Generator {
	return Generator{
		ticketCount: 0,
	}
}

func (g *Generator) GenetateTicket() Ticket {
	g.ticketCount += 1

	tick := NewTicket().WithTicketID(strconv.Itoa(int(g.ticketCount)))
	return tick
}

func (g *Generator) GenetateRandomTicket() Ticket {
	randomRegion := Regions.GetRandom()
	randomProduct := Products.GetRandom()
	randomDifficulty := rand.Intn(MaxDifficulty)

	tick := g.GenetateTicket().
		WithRegion(randomRegion).
		WithProduct(randomProduct).
		WithDifficulty(uint8(randomDifficulty))

	return tick
}
