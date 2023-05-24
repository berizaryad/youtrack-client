package youtrack

import (
	"net/http"
	"time"
)

// Option defines an option for a Client
type Option func(*Client)

type Client struct {
	url        string
	token      string
	httpClient *http.Client
}

func New(url, token string, options ...Option) Client {
	c := Client{
		url:        url,
		token:      token,
		httpClient: &http.Client{},
	}

	for _, opt := range options {
		opt(&c)
	}

	return c
}

// OptionTimeout set timeout for HTTP client.
func OptionTimeout(timeout time.Duration) func(*Client) {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}
