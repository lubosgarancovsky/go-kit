package filter

import (
	"reflect"
	"testing"

	"github.com/lubosgarancovsky/go-kit/rsql"
)

func TestBuildFilter(t *testing.T) {
	ast, err := rsql.New().Parse("name==\"John Doe\"")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	filter, err := BuildFilter(ast, map[string]string{"name": "name"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := &Filter{
		Query: "name = ?",
		Args:  []interface{}{"John Doe"},
	}

	if !reflect.DeepEqual(filter, expected) {
		t.Errorf("Filter does not match expected.\nGot:\n%#v\nExpected:\n%#v", filter, expected)
	}
}

func TestBuildFilterAND(t *testing.T) {
	ast, err := rsql.New().Parse("name==\"John Doe\";age=ge=18")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	filter, err := BuildFilter(ast, map[string]string{"name": "name", "age": "age"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := &Filter{
		Query: "(name = ? AND age >= ?)",
		Args:  []interface{}{"John Doe", "18"},
	}

	if !reflect.DeepEqual(filter, expected) {
		t.Errorf("Filter does not match expected.\nGot:\n%#v\nExpected:\n%#v", filter, expected)
	}
}

func TestBuildFilterOR(t *testing.T) {
	ast, err := rsql.New().Parse("name==\"John Doe\",age=ge=18")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	filter, err := BuildFilter(ast, map[string]string{"name": "name", "age": "age"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := &Filter{
		Query: "(name = ? OR age >= ?)",
		Args:  []interface{}{"John Doe", "18"},
	}

	if !reflect.DeepEqual(filter, expected) {
		t.Errorf("Filter does not match expected.\nGot:\n%#v\nExpected:\n%#v", filter, expected)
	}
}

func TestBuildFilterORAND(t *testing.T) {
	ast, err := rsql.New().Parse("(name==\"John Doe\";age=ge=18),id=in=(\"1\",\"2\",\"3\")")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	filter, err := BuildFilter(ast, map[string]string{"name": "name", "age": "age", "id": "id"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := &Filter{
		Query: "((name = ? AND age >= ?) OR id IN ?)",
		Args:  []interface{}{"John Doe", "18", []string{"1", "2", "3"}},
	}

	if !reflect.DeepEqual(filter, expected) {
		t.Errorf("Filter does not match expected.\nGot:\n%#v\nExpected:\n%#v", filter, expected)
	}
}
