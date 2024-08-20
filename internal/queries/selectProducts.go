package queries

import (
	"github.com/TSE-Coders/tickets/internal/types"
	"github.com/jmoiron/sqlx"
)

const SelectProductsSQL = `
	SELECT * FROM production.product;
	`

type SelectProductsResult struct {
	Products []types.Product
	Err      error
}

type SelectProductsQuery struct {
	SQL    []string
	Result chan SelectProductsResult
}

func NewSelectProductsQuery(result chan SelectProductsResult) *SelectProductsQuery {
	return &SelectProductsQuery{
		SQL:    []string{SelectProductsSQL},
		Result: result,
	}
}

func (q SelectProductsQuery) GetQuery() []string {
	return q.SQL
}

func (q SelectProductsQuery) Execute(dbConnection *sqlx.DB) {
	products := []types.Product{}

	rows, err := dbConnection.Queryx(SelectProductsSQL)
	if err != nil {
		q.Result <- SelectProductsResult{
			Err: err,
		}
		return
	}
	defer rows.Close()

	product := types.Product{}
	if rows.Next() {
		err = rows.StructScan(&product)
		if err != nil {
			q.Result <- SelectProductsResult{
				Products: products,
				Err:      err,
			}
			return
		}
		products = append(products, product)
	}

	q.Result <- SelectProductsResult{
		Products: products,
		Err:      nil,
	}
}
