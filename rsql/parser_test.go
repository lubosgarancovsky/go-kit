package rsql

import (
	"reflect"
	"testing"
)

func TestBasicExpression(t *testing.T) {
	ast, err := New().Parse("name==John")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := &ComparisonNode{
		Field:    "name",
		Operator: "==",
		Value:    []string{"John"},
	}

	if !reflect.DeepEqual(ast, expected) {
		t.Errorf("AST does not match expected.\nGot:\n%#v\nExpected:\n%#v", ast, expected)
	}
}

func TestEscapedBasicExpression(t *testing.T) {
	ast, err := New().Parse("name==\"John Doe III.\"")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := &ComparisonNode{
		Field:    "name",
		Operator: "==",
		Value:    []string{"John Doe III."},
	}

	if !reflect.DeepEqual(ast, expected) {
		t.Errorf("AST does not match expected.\nGot:\n%#v\nExpected:\n%#v", ast, expected)
	}
}

func TestBasicExpressionWithArrayValue(t *testing.T) {
	ast, err := New().Parse("name=in=(\"John Doe III.\",Jane, Mark)")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := &ComparisonNode{
		Field:    "name",
		Operator: "=in=",
		Value:    []string{"John Doe III.", "Jane", "Mark"},
	}

	if !reflect.DeepEqual(ast, expected) {
		t.Errorf("AST does not match expected.\nGot:\n%#v\nExpected:\n%#v", ast, expected)
	}
}

func TestLogicAndExpression(t *testing.T) {
	ast, err := New().Parse("name==John;age>=18")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := &LogicalNode{
		Operator: "AND",
		Children: []Node{
			&ComparisonNode{
				Field:    "name",
				Operator: "==",
				Value:    []string{"John"},
			},
			&ComparisonNode{
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
	ast, err := New().Parse("name==John,age>=18")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := &LogicalNode{
		Operator: "OR",
		Children: []Node{
			&ComparisonNode{
				Field:    "name",
				Operator: "==",
				Value:    []string{"John"},
			},
			&ComparisonNode{
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
	ast, err := New().Parse("(name==John,age>=18);isVerified==true")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := &LogicalNode{
		Operator: "AND",
		Children: []Node{
			&LogicalNode{
				Operator: "OR",
				Children: []Node{
					&ComparisonNode{
						Field:    "name",
						Operator: "==",
						Value:    []string{"John"},
					},
					&ComparisonNode{
						Field:    "age",
						Operator: ">=",
						Value:    []string{"18"},
					},
				},
			},
			&ComparisonNode{
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
	ast, err := New().Parse("isDisabled==false;(name==John,age>=18);isVerified==true")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := &LogicalNode{
		Operator: "AND",
		Children: []Node{
			&ComparisonNode{
				Field:    "isDisabled",
				Operator: "==",
				Value:    []string{"false"},
			},
			&LogicalNode{
				Operator: "OR",
				Children: []Node{
					&ComparisonNode{
						Field:    "name",
						Operator: "==",
						Value:    []string{"John"},
					},
					&ComparisonNode{
						Field:    "age",
						Operator: ">=",
						Value:    []string{"18"},
					},
				},
			},
			&ComparisonNode{
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
	rsqlParser := New()
	rsqlParser.RegisterOperator("<=>")

	ast, err := rsqlParser.Parse("name<=>John")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := &ComparisonNode{
		Field:    "name",
		Operator: "<=>",
		Value:    []string{"John"},
	}

	if !reflect.DeepEqual(ast, expected) {
		t.Errorf("AST does not match expected.\nGot:\n%#v\nExpected:\n%#v", ast, expected)
	}
}
