package common

type Pagination struct {
	TotalRows   int `json:"total_rows"`
	CurrentPage int `json:"current_page"`
	TotalPage   int `json:"total_page"`
	PerPage     int `json:"per_page"`
}
