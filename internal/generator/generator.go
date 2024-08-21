package generator

import (
	"math/rand"
	"strconv"

	"github.com/TSE-Coders/tickets/internal/store"
	"github.com/TSE-Coders/tickets/internal/types"
)

type Generator struct {
	ticketCount uint
	storeConfig store.DBConnectionConfig
	store       *store.DB
}

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

	return g, nil
}

func (g *Generator) GenetateTicket() types.Ticket {
	g.ticketCount += 1

	tick := types.NewTicket().WithId(strconv.Itoa(int(g.ticketCount)))
	return tick
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

	tick := g.GenetateTicket().
		WithOffice(randomOffice.Name).
		WithProduct(randomProduct.Name).
		WithDifficulty(uint8(randomDifficulty))

	return tick, nil
}
