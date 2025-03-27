package query

type PaginateResult struct {
	TotalDocs   int  `json:"total_docs"`
	TotalPages  int  `json:"total_pages"`
	Page        int  `json:"page"`
	Limit       int  `json:"limit"`
	NextPage    int  `json:"next_page"`
	PrevPage    int  `json:"prev_page"`
	HasNextPage bool `json:"has_next_page"`
	HasPrevPage bool `json:"has_prev_page"`
}
