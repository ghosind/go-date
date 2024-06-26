package date_test

import (
	"testing"
	"time"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-date"
)

func TestNew(t *testing.T) {
	a := assert.New(t)

	a.LtNow(time.Since(date.New().Time), time.Microsecond)
	tm := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	a.EqualNow(date.New(tm).Time, tm)
}

func TestDate(t *testing.T) {
	a := assert.New(t)

	a.EqualNow(date.Date(2024, 1, 1, 0, 0, 0, 0).Time, time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC))
	a.EqualNow(
		date.Date(2024, 1, 1, 0, 0, 0, 0, nil).Time,
		time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	)

	tzLA, _ := time.LoadLocation("America/Los_Angeles")
	a.EqualNow(date.Date(2024, 1, 1, 0, 0, 0, 0, tzLA).Time, time.Date(2024, 1, 1, 0, 0, 0, 0, tzLA))
}

func TestNow(t *testing.T) {
	a := assert.New(t)

	now := date.Now()

	a.LtNow(time.Since(now.Time), time.Microsecond)
}

func TestUnix(t *testing.T) {
	a := assert.New(t)

	unixTime := date.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	tm := date.Unix(unixTime.Unix(), 0)
	a.TrueNow(tm.Equal(unixTime.Time))
}

func TestUnixMicro(t *testing.T) {
	a := assert.New(t)

	unixTime := date.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	tm := date.UnixMicro(unixTime.UnixMicro())
	a.TrueNow(tm.Equal(unixTime.Time))
}

func TestUnixMilli(t *testing.T) {
	a := assert.New(t)

	unixTime := date.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	tm := date.UnixMilli(unixTime.UnixMilli())
	a.TrueNow(tm.Equal(unixTime.Time))
}

func TestSince(t *testing.T) {
	a := assert.New(t)

	tm := time.Now().Add(-time.Minute)

	dur := date.Since(tm)
	a.GteNow(dur, time.Minute)
	a.LteNow(dur, time.Minute+time.Millisecond)

	dur = date.Since(date.New(tm))
	a.GteNow(dur, time.Minute)
	a.LteNow(dur, time.Minute+time.Millisecond)
}

func TestUntil(t *testing.T) {
	a := assert.New(t)

	tm := time.Now().Add(time.Minute)

	dur := date.Until(tm)
	a.LteNow(dur, time.Minute)
	a.GteNow(dur, time.Minute-time.Millisecond)

	dur = date.Until(date.New(tm))
	a.LteNow(dur, time.Minute)
	a.GteNow(dur, time.Minute-time.Millisecond)
}

func TestAdd(t *testing.T) {
	a := assert.New(t)

	now := date.Now()
	tm := now.Add(time.Minute)
	a.EqualNow(tm.Sub(now), time.Minute)
}

func TestSub(t *testing.T) {
	a := assert.New(t)

	now := time.Now()
	tm := date.New(now).Add(time.Minute)

	a.Equal(tm.Sub(now), time.Minute)
	a.Equal(tm.Sub(date.New(now)), time.Minute)
}

func TestRound(t *testing.T) {
	a := assert.New(t)

	tm := date.Date(0, 0, 0, 12, 15, 30, 918273645, time.UTC)
	cases := []struct {
		dur    time.Duration
		expect date.Time
	}{
		{time.Nanosecond, date.Date(0, 0, 0, 12, 15, 30, 918273645, time.UTC)},
		{time.Microsecond, date.Date(0, 0, 0, 12, 15, 30, 918274000, time.UTC)},
		{time.Millisecond, date.Date(0, 0, 0, 12, 15, 30, 918000000, time.UTC)},
		{time.Second, date.Date(0, 0, 0, 12, 15, 31, 0, time.UTC)},
		{2 * time.Second, date.Date(0, 0, 0, 12, 15, 30, 0, time.UTC)},
		{time.Minute, date.Date(0, 0, 0, 12, 16, 0, 0, time.UTC)},
		{10 * time.Minute, date.Date(0, 0, 0, 12, 20, 0, 0, time.UTC)},
		{time.Hour, date.Date(0, 0, 0, 12, 0, 0, 0, time.UTC)},
	}

	for _, test := range cases {
		a.TrueNow(tm.Round(test.dur).Equal(test.expect))
	}
}

