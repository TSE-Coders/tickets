package queries

import (
	"github.com/TSE-Coders/tickets/internal/types"
	"github.com/jmoiron/sqlx"
)

const SelectOfficesSQL = `
	SELECT * FROM production.office;
	`

type SelectOfficesResult struct {
	Offices []types.Office
	Err     error
}

type SelectOfficesQuery struct {
	SQL    []string
	Result chan SelectOfficesResult
}

func NewSelectOfficesQuery(result chan SelectOfficesResult) *SelectOfficesQuery {
	return &SelectOfficesQuery{
		SQL:    []string{SelectOfficesSQL},
		Result: result,
	}
}

func (q SelectOfficesQuery) GetQuery() []string {
	return q.SQL
}

func (q SelectOfficesQuery) Execute(dbConnection *sqlx.DB) {
	offices := []types.Office{}
	rows, err := dbConnection.Queryx(SelectOfficesSQL)
	if err != nil {
		q.Result <- SelectOfficesResult{
			Err: err,
		}
		return
	}
	defer rows.Close()

	office := types.Office{}
	for rows.Next() {
		err = rows.StructScan(&office)
		if err != nil {
			q.Result <- SelectOfficesResult{
				Offices: offices,
				Err:     err,
			}
			return
		}
		offices = append(offices, office)
	}
	q.Result <- SelectOfficesResult{
		Offices: offices,
		Err:     nil,
	}
}
