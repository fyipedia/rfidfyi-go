package rfidfyi

// SearchResult represents the API search response.
type SearchResult struct {
	Query   string       `json:"query"`
	Results []SearchItem `json:"results"`
	Total   int          `json:"total"`
}

// SearchItem represents a single search result item.
type SearchItem struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

// TagDetail represents an RFID tag.
type TagDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Family      string `json:"family"`
	Frequency   string `json:"frequency"`
	URL         string `json:"url"`
}

// ReaderDetail represents an RFID reader.
type ReaderDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// FamilyDetail represents an RFID family.
type FamilyDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// FrequencyDetail represents an RFID frequency band.
type FrequencyDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// StandardDetail represents an RFID standard.
type StandardDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// EPCDetail represents an EPC scheme.
type EPCDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// UseCaseDetail represents an RFID use case.
type UseCaseDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// GlossaryTerm represents a glossary term.
type GlossaryTerm struct {
	Term       string `json:"term"`
	Slug       string `json:"slug"`
	Definition string `json:"definition"`
	URL        string `json:"url"`
}

// CompareResult represents a comparison between two RFID tags.
type CompareResult struct {
	TagA interface{} `json:"tag_a"`
	TagB interface{} `json:"tag_b"`
	URL  string      `json:"url"`
}