func TestTruncate(t *testing.T) {
	a := assert.New(t)

	tm := date.Date(0, 0, 0, 12, 15, 30, 918273645, time.UTC)
	cases := []struct {
		dur    time.Duration
		expect date.Time
	}{
		{time.Nanosecond, date.Date(0, 0, 0, 12, 15, 30, 918273645, time.UTC)},
		{time.Microsecond, date.Date(0, 0, 0, 12, 15, 30, 918273000, time.UTC)},
		{time.Millisecond, date.Date(0, 0, 0, 12, 15, 30, 918000000, time.UTC)},
		{time.Second, date.Date(0, 0, 0, 12, 15, 30, 0, time.UTC)},
		{2 * time.Second, date.Date(0, 0, 0, 12, 15, 30, 0, time.UTC)},
		{time.Minute, date.Date(0, 0, 0, 12, 15, 0, 0, time.UTC)},
		{10 * time.Minute, date.Date(0, 0, 0, 12, 10, 0, 0, time.UTC)},
		{time.Hour, date.Date(0, 0, 0, 12, 0, 0, 0, time.UTC)},
	}

	for _, test := range cases {
		a.TrueNow(tm.Truncate(test.dur).Equal(test.expect))
	}
}

func TestAddDate(t *testing.T) {
	a := assert.New(t)

	tm := date.Date(2024, time.January, 1, 12, 30, 30, 0, time.UTC)
	expect := date.Date(2025, time.February, 15, 12, 30, 30, 0, time.UTC)

	a.TrueNow(tm.AddDate(1, 1, 14).Equal(expect))
}

func TestAfter(t *testing.T) {
	a := assert.New(t)

	tm := date.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	after := date.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	before := date.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)

	a.TrueNow(tm.After(before))
	a.TrueNow(tm.After(before.Time))
	a.NotTrueNow(tm.After(after))
	a.NotTrueNow(tm.After(after.Time))
}

func TestBefore(t *testing.T) {
	a := assert.New(t)

	tm := date.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	after := date.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	before := date.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)

	a.TrueNow(tm.Before(after))
	a.TrueNow(tm.Before(after.Time))
	a.NotTrueNow(tm.Before(before))
	a.NotTrueNow(tm.Before(before.Time))
}

func TestCompare(t *testing.T) {
	a := assert.New(t)

	tm := date.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	equal := date.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	after := date.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	before := date.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)

	a.EqualNow(tm.Compare(equal), 0)
	a.EqualNow(tm.Compare(equal.Time), 0)
	a.EqualNow(tm.Compare(after), -1)
	a.EqualNow(tm.Compare(after.Time), -1)
	a.EqualNow(tm.Compare(before), 1)
	a.EqualNow(tm.Compare(before.Time), 1)
}

func TestEqual(t *testing.T) {
	a := assert.New(t)
	tzLA, _ := time.LoadLocation("America/Los_Angeles")

	a.TrueNow(date.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC).
		Equal(date.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)))
	a.NotTrueNow(date.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC).
		Equal(date.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)))
	a.NotTrueNow(date.Date(2024, time.January, 1, 0, 0, 0, 0, tzLA).
		Equal(date.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)))

	a.TrueNow(date.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC).
		Equal(time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)))
	a.NotTrueNow(date.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC).
		Equal(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)))
	a.NotTrueNow(date.Date(2024, time.January, 1, 0, 0, 0, 0, tzLA).
		Equal(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)))

	a.TrueNow(date.Date(2024, time.January, 1, 8, 0, 0, 0, time.UTC).
		Equal(date.Date(2024, time.January, 1, 0, 0, 0, 0, tzLA)))

	a.PanicOfNow(
		func() { date.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC).Equal(1) },
		date.ErrNotTime,
	)
}

func TestTimeHour12(t *testing.T) {
	a := assert.New(t)

	a.EqualNow(date.Date(2024, 1, 1, 0, 0, 0, 0).Hour12(), 12)
	a.EqualNow(date.Date(2024, 1, 1, 1, 0, 0, 0).Hour12(), 1)
	a.EqualNow(date.Date(2024, 1, 1, 12, 0, 0, 0).Hour12(), 12)
	a.EqualNow(date.Date(2024, 1, 1, 13, 0, 0, 0).Hour12(), 1)
}

func TestMicrosecond(t *testing.T) {
	a := assert.New(t)

	a.EqualNow(date.Date(2024, 1, 1, 0, 0, 0, 900).Microsecond(), 0)
	a.EqualNow(date.Date(2024, 1, 1, 0, 0, 0, 9000).Microsecond(), 9)
	a.EqualNow(date.Date(2024, 1, 1, 0, 0, 0, 99000).Microsecond(), 99)
	a.EqualNow(date.Date(2024, 1, 1, 0, 0, 0, 999000).Microsecond(), 999)
	a.EqualNow(date.Date(2024, 1, 1, 0, 0, 0, 9999000).Microsecond(), 9999)
	a.EqualNow(date.Date(2024, 1, 1, 0, 0, 0, 99999000).Microsecond(), 99999)
	a.EqualNow(date.Date(2024, 1, 1, 0, 0, 0, 999999000).Microsecond(), 999999)
	a.EqualNow(date.Date(2024, 1, 1, 0, 0, 0, 9999999000).Microsecond(), 999999)
}

