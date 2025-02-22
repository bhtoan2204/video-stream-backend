package query

type QueryOptions struct {
	Filters map[string]interface{}
	OrderBy string // "created_at DESC", "username ASC"
	Limit   int
	Offset  int
}
