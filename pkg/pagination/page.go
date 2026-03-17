package pagination

const DefaultPageSize = 50
const MaxPageSize = 1000
const DefaultPageStr = "1"
const DefaultPageSizeStr = "50"

type Page[T any] struct {
	TotalItems int64 `json:"total_items"`
	TotalPages int   `json:"total_pages"`
	PerPage    int   `json:"per_page"`
	Page       int   `json:"page"`
	Data       []T   `json:"data"`
}
