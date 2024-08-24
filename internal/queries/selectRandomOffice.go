package queries

import (
	"fmt"

	"github.com/TSE-Coders/tickets/internal/types"
	"github.com/jmoiron/sqlx"
)

const SelectRandomOfficeSQL = `
	SELECT * 
	FROM production.office 
	ORDER BY random() 
	LIMIT 1;
	`

type SelectRandomOfficeResult struct {
	Office types.Office
	Err    error
}

type SelectRandomOfficeQuery struct {
	SQL    []string
	Result chan SelectRandomOfficeResult
}

func NewSelectRandomOfficeQuery(result chan SelectRandomOfficeResult) *SelectRandomOfficeQuery {
	return &SelectRandomOfficeQuery{
		SQL:    []string{SelectRandomOfficeSQL},
		Result: result,
	}
}

func (q SelectRandomOfficeQuery) GetQuery() []string {
	return q.SQL
}

func (q SelectRandomOfficeQuery) Execute(dbConnection *sqlx.DB) {
	rows, err := dbConnection.Queryx(SelectRandomOfficeSQL)
	if err != nil {
		q.Result <- SelectRandomOfficeResult{
			Err: fmt.Errorf("failed to execute query: %s", err.Error()),
		}
		return
	}
	defer rows.Close()

	office := types.Office{}
	for rows.Next() {
		err = rows.StructScan(&office)
		if err != nil {
			q.Result <- SelectRandomOfficeResult{
				Office: office,
				Err:    err,
			}
			return
		}
	}

	q.Result <- SelectRandomOfficeResult{
		Office: office,
		Err:    nil,
	}
}