func TestMillisecond(t *testing.T) {
	a := assert.New(t)

	a.EqualNow(date.Date(2024, 1, 1, 0, 0, 0, 900).Millisecond(), 0)
	a.EqualNow(date.Date(2024, 1, 1, 0, 0, 0, 900000).Millisecond(), 0)
	a.EqualNow(date.Date(2024, 1, 1, 0, 0, 0, 9000000).Millisecond(), 9)
	a.EqualNow(date.Date(2024, 1, 1, 0, 0, 0, 99000000).Millisecond(), 99)
	a.EqualNow(date.Date(2024, 1, 1, 0, 0, 0, 999000000).Millisecond(), 999)
	a.EqualNow(date.Date(2024, 1, 1, 0, 0, 0, 9999000000).Millisecond(), 999)
}

func TestIn(t *testing.T) {
	a := assert.New(t)

	tzLA, _ := time.LoadLocation("America/Los_Angeles")

	utc := date.Date(2024, time.January, 1, 8, 0, 0, 0, time.UTC)
	tm := utc.In(tzLA)

	a.TrueNow(tm.Equal(utc))
	a.EqualNow(utc.Format("YYYY-MM-DD HH:mm:ss Z"), "2024-01-01 08:00:00 +00:00")
	a.EqualNow(tm.Format("YYYY-MM-DD HH:mm:ss Z"), "2024-01-01 00:00:00 -08:00")
}

func TestLocal(t *testing.T) {
	a := assert.New(t)

	utc := date.Date(2024, time.January, 1, 0, 0, 0, 0)
	tm := utc.Local()

	a.TrueNow(tm.Equal(utc))
}

func TestUTC(t *testing.T) {
	a := assert.New(t)

	tzLA, _ := time.LoadLocation("America/Los_Angeles")

	tm := date.Date(2024, time.January, 1, 0, 0, 0, 0, tzLA)
	utc := tm.UTC()

	a.TrueNow(tm.Equal(utc))
	a.EqualNow(utc.Format("YYYY-MM-DD HH:mm:ss Z"), "2024-01-01 08:00:00 +00:00")
	a.EqualNow(tm.Format("YYYY-MM-DD HH:mm:ss Z"), "2024-01-01 00:00:00 -08:00")
}

func TestZoneBounds(t *testing.T) {
	a := assert.New(t)

	expectedStart, expectedEnd := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC).ZoneBounds()
	start, end := date.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC).ZoneBounds()
	a.TrueNow(start.Equal(expectedStart))
	a.TrueNow(end.Equal(expectedEnd))
}

func TestStartOfYear(t *testing.T) {
	a := assert.New(t)

	tzLA, _ := time.LoadLocation("America/Los_Angeles")

	a.TrueNow(date.Date(2006, 2, 2, 15, 4, 5, 0).
		StartOfYear().
		Equal(date.Date(2006, 1, 1, 0, 0, 0, 0)))
	a.TrueNow(date.Date(2006, 2, 2, 15, 4, 5, 0, tzLA).
		StartOfYear().
		Equal(date.Date(2006, 1, 1, 0, 0, 0, 0, tzLA)))
}

func TestStartOfHalfYear(t *testing.T) {
	a := assert.New(t)

	a.TrueNow(date.Date(2006, 1, 3, 15, 4, 5, 0).
		StartOfHalfYear().
		Equal(date.Date(2006, 1, 1, 0, 0, 0, 0)))
	a.TrueNow(date.Date(2006, 2, 3, 15, 4, 5, 0).
		StartOfHalfYear().
		Equal(date.Date(2006, 1, 1, 0, 0, 0, 0)))
	a.TrueNow(date.Date(2006, 7, 3, 15, 4, 5, 0).
		StartOfHalfYear().
		Equal(date.Date(2006, 7, 1, 0, 0, 0, 0)))
	a.TrueNow(date.Date(2006, 12, 3, 15, 4, 5, 0).
		StartOfHalfYear().
		Equal(date.Date(2006, 7, 1, 0, 0, 0, 0)))
}

