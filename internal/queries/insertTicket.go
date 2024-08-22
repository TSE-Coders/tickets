package queries

import (
	"github.com/TSE-Coders/tickets/internal/types"
	"github.com/jmoiron/sqlx"
)

const InsertTicketSQL = `
	INSERT INTO production.ticket (
		office,
		difficulty,
		product
	) VALUES (
		:office,
		:difficulty,
		:product
	 );
	`

type InsertTicketResult struct {
	Err error
}

type InsertTicketQuery struct {
	SQL    []string
	Ticket types.Ticket
	Result chan InsertTicketResult
}

func NewInsertTicketQuery(result chan InsertTicketResult, ticket types.Ticket) *InsertTicketQuery {
	return &InsertTicketQuery{
		SQL:    []string{InsertTicketSQL},
		Ticket: ticket,
		Result: result,
	}
}

func (q InsertTicketQuery) GetQuery() []string {
	return q.SQL
}

func (q InsertTicketQuery) Execute(dbConnection *sqlx.DB) {
	rows, err := dbConnection.NamedQuery(q.SQL[0], q.Ticket)
	if err != nil {
		q.Result <- InsertTicketResult{
			Err: err,
		}
	}
	defer rows.Close()
	q.Result <- InsertTicketResult{
		Err: nil,
	}
}
