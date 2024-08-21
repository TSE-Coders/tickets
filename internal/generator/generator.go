package generator

import (
	"fmt"
	"math/rand"

	"github.com/TSE-Coders/tickets/internal/store"
	"github.com/TSE-Coders/tickets/internal/types"
)

type Generator struct {
	TicketCount uint
	storeConfig store.DBConnectionConfig
	store       *store.DB
}

func NewGenerator(storeConfig store.DBConnectionConfig) (Generator, error) {
	g := Generator{
		TicketCount: 0,
		storeConfig: storeConfig,
	}

	store, err := store.NewDBConnection(storeConfig)
	if err != nil {
		return g, err
	}

	g.store = store

	return g, nil
}

func (g *Generator) GenetateTicket() types.Ticket {
	g.TicketCount += 1

	ticket := types.NewTicket()

	return ticket
}

func (g *Generator) GenetateRandomTicket() (types.Ticket, error) {
	randomOffice, err := g.store.GetRandomOffice()
	if err != nil {
		return types.Ticket{}, err
	}
	randomProduct, err := g.store.GetRandomProduct()
	if err != nil {
		return types.Ticket{}, err
	}
	randomDifficulty := rand.Intn(types.MaxTicketDifficulty)

	ticket := g.GenetateTicket().
		WithOffice(randomOffice.Name).
		WithProduct(randomProduct.Name).
		WithDifficulty(uint8(randomDifficulty))

	if err := g.store.InsertTicket(ticket); err != nil {
		return ticket, fmt.Errorf("failed to insert ticket: %s", err.Error())
	}

	return ticket, nil
}
