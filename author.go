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

// Returns the name and a list of books by the author.
// TODO Find why this returns empty
func (c *Client) GetAuthor(ctx context.Context, author string, options *PageOptions) (Author, error) {
	url := c.baseURL.JoinPath(authorPath, author)
	addPageQueryParams(url, options)

	var response Author
	err := c.get(ctx, url.String(), &response)

	return response, err
}

func (c *Client) QueryAuthors(ctx context.Context, query string, options *PageOptions) (AuthorQuery, error) {
	url := c.baseURL.JoinPath(authorQueryPath, query)
	addPageQueryParams(url, options)

	var response AuthorQuery
	err := c.get(ctx, url.String(), &response)

	return response, err
}
