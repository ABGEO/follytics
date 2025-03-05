package pagination

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	DefaultPage   = 1
	DefaultOffset = 0
	DefaultLimit  = 10

	MaxLimit = 100

	LimitQueryParam  = "limit"
	PageQueryParam   = "page"
	OffsetQueryParam = "offset"
)

// Paginator interface defines pagination contract.
type Paginator interface {
	Apply(tx *gorm.DB) *gorm.DB
	GetMetadata() *Metadata
}

// Metadata represents pagination information.
type Metadata struct {
	Page       int `example:"1"   json:"page"`
	Limit      int `example:"10"  json:"limit"`
	TotalItems int `example:"100" json:"totalItems"`
	TotalPages int `example:"10"  json:"totalPages"`
}

// Pagination provides pagination functionality for Gorm queries.
type Pagination struct {
	limit  int
	page   int
	offset int

	totalItems int64
	totalPages int
}

var _ Paginator = (*Pagination)(nil)

// New creates a new GormPaginator with default values.
func New() *Pagination {
	return &Pagination{
		limit:  DefaultLimit,
		page:   DefaultPage,
		offset: DefaultOffset,
	}
}

// FromContext populates pagination parameters from Gin context.
func (p *Pagination) FromContext(ctx *gin.Context) *Pagination {
	if page, err := strconv.Atoi(ctx.Query(PageQueryParam)); err == nil {
		p.WithPage(page)
	}

	if limit, err := strconv.Atoi(ctx.Query(LimitQueryParam)); err == nil {
		p.WithLimit(limit)
	}

	if offset, err := strconv.Atoi(ctx.Query(OffsetQueryParam)); err == nil {
		p.WithOffset(offset)
	}

	return p
}

// WithLimit sets the number of items per page with validation.
func (p *Pagination) WithLimit(limit int) *Pagination {
	switch {
	case limit <= 0:
		p.limit = DefaultLimit
	case limit > MaxLimit:
		p.limit = MaxLimit
	default:
		p.limit = limit
	}

	return p
}

// WithPage sets the current page number.
func (p *Pagination) WithPage(page int) *Pagination {
	if page > 0 {
		p.page = page
	} else {
		p.page = DefaultPage
	}

	return p
}

// WithOffset sets the starting offset for pagination.
func (p *Pagination) WithOffset(offset int) *Pagination {
	if offset >= 0 {
		p.offset = offset
	} else {
		p.offset = DefaultOffset
	}

	return p
}

// Apply modifies the Gorm query to implement pagination.
func (p *Pagination) Apply(tx *gorm.DB) *gorm.DB {
	tx.Session(&gorm.Session{}).Count(&p.totalItems)

	if p.offset == 0 && p.page > 0 {
		p.offset = (p.page - 1) * p.limit
	}

	p.totalPages = int(math.Ceil(float64(p.totalItems) / float64(p.limit)))

	return tx.Offset(p.offset).Limit(p.limit)
}

// GetMetadata returns pagination metadata.
func (p *Pagination) GetMetadata() *Metadata {
	return &Metadata{
		Page:       p.page,
		Limit:      p.limit,
		TotalItems: int(p.totalItems),
		TotalPages: p.totalPages,
	}
}
