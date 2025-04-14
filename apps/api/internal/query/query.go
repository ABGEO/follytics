package query

import (
	"gorm.io/gorm"

	"github.com/abgeo/follytics/internal/query/filter"
	"github.com/abgeo/follytics/internal/query/pagination"
)

type Querier interface {
	Apply(tx *gorm.DB) *gorm.DB
}

type Query struct {
	paginator pagination.Paginator
	filterer  filter.Filterer
}

var _ Querier = (*Query)(nil)

func New() *Query {
	return &Query{}
}

func NewWithPaginator(paginator pagination.Paginator) *Query {
	return &Query{
		paginator: paginator,
	}
}

func NewWithFilterer(filterer filter.Filterer) *Query {
	return &Query{
		filterer: filterer,
	}
}

func (q *Query) HasPaginator() bool {
	return q.paginator != nil
}

func (q *Query) WithPaginator(paginator pagination.Paginator) *Query {
	q.paginator = paginator

	return q
}

func (q *Query) HasFilterer() bool {
	return q.filterer != nil
}

func (q *Query) WithFilterer(filterer filter.Filterer) *Query {
	q.filterer = filterer

	return q
}

func (q *Query) Apply(tx *gorm.DB) *gorm.DB {
	if q.HasFilterer() {
		tx = q.filterer.Apply(tx)
	}

	if q.HasPaginator() {
		tx = q.paginator.Apply(tx)
	}

	return tx
}
