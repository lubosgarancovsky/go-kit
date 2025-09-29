package filter

import (
	"fmt"
	"strings"

	"github.com/lubosgarancovsky/go-kit/rsql"
)

type Filter struct {
	Query string
	Args  []interface{}
}

var rsqlToSQL = map[string]string{
	"==":     "=",
	"!=":     "!=",
	"=gt=":   ">",
	"=lt=":   "<",
	"=ge=":   ">=",
	"=le=":   "<=",
	"=in=":   "IN",
	"=out=":  "NOT IN",
	"=like=": "LIKE",
}

func BuildFilter(node rsql.Node, fieldMap map[string]string) (*Filter, error) {
	return rsqlNode(node, fieldMap)
}

func rsqlNode(node rsql.Node, fieldMap map[string]string) (*Filter, error) {
	if val, ok := node.(*rsql.LogicalNode); ok {
		return logicalNode(val, fieldMap)
	}

	if val, ok := node.(*rsql.ComparisonNode); ok {
		return comparisonNode(val, fieldMap)
	}

	return nil, fmt.Errorf("unsupported node type")
}

func logicalNode(node *rsql.LogicalNode, fieldMap map[string]string) (*Filter, error) {
	var results []*Filter

	for _, child := range node.Children {
		result, err := rsqlNode(child, fieldMap)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	connector := " AND "
	if node.Operator == rsql.TokenOr {
		connector = " OR "
	}

	queries := make([]string, 0, len(results))
	args := make([]interface{}, 0, len(results))
	for condition := range results {
		queries = append(queries, results[condition].Query)
		args = append(args, results[condition].Args...)
	}

	return &Filter{
		Query: "(" + strings.Join(queries, connector) + ")",
		Args:  args,
	}, nil
}

func comparisonNode(node *rsql.ComparisonNode, fieldMap map[string]string) (*Filter, error) {
	operator, ok := rsqlToSQL[node.Operator]
	if !ok {
		return nil, fmt.Errorf("unsupported comparison operator %s", node.Operator)
	}

	field, ok := fieldMap[node.Field]
	if !ok {
		return nil, fmt.Errorf("field %s is not filterable", node.Field)
	}

	var args []interface{}
	switch operator {
	case "IN", "NOT IN":
		args = []interface{}{node.Value}
	default:
		if len(node.Value) == 0 {
			return nil, fmt.Errorf("no value provided for field %s", node.Field)
		}
		args = []interface{}{node.Value[0]}
	}

	return &Filter{
		Query: fmt.Sprintf("%s %s ?", field, operator),
		Args:  args,
	}, nil
}
