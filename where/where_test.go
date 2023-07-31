package where_test

import (
	"testing"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/where"
)

func TestWhere(t *testing.T) {
	query := map[string]any{
		"and": []map[string]any{
			{
				"title": map[string]any{
					"equal": "Hello World",
				},
			},
		},
		"or": []map[string]any{
			{
				"title": map[string]any{
					"equal": "Hello World",
				},
				"not": map[string]any{
					"title": map[string]any{
						"equal": "Hello World",
					},
				},
			},
		},
		"not": map[string]any{
			"title": map[string]any{
				"equal": "Hello World",
			},
		},
		"title": map[string]any{
			"equal":  "Hello World",
			"isNull": true,
		},
	}

	filter, err := where.Do("mysql", query)
	if err != nil {
		t.Error(err)
	}

	if filter.And[0].Query != "title = ?" {
		t.Errorf("query is not correct: '%s'", filter.And[0].Query)
	}

	if filter.And[0].Args[0] != "Hello World" {
		t.Errorf("args is not correct: '%s'", filter.And[0].Args[0])
	}

	if filter.Or[0].Query != "title = ?" {
		t.Errorf("query is not correct: '%s'", filter.Or[0].Query)
	}

	if filter.Or[0].Args[0] != "Hello World" {
		t.Errorf("args is not correct: '%s'", filter.Or[0].Args[0])
	}

	if filter.Or[0].Not.Query != "title = ?" {
		t.Errorf("query is not correct: '%s'", filter.Or[0].Not.Query)
	}

	if filter.Or[0].Not.Args[0] != "Hello World" {
		t.Errorf("args is not correct: '%s'", filter.Or[0].Not.Args[0])
	}

	if !(filter.Query == "title = ? AND title IS NULL" ||
		filter.Query == "title IS NULL AND title = ?") {
		t.Errorf("query is not correct: '%s'", filter.Query)
	}

	if filter.Args[0] != "Hello World" {
		t.Errorf("args is not correct: '%s'", filter.Args[0])
	}
}
