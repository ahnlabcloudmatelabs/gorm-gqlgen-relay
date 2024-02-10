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
var tables = &map[string]string{
	"field1": "sample",
}

func TestOrder(t *testing.T) {
	query, err := order.By("sample", nil, input, false)
	if err != nil {
		t.Fatal(err)
	}

	if len(query) != 2 {
		t.Fatal("query length should be 2")
	}

	if !slices.Contains(query, "\"sample\".\"field1\" ASC") {
		t.Fatal("query should contain \"sample\".\"field1\" ASC", "\nactual:", query)
	}

	if !slices.Contains(query, "\"sample\".\"field2\" DESC") {
		t.Fatal("query should contain \"sample\".\"field2\" DESC", "\nactual:", query)
	}
}

func TestReverseOrder(t *testing.T) {
	query, err := order.By("", nil, input, true)
	if err != nil {
		t.Fatal(err)
	}

	if len(query) != 2 {
		t.Fatal("query length should be 2")
	}

	if !slices.Contains(query, "\"field1\" DESC") {
		t.Fatal("query should contain \"field1\" DESC", "\nactual:", query)
	}

	if !slices.Contains(query, "\"field2\" ASC") {
		t.Fatal("query should contain \"field2\" ASC", "\nactual:", query)
	}
}

func TestOrder_With_Tables(t *testing.T) {
	query, err := order.By("", tables, input, false)
	if err != nil {
		t.Fatal(err)
	}

	if len(query) != 2 {
		t.Fatal("query length should be 2")
	}

	if !slices.Contains(query, "\"sample\".\"field1\" ASC") {
		t.Fatal("query should contain \"sample\".\"field1\" ASC")
	}

	if !slices.Contains(query, "\"field2\" DESC") {
		t.Fatal("query should contain \"field2\" DESC")
	}
}
