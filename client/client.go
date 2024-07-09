package httpclient

import (
	"net/http"
)

type Client struct {
	httpURL    string
	ginURL     string
	httpPort   string
	ginPort    string
	endPoint   string
	httpClient *http.Client
}

type Option func(c *Client)

type JSONResponse struct {
	Message string `json:"message"`
}

func NewClient(options ...Option) *Client {
	client := &Client{
		httpClient: http.DefaultClient,
		httpURL:    "http://localhost:",
		ginURL:     "http://localhost:",
		httpPort:   "8090",
		ginPort:    "8080",
		endPoint:   "/datetime",
	}

	for _, option := range options {
		option(client)
	}
	return client
}

func CustomURL(httpURL string, ginURL string) Option {
	return func(c *Client) {
		c.ginURL = ginURL
		c.httpURL = httpURL
	}
}

func CustomPort(httpPort string, ginPort string) Option {
	return func(c *Client) {
		c.httpPort = httpPort
		c.ginPort = ginPort
	}
}

func CustomEndPoint(endPoint string) Option {
	return func(c *Client) {
		c.endPoint = endPoint
	}
}

func CustomClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}
