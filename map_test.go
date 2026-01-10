package go_kit

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	map1 := map[string]string{"name": "John", "age": "20"}
	map2 := map[string]string{"age": "30", "city": "London"}

	result := MergeMaps(map1, map2)

	expected := map[string]string{"name": "John", "age": "30", "city": "London"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Merged map does not match expected.\nGot:\n%#v\nExpected:\n%#v", result, expected)
	}
}
