package date_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-date"
)

func ExampleParse() {
	time.Local = time.UTC // force local to UTC, just for the example
	tm, err := date.Parse("YYYY-MM-DDTHH:mm:ss", "2006-01-02T15:04:05")
	fmt.Println(tm)
	fmt.Println(err)
	// Output:
	// 2006-01-02 15:04:05 +0000 UTC
	// <nil>
}

func TestParse(t *testing.T) {
	a := assert.New(t)

	cases := []struct {
		expect date.Time
		layout string
		str    string
	}{
		{date.Date(0, 1, 1, 0, 0, 0, 0, time.Local), "", ""},
		{date.Date(0, 1, 1, 0, 0, 0, 0, time.Local), "oo", "oo"},
		{date.Date(1970, 1, 1, 0, 0, 0, 0, time.Local), "YYYY-MM-DD HH:mm:ss", "1970-01-01 00:00:00"},
		{
			date.Date(2024, 12, 10, 13, 30, 30, 0, time.Local),
			"YYYY-MM-DD HH:mm:ss", "2024-12-10 13:30:30",
		},
		{date.Date(1970, 1, 1, 0, 0, 0, 0, time.Local), "YY-M-D H:m:s", "70-1-1 0:0:0"},
		{date.Date(2024, 12, 10, 13, 30, 30, 0, time.Local), "YY-M-D H:m:s", "24-12-10 13:30:30"},
		{
			date.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
			"YYYY-MMMM-DD HH:mm:ss", "2024-January-01 00:00:00",
		},
		{
			date.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
			"YYYY-MMM-DD HH:mm:ss", "2024-Jan-01 00:00:00",
		},
		{
			date.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
			"YYYY-MM-DD hh:mm:ss a", "2024-01-01 12:00:00 am",
		},
		{
			date.Date(2024, 1, 1, 12, 0, 0, 0, time.Local),
			"YYYY-MM-DD hh:mm:ss a", "2024-01-01 00:00:00 pm"},
		{
			date.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
			"YYYY-MM-DD h:mm:ss A", "2024-01-01 0:00:00 AM",
		},
		{
			date.Date(2024, 1, 1, 22, 0, 0, 0, time.Local),
			"YYYY-MM-DD h:mm:ss A", "2024-01-01 10:00:00 PM",
		},
		{date.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC), "YYYY-MM-DD ZZ", "2024-01-01 +0800"},
		{date.Date(2023, 12, 31, 16, 0, 0, 0, time.UTC), "YYYY-MM-DD ZZ", "2024-01-01 -0800"},
		{date.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC), "YYYY-MM-DD Z", "2024-01-01 +08:00"},
		{date.Date(2023, 12, 31, 16, 0, 0, 0, time.UTC), "YYYY-MM-DD Z", "2024-01-01 -08:00"},
		{
			date.Date(2024, 1, 1, 0, 0, 0, 999000000, time.Local),
			"YYYY-MM-DD HH:mm:ss.SSS", "2024-01-01 00:00:00.999",
		},
		{
			date.Date(2024, 1, 1, 0, 0, 0, 990000000, time.Local),
			"YYYY-MM-DD HH:mm:ss.SS", "2024-01-01 00:00:00.99",
		},
		{
			date.Date(2024, 1, 1, 0, 0, 0, 900000000, time.Local),
			"YYYY-MM-DD HH:mm:ss.S", "2024-01-01 00:00:00.9",
		},
		{
			date.Date(2006, time.January, 2, 15, 4, 5, 0, time.Local),
			"2006-01-02T15:04:05", "2006-01-02T15:04:05",
		},
	}

	for _, test := range cases {
		tm, err := date.Parse(test.layout, test.str)
		a.NilNow(err)
		a.NotNilNow(tm)
		a.TrueNow(tm.Equal(test.expect.Time))
	}
}

