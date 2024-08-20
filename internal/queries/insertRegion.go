package queries

import (
	"github.com/TSE-Coders/tickets/internal/types"
	"github.com/jmoiron/sqlx"
)

const InsertRegionSQL = `
	INSERT INTO production.region (
		name
	) VALUES (
		:name
	 );
	`

type InsertRegionResult struct {
	Err error
}

type InsertRegionQuery struct {
	SQL    []string
	Region types.Region
	Result chan InsertRegionResult
}

func NewInsertRegionQuery(result chan InsertRegionResult, region types.Region) *InsertRegionQuery {
	return &InsertRegionQuery{
		SQL:    []string{InsertRegionSQL},
		Region: region,
		Result: result,
	}
}

func (q InsertRegionQuery) GetQuery() []string {
	return q.SQL
}

func (q InsertRegionQuery) Execute(dbConnection *sqlx.DB) {
	_, err := dbConnection.NamedQuery(q.SQL[0], q.Region)
	if err != nil {
		q.Result <- InsertRegionResult{
			Err: err,
		}
	}
	q.Result <- InsertRegionResult{
		Err: nil,
	}
}
