package order_test

import (
	"testing"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/order"
	"golang.org/x/exp/slices"
)

var input = map[string]any{
	"field1": "ASC",
	"field2": "DESC",
}

func TestOrder(t *testing.T) {
	query, err := order.By(input, false)
	if err != nil {
		t.Fatal(err)
	}

	if len(query) != 2 {
		t.Fatal("query length should be 2")
	}

	if !slices.Contains(query, "field1 ASC") {
		t.Fatal("query should contain field1 ASC")
	}

	if !slices.Contains(query, "field2 DESC") {
		t.Fatal("query should contain field2 DESC")
	}
}

func TestReverseOrder(t *testing.T) {
	query, err := order.By(input, true)
	if err != nil {
		t.Fatal(err)
	}

	if len(query) != 2 {
		t.Fatal("query length should be 2")
	}

	if !slices.Contains(query, "field1 DESC") {
		t.Fatal("query should contain field1 DESC")
	}

	if !slices.Contains(query, "field2 ASC") {
		t.Fatal("query should contain field2 ASC")
	}
}
