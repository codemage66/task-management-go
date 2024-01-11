package types

import (
	"strconv"
)

type JSONResult struct {
	Code    int         `json:"code"    example:"200"`
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data"`
}

type JSONError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Pointer string `json:"pointer" example:"'field_name'"`
}

type JSONResultWithPaginate struct {
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
	Length   int         `json:"length"`
	Paginate *Pageable   `json:"paginate"`
}

type Pageable struct {
	Offset  uint64 `json:"offset"`
	Limit   uint64 `json:"limit"`
	Total   int64  `json:"total"`
	Prev    int64  `json:"prev"`
	Next    int64  `json:"next"`
	HasNext bool   `json:"has_next"`
	HasPrev bool   `json:"has_prev"`
}

func (p Pageable) Validate() (uint64, uint64) {
	if p.Limit < 1 {
		p.Limit = 10
	}

	return p.Offset, p.Limit
}

// Parse parses the limit and offset query parameters
func (p Pageable) Parse(limit, offset string) Pageable {
	if offset == "" {
		offset = "0"
	}

	if limit == "" {
		limit = "10"
	}

	l, _ := strconv.ParseUint(limit, 10, 64)

	if l > 50 {
		l = 50
	}

	p.Limit = l

	o, _ := strconv.ParseUint(offset, 10, 64)
	p.Offset = o

	return p
}

func (p *Pageable) Calc() {
	p.HasNextPage()
	p.HasPrevPage()
	p.NextPage()
	p.PrevPage()
}

func (p *Pageable) NextPage() {
	if p.Total == 0 {
		p.Next = 0
		return
	}
	p.Next = int64((p.Offset + p.Limit))
	if p.Next >= p.Total {
		p.Next = p.Total
	}
}

func (p *Pageable) PrevPage() {
	p.Prev = int64((p.Offset - p.Limit)) + 1
	if p.Prev < 0 {
		p.Prev = 0
	}
}

func (p *Pageable) HasNextPage() {
	p.HasNext = p.Next > 0 && p.Next <= p.Total
}

func (p *Pageable) HasPrevPage() {
	p.HasPrev = p.Prev > 0 && p.Prev <= p.Total
}
