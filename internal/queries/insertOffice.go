package queries

import (
	"github.com/TSE-Coders/tickets/internal/types"
	"github.com/jmoiron/sqlx"
)

const InsertOfficeSQL = `
	INSERT INTO production.office (
		name
	) VALUES (
		:name
	 );
	`

type InsertOfficeResult struct {
	Err error
}

type InsertOfficeQuery struct {
	SQL    []string
	Office types.Office
	Result chan InsertOfficeResult
}

func NewInsertOfficeQuery(result chan InsertOfficeResult, office types.Office) *InsertOfficeQuery {
	return &InsertOfficeQuery{
		SQL:    []string{InsertOfficeSQL},
		Office: office,
		Result: result,
	}
}

func (q InsertOfficeQuery) GetQuery() []string {
	return q.SQL
}

func (q InsertOfficeQuery) Execute(dbConnection *sqlx.DB) {
	rows, err := dbConnection.NamedQuery(InsertOfficeSQL, q.Office)
	if err != nil {
		q.Result <- InsertOfficeResult{
			Err: err,
		}
	}
	defer rows.Close()
	q.Result <- InsertOfficeResult{
		Err: nil,
	}
}
