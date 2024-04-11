package date_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-date"
)

func ExampleFormat() {
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

	layout := "YYYY YY MMMM MMM MM M DD D dddd ddd d HH H hh h mm m ss s SSS SS S A a Z ZZ \\Ho"
	cases := []struct {
		tm     date.Time
		expect string
	}{
		{
			date.Date(2024, 1, 1, 0, 0, 0, 0),
			"2024 24 January Jan 01 1 01 1 Monday Mon 1 00 0 12 12 00 0 00 0 000 00 0 AM am +00:00 +0000 Ho",
		},
		{
			date.Date(2024, 10, 1, 0, 0, 0, 0),
			"2024 24 October Oct 10 10 01 1 Tuesday Tue 2 00 0 12 12 00 0 00 0 000 00 0 AM am +00:00 +0000 Ho",
		},
		{
			date.Date(2024, 1, 1, 1, 0, 0, 0),
			"2024 24 January Jan 01 1 01 1 Monday Mon 1 01 1 01 1 00 0 00 0 000 00 0 AM am +00:00 +0000 Ho",
		},
		{
			date.Date(2024, 1, 1, 12, 0, 0, 0),
			"2024 24 January Jan 01 1 01 1 Monday Mon 1 12 12 12 12 00 0 00 0 000 00 0 PM pm +00:00 +0000 Ho",
		},
		{
			date.Date(2024, 1, 1, 13, 0, 0, 0),
			"2024 24 January Jan 01 1 01 1 Monday Mon 1 13 13 01 1 00 0 00 0 000 00 0 PM pm +00:00 +0000 Ho",
		},
		{
			date.Date(2024, 1, 1, 0, 0, 0, 0, tzLA),
			"2024 24 January Jan 01 1 01 1 Monday Mon 1 00 0 12 12 00 0 00 0 000 00 0 AM am -08:00 -0800 Ho",
		},
		{
			date.Date(2024, 1, 1, 0, 0, 0, 0, tzSH),
			"2024 24 January Jan 01 1 01 1 Monday Mon 1 00 0 12 12 00 0 00 0 000 00 0 AM am +08:00 +0800 Ho",
		},
	}

	for _, test := range cases {
		testFormat(a, test.tm, layout, test.expect)
	}
}

func testFormat(a *assert.Assertion, time date.Time, layout, expect string) {
	a.Helper()

	str := time.Format(layout)

	a.EqualNow(str, expect)
}
