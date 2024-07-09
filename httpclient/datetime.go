package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (c *Client) GetDateTime(ctx context.Context) (string, error) {
	currentTime := time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC).Format("Wed Feb 25 11:06:39 PST 2015")
	var JSON JSONResponse

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/datetime", nil)

	if err != nil {
		return currentTime, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return currentTime, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return currentTime, fmt.Errorf("unexpected status code")
	}

	err = json.NewDecoder(res.Body).Decode(&JSON)
	if err != nil {
		return currentTime, err
	}

	return JSON.Message, nil
}
