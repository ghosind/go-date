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

// Unix returns the local Time corresponding to the given Unix time, sec seconds and nsec
// nanoseconds since January 1, 1970 UTC. It is valid to pass nsec outside the range
// [0, 999999999]. Not all sec values have a corresponding time value. One such value is 1<<63-1
// (the largest int64 value).
func Unix(sec, nsec int64) Time {
	tm := time.Unix(sec, nsec)
	return New(tm)
}

// UnixMicro returns the local Time corresponding to the given Unix time, usec microseconds since
// January 1, 1970 UTC.
func UnixMicro(usec int64) Time {
	tm := time.UnixMicro(usec)
	return New(tm)
}

// UnixMilli returns the local Time corresponding to the given Unix time, msec milliseconds since
// January 1, 1970 UTC.
func UnixMilli(msec int64) Time {
	tm := time.UnixMilli(msec)
	return New(tm)
}

// Since returns the time elapsed since t. It is shorthand for time.Now().Sub(t).
func Since(t any) time.Duration {
	tm := getTime(t)
	return time.Since(tm)
}

// Until returns the duration until t. It is shorthand for t.Sub(time.Now()).
func Until(t any) time.Duration {
	tm := getTime(t)
	return time.Until(tm)
}

// Add returns the time t+d.
func (t Time) Add(d time.Duration) Time {
	t.Time = t.Time.Add(d)
	return t
}

// Sub returns the duration t-u. If the result exceeds the maximum (or minimum) value that can be
// stored in a Duration, the maximum (or minimum) duration will be returned. To compute t-d for a
// duration d, use t.Add(-d).
func (t Time) Sub(u any) time.Duration {
	tm := getTime(u)
	return t.Time.Sub(tm)
}

// Round returns the result of rounding t to the nearest multiple of d (since the zero time). The
// rounding behavior for halfway values is to round up. If d <= 0, Round returns t stripped of any
// monotonic clock reading but otherwise unchanged.
//
// Round operates on the time as an absolute duration since the zero time; it does not operate on
// the presentation form of the time. Thus, Round(Hour) may return a time with a non-zero minute,
// depending on the time's Location.
func (t Time) Round(d time.Duration) Time {
	t.Time = t.Time.Round(d)
	return t
}

// Truncate returns the result of rounding t down to a multiple of d (since the zero time). If
// d <= 0, Truncate returns t stripped of any monotonic clock reading but otherwise unchanged.
//
// Truncate operates on the time as an absolute duration since the zero time; it does not operate
// on the presentation form of the time. Thus, Truncate(Hour) may return a time with a non-zero
// minute, depending on the time's Location.
func (t Time) Truncate(d time.Duration) Time {
	t.Time = t.Time.Truncate(d)
	return t
}

// AddDate returns the time corresponding to adding the given number of years, months, and days to
// t. For example, AddDate(-1, 2, 3) applied to January 1, 2011 returns March 4, 2010.
func (t Time) AddDate(years int, months int, days int) Time {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	return Date(year+years, month+time.Month(months), day+days, hour, min, sec,
		t.Nanosecond(), t.Location())
}

// After reports whether the time instant t is after u.
func (t Time) After(u any) bool {
	tm := getTime(u)
	return t.Time.After(tm)
}

// Before reports whether the time instant t is before u.
func (t Time) Before(u any) bool {
	tm := getTime(u)
	return t.Time.Before(tm)
}

// Compare compares the time instant t with u. If t is before u, it returns -1; if t is after u, it
// returns +1; if they're the same, it returns 0.
func (t Time) Compare(u any) int {
	tm := getTime(u)

	return t.Time.Compare(tm)
}

// Equal reports whether t and u represent the same time instant. Two times can be equal even if
// they are in different locations. For example, 6:00 +0200 and 4:00 UTC are Equal.
func (t Time) Equal(u any) bool {
	tm := getTime(u)

	return t.Time.Equal(tm)
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

// In returns a copy of t representing the same time instant, but with the copy's location
// information set to loc for display purposes.
//
// In panics if loc is nil.
func (t Time) In(loc *time.Location) Time {
	t.Time = t.Time.In(loc)
	return t
}

// Local returns t with the location set to local time.
func (t Time) Local() Time {
	t.Time = t.Time.Local()
	return t
}

// UTC returns t with the location set to UTC.
func (t Time) UTC() Time {
	t.Time = t.Time.UTC()
	return t
}

// ZoneBounds returns the bounds of the time zone in effect at time t. The zone begins at start and
// the next zone begins at end. If the zone begins at the beginning of time, start will be returned
// as a zero Time. If the zone goes on forever, end will be returned as a zero Time. The Location
// of the returned times will be the same as t.
func (t Time) ZoneBounds() (Time, Time) {
	startTm, endTm := t.Time.ZoneBounds()
	return New(startTm), New(endTm)
}
