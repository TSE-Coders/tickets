package queries

import (
	"fmt"

	"github.com/TSE-Coders/tickets/internal/types"
	"github.com/jmoiron/sqlx"
)

const SelectRandomProductSQL = `
	SELECT * 
	FROM production.product 
	ORDER BY random() 
	LIMIT 1;
	`

type SelectRandomProductResult struct {
	Product types.Product
	Err     error
}

type SelectRandomProductQuery struct {
	SQL    []string
	Result chan SelectRandomProductResult
}

func NewSelectRandomProductQuery(result chan SelectRandomProductResult) *SelectRandomProductQuery {
	return &SelectRandomProductQuery{
		SQL:    []string{SelectRandomProductSQL},
		Result: result,
	}
}

func (q SelectRandomProductQuery) GetQuery() []string {
	return q.SQL
}

func (q SelectRandomProductQuery) Execute(dbConnection *sqlx.DB) {
	rows, err := dbConnection.Queryx(SelectRandomProductSQL)
	if err != nil {
		q.Result <- SelectRandomProductResult{
			Err: fmt.Errorf("failed to execute query: %s", err.Error()),
		}
		return
	}
	defer rows.Close()

	product := types.Product{}
	for rows.Next() {
		err = rows.StructScan(&product)
		if err != nil {
			q.Result <- SelectRandomProductResult{
				Product: product,
				Err:     err,
			}
			return
		}
	}

	q.Result <- SelectRandomProductResult{
		Product: product,
		Err:     nil,
	}
}
