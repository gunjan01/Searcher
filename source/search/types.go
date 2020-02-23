package search

// GetLiteraryBooksRequest is the request for search.
type GetLiteraryBooksRequest struct {
	Query    string  `json:"query"`
	Author   *string `json:"author"`
	Title    *string `json:"title"`
	Location *int64  `json:"location"`
}

// BookDocument hold the book
type BookDocument struct {
	Author   string `json:"author"`
	Title    string `json:"title"`
	Location int64  `json:"location"`
	Text     string `json:"text"`
}
