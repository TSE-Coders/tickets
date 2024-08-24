package queries

import (
	"fmt"

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
	 ) RETURNING id;
	`

type InsertTicketResult struct {
	Id  int
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
	rows, err := dbConnection.NamedQuery(InsertTicketSQL, q.Ticket)
	if err != nil {
		q.Result <- InsertTicketResult{
			Err: err,
		}
	}
	defer rows.Close()

	newTicketId := struct {
		Id int `db:"id"`
	}{}
	for rows.Next() {
		err = rows.StructScan(&newTicketId)
		if err != nil {
			q.Result <- InsertTicketResult{
				Err: fmt.Errorf("failed to get id from the inserted ticket"),
			}
		}
	}

	q.Result <- InsertTicketResult{
		Id:  newTicketId.Id,
		Err: nil,
	}
}
