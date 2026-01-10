package go_kit

import (
	"reflect"
	"testing"

	rsql2 "github.com/lubosgarancovsky/go-kit/internal/rsql"
)

func TestBasicExpression(t *testing.T) {
	ast, err := NewRSQLParser().Parse("name==John")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := &rsql2.ComparisonNode{
		Field:    "name",
		Operator: "==",
		Value:    []string{"John"},
	}

	if !reflect.DeepEqual(ast, expected) {
		t.Errorf("AST does not match expected.\nGot:\n%#v\nExpected:\n%#v", ast, expected)
	}
}

func TestEscapedBasicExpression(t *testing.T) {
	ast, err := NewRSQLParser().Parse("name==\"John Doe III.\"")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := &rsql2.ComparisonNode{
		Field:    "name",
		Operator: "==",
		Value:    []string{"John Doe III."},
	}

	if !reflect.DeepEqual(ast, expected) {
		t.Errorf("AST does not match expected.\nGot:\n%#v\nExpected:\n%#v", ast, expected)
	}
}

func TestBasicExpressionWithArrayValue(t *testing.T) {
	ast, err := NewRSQLParser().Parse("name=in=(\"John Doe III.\",Jane, Mark)")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := &rsql2.ComparisonNode{
		Field:    "name",
		Operator: "=in=",
		Value:    []string{"John Doe III.", "Jane", "Mark"},
	}

	if !reflect.DeepEqual(ast, expected) {
		t.Errorf("AST does not match expected.\nGot:\n%#v\nExpected:\n%#v", ast, expected)
	}
}

func TestLogicAndExpression(t *testing.T) {
	ast, err := NewRSQLParser().Parse("name==John;age>=18")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := &rsql2.LogicalNode{
		Operator: "AND",
		Children: []rsql2.Node{
			&rsql2.ComparisonNode{
				Field:    "name",
				Operator: "==",
				Value:    []string{"John"},
			},
			&rsql2.ComparisonNode{
				Field:    "age",
				Operator: ">=",
				Value:    []string{"18"},
			},
		},
	}

	if !reflect.DeepEqual(ast, expected) {
		t.Errorf("AST does not match expected.\nGot:\n%#v\nExpected:\n%#v", ast, expected)
	}
}

func TestLogicOrExpression(t *testing.T) {
	ast, err := NewRSQLParser().Parse("name==John,age>=18")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := &rsql2.LogicalNode{
		Operator: "OR",
		Children: []rsql2.Node{
			&rsql2.ComparisonNode{
				Field:    "name",
				Operator: "==",
				Value:    []string{"John"},
			},
			&rsql2.ComparisonNode{
				Field:    "age",
				Operator: ">=",
				Value:    []string{"18"},
			},
		},
	}

	if !reflect.DeepEqual(ast, expected) {
		t.Errorf("AST does not match expected.\nGot:\n%#v\nExpected:\n%#v", ast, expected)
	}
}

func TestLogicOrAndExpression(t *testing.T) {
	ast, err := NewRSQLParser().Parse("(name==John,age>=18);isVerified==true")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := &rsql2.LogicalNode{
		Operator: "AND",
		Children: []rsql2.Node{
			&rsql2.LogicalNode{
				Operator: "OR",
				Children: []rsql2.Node{
					&rsql2.ComparisonNode{
						Field:    "name",
						Operator: "==",
						Value:    []string{"John"},
					},
					&rsql2.ComparisonNode{
						Field:    "age",
						Operator: ">=",
						Value:    []string{"18"},
					},
				},
			},
			&rsql2.ComparisonNode{
				Field:    "isVerified",
				Operator: "==",
				Value:    []string{"true"},
			},
		},
	}

	if !reflect.DeepEqual(ast, expected) {
		t.Errorf("AST does not match expected.\nGot:\n%#v\nExpected:\n%#v", ast, expected)
	}
}

func TestAdvancedExpression(t *testing.T) {
	ast, err := NewRSQLParser().Parse("isDisabled==false;(name==John,age>=18);isVerified==true")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := &rsql2.LogicalNode{
		Operator: "AND",
		Children: []rsql2.Node{
			&rsql2.ComparisonNode{
				Field:    "isDisabled",
				Operator: "==",
				Value:    []string{"false"},
			},
			&rsql2.LogicalNode{
				Operator: "OR",
				Children: []rsql2.Node{
					&rsql2.ComparisonNode{
						Field:    "name",
						Operator: "==",
						Value:    []string{"John"},
					},
					&rsql2.ComparisonNode{
						Field:    "age",
						Operator: ">=",
						Value:    []string{"18"},
					},
				},
			},
			&rsql2.ComparisonNode{
				Field:    "isVerified",
				Operator: "==",
				Value:    []string{"true"},
			},
		},
	}

	if !reflect.DeepEqual(ast, expected) {
		t.Errorf("AST does not match expected.\nGot:\n%#v\nExpected:\n%#v", ast, expected)
	}
}

func TestCustomOperator(t *testing.T) {
	rsqlParser := NewRSQLParser()
	rsqlParser.RegisterOperator("<=>")

	ast, err := rsqlParser.Parse("name<=>John")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := &rsql2.ComparisonNode{
		Field:    "name",
		Operator: "<=>",
		Value:    []string{"John"},
	}

	if !reflect.DeepEqual(ast, expected) {
		t.Errorf("AST does not match expected.\nGot:\n%#v\nExpected:\n%#v", ast, expected)
	}
}
