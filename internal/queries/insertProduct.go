package queries

import (
	"github.com/TSE-Coders/tickets/internal/types"
	"github.com/jmoiron/sqlx"
)

const InsertProductSQL = `
	INSERT INTO production.Product (
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
	rows, err := dbConnection.NamedQuery(q.SQL[0], q.Product)
	if err != nil {
		q.Result <- InsertProductResult{
			Err: err,
		}
	}
	defer rows.Close()
	q.Result <- InsertProductResult{
		Err: nil,
	}
}
