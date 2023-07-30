package cursor_test

import (
	"testing"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/cursor"
)

func TestAfter(t *testing.T) {
	after := "eyJpZCI6IjEifQ=="
	queries, args, err := cursor.After(&after, nil, "id")
	if err != nil {
		t.Error(err)
	}

	if len(queries) != 1 {
		t.Error("Expected 1 query")
	}

	if len(args) != 1 {
		t.Error("Expected 1 argument")
	}

	if queries[0] != "id > ?" {
		t.Error("Expected id > ?, got", queries[0])
	}

	if args[0] != "1" {
		t.Error("Expected 1, got", args[0])
	}
}

func TestBefore(t *testing.T) {
	before := "eyJpZCI6IjEifQ=="
	queries, args, err := cursor.Before(&before, map[string]interface{}{"id": "DESC"}, "id")
	if err != nil {
		t.Error(err)
	}

	if len(queries) != 1 {
		t.Error("Expected 1 query")
	}

	if len(args) != 1 {
		t.Error("Expected 1 argument")
	}

	if queries[0] != "id > ?" {
		t.Error("Expected id > ?, got", queries[0])
	}

	if args[0] != "1" {
		t.Error("Expected 1, got", args[0])
	}
}
