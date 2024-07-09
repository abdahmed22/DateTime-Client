package httpclient

import "net/http"

type Client struct {
	httpClient *http.Client
}

type JSONResponse struct {
	Message string `json:"message"`
}

func NewClient() *Client {
	client := &Client{
		httpClient: http.DefaultClient,
	}
	return client
}
