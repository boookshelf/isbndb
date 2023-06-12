package isbndb

import (
	"context"
)

type Author struct {
	Author string `json:"author,omitempty"`
	Books  []Book `json:"books,omitempty"`
}

type AuthorQuery struct {
	Total   int
	Authors []string
}

func (c *Client) GetAuthor(ctx context.Context, author string) (Author, error) {
	url := c.baseURL.JoinPath(author)
	var response Author
	err := c.get(ctx, url, response)

	return response, err
}

func (c *Client) QueryAuthors(ctx context.Context, query string) (AuthorQuery, error) {
	return AuthorQuery{}, nil
}
