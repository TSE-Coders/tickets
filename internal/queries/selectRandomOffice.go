package queries

const SelectRandomOfficeSQL = `
	SELECT * FROM production.office
	WHERE id >= floor(random() * (SELECT max(id) FROM production.office))
	ORDER BY id
	LIMIT 1;
	`

type SelectRandomOfficeResult struct {
	Err error
}

type SelectRandomOfficeQuery struct {
	SQL    []string
	Result chan SelectRandomOfficeResult
}

func NewSelectOfficeRandomQuery(result chan SelectRandomOfficeResult) *SelectRandomOfficeQuery {
	return &SelectRandomOfficeQuery{
		SQL:    []string{SelectRandomOfficeSQL},
		Result: result,
	}
}
