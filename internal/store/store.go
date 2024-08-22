package store

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/TSE-Coders/tickets/internal/queries"
	"github.com/TSE-Coders/tickets/internal/types"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	Connection  *sqlx.DB
	QueryBuffer chan QueryExecutor
}

func NewDBConnection(config DBConnectionConfig) (*DB, error) {
	connectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", config.Host, config.Port, config.User, config.DatabaseName, config.Password)
	connection, err := sqlx.Connect(config.DatabaseDriver, connectString)
	if err != nil {
		return nil, err
	}
	connection.SetMaxOpenConns(20)
	connection.SetMaxIdleConns(20)
	connection.SetConnMaxLifetime(5 * time.Minute)

	db := &DB{
		Connection:  connection,
		QueryBuffer: make(chan QueryExecutor, config.QueryBufferSize),
	}

	go db.connectionLoop()

	fmt.Println("DB Connection created")
	return db, nil
}

func NewDefaultDBConnection() (*DB, error) {
	defaultConfig := NewDBConnectionConfig()
	dbConnection, err := NewDBConnection(defaultConfig)
	if err != nil {
		return nil, err
	}

	return dbConnection, nil
}

func (db *DB) connectionLoop() {
	for q := range db.QueryBuffer {
		slog.Debug("Queries received", "queries", q.GetQuery())
		q.Execute(db.Connection)
	}
}

func (db *DB) loadDBQuery(query QueryExecutor) {
	db.QueryBuffer <- query
	queryBufferLength := len(db.QueryBuffer)
	slog.Debug("query buffer size", "queued_queries_count", queryBufferLength)
}

func (db *DB) InsertOffice(office types.Office) error {
	resultChan := make(chan queries.InsertOfficeResult)
	query := queries.NewInsertOfficeQuery(resultChan, office)
	db.loadDBQuery(query)

	result := <-resultChan
	return result.Err
}

func (db *DB) InsertProduct(product types.Product) error {
	resultChan := make(chan queries.InsertProductResult)
	query := queries.NewInsertProductQuery(resultChan, product)
	db.loadDBQuery(query)

	result := <-resultChan
	return result.Err
}

func (db *DB) InsertTicket(ticket types.Ticket) error {
	resultChan := make(chan queries.InsertTicketResult)
	query := queries.NewInsertTicketQuery(resultChan, ticket)
	db.loadDBQuery(query)

	result := <-resultChan
	return result.Err
}

func (db *DB) GetAllOffices() ([]types.Office, error) {
	resultChan := make(chan queries.SelectOfficesResult)
	query := queries.NewSelectOfficesQuery(resultChan)
	db.loadDBQuery(query)

	result := <-resultChan
	return result.Offices, result.Err
}

func (db *DB) GetAllProducts() ([]types.Product, error) {
	resultChan := make(chan queries.SelectProductsResult)
	query := queries.NewSelectProductsQuery(resultChan)
	db.loadDBQuery(query)

	result := <-resultChan
	return result.Products, result.Err
}

func (db *DB) GetRandomOffice() (types.Office, error) {
	resultChan := make(chan queries.SelectRandomOfficeResult)
	query := queries.NewSelectRandomOfficeQuery(resultChan)
	db.loadDBQuery(query)

	result := <-resultChan
	return result.Office, result.Err
}

func (db *DB) GetRandomProduct() (types.Product, error) {
	resultChan := make(chan queries.SelectRandomProductResult)
	query := queries.NewSelectRandomProductQuery(resultChan)
	db.loadDBQuery(query)

	result := <-resultChan
	return result.Product, result.Err
}
