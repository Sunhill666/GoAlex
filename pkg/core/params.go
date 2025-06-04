package core

import (
	"fmt"
	"net/url"
	"strings"
)

type PaginationParams struct {
	Page    int
	PerPage int
	Cursor  string
}

func (p *PaginationParams) ToQuery() url.Values {
	q := url.Values{}
	if p.Page > 0 {
		q.Set("page", fmt.Sprintf("%d", p.Page))
	}
	if p.PerPage > 0 {
		q.Set("per-page", fmt.Sprintf("%d", p.PerPage))
	}
	if p.Cursor != "" {
		q.Set("cursor", p.Cursor)
	}
	return q
}

type QueryParams struct {
	Pagination *PaginationParams
	Filter     map[string]any
}

func (q *QueryParams) ToQuery() url.Values {
	query := url.Values{}
	if q.Pagination != nil {
		for k, vs := range q.Pagination.ToQuery() {
			for _, v := range vs {
				query.Add(k, v)
			}
		}
	}
	if q.Filter != nil {
		var sb strings.Builder
		first := true
		for k, v := range q.Filter {
			if !first {
				sb.WriteString(",")
			}
			escaped := strings.ReplaceAll(fmt.Sprint(v), "+", " ")
			sb.WriteString(fmt.Sprintf("%s:%s", k, escaped))
			first = false
		}
		query.Set("filter", sb.String())
	}
	return query
}
