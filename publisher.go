package isbndb

type Publisher struct {
	Name  string `json:"name,omitempty"`
	Books []struct {
		ISBN string `json:"isbn,omitempty"`
	} `json:"books,omitempty"`
}
