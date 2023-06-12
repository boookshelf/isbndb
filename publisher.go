package isbndb

import "context"

const (
	publisherPath = "publisher"
)

type Publisher struct {
	Name  string `json:"name,omitempty"`
	Books []struct {
		ISBN string `json:"isbn,omitempty"`
	} `json:"books,omitempty"`
}

func (c *Client) GetPublisher(ctx context.Context, publisherName string) (Publisher, error) {
	url := c.baseURL.JoinPath(publisherPath, publisherName)
	var result Publisher
	err := c.get(ctx, url.String(), &result)

	return result, err
}
