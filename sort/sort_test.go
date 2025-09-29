package sort

import (
	"reflect"
	"testing"
)

func TestBuildSortDESC(t *testing.T) {
	sortString := "name;"

	sort, err := BuildSort(sortString, map[string]string{"name": "name"})
	if err != nil {
		t.Fatal(err)
	}

	expected := []Sort{Sort{Field: "name", Direction: "DESC"}}

	if !reflect.DeepEqual(sort, expected) {
		t.Errorf("Filter does not match expected.\nGot:\n%#v\nExpected:\n%#v", sort, expected)
	}
}

func TestBuildSortASC(t *testing.T) {
	sortString := "name:"

	sort, err := BuildSort(sortString, map[string]string{"name": "name"})
	if err != nil {
		t.Fatal(err)
	}

	expected := []Sort{Sort{Field: "name", Direction: "ASC"}}

	if !reflect.DeepEqual(sort, expected) {
		t.Errorf("Filter does not match expected.\nGot:\n%#v\nExpected:\n%#v", sort, expected)
	}
}

func TestBuildSortMultiple(t *testing.T) {
	sortString := "name;,age:"

	sort, err := BuildSort(sortString, map[string]string{"name": "name", "age": "age"})
	if err != nil {
		t.Fatal(err)
	}

	expected := []Sort{{Field: "name", Direction: "DESC"}, {Field: "age", Direction: "ASC"}}

	if !reflect.DeepEqual(sort, expected) {
		t.Errorf("Filter does not match expected.\nGot:\n%#v\nExpected:\n%#v", sort, expected)
	}
}
