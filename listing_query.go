package go_kit

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/lubosgarancovsky/go-kit/internal/rsql"
)

const DEFAULT_PAGE = 1
const DEFAULT_PAGE_SIZE = 10

type QueryParams struct {
	Page     *int    `form:"page"`
	PageSize *int    `form:"pageSize"`
	Sort     *string `form:"sort"`
	Filter   *string `form:"filter"`
}

type ListingQuery struct {
	Limit  int
	Offset int
	Page   int
	Filter *rsql.Filter
	Sort   *[]Sort
}

type Page[T any] struct {
	Items      []T   `json:"items"`
	Page       int   `json:"page"`
	PageSize   int   `json:"pageSize"`
	TotalCount int64 `json:"totalCount"`
}

type List[T any] struct {
	Items []T `json:"items"`
}

func NewListingQuery(qp *QueryParams, parser *Parser, listingAttr interface{}) (*ListingQuery, error) {
	lq := &ListingQuery{}

	if qp.Page == nil || *qp.Page <= 0 {
		lq.Page = DEFAULT_PAGE
	} else {
		lq.Page = *qp.Page
	}

	if qp.PageSize == nil || *qp.PageSize <= 0 {
		lq.Limit = DEFAULT_PAGE_SIZE
	} else {
		lq.Limit = *qp.PageSize
	}

	lq.Offset = (lq.Page - 1) * lq.Limit

	filterMap, sortMap, err := parseListingAttribute(listingAttr)
	if err != nil {
		return nil, err
	}

	if qp.Filter != nil && *qp.Filter != "" {
		ast, err := parser.Parse(*qp.Filter)
		if err != nil {
			return nil, err
		}

		fil, err := rsql.BuildFilter(ast, filterMap)
		if err != nil {
			return nil, err
		}

		lq.Filter = fil
	}

	if qp.Sort != nil && len(*qp.Sort) > 0 {
		srt, err := BuildSort(*qp.Sort, sortMap)
		if err != nil {
			return nil, err
		}
		lq.Sort = &srt
	}

	return lq, nil
}

func parseListingAttribute(attribute interface{}) (map[string]string, map[string]string, error) {
	filterMap := make(map[string]string)
	sortMap := make(map[string]string)

	t := reflect.TypeOf(attribute)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return filterMap, sortMap, fmt.Errorf("unexpected type %T of listing attribute parameter", attribute)
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("rsql")
		if tag == "" {
			continue
		}

		parts := strings.Split(tag, ",")
		fieldName := ToSnakeCase(field.Name) // default field mapping
		for _, part := range parts {
			part = strings.TrimSpace(part)
			if strings.HasPrefix(part, "field:") {
				fieldName = strings.TrimPrefix(part, "field:")
				continue
			}
			switch part {
			case "filter":
				filterMap[ToCamelCase(field.Name)] = fieldName
			case "sort":
				sortMap[ToCamelCase(field.Name)] = fieldName
			}
		}
	}

	return filterMap, sortMap, nil
}
