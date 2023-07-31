package relay

type Connection[Model any] struct {
	TotalCount int64          `json:"totalCount"`
	Edges      []*Edge[Model] `json:"edges"`
	PageInfo   *PageInfo      `json:"pageInfo"`
}

type Edge[T any] struct {
	Cursor string `json:"cursor"`
	Node   *T     `json:"node"`
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor,omitempty"`
	EndCursor       *string `json:"endCursor,omitempty"`
}
