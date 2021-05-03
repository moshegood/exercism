package clock

// type Clock                      // define the clock type
// New(hour, minute int) Clock     // a "constructor"
// (Clock) String() string         // a "stringer"
// (Clock) Add(minutes int) Clock
// (Clock) Subtract(minutes int) Clock

import (
	"fmt"
)

type Clock int

func New(hour, minute int) Clock {
	return normalizedClock(60*hour + minute)
}

func (c Clock) Add(minutes int) Clock {
	return normalizedClock(int(c) + minutes)
}

func (c Clock) Subtract(minutes int) Clock {
	return c.Add(-minutes)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60%24, c%60)
}

func normalizedClock(c int) Clock {
	c %= (24 * 60)
	if c < 0 {
		c += 24 * 60
	}
	return Clock(c)
}
