// Package httpclient provides a testing enviroment to test datetime server.
package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/cenkalti/backoff"
)

// GetHTTPDateTime mimics a user performing a request to a certain endpoint to the http server
func (c *Client) GetHTTPDateTime(ctx context.Context) (string, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.httpURL+c.httpPort+c.endPoint, nil)

	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	res, err := c.SendRequest(req, 5)

	if err != nil {
		return "", fmt.Errorf("faild to send request: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code")
	}

	currentTime, err := io.ReadAll(res.Body)

	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	return string(currentTime), nil
}

// GetGinDateTime mimics a user performing a request to a certain endpoint to the gin server
func (c *Client) GetGinDateTime(ctx context.Context) (string, error) {
	var JSON JSONResponse

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.ginURL+c.ginPort+c.endPoint, nil)

	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	res, err := c.SendRequest(req, 5)

	if err != nil {
		return "", fmt.Errorf("faild to send request: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code")
	}

	err = json.NewDecoder(res.Body).Decode(&JSON)
	if err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return JSON.Message, nil
}

// SendRequest keeps sending request till the server responds for n seconds
func (c *Client) SendRequest(req *http.Request, n int) (*http.Response, error) {
	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = time.Duration(n) * time.Second

	var (
		res    *http.Response
		resErr error
	)

	retryable := func() error {
		res, resErr = c.httpClient.Do(req)
		if resErr != nil {
			return fmt.Errorf("error after retrying: %w", resErr)
		}
		return nil
	}

	notify := func(err error, t time.Duration) {

	}

	err := backoff.RetryNotify(retryable, b, notify)

	if err != nil {
		return res, err
	}

	return res, nil
}
