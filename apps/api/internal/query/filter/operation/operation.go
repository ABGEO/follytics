package operation

import (
	"gorm.io/gorm/clause"
)

type (
	Operation string
	Handler   func(column string, params ...string) clause.Expression
)

func GetHandlers() map[Operation]Handler {
	return map[Operation]Handler{
		Eq: eqHandler,
	}
}
