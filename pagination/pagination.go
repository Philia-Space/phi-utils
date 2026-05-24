// Package pagination provides offset and cursor-based pagination helpers.
package pagination

import (
	"fmt"
	"math"
)

// OffsetPagination represents traditional page-based pagination.
type OffsetPagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

// CursorPagination represents cursor-based (keyset) pagination.
type CursorPagination struct {
	Cursor string `json:"cursor"`
	Limit  int    `json:"limit"`
}

// PageResponse is the standard paginated response envelope.
type PageResponse struct {
	Data       interface{} `json:"data"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	TotalPages int         `json:"total_pages"`
	HasNext    bool        `json:"has_next"`
	HasPrev    bool        `json:"has_prev"`
	NextCursor string      `json:"next_cursor,omitempty"`
}

// NewPageResponse creates a paginated response from offset pagination.
func NewPageResponse(data interface{}, total int64, page, limit int) PageResponse {
	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	return PageResponse{
		Data:       data,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
		HasNext:    page < totalPages,
		HasPrev:    page > 1,
	}
}

// NewCursorResponse creates a paginated response from cursor pagination.
func NewCursorResponse(data interface{}, total int64, limit int, nextCursor string) PageResponse {
	return PageResponse{
		Data:       data,
		Total:      total,
		Limit:      limit,
		HasNext:    nextCursor != "",
		NextCursor: nextCursor,
	}
}

// Offset returns the SQL OFFSET value.
func (p OffsetPagination) Offset() int {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit < 1 {
		p.Limit = 20
	}
	return (p.Page - 1) * p.Limit
}

// Normalize ensures valid pagination values.
func (p *OffsetPagination) Normalize() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit < 1 {
		p.Limit = 20
	}
	if p.Limit > 1000 {
		p.Limit = 1000
	}
}

// Normalize ensures valid cursor pagination values.
func (p *CursorPagination) Normalize() {
	if p.Limit < 1 {
		p.Limit = 20
	}
	if p.Limit > 1000 {
		p.Limit = 1000
	}
}

// CursorEncoder is a function type for encoding cursor values.
type CursorEncoder func(values ...interface{}) string

// CursorDecoder is a function type for decoding cursor values.
type CursorDecoder func(cursor string) ([]interface{}, error)

// DefaultLimit is the default page size.
const DefaultLimit = 20

// MaxLimit is the maximum allowed page size.
const MaxLimit = 1000

// ValidateOffset checks that page and limit are within bounds.
func ValidateOffset(page, limit int) (int, int, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = DefaultLimit
	}
	if limit > MaxLimit {
		return 0, 0, fmt.Errorf("limit exceeds maximum of %d", MaxLimit)
	}
	return page, limit, nil
}
