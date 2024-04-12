package date_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-date"
)

func ExampleTime_Format() {
	tm := date.Date(2006, time.January, 2, 15, 4, 5, 0)
	fmt.Println(tm.Format("YYYY-MM-DDTHH:mm:ss"))
	// Output:
	// 2006-01-02T15:04:05
}

func TestAppendFormat(t *testing.T) {
	a := assert.New(t)
	text := []byte("Time: ")
	tm := date.Date(2024, time.January, 1, 12, 30, 15, 0)

	text = tm.AppendFormat(text, "YYYY-MM-DD HH:mm:ss")
	a.EqualNow(string(text), "Time: 2024-01-01 12:30:15")
}

func TestFormat(t *testing.T) {
	a := assert.New(t)

	tzLA, _ := time.LoadLocation("America/Los_Angeles")
	tzSH, _ := time.LoadLocation("Asia/Shanghai")

	cases := []struct {
		tm     date.Time
		layout string
		expect string
	}{
		{
			date.Date(2006, time.January, 2, 15, 4, 5, 0),
			"YYYY-MM-DDTHH:mm:ss", "2006-01-02T15:04:05",
		},
		{
			date.Date(2006, time.January, 2, 15, 4, 5, 0),
			"YY-M-DTH:m:s", "06-1-2T15:4:5",
		},
		{
			date.Date(2006, time.January, 2, 15, 4, 5, 0),
			"dddd, DD MMMM, YYYY", "Monday, 02 January, 2006",
		},
		{
			date.Date(2006, time.January, 2, 15, 4, 5, 0),
			"ddd, DD MMM, YYYY", "Mon, 02 Jan, 2006",
		},
		{
			date.Date(2006, time.January, 2, 15, 4, 5, 0),
			"d, DD MMM, YYYY", "1, 02 Jan, 2006",
		},
		{
			date.Date(2006, time.January, 2, 15, 4, 5, 999000000),
			"hh:mm:ss.SSS A", "03:04:05.999 PM",
		},
		{
			date.Date(2006, time.January, 2, 3, 4, 5, 999000000),
			"hh:mm:ss.SS A", "03:04:05.99 AM",
		},
		{
			date.Date(2006, time.January, 2, 15, 4, 5, 999000000),
			"h:m:s.S a", "3:4:5.9 pm",
		},
		{
			date.Date(2006, time.January, 2, 3, 4, 5, 0),
			"h:m:s a", "3:4:5 am",
		},
		{
			date.Date(2006, time.January, 2, 15, 4, 5, 0, tzLA),
			"YYYY-MM-DD HH:mm:ss Z", "2006-01-02 15:04:05 -08:00",
		},
		{
			date.Date(2006, time.January, 2, 15, 4, 5, 0, tzSH),
			"YYYY-MM-DD HH:mm:ss ZZ", "2006-01-02 15:04:05 +0800",
		},
		{
			date.Date(2006, time.January, 2, 15, 4, 5, 0),
			"2006-01-02T15:04:05", "2006-01-02T15:04:05",
		},
	}

	for _, test := range cases {
		testFormat(a, test.tm, test.layout, test.expect)
	}
}

func testFormat(a *assert.Assertion, time date.Time, layout, expect string) {
	a.Helper()

	str := time.Format(layout)

	a.EqualNow(str, expect)
}
