package timer

import (
	"time"
)

// Hourglass calculate left time
type Hourglass struct {
	endTime time.Time
}

// NewHourglass create hourglass
func NewHourglass(nanoseconds time.Duration) *Hourglass {
	t := new(Hourglass)
	t.endTime = time.Now().Add(nanoseconds)
	return t
}

// LeftNanoseconds get left time for nanoseconds
func (hourglass *Hourglass) LeftNanoseconds() time.Duration {
	return time.Duration(hourglass.endTime.Sub(time.Now()).Nanoseconds())
}
