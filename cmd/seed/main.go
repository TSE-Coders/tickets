package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "embed"

	"github.com/TSE-Coders/tickets/internal/store"
	"github.com/TSE-Coders/tickets/internal/types"
)

const (
	InsertSeedStatusSQL = `
	INSERT INTO production.seed (
		seeded
	) VALUES (
		:status
	 );
	`
	SelectSeedStatusSQL = `
	SELECT * FROM production.seed;
	`
)

type Seed struct {
	Status bool `json:"status"`
}

//go:embed data.json
var embeddedData []byte

type SeedData struct {
	Offices  []types.Office  `json:"offices"`
	Products []types.Product `json:"products"`
}

func seedDB(db *store.DB, data SeedData) error {
	rows, err := db.Connection.Queryx(SelectSeedStatusSQL)
	if err != nil {
		return fmt.Errorf("failed to check seed table: %s", err.Error())
	}
	defer rows.Close()

	if rows.Next() {
		fmt.Println("database already seeded")
		os.Exit(0)
	}

	for _, office := range data.Offices {
		err := db.InsertOffice(office)
		if err != nil {
			return fmt.Errorf("failed to insert office: %s", err.Error())
		}
	}

	for _, product := range data.Products {
		err := db.InsertProduct(product)
		if err != nil {
			return fmt.Errorf("failed to insert product: %s", err.Error())
		}
	}

	seedStatus := Seed{
		Status: true,
	}
	_, err = db.Connection.NamedQuery(InsertSeedStatusSQL, seedStatus)
	if err != nil {
		return fmt.Errorf("failed to update seed table: %s", err.Error())
	}

	return nil
}

func getSeedData() SeedData {
	seedData := SeedData{}

	err := json.Unmarshal(embeddedData, &seedData)
	if err != nil {
		log.Fatalf("failed to unmarshal the data: %s", err.Error())
	}

	return seedData
}

func main() {
	db, err := store.NewDefaultDBConnection()
	if err != nil {
		log.Fatalf("failed to connect to the database: %s", err.Error())
	}
	seedData := getSeedData()
	err = seedDB(db, seedData)
	if err != nil {
		log.Fatalf("failed to seed database: %s", err)
	}

	fmt.Println("database seeded")
}
