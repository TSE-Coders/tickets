package generator

import (
	"math/rand"
	"strconv"

	"github.com/TSE-Coders/tickets/internal/store"
)

func NewGenerator(storeConfig store.DBConnectionConfig) (Generator, error) {
	g := Generator{
		ticketCount: 0,
		storeConfig: storeConfig,
	}

	store, err := store.NewDBConnection(storeConfig)
	if err != nil {
		return g, err
	}

	g.store = store

	g.loadAvailableProducts()
	g.loadAvailableRegions()

	return g, nil
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
	products, err := g.store.GetAllProducts()
	if err != nil {
		return nil
	}
	for _, product := range products {
		Products = append(Products, product.Name)
	}

	return nil
}

func (g *Generator) loadAvailableRegions() error {
	regions, err := g.store.GetAllRegions()
	if err != nil {
		return nil
	}
	for _, region := range regions {
		Regions = append(Regions, region.Name)
	}

	return nil
}
