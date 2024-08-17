package store

import (
	"fmt"
	"io/fs"
	"log/slog"
	"os"

	"github.com/TSE-Coders/tickets/internal/queries"
	"github.com/TSE-Coders/tickets/internal/types"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
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

func (db *DB) SeedDB(regions []types.Region, products []types.Product) error {
	return nil
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

func MigrateDB(migrations fs.FS) error {
	config := NewDBConnectionConfig().WithPassword("password")
	db, err := NewDBConnection(config)
	if err != nil {
		return err
	}
	defer db.Connection.Close()

	goose.SetBaseFS(migrations)
	if err = goose.SetDialect(config.DatabaseDriver); err != nil {
		slog.Error("Failed to select a dialect", "error", err)
		os.Exit(1)
	}

	if err = goose.Up(db.Connection.DB, "embed/migrations"); err != nil {
		slog.Error("Failed to apply migrations", "error", err)
		os.Exit(1)
	}

	return nil
}
