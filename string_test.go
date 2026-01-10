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