func TestParseWithError(t *testing.T) {
	a := assert.New(t)

	cases := []struct {
		layout        string
		str           string
		expectedError string
	}{
		{"oo", "o", `parsing time "o" as "oo": cannot parse "o" as ""`},
		{"oo", "ttt", `parsing time "ttt" as "oo": cannot parse "o" as "t"`},
		{"YYYY", "2xxx", `parsing time "2xxx" as "YYYY": cannot parse "YYYY" as "2xxx"`},
		{"YY", "2x", `parsing time "2x" as "YY": cannot parse "YY" as "2x"`},
		{"MMM", "XXX", `parsing time "XXX" as "MMM": cannot parse "MMM" as "XXX"`},
		{"MMMM", "unknown", `parsing time "unknown" as "MMMM": cannot parse "MMMM" as "unknown"`},
		{"S", "X", `parsing time "X" as "S": cannot parse "S" as "X"`},
		{"SS", "X", `parsing time "X" as "SS": cannot parse "SS" as "X"`},
		{"A", "p", `parsing time "p" as "A": cannot parse "A" as "p"`},
		{"A", "am", `parsing time "am" as "A": cannot parse "A" as "am"`},
		{"a", "p", `parsing time "p" as "a": cannot parse "a" as "p"`},
		{"a", "AM", `parsing time "AM" as "a": cannot parse "a" as "AM"`},
		{"Z", "+08", `parsing time "+08" as "Z": cannot parse "Z" as "+08"`},
		{"Z", "x08:00", `parsing time "x08:00" as "Z": cannot parse "Z" as "x08:00"`},
		{"ZZ", "+08", `parsing time "+08" as "ZZ": cannot parse "ZZ" as "+08"`},
		{"ZZ", "x0800", `parsing time "x0800" as "ZZ": cannot parse "ZZ" as "x0800"`},
	}

	for _, test := range cases {
		_, err := date.Parse(test.layout, test.str)
		a.NotNilNow(err)
		a.EqualNow(err.Error(), test.expectedError)
	}
}

func TestParseInLocation(t *testing.T) {
	a := assert.New(t)

	tzLA, _ := time.LoadLocation("America/Los_Angeles")
	tzSH, _ := time.LoadLocation("Asia/Shanghai")

	tm, err := date.ParseInLocation("YYYY-MM-DD", "2024-01-01", time.UTC)
	a.NilNow(err)
	a.TrueNow(tm.Equal(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)))

	tm, err = date.ParseInLocation("YYYY-MM-DD", "2024-01-01", tzLA)
	a.NilNow(err)
	a.TrueNow(tm.Equal(time.Date(2024, 1, 1, 0, 0, 0, 0, tzLA)))

	tm, err = date.ParseInLocation("YYYY-MM-DD", "2024-01-01", tzSH)
	a.NilNow(err)
	a.TrueNow(tm.Equal(time.Date(2024, 1, 1, 0, 0, 0, 0, tzSH)))
}

func TestParseInLocationName(t *testing.T) {
	a := assert.New(t)

	tzLA, _ := time.LoadLocation("America/Los_Angeles")
	tzSH, _ := time.LoadLocation("Asia/Shanghai")

	tm, err := date.ParseInLocationName("YYYY-MM-DD", "2024-01-01", "UTC")
	a.NilNow(err)
	a.TrueNow(tm.Equal(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)))

	tm, err = date.ParseInLocationName("YYYY-MM-DD", "2024-01-01", "America/Los_Angeles")
	a.NilNow(err)
	a.TrueNow(tm.Equal(time.Date(2024, 1, 1, 0, 0, 0, 0, tzLA)))

	tm, err = date.ParseInLocationName("YYYY-MM-DD", "2024-01-01", "Asia/Shanghai")
	a.NilNow(err)
	a.TrueNow(tm.Equal(time.Date(2024, 1, 1, 0, 0, 0, 0, tzSH)))

	_, err = date.ParseInLocationName("YYYY-MM-DD", "2024-01-01", "Unknown")
	a.NotNilNow(err)
}
