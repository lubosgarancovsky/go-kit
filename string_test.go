package go_kit

import "testing"

func TestPascalToSnake(t *testing.T) {
	snake := ToSnakeCase("TestString")
	if snake != "test_string" {
		t.Errorf("Expected test_string, got %s", snake)
	}
}

func TestCamelToSnake(t *testing.T) {
	snake := ToSnakeCase("testString")
	if snake != "test_string" {
		t.Errorf("Expected test_string, got %s", snake)
	}
}

func TestSnakeToSnake(t *testing.T) {
	snake := ToSnakeCase("test_string")
	if snake != "test_string" {
		t.Errorf("Expected test_string, got %s", snake)
	}
}

func TestScreamingToSnake(t *testing.T) {
	snake := ToSnakeCase("TEST_STRING")
	if snake != "test_string" {
		t.Errorf("Expected test_string, got %s", snake)
	}
}

func TestKebabToSnake(t *testing.T) {
	snake := ToSnakeCase("test-string")
	if snake != "test_string" {
		t.Errorf("Expected test_string, got %s", snake)
	}
}

func TestSnakeToCamel(t *testing.T) {
	camel := ToCamelCase("test_string")
	if camel != "testString" {
		t.Errorf("Expected testString, got %s", camel)
	}
}

func TestPascalToCamel(t *testing.T) {
	camel := ToCamelCase("TestString")
	if camel != "testString" {
		t.Errorf("Expected testString, got %s", camel)
	}
}

func TestCamelToCamel(t *testing.T) {
	camel := ToCamelCase("testString")
	if camel != "testString" {
		t.Errorf("Expected testString, got %s", camel)
	}
}

func ScreamingToCamel(t *testing.T) {
	camel := ToCamelCase("TEST_STRING")
	if camel != "testString" {
		t.Errorf("Expected testString, got %s", camel)
	}
}

func KebabToCamel(t *testing.T) {
	camel := ToCamelCase("test-string")
	if camel != "testString" {
		t.Errorf("Expected testString, got %s", camel)
	}
}
