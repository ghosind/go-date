package date

import (
	"testing"
	"time"

	"github.com/ghosind/go-assert"
)

func TestNextLayoutToken(t *testing.T) {
	a := assert.New(t)
	layout := "YYYY YY MMMM MMM MM M DD D dddd ddd d HH H hh h mm m ss s SSS SS S A a Z ZZ \\Ho"
	expectedTokens := []int{
		layoutTokenYearLong, layoutTokenNone,
		layoutTokenYear, layoutTokenNone,
		layoutTokenMonthFull, layoutTokenNone,
		layoutTokenMonthAbbr, layoutTokenNone,
		layoutTokenMonthLong, layoutTokenNone,
		layoutTokenMonth, layoutTokenNone,
		layoutTokenDayLong, layoutTokenNone,
		layoutTokenDay, layoutTokenNone,
		layoutTokenDayOfWeekFull, layoutTokenNone,
		layoutTokenDayOfWeekAbbr, layoutTokenNone,
		layoutTokenDayOfWeek, layoutTokenNone,
		layoutTokenHourLong, layoutTokenNone,
		layoutTokenHour, layoutTokenNone,
		layoutTokenHour12Long, layoutTokenNone,
		layoutTokenHour12, layoutTokenNone,
		layoutTokenMinuteLong, layoutTokenNone,
		layoutTokenMinute, layoutTokenNone,
		layoutTokenSecondLong, layoutTokenNone,
		layoutTokenSecond, layoutTokenNone,
		layoutTokenMillisecond, layoutTokenNone,
		layoutTokenMillisecondTen, layoutTokenNone,
		layoutTokenMillisecondHundred, layoutTokenNone,
		layoutTokenPMUpper, layoutTokenNone,
		layoutTokenPMLower, layoutTokenNone,
		layoutTokenTZColon, layoutTokenNone,
		layoutTokenTZ, layoutTokenNone,
		layoutTokenNone, layoutTokenNone,
		layoutTokenEnd,
	}

	for _, expected := range expectedTokens {
		token, _, suffix := nextLayoutToken(layout)
		layout = suffix
		a.EqualNow(token, expected)
	}

	a.EqualNow(layout, "")
}

func TestNextLayoutTokenWithBuiltinLayout(t *testing.T) {
	a := assert.New(t)
	layout := "2006 06 January Jan 01 1 02 2 Monday Mon 15 03 3 04 4 05 5 PM 0"
	expectedTokens := []int{
		layoutTokenYearLong, layoutTokenNone,
		layoutTokenYear, layoutTokenNone,
		layoutTokenMonthFull, layoutTokenNone,
		layoutTokenMonthAbbr, layoutTokenNone,
		layoutTokenMonthLong, layoutTokenNone,
		layoutTokenMonth, layoutTokenNone,
		layoutTokenDayLong, layoutTokenNone,
		layoutTokenDay, layoutTokenNone,
		layoutTokenDayOfWeekFull, layoutTokenNone,
		layoutTokenDayOfWeekAbbr, layoutTokenNone,
		layoutTokenHourLong, layoutTokenNone,
		layoutTokenHour12Long, layoutTokenNone,
		layoutTokenHour12, layoutTokenNone,
		layoutTokenMinuteLong, layoutTokenNone,
		layoutTokenMinute, layoutTokenNone,
		layoutTokenSecondLong, layoutTokenNone,
		layoutTokenSecond, layoutTokenNone,
		layoutTokenPMUpper, layoutTokenNone,
		layoutTokenNone,
		layoutTokenEnd,
	}

	for _, expected := range expectedTokens {
		token, _, suffix := nextLayoutToken(layout)
		layout = suffix
		a.EqualNow(token, expected)
	}

	a.EqualNow(layout, "")
}

func TestGetTime(t *testing.T) {
	a := assert.New(t)

	tm := Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)

	a.Equal(getTime(tm), tm.Time)
	a.Equal(getTime(tm.Time), tm.Time)
	a.PanicOfNow(func() { getTime(1) }, ErrNotTime)
}

func TestAppendIntToBuffer(t *testing.T) {
	a := assert.New(t)

	testAppendIntToBuffer(a, 123, 0, "123")
	testAppendIntToBuffer(a, 123, 4, "0123")
	testAppendIntToBuffer(a, -123, 0, "-123")
	testAppendIntToBuffer(a, -123, 5, "-0123")
}

func testAppendIntToBuffer(a *assert.Assertion, val, width int, expect string) {
	a.Helper()

	buf := make([]byte, 0, len(expect))
	buf = appendIntToBuffer(buf, val, width)

	a.Equal(string(buf), expect)
}

func TestLookup(t *testing.T) {
	a := assert.New(t)

	cases := []struct {
		str      string
		expected int
	}{
		{"Jan", 0},
		{"JAN", 0},
		{"jan", 0},
		{"Apr", 3},
		{"DecApr", 11},
		{"DECAPR", 11},
		{"TEST", -1},
		{"T", -1},
	}

	for _, test := range cases {
		i, _, err := lookup(abbrMonthNames, test.str)
		if test.expected >= 0 {
			a.NilNow(err)
		} else {
			a.NotNilNow(err)
		}

		a.EqualNow(i, test.expected)
	}
}

func TestReadNum(t *testing.T) {
	a := assert.New(t)

	testReadNum(a, "1234", 12, 2, true, false)
	testReadNum(a, "1234", 1234, 4, true, false)
	testReadNum(a, "12", 12, 4, false, false)
	testReadNum(a, "12", -1, 4, true, true)
	testReadNum(a, "xx", -1, 2, true, true)
}

func testReadNum(
	a *assert.Assertion,
	value string,
	expected, width int,
	fixed bool,
	hasError bool,
) {
	a.Helper()

	i, s, err := readNum(value, width, fixed)
	if hasError {
		a.NotNilNow(err)
		a.EqualNow(s, value)
	} else {
		a.NilNow(err)
		a.EqualNow(i, expected)
	}
}
