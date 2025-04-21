// Package gigasecond provides functionality to calculate dates that are exactly one gigasecond
// (10^9 seconds) after a given time.
package gigasecond

import (
	"time"
)

// AddGigasecond takes a date and returns a new date after one gigasecond has passed.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
