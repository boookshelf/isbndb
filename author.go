package isbndb

import (
	"context"
)

const (
	authorPath      = "author"
	authorQueryPath = "authors"
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
	url := c.baseURL.JoinPath(authorPath, author)
	var response Author
	err := c.get(ctx, url.String(), &response)

	return response, err
}

func (c *Client) QueryAuthors(ctx context.Context, query string) (AuthorQuery, error) {
	url := c.baseURL.JoinPath(authorQueryPath, query)
	var response AuthorQuery
	err := c.get(ctx, url.String(), &response)

	return response, err
}
