// Package gigasecond returns a time.Time that's one gigasecond in the future.
package gigasecond

import "time"

const testVersion = 4 // find the value in gigasecond_test.go
// A gigasecond is 10^9 (1,000,000,000) seconds.
const gigasecond = time.Second * 1000000000

// AddGigasecond adds a gigasecond to a time.Time.
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(gigasecond)
	return t
}
