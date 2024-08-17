package store

import (
	"github.com/jmoiron/sqlx"
)

type QueryExecutor interface {
	GetQuery() []string
	Execute(*sqlx.DB)
}
