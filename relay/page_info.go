package relay

func (p *PageInfo) SetHasPreviousPage(totalCount, edgesLen int, after *string) {
	if totalCount == 0 || totalCount == edgesLen {
		return
	}

	if after != nil {
		p.HasPreviousPage = true
		return
	}
}

func (p *PageInfo) SetHasNextPage(remainingCount, edgesLen int, first, last *int, before, after *string) {
	if remainingCount == 0 || remainingCount == edgesLen {
		return
	}

	if before != nil {
		p.HasNextPage = true
		return
	}

	if first != nil && *first == edgesLen {
		p.HasNextPage = true
		return
	}

	if last != nil && *last == edgesLen {
		p.HasNextPage = true
		return
	}

	if first != nil && *first > edgesLen {
		return
	}

	if last != nil && *last > edgesLen {
		return
	}

	if after == nil && first != nil && *first < remainingCount {
		p.HasNextPage = true
		return
	}

	if after == nil && last != nil && *last < remainingCount {
		p.HasNextPage = true
		return
	}
}
