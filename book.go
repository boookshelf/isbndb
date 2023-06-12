package isbndb

import "context"

const (
	bookPath = "book"
)

type Book struct {
	Book struct {
		Title         string   `json:"title,omitempty"`
		TitleLong     string   `json:"title_long,omitempty"`
		ISBN          string   `json:"isbn,omitempty"`
		ISBN13        string   `json:"isbn_13,omitempty"`
		DeweyDecimal  string   `json:"dewey_decimal,omitempty"`
		Binding       string   `json:"binding,omitempty"`
		Publisher     string   `json:"publisher,omitempty"`
		Language      string   `json:"language,omitempty"`
		DatePublished string   `json:"date_published,omitempty"`
		Edition       string   `json:"edition,omitempty"`
		Pages         int      `json:"pages,omitempty"`
		Dimensions    string   `json:"dimensions,omitempty"`
		Overview      string   `json:"overview,omitempty"`
		Image         string   `json:"image,omitempty"`
		MSRP          float32  `json:"msrp,omitempty"`
		Excerpt       string   `json:"excerpt,omitempty"`
		Synopsis      string   `json:"synopsis,omitempty"`
		Authors       []string `json:"authors,omitempty"`
		Subjects      []string `json:"subjects,omitempty"`
		Reviews       []string `json:"reviews,omitempty"`
	} `json:"book,omitempty"`
}

func (c *Client) GetBook(ctx context.Context, isbn string) (Book, error) {
	url := c.baseURL.JoinPath(bookPath, isbn)
	var book Book
	err := c.get(ctx, url.String(), &book)

	return book, err
}
