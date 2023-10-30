package entity

type ResponsePagination[P, D any] struct {
	Pagination P `json:"pagination,omitempty"`
	Data       D `json:"data,omitempty"`
}

type Pagination struct {
	Total       uint64 `json:"total,omitempty"`
	Limit       uint64 `json:"limit,omitempty"`
	CurrentPage uint64 `json:"current_page,omitempty"`
	TotalPages  uint64 `json:"total_pages,omitempty"`
}
