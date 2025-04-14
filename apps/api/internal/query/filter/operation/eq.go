package operation

import "gorm.io/gorm/clause"

const Eq Operation = "eq"

func eqHandler(column string, params ...string) clause.Expression {
	return clause.Eq{
		Column: column,
		Value:  params[0],
	}
}
