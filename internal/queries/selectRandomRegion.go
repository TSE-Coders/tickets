package queries

const SelectRandomRegionSQL = `
	SELECT * FROM tickets.region
	WHERE id >= floor(random() * (SELECT max(id) FROM tickets.region))
	ORDER BY id
	LIMIT 1;
	`

type SelectRandomRegionResult struct {
	Err error
}

type SelectRandomRegionQuery struct {
	SQL    []string
	Result chan SelectRandomRegionResult
}

func NewSelectRegionRandomQuery(result chan SelectRandomRegionResult) *SelectRandomRegionQuery {
	return &SelectRandomRegionQuery{
		SQL:    []string{SelectRandomRegionSQL},
		Result: result,
	}
}
