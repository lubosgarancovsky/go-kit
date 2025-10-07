package list

import (
	"github.com/lubosgarancovsky/go-kit/filter"
	"github.com/lubosgarancovsky/go-kit/sort"
)

type QueryParms struct {
	Page     int    `form:"page" default:"1"`
	PageSize int    `form:"pageSize" default:"10"`
	Sort     string `form:"sort" default:""`
	Filter   string `form:"filter" default:""`
}

type ListingQuery struct {
	Limit  int
	Offset int
	Filter *filter.Filter
	Sort   []sort.Sort
	Page   int
}

type Page[T any] struct {
	Items      []T   `json:"items"`
	Page       int   `json:"page"`
	PageSize   int   `json:"pageSize"`
	TotalCount int64 `json:"totalCount"`
}
