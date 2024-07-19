// Package httpclient provides a testing enviroment to test datetime server.
package httpclient

import (
	"net/http"
	"os"
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

// NewClient creates a new client
func NewClient(options ...Option) *Client {

	client := &Client{
		httpClient: http.DefaultClient,
		httpURL:    os.Getenv("HTTPURL"),
		ginURL:     os.Getenv("GINURL"),
		httpPort:   os.Getenv("HTTPPORT"),
		ginPort:    os.Getenv("GINPORT"),
		endPoint:   os.Getenv("ENDPOINT"),
	}

	for _, option := range options {
		option(client)
	}
	return client
}

// CustomURL provides the option to change default URLs for gin and http servers
func CustomURL(httpURL string, ginURL string) Option {
	return func(c *Client) {
		c.ginURL = ginURL
		c.httpURL = httpURL
	}
}

// CustomPort provides the option to change default Port numbers for gin and http servers
func CustomPort(httpPort string, ginPort string) Option {
	return func(c *Client) {
		c.httpPort = httpPort
		c.ginPort = ginPort
	}
}

// CustomEndPoint provides the option to change default endpoint
func CustomEndPoint(endPoint string) Option {
	return func(c *Client) {
		c.endPoint = endPoint
	}
}

// CustomClient provides the option to change default Client
func CustomClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}
