package isbndb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	baseURL = "https://api2.isbndb.com"
)

type Client struct {
	baseURL *url.URL
	http    *http.Client
}

func New(httpClient *http.Client) *Client {
	url, _ := url.Parse(baseURL)

	return &Client{
		baseURL: url,
		http:    httpClient,
	}
}

func (c *Client) get(ctx context.Context, url *url.URL, result interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return err
	}

	return c.do(req, result)
}

func (c *Client) post() {

}

func (c *Client) do(req *http.Request, result interface{}) error {
	response, err := c.http.Do(req)
	if err != nil {
		return err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return fmt.Errorf("Status code: %v, Error: %v", response.StatusCode, response.Status)
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return err
	}

	return nil
}
