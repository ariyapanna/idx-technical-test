package filter 

type TodoFilter struct {
	Page       int
	Limit      int
	Search     string
	SortBy     string
	SortOrder  string
}