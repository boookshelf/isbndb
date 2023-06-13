package isbndb

import "context"

const (
	publisherPath      = "publisher"
	publisherQueryPath = "publishers"
)

type Publisher struct {
	Name  string `json:"name,omitempty"`
	Books []struct {
		ISBN string `json:"isbn,omitempty"`
	} `json:"books,omitempty"`
}

// Returns details and a list of books by the publisher.
func (c *Client) GetPublisher(ctx context.Context, publisherName string) (Publisher, error) {
	url := c.baseURL.JoinPath(publisherPath, publisherName).String()
	var result Publisher
	err := c.get(ctx, url, &result)

	return result, err
}

// This returns a list of publishers that match the given query
func (c *Client) QueryPublishers(ctx context.Context, query string) {

}
