package isbndb

type Subject struct {
	Subject string `json:"subject,omitempty"`
	Parent  string `json:"parent,omitempty"`
}
