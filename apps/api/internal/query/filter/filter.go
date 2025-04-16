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
	for _, item := range ctx.QueryArray(QueryParamIndex) {
		if item == "" {
			continue
		}

		parts := strings.Split(item, Separator)
		if len(parts) < 2 {
			continue
		}

		var params []string
		column := parts[0]
		op := operation.Operation(parts[1])

		if len(parts) > 2 {
			params = parts[2:]
		}

		f.rules[column] = rule{
			operation:  op,
			parameters: params,
		}
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
