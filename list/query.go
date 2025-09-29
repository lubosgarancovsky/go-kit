package list

import (
	"github.com/lubosgarancovsky/go-kit/filter"
	"github.com/lubosgarancovsky/go-kit/sort"
)

type QueryParms struct {
	Page     int    `query:"page" default:"1"`
	PageSize int    `query:"page" default:"10"`
	Sort     string `query:"sort" default:""`
	Filter   string `query:"filter" default:""`
}

type ListingQuery struct {
	Limit  int
	Offset int
	Filter filter.Filter
	Sort   sort.Sort
}

type Page[T any] struct {
	Items      []T `json:"items"`
	Page       int `json:"page"`
	PageSize   int `json:"pageSize"`
	TotalCount int `json:"totalCount"`
}
