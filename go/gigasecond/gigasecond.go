// Package gigasecond contains the implementation of the AddGigasecond function
package gigasecond

import (
	"time"
)

const Gigasecond = 1e9 * time.Second

// AddGigasecond takes in a time and return time after adding a Gigasecond to it.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(Gigasecond)
}
