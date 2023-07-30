package cursor_test

import (
	"testing"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/cursor"
)

type Row struct {
	ID     string
	Author Author
}

type Author struct {
	ID   string
	Name string
}

var row = &Row{
	ID: "1",
	Author: Author{
		ID:   "1",
		Name: "John Doe",
	},
}

func TestCursor(t *testing.T) {
	cur, err := cursor.Create(row, nil, "id")
	if err != nil {
		t.Fatal(err)
	}

	if cur != "eyJpZCI6IjEifQ==" {
		t.Fatalf("Expected cursor to be 'eyJpZCI6IjEifQ==', got '%s'", cur)
	}
}
