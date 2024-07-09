package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/cenkalti/backoff"
)

func (c *Client) GetHTTPDateTime(ctx context.Context) (string, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.httpURL+c.httpPort+c.endPoint, nil)

	if err != nil {
		return "", err
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 3 * time.Second

	var (
		res    *http.Response
		resErr error
	)

	retryable := func() error {
		res, resErr = c.httpClient.Do(req)
		return resErr
	}

	notify := func(err error, t time.Duration) {
		log.Printf("error: %v happened at time: %v", err, t)
	}

	err = backoff.RetryNotify(retryable, b, notify)
	if err != nil {
		log.Fatalf("error after retrying: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code")
	}

	currentTime, err := io.ReadAll(res.Body)

	defer res.Body.Close()

	if err != nil {
		return "", err
	}

	return string(currentTime), nil
}

func (c *Client) GetGinDateTime(ctx context.Context) (string, error) {
	var JSON JSONResponse

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.ginURL+c.ginPort+c.endPoint, nil)

	if err != nil {
		return "", err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code")
	}

	err = json.NewDecoder(res.Body).Decode(&JSON)
	if err != nil {
		return "", err
	}

	return JSON.Message, nil
}
