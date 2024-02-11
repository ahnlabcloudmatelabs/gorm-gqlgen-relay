package where_test

import (
	"testing"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/utils"
	"github.com/cloudmatelabs/gorm-gqlgen-relay/where"
)

var query = map[string]any{
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

func TestWhereWithTable(t *testing.T) {
	filter, err := where.Do("postgres", "test", nil, nil, query)
	if err != nil {
		t.Error(err)
	}

	if filter.And[0].Query != `"test"."title" = ?` {
		t.Errorf("query is not correct: '%s'", filter.And[0].Query)
	}

	if filter.And[0].Args[0] != "Hello World" {
		t.Errorf("args is not correct: '%s'", filter.And[0].Args[0])
	}

	if filter.Or[0].Query != `"test"."title" = ?` {
		t.Errorf("query is not correct: '%s'", filter.Or[0].Query)
	}

	if filter.Or[0].Args[0] != "Hello World" {
		t.Errorf("args is not correct: '%s'", filter.Or[0].Args[0])
	}

	if filter.Or[0].Not.Query != `"test"."title" = ?` {
		t.Errorf("query is not correct: '%s'", filter.Or[0].Not.Query)
	}

	if filter.Or[0].Not.Args[0] != "Hello World" {
		t.Errorf("args is not correct: '%s'", filter.Or[0].Not.Args[0])
	}

	if !(filter.Query == `"test"."title" = ? AND "test"."title" IS NULL` ||
		filter.Query == `"test"."title" IS NULL AND "test"."title" = ?`) {
		t.Errorf("query is not correct: '%s'", filter.Query)
	}

	if filter.Args[0] != "Hello World" {
		t.Errorf("args is not correct: '%s'", filter.Args[0])
	}
}

func TestWhereWithTables(t *testing.T) {
	filter, err := where.Do("postgres", "", &map[string]string{"title": "sample"}, nil, query)
	if err != nil {
		t.Error(err)
	}

	if filter.And[0].Query != `"sample"."title" = ?` {
		t.Errorf("query is not correct: '%s'", filter.And[0].Query)
	}

	if filter.And[0].Args[0] != "Hello World" {
		t.Errorf("args is not correct: '%s'", filter.And[0].Args[0])
	}

	if filter.Or[0].Query != `"sample"."title" = ?` {
		t.Errorf("query is not correct: '%s'", filter.Or[0].Query)
	}

	if filter.Or[0].Args[0] != "Hello World" {
		t.Errorf("args is not correct: '%s'", filter.Or[0].Args[0])
	}

	if filter.Or[0].Not.Query != `"sample"."title" = ?` {
		t.Errorf("query is not correct: '%s'", filter.Or[0].Not.Query)
	}

	if filter.Or[0].Not.Args[0] != "Hello World" {
		t.Errorf("args is not correct: '%s'", filter.Or[0].Not.Args[0])
	}

	if !(filter.Query == `"sample"."title" = ? AND "sample"."title" IS NULL` ||
		filter.Query == `"sample"."title" IS NULL AND "sample"."title" = ?`) {
		t.Errorf("query is not correct: '%s'", filter.Query)
	}

	if filter.Args[0] != "Hello World" {
		t.Errorf("args is not correct: '%s'", filter.Args[0])
	}
}

func TestWhereWithTablesWhenNoMatchesColumns(t *testing.T) {
	filter, err := where.Do("postgres", "", &map[string]string{"created_at": "sample"}, nil, query)
	if err != nil {
		t.Error(err)
	}

	if filter.And[0].Query != `"title" = ?` {
		t.Errorf("query is not correct: '%s'", filter.And[0].Query)
	}

	if filter.And[0].Args[0] != "Hello World" {
		t.Errorf("args is not correct: '%s'", filter.And[0].Args[0])
	}

	if filter.Or[0].Query != `"title" = ?` {
		t.Errorf("query is not correct: '%s'", filter.Or[0].Query)
	}

	if filter.Or[0].Args[0] != "Hello World" {
		t.Errorf("args is not correct: '%s'", filter.Or[0].Args[0])
	}

	if filter.Or[0].Not.Query != `"title" = ?` {
		t.Errorf("query is not correct: '%s'", filter.Or[0].Not.Query)
	}

	if filter.Or[0].Not.Args[0] != "Hello World" {
		t.Errorf("args is not correct: '%s'", filter.Or[0].Not.Args[0])
	}

	if !(filter.Query == `"title" = ? AND "title" IS NULL` ||
		filter.Query == `"title" IS NULL AND "title" = ?`) {
		t.Errorf("query is not correct: '%s'", filter.Query)
	}

	if filter.Args[0] != "Hello World" {
		t.Errorf("args is not correct: '%s'", filter.Args[0])
	}
}

func TestWhereWithTableAndPrefix(t *testing.T) {
	filter, err := where.Do("postgres", "test", nil, utils.ToPointer("dev"), query)
	if err != nil {
		t.Error(err)
	}

	if filter.And[0].Query != `"dev"."test"."title" = ?` {
		t.Errorf("query is not correct: '%s'", filter.And[0].Query)
	}

	if filter.And[0].Args[0] != "Hello World" {
		t.Errorf("args is not correct: '%s'", filter.And[0].Args[0])
	}

	if filter.Or[0].Query != `"dev"."test"."title" = ?` {
		t.Errorf("query is not correct: '%s'", filter.Or[0].Query)
	}

	if filter.Or[0].Args[0] != "Hello World" {
		t.Errorf("args is not correct: '%s'", filter.Or[0].Args[0])
	}

	if filter.Or[0].Not.Query != `"dev"."test"."title" = ?` {
		t.Errorf("query is not correct: '%s'", filter.Or[0].Not.Query)
	}

	if filter.Or[0].Not.Args[0] != "Hello World" {
		t.Errorf("args is not correct: '%s'", filter.Or[0].Not.Args[0])
	}

	if !(filter.Query == `"dev"."test"."title" = ? AND "dev"."test"."title" IS NULL` ||
		filter.Query == `"dev"."test"."title" IS NULL AND "dev"."test"."title" = ?`) {
		t.Errorf("query is not correct: '%s'", filter.Query)
	}

	if filter.Args[0] != "Hello World" {
		t.Errorf("args is not correct: '%s'", filter.Args[0])
	}
}
