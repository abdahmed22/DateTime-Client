package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetHTTPDateTime(ctx context.Context) (string, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.httpURL+c.httpPort+c.endPoint, nil)

	if err != nil {
		return "", err
	}

	// b := backoff.NewExponentialBackOff()
	// b.MaxElapsedTime = 3 * time.Second

	// var (
	// 	res    *http.Response
	// 	resErr error
	// )

	// retryable := func() error {
	// 	res, resErr = c.httpClient.Do(req)
	// 	return resErr
	// }

	// defer res.Body.Close()

	// err = backoff.Retry(retryable, b)
	res, err := c.httpClient.Do(req)

	defer res.Body.Close()

	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code")
	}

	currentTime, err := io.ReadAll(res.Body)

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
