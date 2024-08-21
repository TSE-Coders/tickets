package generator

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/TSE-Coders/tickets/internal/store"
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

	err = g.loadAvailableProducts()
	if err != nil {
		return g, err
	}
	err = g.loadAvailableOffices()
	if err != nil {
		return g, err
	}

	return g, nil
}

func (g *Generator) GenetateTicket() Ticket {
	g.ticketCount += 1

	tick := NewTicket().WithTicketID(strconv.Itoa(int(g.ticketCount)))
	return tick
}

func (g *Generator) GenetateRandomTicket() Ticket {
	randomOffice := GetRandomOffice()
	randomProduct := GetRandomProduct()
	randomDifficulty := rand.Intn(MaxDifficulty)

	tick := g.GenetateTicket().
		WithOffice(randomOffice).
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

	if len(products) == 0 {
		return fmt.Errorf("no products loaded")
	}

	return nil
}

func (g *Generator) loadAvailableOffices() error {
	offices, err := g.store.GetAllOffices()
	if err != nil {
		return nil
	}
	for _, office := range offices {
		Offices = append(Offices, office.Name)
	}

	if len(offices) == 0 {
		return fmt.Errorf("no offices loaded")
	}

	return nil
}