func TestStartOfQuarter(t *testing.T) {
	a := assert.New(t)

	a.TrueNow(date.Date(2006, 1, 15, 12, 30, 30, 999).
		StartOfQuarter().
		Equal(date.Date(2006, 1, 1, 0, 0, 0, 0)))
	a.TrueNow(date.Date(2006, 2, 15, 12, 30, 30, 999).
		StartOfQuarter().
		Equal(date.Date(2006, 1, 1, 0, 0, 0, 0)))
	a.TrueNow(date.Date(2006, 4, 15, 12, 30, 30, 999).
		StartOfQuarter().
		Equal(date.Date(2006, 4, 1, 0, 0, 0, 0)))
	a.TrueNow(date.Date(2006, 12, 15, 12, 30, 30, 999).
		StartOfQuarter().
		Equal(date.Date(2006, 10, 1, 0, 0, 0, 0)))
}

func TestStartOfMonth(t *testing.T) {
	a := assert.New(t)

	tzLA, _ := time.LoadLocation("America/Los_Angeles")

	a.TrueNow(date.Date(2006, 2, 2, 15, 4, 5, 0).
		StartOfMonth().
		Equal(date.Date(2006, 2, 1, 0, 0, 0, 0)))
	a.TrueNow(date.Date(2006, 2, 2, 15, 4, 5, 0, tzLA).
		StartOfMonth().
		Equal(date.Date(2006, 2, 1, 0, 0, 0, 0, tzLA)))
}

func TestStartOfDay(t *testing.T) {
	a := assert.New(t)

	tzLA, _ := time.LoadLocation("America/Los_Angeles")

	a.TrueNow(date.Date(2006, 2, 2, 15, 4, 5, 0).
		StartOfDay().
		Equal(date.Date(2006, 2, 2, 0, 0, 0, 0)))
	a.TrueNow(date.Date(2006, 2, 2, 15, 4, 5, 0, tzLA).
		StartOfDay().
		Equal(date.Date(2006, 2, 2, 0, 0, 0, 0, tzLA)))
}

func TestStartOfHour(t *testing.T) {
	a := assert.New(t)

	tzLA, _ := time.LoadLocation("America/Los_Angeles")

	a.TrueNow(date.Date(2006, 2, 3, 15, 4, 5, 999).
		StartOfHour().
		Equal(date.Date(2006, 2, 3, 15, 0, 0, 0)))
	a.TrueNow(date.Date(2006, 2, 3, 15, 4, 5, 999, tzLA).
		StartOfHour().
		Equal(date.Date(2006, 2, 3, 15, 0, 0, 0, tzLA)))
}

func TestStartOfMinute(t *testing.T) {
	a := assert.New(t)

	tzLA, _ := time.LoadLocation("America/Los_Angeles")

	a.TrueNow(date.Date(2006, 2, 3, 15, 30, 55, 999).
		StartOfMinute().
		Equal(date.Date(2006, 2, 3, 15, 30, 0, 0)))
	a.TrueNow(date.Date(2006, 2, 3, 15, 30, 55, 999, tzLA).
		StartOfMinute().
		Equal(date.Date(2006, 2, 3, 15, 30, 0, 0, tzLA)))
}

func TestStartOfSecond(t *testing.T) {
	a := assert.New(t)

	tzLA, _ := time.LoadLocation("America/Los_Angeles")

	a.TrueNow(date.Date(2006, 2, 3, 15, 30, 55, 999).
		StartOfSecond().
		Equal(date.Date(2006, 2, 3, 15, 30, 55, 0)))
	a.TrueNow(date.Date(2006, 2, 3, 15, 30, 55, 999, tzLA).
		StartOfSecond().
		Equal(date.Date(2006, 2, 3, 15, 30, 55, 0, tzLA)))
}

func TestEndOfYear(t *testing.T) {
	a := assert.New(t)

	tzLA, _ := time.LoadLocation("America/Los_Angeles")

	a.TrueNow(date.Date(2006, 2, 2, 15, 4, 5, 0).
		EndOfYear().
		Equal(date.Date(2006, 12, 31, 23, 59, 59, 999999999)))
	a.TrueNow(date.Date(2006, 2, 2, 15, 4, 5, 0, tzLA).
		EndOfYear().
		Equal(date.Date(2006, 12, 31, 23, 59, 59, 999999999, tzLA)))
}

