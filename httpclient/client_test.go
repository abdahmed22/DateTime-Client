package httpclient

import (
	"context"
	"slices"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClientCanHitAPI(t *testing.T) {

	now := time.Now()

	firstPossibleValue := time.Date(now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second()-1,
		now.Nanosecond(), now.Location()).Format("2006-01-02 15:04")

	secondPossibleValue := time.Date(now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second()-2,
		now.Nanosecond(), now.Location()).Format("2006-01-02 15:04")

	thirdPossibleValue := time.Date(now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second()-3,
		now.Nanosecond(), now.Location()).Format("2006-01-02 15:04")

	possibleValues := []string{firstPossibleValue, secondPossibleValue, thirdPossibleValue}
	t.Run("happy path - can hit the api and return date & time", func(*testing.T) {

		myClient := NewClient()
		returnedDateTime, err := myClient.GetDateTime(context.Background())

		assert.NoError(t, err)

		if !slices.Contains(possibleValues, returnedDateTime) {
			t.Fail()
		}
	})
}
