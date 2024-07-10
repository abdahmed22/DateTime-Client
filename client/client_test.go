package httpclient

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClientCanHitDateTimeServer(t *testing.T) {

	now := time.Now()

	firstPossibleValue := time.Date(now.Year(), now.Month(), now.Day(),
		now.Hour()-3, now.Minute(), now.Second()-1,
		now.Nanosecond(), now.Location()).Format("2006-01-02 15:04")

	secondPossibleValue := time.Date(now.Year(), now.Month(), now.Day(),
		now.Hour()-3, now.Minute(), now.Second()-2,
		now.Nanosecond(), now.Location()).Format("2006-01-02 15:04")

	thirdPossibleValue := time.Date(now.Year(), now.Month(), now.Day(),
		now.Hour()-3, now.Minute(), now.Second()-3,
		now.Nanosecond(), now.Location()).Format("2006-01-02 15:04")

	possibleValues := []string{firstPossibleValue, secondPossibleValue, thirdPossibleValue}
	t.Run("can hit the httpserver and return date & time", func(*testing.T) {

		myClient := NewClient()
		returnedDateTime, err := myClient.GetHTTPDateTime(context.Background())

		assert.NoError(t, err)

		if !slices.Contains(possibleValues, returnedDateTime) {
			t.Fail()
		}
	})

	t.Run("can hit the ginserver and return date & time", func(*testing.T) {

		myClient := NewClient()
		returnedDateTime, err := myClient.GetGinDateTime(context.Background())

		assert.NoError(t, err)

		if !slices.Contains(possibleValues, returnedDateTime) {
			t.Fail()
		}
	})
}

func TestOptionFunctions(t *testing.T) {
	t.Run("happy path - can add custom URLS using option function", func(*testing.T) {

		myClient := NewClient(
			CustomURL("http_url", "gin_url"),
		)

		assert.Equal(t, "http_url", myClient.httpURL)
		assert.Equal(t, "gin_url", myClient.ginURL)
	})

	t.Run("happy path - can add custom port numbers using option function", func(*testing.T) {

		myClient := NewClient(
			CustomPort("http_port", "gin_port"),
		)

		assert.Equal(t, "http_port", myClient.httpPort)
		assert.Equal(t, "gin_port", myClient.ginPort)
	})

	t.Run("happy path - can add custom client using option function", func(*testing.T) {

		myClient := NewClient(
			CustomEndPoint("/endpoint"),
		)

		assert.Equal(t, "/endpoint", myClient.endPoint)
	})

	t.Run("happy path - can add custom client using option function", func(*testing.T) {

		myClient := NewClient(
			CustomClient(&http.Client{
				Timeout: 5 * time.Second,
			}),
		)

		assert.Equal(t, 5*time.Second, myClient.httpClient.Timeout)
	})

}

func TestClientCanHitHTTPMockServer(t *testing.T) {
	t.Run("can hit the mockserver and return date & time", func(*testing.T) {

		mockServer := httptest.NewServer(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintf(w, "2006-01-02 15:04")
				}),
		)

		defer mockServer.Close()

		myClient := NewClient(
			CustomURL(mockServer.URL, mockServer.URL),
			CustomPort("", ""),
		)

		returnedDateTime, err := myClient.GetHTTPDateTime(context.Background())

		assert.NoError(t, err)

		assert.Equal(t, "2006-01-02 15:04", returnedDateTime)

	})

	t.Run("can handle 500 status code", func(*testing.T) {

		mockServer := httptest.NewServer(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
				}),
		)

		defer mockServer.Close()

		myClient := NewClient(
			CustomURL(mockServer.URL, mockServer.URL),
			CustomPort("", ""),
		)

		_, err := myClient.GetHTTPDateTime(context.Background())

		assert.Error(t, err)

	})
}

func TestClientCanHitGinMockServer(t *testing.T) {
	t.Run("can hit the mockserver and return date & time", func(*testing.T) {

		mockServer := httptest.NewServer(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintf(w, `{"message": "2006-01-02 15:04"}`)
				}),
		)

		defer mockServer.Close()

		myClient := NewClient(
			CustomURL(mockServer.URL, mockServer.URL),
			CustomPort("", ""),
		)

		returnedDateTime, err := myClient.GetGinDateTime(context.Background())

		assert.NoError(t, err)

		assert.Equal(t, "2006-01-02 15:04", returnedDateTime)

	})

	t.Run("can handle 500 status code", func(*testing.T) {

		mockServer := httptest.NewServer(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
				}),
		)

		defer mockServer.Close()

		myClient := NewClient(
			CustomURL(mockServer.URL, mockServer.URL),
			CustomPort("", ""),
		)

		_, err := myClient.GetGinDateTime(context.Background())

		assert.Error(t, err)

	})

	t.Run("can handle wrong json format", func(*testing.T) {

		mockServer := httptest.NewServer(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintf(w, `{"message": "2006-01-02 15:04"`)
				}),
		)

		defer mockServer.Close()

		myClient := NewClient(
			CustomURL(mockServer.URL, mockServer.URL),
			CustomPort("", ""),
		)

		_, err := myClient.GetGinDateTime(context.Background())

		assert.Error(t, err)

	})
}
