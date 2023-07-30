package edge_test

import (
	"testing"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/edge"
)

type Row struct {
	ID     string
	Author Author
}

type Author struct {
	ID   string
	Name string
}

var row = []*Row{
	{
		ID: "1",
		Author: Author{
			ID:   "1",
			Name: "John Doe",
		},
	},
	{
		ID: "2",
		Author: Author{
			ID:   "2",
			Name: "Jane Doe",
		},
	},
}

func TestEdge(t *testing.T) {
	edges, err := edge.Convert(row, []string{"id"}, "id")
	if err != nil {
		t.Fatal(err)
	}

	if len(edges) != 2 {
		t.Fatalf("Expected 2 edges, got %d", len(edges))
	}

	if edges[0].Cursor != "eyJpZCI6IjEifQ==" {
		t.Fatalf("Expected cursor to be 'eyJpZCI6IjEifQ==', got '%s'", edges[0].Cursor)
	}
}
