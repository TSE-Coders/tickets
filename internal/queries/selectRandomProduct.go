package queries

const SelectRandomProductSQL = `
	SELECT * FROM tickets.product
	WHERE id >= floor(random() * (SELECT max(id) FROM tickets.product))
	ORDER BY id
	LIMIT 1;
	`

type SelectRandomProductResult struct {
	Err error
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
