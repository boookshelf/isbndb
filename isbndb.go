package isbndb

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

const (
	baseURL    = "https://api2.isbndb.com"
	premiumURL = "https://api.premium.isbndb.com"
	proURL     = "https://api.pro.isbndb.com"
)

type Client struct {
	baseURL *url.URL
	http    *http.Client
	api_key string
}

func New(httpClient *http.Client) *Client {
	url, _ := url.Parse(baseURL)

	return &Client{
		baseURL: url,
		http:    httpClient,
		api_key: os.Getenv("ISBNDB_API_KEY"),
	}
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

	if response.StatusCode < http.StatusOK || response.StatusCode > 299 {
		return fmt.Errorf("status code: %v, error: %v", response.StatusCode, response.Status)
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return err
	}

	return nil
}
