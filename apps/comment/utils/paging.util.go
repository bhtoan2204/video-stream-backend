package utils

// func BuildPaginateResult(totalDocs int, page int, limit int) *query.PaginateResult {
// 	totalPages := totalDocs / limit
// 	if totalDocs%limit != 0 {
// 		totalPages++
// 	}

// 	var nextPage, prevPage int
// 	hasNextPage := page < totalPages
// 	hasPrevPage := page > 1

// 	if hasNextPage {
// 		nextPage = page + 1
// 	}
// 	if hasPrevPage {
// 		prevPage = page - 1
// 	}

// 	return &query.PaginateResult{
// 		TotalDocs:   totalDocs,
// 		TotalPages:  totalPages,
// 		Page:        page,
// 		Limit:       limit,
// 		NextPage:    nextPage,
// 		PrevPage:    prevPage,
// 		HasNextPage: hasNextPage,
// 		HasPrevPage: hasPrevPage,
// 	}
// }
