package filter

import (
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/abgeo/follytics/internal/query/filter/operation"
)

const (
	QueryParamIndex = "filter"
	Separator       = "||"
)

// Filterer interface defines filter contract.
type Filterer interface {
	Apply(tx *gorm.DB) *gorm.DB
}

type rule struct {
	operation  operation.Operation
	parameters []string
}

// Filter provides filter functionality for Gorm queries.
type Filter struct {
	handlers map[operation.Operation]operation.Handler
	rules    map[string]rule
}

var _ Filterer = (*Filter)(nil)

// New creates a new Filterer with default values.
func New() *Filter {
	return &Filter{
		handlers: operation.GetHandlers(),
		rules:    make(map[string]rule),
	}
}

// FromContext populates filter parameters from Gin context.
func (f *Filter) FromContext(ctx *gin.Context) *Filter {
	for column, value := range ctx.QueryMap(QueryParamIndex) {
		if value == "" {
			continue
		}

		valueParts := strings.Split(value, Separator)
		ruleInstance := rule{
			operation: operation.Operation(valueParts[0]),
		}

		if len(valueParts) > 1 {
			ruleInstance.parameters = valueParts[1:]
		}

		f.rules[column] = ruleInstance
	}

	return f
}

// Apply modifies the Gorm query to apply filters.
func (f *Filter) Apply(tx *gorm.DB) *gorm.DB {
	clauses := make([]clause.Expression, 0)

	for column, item := range f.rules {
		if handler, ok := f.handlers[item.operation]; ok {
			clauses = append(clauses, handler(column, item.parameters...))
		}
	}

	return tx.Clauses(clauses...)
}
