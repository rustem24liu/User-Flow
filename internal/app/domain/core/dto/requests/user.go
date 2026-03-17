package requests

type GetDTO struct {
	Keyword string `json:"keyword"`
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
}
