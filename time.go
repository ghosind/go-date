package date

import (
	"time"
)

type Time struct {
	time time.Time
}

// New creates and returns a new Time instance. It'll initialize by the parameter, or set the time
// to now.
func New(t ...time.Time) *Time {
	ins := new(Time)

	if len(t) > 0 {
		ins.time = t[0]
	} else {
		ins.time = time.Now()
	}

	return ins
}

// Now returns the current time.
func Now() *Time {
	return New()
}

// Date creates and returns a new Time by the specific values. The location is an optional
// parameter, default time.UTC.
func Date(year, month, day, hour, min, sec, nsec int, loc ...*time.Location) *Time {
	ins := new(Time)

	location := time.UTC
	if len(loc) > 0 && loc[0] != nil {
		location = loc[0]
	}

	ins.time = time.Date(year, time.Month(month), day, hour, min, sec, nsec, location)

	return ins
}

// Year returns the year in which the time occurs.
func (t *Time) Year() int {
	return t.time.Year()
}

// YearDay returns the day of the year specified by th time, in the range [1, 365] for non-leap
// years, and [1, 366] in leap years.
func (t *Time) YearDay() int {
	return t.time.YearDay()
}

// Month returns the month of the year specified by the time.
func (t *Time) Month() time.Month {
	return t.time.Month()
}

// Day returns the day of the month specified by the time.
func (t *Time) Day() int {
	return t.time.Day()
}

// Date returns the year, month, and day in which the time occurs.
func (t *Time) Date() (year int, month time.Month, day int) {
	return t.time.Date()
}

// Weekday returns the day of the week specified by the time, in the range [0, 6] (Sunday = 0).
func (t *Time) Weekday() time.Weekday {
	return t.time.Weekday()
}

// ISOWeek returns the ISO 8601 year and week number in which the time occurs. Week ranges from 1
// to 53. Jan 01 to Jan 03 of year n might belong to week 52 or 53 of year n-1, and Dec 29 to
// Dec 31 might belong to week 1 of year n+1.
func (t *Time) ISOWeek() (year, week int) {
	return t.time.ISOWeek()
}

// Hour returns he hour offset within the day specified by the time, in the range [0, 23].
func (t *Time) Hour() int {
	return t.time.Hour()
}

// Minute returns he minute offset within the hour specified by the time, in the range [0, 59].
func (t *Time) Minute() int {
	return t.time.Minute()
}

// Second returns he second offset within the minute specified by the time, in the range [0, 59].
func (t *Time) Second() int {
	return t.time.Second()
}

// Clock returns the hour, minute, and second within the day specified by the time.
func (t *Time) Clock() (hour, min, sec int) {
	return t.time.Clock()
}

// Millisecond returns the millisecond offset within the second specified by the time, in the
// range [0, 999].
func (t *Time) Millisecond() int {
	return t.time.Nanosecond() / int(time.Millisecond)
}

// Microsecond returns the microsecond offset within the second specified by the time, in
// the range [0, 999999].
func (t *Time) Microsecond() int {
	return t.time.Nanosecond() / int(time.Microsecond)
}

// Nanosecond returns the nanosecond offset within the second specified by the time, in the
// range [0, 999999999].
func (t *Time) Nanosecond() int {
	return t.time.Nanosecond()
}

// Unix returns the time as a Unix time, the number of seconds elapsed since January 1, 1970 UTC.
// The result does not depend on the location associated with the time. Unix-like operating
// systems often record time as a 32-bit count of seconds, but since the method here returns
// a 64-bit value it is valid for billions of years into the past or future.
func (t *Time) Unix() int64 {
	return t.time.Unix()
}

// UnixMicro returns the time as a Unix time, the number of microseconds elapsed since January 1,
// 1970 UTC. The result is undefined if the Unix time in microseconds cannot be represented by an
// int64 (a date before year -290307 or after year 294246). The result does not depend on the
// location associated with the time.
func (t *Time) UnixMicro() int64 {
	return t.time.UnixMicro()
}

// UnixMilli returns the time as a Unix time, the number of milliseconds elapsed since January 1,
// 1970 UTC. The result is undefined if the Unix time in milliseconds cannot be represented by an
// int64 (a date more than 292 million years before or after 1970). The result does not depend on
// the location associated with the time.
func (t *Time) UnixMilli() int64 {
	return t.time.UnixMilli()
}

// UnixNano returns the time as a Unix time, the number of nanoseconds elapsed since January 1,
// 1970 UTC. The result is undefined if the Unix time in nanoseconds cannot be represented by an
// int64 (a date before the year 1678 or after 2262). Note that this means the result of calling
// UnixNano on the zero Time is undefined. The result does not depend on the location associated
// with the time.
func (t *Time) UnixNano() int64 {
	return t.time.UnixNano()
}

// String returns the time formatted using the format string
//
//	"2006-01-02 15:04:05.999999999 -0700 MST"
//
// If the time has a monotonic clock reading, the returned string includes a final field
// "m=±<value>", where value is the monotonic clock reading formatted as a decimal number of
// seconds.
// The returned string is meant for debugging; for a stable serialized representation, use
// t.MarshalText, t.MarshalBinary, or t.Format with an explicit format string.
func (t *Time) String() string {
	return t.time.String()
}

// GoString implements fmt.GoStringer and formats t to be printed in Go source code.
func (t *Time) GeoString() string {
	return t.time.GoString()
}

// GobEncode implements the gob.GobEncoder interface.
func (t Time) GobEncode() ([]byte, error) {
	return t.time.GobEncode()
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (t *Time) MarshalBinary() ([]byte, error) {
	return t.time.MarshalBinary()
}

// MarshalJSON implements the json.Marshaler interface. The time is a quoted string in the
// RFC 3339 format with sub-second precision. If the timestamp cannot be represented as valid
// RFC 3339 (e.g., the year is out of range), then an error is reported.
func (t *Time) MarshalJSON() ([]byte, error) {
	return t.time.MarshalJSON()
}

// MarshalText implements the encoding.TextMarshaler interface. The time is formatted in RFC 3339
// format with sub-second precision. If the timestamp cannot be represented as valid RFC 3339
// (e.g., the year is out of range), then an error is reported.
func (t *Time) MarshalText() ([]byte, error) {
	return t.time.MarshalText()
}

// Location returns the time zone information associated with the time.
func (t Time) Location() *time.Location {
	return t.time.Location()
}

// IsDST reports whether the time in the configured location is in Daylight Savings Time.
func (t *Time) IsDST() bool {
	return t.time.IsDST()
}

// IsZero reports whether the time represents the zero time instant, January 1, year 1,
// 00:00:00 UTC.
func (t *Time) IsZero() bool {
	return t.time.IsZero()
}

// Zone computes the time zone in effect at the time, returning the abbreviated name of the zone
// (such as "CET") and its offset in seconds east of UTC.
func (t *Time) Zone() (name string, offset int) {
	return t.time.Zone()
}
