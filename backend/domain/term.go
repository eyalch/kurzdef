package domain

// Term represents a single term which has a definition
type Term struct {
	Title      string `json:"title"`
	Definition string `json:"definition"`
	Slug       string `json:"slug"`
	Author     string `json:"author"`
	SourceURL  string `json:"source_url"`
}

// TermUsecase is an interface which represents all the possible usecases
// regarding terms
type TermUsecase interface {
	Get() error
	Create() error
}
