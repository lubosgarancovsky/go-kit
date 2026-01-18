package go_kit

import (
	"reflect"
	"testing"

	"github.com/lubosgarancovsky/go-kit/internal/rsql"
)

type testListingAttributes struct {
	Name      string `rsql:"filter,sort"`
	Amount    string `rsql:"filter"`
	CreatedAt string `rsql:"sort"`
}

func TestNewListingQueryDefaults(t *testing.T) {
	qp := &QueryParams{
		Page:     nil,
		PageSize: nil,
		Sort:     nil,
		Filter:   nil,
	}

	parser := NewRSQLParser()

	lq, err := NewListingQuery(qp, parser, &testListingAttributes{})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if lq.Page != 1 {
		t.Errorf("Unexpected page number. Expected 1, got %d", lq.Page)
	}

	if lq.Limit != 10 {
		t.Errorf("Unexpected limit number. Expected 1, got %d", lq.Limit)
	}

	if lq.Sort != nil {
		t.Errorf("Unexpected sort value. Expected nil, got %T", lq.Sort)
	}

	if lq.Filter != nil {
		t.Errorf("Unexpected filter value. Expected nil, got %T", lq.Filter)
	}
}

func TestNewListingQuery(t *testing.T) {
	page := 3
	pageSize := 25
	filter := "name==\"John Doe\""
	sort := "name;"

	qp := &QueryParams{
		Page:     &page,
		PageSize: &pageSize,
		Sort:     &sort,
		Filter:   &filter,
	}
	expectedSort := []Sort{{Field: "name", Direction: "DESC"}}
	expectedFilter := rsql.Filter{Query: "name = ?", Args: []interface{}{"John Doe"}}

	parser := NewRSQLParser()

	lq, err := NewListingQuery(qp, parser, &testListingAttributes{})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if lq.Page != 3 {
		t.Errorf("Unexpected page number. Expected 3, got %d", lq.Page)
	}

	if lq.Limit != 25 {
		t.Errorf("Unexpected limit number. Expected 25, got %d", lq.Limit)
	}

	if lq.Offset != 50 {
		t.Errorf("Unexpected offset number. Expected 50, got %d", lq.Offset)
	}

	if !reflect.DeepEqual(*lq.Filter, expectedFilter) {
		t.Errorf("Filter does not match expected.\nGot:\n%#v\nExpected:\n%#v", *lq.Filter, expectedFilter)
	}

	if !reflect.DeepEqual(*lq.Sort, expectedSort) {
		t.Errorf("Sort does not match expected.\nGot:\n%#v\nExpected:\n%#v", *lq.Sort, expectedSort)
	}
}

func TestAttributeParsing(t *testing.T) {
	type attributes struct {
		ID        string `rsql:"filter"`
		FirstName string `rsql:"filter,sort"`
		Age       string `rsql:"filter,sort"`
		IsActive  bool   `rsql:"filter"`
		UserID    string `rsql:"field:userId,filter"`
		ParentID  string `rsql:"field:parent.id,filter"`
	}

	expectedFilter := map[string]string{"id": "id", "firstName": "first_name", "age": "age", "isActive": "is_active", "userId": "userId", "parentId": "parent.id"}
	expectedSort := map[string]string{"firstName": "first_name", "age": "age"}

	filterMap, sortMap, err := parseListingAttribute(&attributes{})
	if err != nil {
		t.Errorf("Got an unexpected error %#v", err)
	}

	if !reflect.DeepEqual(filterMap, expectedFilter) {
		t.Errorf("Filter attributes do not match expected.\nGot:\n%#v\nExpected:\n%#v", filterMap, expectedFilter)
	}

	if !reflect.DeepEqual(sortMap, expectedSort) {
		t.Errorf("Sort attributes do not match expected.\nGot:\n%#v\nExpected:\n%#v", sortMap, expectedSort)
	}
}

func TestAttributeParsing2(t *testing.T) {
	type attributes struct {
		IsStarred bool   `rsql:"filter"`
		Role      string `rsql:"filter,sort"`
		JoinedAt  string `rsql:"filter,sort"`
		UserID    string `rsql:"field:users.id,filter"`
		FirstName string `rsql:"field:users.first_name,filter,sort"`
		LastName  string `rsql:"field:users.last_name,filter,sort"`
		Username  string `rsql:"field:users.username,filter,sort"`
		Name      string `rsql:"field:LOWER(users.first_name || ' ' || users.last_name),filter,sort"`
	}

	expectedFilter := map[string]string{
		"isStarred": "is_starred",
		"role":      "role",
		"joinedAt":  "joined_at",
		"userId":    "users.id",
		"firstName": "users.first_name",
		"lastName":  "users.last_name",
		"username":  "users.username",
		"name":      "LOWER(users.first_name || ' ' || users.last_name)",
	}
	expectedSort := map[string]string{
		"role":      "role",
		"joinedAt":  "joined_at",
		"firstName": "users.first_name",
		"lastName":  "users.last_name",
		"username":  "users.username",
		"name":      "LOWER(users.first_name || ' ' || users.last_name)",
	}

	filterMap, sortMap, err := parseListingAttribute(&attributes{})
	if err != nil {
		t.Errorf("Got an unexpected error %#v", err)
	}

	if !reflect.DeepEqual(filterMap, expectedFilter) {
		t.Errorf("Filter attributes do not match expected.\nGot:\n%#v\nExpected:\n%#v", filterMap, expectedFilter)
	}

	if !reflect.DeepEqual(sortMap, expectedSort) {
		t.Errorf("Sort attributes do not match expected.\nGot:\n%#v\nExpected:\n%#v", sortMap, expectedSort)
	}
}
