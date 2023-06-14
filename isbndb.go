package isbndb

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

// TODO Look into pagination

const (
	BaseURL    = "https://api2.isbndb.com"
	PremiumURL = "https://api.premium.isbndb.com"
	ProURL     = "https://api.pro.isbndb.com"
)

type Client struct {
	baseURL *url.URL
	http    *http.Client
	api_key string
}

type ClientOptions func(*Client)

type PageOptions struct {
	// The number of page to retrieve, please note the API will not return more than 10,000 results no matter how you paginate them
	Page int
	// How many items should be returned per page, maximum of 1,000
	PageSize int
}

type StatusCodeError struct {
	StatusCode int
}

func (s StatusCodeError) Error() string {
	return fmt.Sprintf("Unexpected status code: %d", s.StatusCode)
}

func New(opts ...ClientOptions) *Client {
	url, _ := url.Parse(BaseURL)

	client := &Client{
		baseURL: url,
		http:    http.DefaultClient,
		api_key: os.Getenv("ISBNDB_API_KEY"),
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

func WithHttpClient(http *http.Client) ClientOptions {
	return func(c *Client) {
		c.http = http
	}
}

func WithAPIKey(apiKey string) ClientOptions {
	return func(c *Client) {
		c.api_key = apiKey
	}
}

func WithURL(customURL *url.URL) ClientOptions {
	return func(c *Client) {
		c.baseURL = customURL
	}
}

func (c *Client) SetURL(_url string) {
	newURL, _ := url.Parse(_url)
	c.baseURL = newURL
}

func (c *Client) get(ctx context.Context, url string, result interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	return c.do(req, result)
}

func (c *Client) post(ctx context.Context, url string, body map[string]interface{}, result interface{}) error {
	marshalledBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(marshalledBody))
	if err != nil {
		return err
	}

	return c.do(req, result)
}

func (c *Client) do(req *http.Request, result interface{}) error {
	req.Header.Add("Authorization", c.api_key)
	req.Header.Add("Content-Type", "application/json")

	response, err := c.http.Do(req)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return StatusCodeError{
			StatusCode: response.StatusCode,
		}
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return err
	}

	return nil
}

func addPageQueryParams(url *url.URL, options *PageOptions) {
	if options == nil {
		return
	}

	queryParams := url.Query()
	queryParams.Add("page", strconv.Itoa(options.Page))
	queryParams.Add("pageSize", strconv.Itoa(options.PageSize))
	url.RawQuery = queryParams.Encode()
}
