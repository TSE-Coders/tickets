package generator

import (
	"math/rand"
	"strconv"

	"github.com/TSE-Coders/tickets/internal/store"
)

func NewGenerator(db *store.DB) Generator {
	g := Generator{
		ticketCount: 0,
		db:          db,
	}
	g.loadAvailableProducts()
	g.loadAvailableRegions()

	return g
}

func (g *Generator) GenetateTicket() Ticket {
	g.ticketCount += 1

	tick := NewTicket().WithTicketID(strconv.Itoa(int(g.ticketCount)))
	return tick
}

func (g *Generator) GenetateRandomTicket() Ticket {
	randomRegion := GetRandomRegion()
	randomProduct := GetRandomProduct()
	randomDifficulty := rand.Intn(MaxDifficulty)

	tick := g.GenetateTicket().
		WithRegion(randomRegion).
		WithProduct(randomProduct).
		WithDifficulty(uint8(randomDifficulty))

	return tick
}

func (g *Generator) loadAvailableProducts() error {
	products, err := g.db.GetAllProducts()
	if err != nil {
		return nil
	}
	for _, product := range products {
		Products = append(Products, product.Name)
	}

	return nil
}

func (g *Generator) loadAvailableRegions() error {
	regions, err := g.db.GetAllRegions()
	if err != nil {
		return nil
	}
	for _, region := range regions {
		Regions = append(Regions, region.Name)
	}

	return nil
}
