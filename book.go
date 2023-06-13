package isbndb

import "context"

const (
	bookPath      = "book"
	bookQueryPath = "books"
)

type Book struct {
	Title         string      `json:"title,omitempty"`
	TitleLong     string      `json:"title_long,omitempty"`
	ISBN          string      `json:"isbn,omitempty"`
	ISBN13        string      `json:"isbn_13,omitempty"`
	DeweyDecimal  string      `json:"dewey_decimal,omitempty"`
	Binding       string      `json:"binding,omitempty"`
	Publisher     string      `json:"publisher,omitempty"`
	Language      string      `json:"language,omitempty"`
	DatePublished interface{} `json:"date_published,omitempty"` // TODO fix this typing
	Edition       string      `json:"edition,omitempty"`
	Pages         int         `json:"pages,omitempty"`
	Dimensions    string      `json:"dimensions,omitempty"`
	Overview      string      `json:"overview,omitempty"`
	Image         string      `json:"image,omitempty"`
	MSRP          float32     `json:"msrp,omitempty"`
	Excerpt       string      `json:"excerpt,omitempty"`
	Synopsis      string      `json:"synopsis,omitempty"`
	Authors       []string    `json:"authors,omitempty"`
	Subjects      []string    `json:"subjects,omitempty"`
	Reviews       []string    `json:"reviews,omitempty"`
}

// Returns the book details
func (c *Client) GetBook(ctx context.Context, isbn string) (Book, error) {
	url := c.baseURL.JoinPath(bookPath, isbn).String()
	var result struct {
		Book Book `json:"book,omitempty"`
	}
	err := c.get(ctx, url, &result)

	return result.Book, err
}
