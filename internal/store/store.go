package store

import (
	"fmt"
	"log/slog"

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
	connectString := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", config.Host, config.User, config.DatabaseName, config.Password)
	connection, err := sqlx.Connect(config.DatabaseDriver, connectString)
	if err != nil {
		return nil, err
	}

	db := &DB{
		Connection:  connection,
		QueryBuffer: make(chan QueryExecutor, config.QueryBufferSize),
	}

	go db.connectionLoop()

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

func (db *DB) InsertRegion(region types.Region) error {
	resultChan := make(chan queries.InsertRegionResult)
	query := queries.NewInsertRegionQuery(resultChan, region)
	db.QueryBuffer <- query

	queryBufferLength := len(db.QueryBuffer)
	slog.Debug("query buffer size", "queued_queries_count", queryBufferLength)

	result := <-resultChan
	return result.Err
}

func (db *DB) InsertProduct(product types.Product) error {
	resultChan := make(chan queries.InsertProductResult)
	query := queries.NewInsertProductQuery(resultChan, product)
	db.QueryBuffer <- query

	queryBufferLength := len(db.QueryBuffer)
	slog.Debug("query buffer size", "queued_queries_count", queryBufferLength)

	result := <-resultChan
	return result.Err
}
