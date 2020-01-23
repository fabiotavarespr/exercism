package gigasecond

import "time"

const giga = 1e9 * time.Second

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(giga)
}
