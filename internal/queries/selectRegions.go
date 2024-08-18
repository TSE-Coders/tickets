package queries

import (
	"github.com/TSE-Coders/tickets/internal/types"
	"github.com/jmoiron/sqlx"
)

const SelectRegionsSQL = `
	SELECT * FROM tickets.region;
	`

type SelectRegionsResult struct {
	Regions []types.Region
	Err     error
}

type SelectRegionsQuery struct {
	SQL    []string
	Result chan SelectRegionsResult
}

func NewSelectRegionsQuery(result chan SelectRegionsResult) *SelectRegionsQuery {
	return &SelectRegionsQuery{
		SQL:    []string{SelectRegionsSQL},
		Result: result,
	}
}

func (q SelectRegionsQuery) GetQuery() []string {
	return q.SQL
}

func (q SelectRegionsQuery) Execute(dbConnection *sqlx.DB) {
	regions := []types.Region{}
	rows, err := dbConnection.Queryx(SelectRegionsSQL)
	if err != nil {
		q.Result <- SelectRegionsResult{
			Err: err,
		}
		return
	}
	defer rows.Close()

	region := types.Region{}
	if rows.Next() {
		err = rows.StructScan(&region)
		if err != nil {
			q.Result <- SelectRegionsResult{
				Regions: regions,
				Err:     err,
			}
			return
		}
		regions = append(regions, region)
	}
	q.Result <- SelectRegionsResult{
		Regions: regions,
		Err:     nil,
	}
}
