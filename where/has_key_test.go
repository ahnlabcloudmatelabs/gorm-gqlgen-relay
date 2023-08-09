package where_test

import (
	"testing"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/where"
)

func TestHasKey(t *testing.T) {
	input := map[string]any{
		"target": "exists",
	}

	hasKey, err := where.HasKey(input, "target")
	if err != nil {
		t.Error(err)
	}

	if !hasKey {
		t.Error("HasKey should return true")
	}
}

func TestHasKey_With_And(t *testing.T) {
	input := map[string]any{
		"and": []any{
			map[string]any{
				"target": "exists",
			},
		},
	}

	hasKey, err := where.HasKey(input, "target")
	if err != nil {
		t.Error(err)
	}

	if !hasKey {
		t.Error("HasKey should return true")
	}
}

func TestHasKey_With_Or(t *testing.T) {
	input := map[string]any{
		"or": []any{
			map[string]any{
				"target": "exists",
			},
		},
	}

	hasKey, err := where.HasKey(input, "target")
	if err != nil {
		t.Error(err)
	}

	if !hasKey {
		t.Error("HasKey should return true")
	}
}

func TestHasKey_With_Not(t *testing.T) {
	input := map[string]any{
		"not": map[string]any{
			"target": "exists",
		},
	}

	hasKey, err := where.HasKey(input, "target")
	if err != nil {
		t.Error(err)
	}

	if !hasKey {
		t.Error("HasKey should return true")
	}
}