func TestEndOfHalfYear(t *testing.T) {
	a := assert.New(t)

	a.TrueNow(date.Date(2006, 1, 3, 15, 4, 5, 0).
		EndOfHalfYear().
		Equal(date.Date(2006, 6, 30, 23, 59, 59, 999999999)))
	a.TrueNow(date.Date(2006, 2, 3, 15, 4, 5, 0).
		EndOfHalfYear().
		Equal(date.Date(2006, 6, 30, 23, 59, 59, 999999999)))
	a.TrueNow(date.Date(2006, 7, 3, 15, 4, 5, 0).
		EndOfHalfYear().
		Equal(date.Date(2006, 12, 31, 23, 59, 59, 999999999)))
	a.TrueNow(date.Date(2006, 12, 3, 15, 4, 5, 0).
		EndOfHalfYear().
		Equal(date.Date(2006, 12, 31, 23, 59, 59, 999999999)))
}

func TestEndOfQuarter(t *testing.T) {
	a := assert.New(t)

	a.TrueNow(date.Date(2006, 1, 15, 12, 30, 30, 999).
		EndOfQuarter().
		Equal(date.Date(2006, 3, 31, 23, 59, 59, 999999999)))
	a.TrueNow(date.Date(2006, 2, 15, 12, 30, 30, 999).
		EndOfQuarter().
		Equal(date.Date(2006, 3, 31, 23, 59, 59, 999999999)))
	a.TrueNow(date.Date(2006, 4, 15, 12, 30, 30, 999).
		EndOfQuarter().
		Equal(date.Date(2006, 6, 30, 23, 59, 59, 999999999)))
	a.TrueNow(date.Date(2006, 12, 15, 12, 30, 30, 999).
		EndOfQuarter().
		Equal(date.Date(2006, 12, 31, 23, 59, 59, 999999999)))
}

func TestEndOfMonth(t *testing.T) {
	a := assert.New(t)

	a.TrueNow(date.Date(2006, 1, 15, 15, 4, 5, 0).
		EndOfMonth().
		Equal(date.Date(2006, 1, 31, 23, 59, 59, 999999999)))
	a.TrueNow(date.Date(2006, 2, 15, 15, 4, 5, 0).
		EndOfMonth().
		Equal(date.Date(2006, 2, 28, 23, 59, 59, 999999999)))
	a.TrueNow(date.Date(2008, 2, 15, 15, 4, 5, 0).
		EndOfMonth().
		Equal(date.Date(2008, 2, 29, 23, 59, 59, 999999999)))
	a.TrueNow(date.Date(2006, 6, 15, 15, 4, 5, 0).
		EndOfMonth().
		Equal(date.Date(2006, 6, 30, 23, 59, 59, 999999999)))
}

func TestEndOfDay(t *testing.T) {
	a := assert.New(t)

	tzLA, _ := time.LoadLocation("America/Los_Angeles")

	a.TrueNow(date.Date(2006, 2, 2, 15, 4, 5, 0).
		EndOfDay().
		Equal(date.Date(2006, 2, 2, 23, 59, 59, 999999999)))
	a.TrueNow(date.Date(2006, 2, 2, 15, 4, 5, 0, tzLA).
		EndOfDay().
		Equal(date.Date(2006, 2, 2, 23, 59, 59, 999999999, tzLA)))
}

func TestEndOfHour(t *testing.T) {
	a := assert.New(t)

	tzLA, _ := time.LoadLocation("America/Los_Angeles")

	a.TrueNow(date.Date(2006, 2, 3, 15, 4, 5, 999).
		EndOfHour().
		Equal(date.Date(2006, 2, 3, 15, 59, 59, 999999999)))
	a.TrueNow(date.Date(2006, 2, 3, 15, 4, 5, 999, tzLA).
		EndOfHour().
		Equal(date.Date(2006, 2, 3, 15, 59, 59, 999999999, tzLA)))
}

func TestEndOfMinute(t *testing.T) {
	a := assert.New(t)

	tzLA, _ := time.LoadLocation("America/Los_Angeles")

	a.TrueNow(date.Date(2006, 2, 3, 15, 30, 55, 999).
		EndOfMinute().
		Equal(date.Date(2006, 2, 3, 15, 30, 59, 999999999)))
	a.TrueNow(date.Date(2006, 2, 3, 15, 30, 55, 999, tzLA).
		EndOfMinute().
		Equal(date.Date(2006, 2, 3, 15, 30, 59, 999999999, tzLA)))
}

func TestEndOfSecond(t *testing.T) {
	a := assert.New(t)

	tzLA, _ := time.LoadLocation("America/Los_Angeles")

	a.TrueNow(date.Date(2006, 2, 3, 15, 30, 55, 999).
		EndOfSecond().
		Equal(date.Date(2006, 2, 3, 15, 30, 55, 999999999)))
	a.TrueNow(date.Date(2006, 2, 3, 15, 30, 55, 999, tzLA).
		EndOfSecond().
		Equal(date.Date(2006, 2, 3, 15, 30, 55, 999999999, tzLA)))
}
