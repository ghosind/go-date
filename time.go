package date

import (
	"time"
)

type Time struct {
	time.Time
}

// New creates and returns a new Time. It'll initialize by the parameter, or set the time to now.
func New(t ...time.Time) Time {
	tm := Time{}

	if len(t) > 0 {
		tm.Time = t[0]
	} else {
		tm.Time = time.Now()
	}

	return tm
}

// Date creates and returns a new Time by the specific values. The location is an optional
// parameter, default time.UTC.
func Date(year int, month time.Month, day, hour, min, sec, nsec int, loc ...*time.Location) Time {
	tm := Time{}

	location := time.UTC
	if len(loc) > 0 && loc[0] != nil {
		location = loc[0]
	}

	tm.Time = time.Date(year, time.Month(month), day, hour, min, sec, nsec, location)

	return tm
}

// Now returns the current time.
func Now() Time {
	return New()
}

// Hour12 returns the 12-hours clock hour offset within the day specified by the time, in the
// range [1, 12]. See https://en.wikipedia.org/wiki/12-hour_clock for more details about the
// value.
func (t Time) Hour12() int {
	hour := t.Hour() % 12
	if hour == 0 {
		hour = 12
	}
	return hour
}

// Microsecond returns the microsecond offset within the second specified by the time, in the
// range [0, 999999].
func (t Time) Microsecond() int {
	return t.Nanosecond() / int(time.Microsecond)
}

// Millisecond returns the millisecond offset within the second specified by the time, in the
// range [0, 999].
func (t Time) Millisecond() int {
	return t.Nanosecond() / int(time.Millisecond)
}
