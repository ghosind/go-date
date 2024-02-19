package date

import (
	"testing"
	"time"

	"github.com/ghosind/go-assert"
)

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
