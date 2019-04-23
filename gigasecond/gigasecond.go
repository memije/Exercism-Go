package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond calculates the moment when someone has lived for 10^9 seconds
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1000000000 * time.Second)
}
