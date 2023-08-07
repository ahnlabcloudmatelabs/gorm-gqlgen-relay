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

func (p *PageInfo) SetHasNextPage(totalCount, edgesLen int, first, last *int, before, after *string) {
	if totalCount == 0 || totalCount == edgesLen {
		return
	}

	if before != nil {
		p.HasNextPage = true
		return
	}

	if first != nil && *first > edgesLen {
		return
	}

	if last != nil && *last > edgesLen {
		return
	}

	if after == nil && first != nil && *first < totalCount {
		p.HasNextPage = true
		return
	}

	if after == nil && last != nil && *last < totalCount {
		p.HasNextPage = true
		return
	}

}
