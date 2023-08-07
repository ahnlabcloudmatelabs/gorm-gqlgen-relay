package relay_test

import (
	"testing"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/relay"
)

func TestSetPrevious_WithoutAfter(t *testing.T) {
	p := relay.PageInfo{}
	p.SetHasPreviousPage(10, 5, nil)

	if p.HasPreviousPage {
		t.Errorf("Expected HasPreviousPage to be false")
	}
}

func TestSetPrevious_WithAfter(t *testing.T) {
	after := "eyJpZCI6MX0="
	p := relay.PageInfo{}
	p.SetHasPreviousPage(10, 5, &after)

	if !p.HasPreviousPage {
		t.Errorf("Expected HasPreviousPage to be true")
	}
}

func TestSetPrevious_WithAfter_TotalCountEqualsEdgesLen(t *testing.T) {
	after := "eyJpZCI6MX0="
	p := relay.PageInfo{}
	p.SetHasPreviousPage(10, 10, &after)

	if p.HasPreviousPage {
		t.Errorf("Expected HasPreviousPage to be false")
	}
}

func TestSetNextPage_WithBefore(t *testing.T) {
	before := "eyJpZCI6MX0="
	p := relay.PageInfo{}
	p.SetHasNextPage(10, 5, nil, nil, &before, nil)

	if !p.HasNextPage {
		t.Errorf("Expected HasNextPage to be true")
	}
}

func TestSetNextPage_TotalCountEqualsEdgesLen(t *testing.T) {
	p := relay.PageInfo{}
	p.SetHasNextPage(10, 10, nil, nil, nil, nil)

	if p.HasNextPage {
		t.Errorf("Expected HasNextPage to be false")
	}
}

func TestSetNextPage_First_GreaterThan_EdgesLen(t *testing.T) {
	first := 10
	p := relay.PageInfo{}
	p.SetHasNextPage(10, 5, &first, nil, nil, nil)

	if p.HasNextPage {
		t.Errorf("Expected HasNextPage to be false")
	}
}

func TestSetNextPage_Last_GreaterThan_EdgesLen(t *testing.T) {
	last := 10
	p := relay.PageInfo{}
	p.SetHasNextPage(10, 5, nil, &last, nil, nil)

	if p.HasNextPage {
		t.Errorf("Expected HasNextPage to be false")
	}
}

func TestSetNextPage_After_IsNil_First_IsNil_TotalCount_GreaterThan_First(t *testing.T) {
	first := 5
	p := relay.PageInfo{}
	p.SetHasNextPage(10, 5, &first, nil, nil, nil)

	if !p.HasNextPage {
		t.Errorf("Expected HasNextPage to be true")
	}
}

func TestSetNextPage_After_IsNil_Last_IsNil_TotalCount_GreaterThan_Last(t *testing.T) {
	last := 5
	p := relay.PageInfo{}
	p.SetHasNextPage(10, 5, nil, &last, nil, nil)

	if !p.HasNextPage {
		t.Errorf("Expected HasNextPage to be true")
	}
}

func TestSetNextPage_First_Equals_EdgesLen(t *testing.T) {
	first := 5
	p := relay.PageInfo{}
	p.SetHasNextPage(10, 5, &first, nil, nil, nil)

	if !p.HasNextPage {
		t.Errorf("Expected HasNextPage to be true")
	}
}

func TestSetNextPage_First_Equals_EdgesLen_With_After(t *testing.T) {
	first := 5
	after := "eyJpZCI6MX0="
	p := relay.PageInfo{}
	p.SetHasNextPage(10, 5, &first, nil, nil, &after)

	if !p.HasNextPage {
		t.Errorf("Expected HasNextPage to be true")
	}
}

func TestSetNextPage_Last_Equals_EdgesLen(t *testing.T) {
	last := 5
	p := relay.PageInfo{}
	p.SetHasNextPage(10, 5, nil, &last, nil, nil)

	if !p.HasNextPage {
		t.Errorf("Expected HasNextPage to be true")
	}
}

func TestSetNextPage_Last_Equals_EdgesLen_With_After(t *testing.T) {
	last := 5
	after := "eyJpZCI6MX0="
	p := relay.PageInfo{}
	p.SetHasNextPage(10, 5, nil, &last, nil, &after)

	if !p.HasNextPage {
		t.Errorf("Expected HasNextPage to be true")
	}
}
