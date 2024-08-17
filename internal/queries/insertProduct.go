package queries

import (
	"fmt"

	"github.com/TSE-Coders/tickets/internal/types"
	"github.com/jmoiron/sqlx"
)

const InsertProductSQL = `
	INSERT INTO tickets.Product (
		name
	) VALUES (
		:name
	 );
	`

type InsertProductResult struct {
	Err error
}

type InsertProductQuery struct {
	SQL     []string
	Product types.Product
	Result  chan InsertProductResult
}

func NewInsertProductQuery(result chan InsertProductResult, product types.Product) *InsertProductQuery {
	return &InsertProductQuery{
		SQL:     []string{InsertProductSQL},
		Product: product,
		Result:  result,
	}
}

func (q InsertProductQuery) GetQuery() []string {
	return q.SQL
}

func (q InsertProductQuery) Execute(dbConnection *sqlx.DB) {
	_, err := dbConnection.NamedQuery(q.SQL[0], q.Product)
	if err != nil {
		q.Result <- InsertProductResult{
			Err: fmt.Errorf("failed to insert region: %q", err),
		}
	}
	q.Result <- InsertProductResult{
		Err: nil,
	}
}
