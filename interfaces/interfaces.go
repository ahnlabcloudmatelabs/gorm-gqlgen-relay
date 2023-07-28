package interfaces

type Connection[Model any] struct {
	TotalCount int            `json:"totalCount"`
	Edges      []*Edge[Model] `json:"edges"`
	PageInfo   *PageInfo      `json:"pageInfo"`
}

type Edge[Model any] struct {
	Node   *Model `json:"node"`
	Cursor string `json:"cursor"`
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor,omitempty"`
	EndCursor       *string `json:"endCursor,omitempty"`
}
