package date

import (
	"strconv"
	"testing"
	"time"

	"github.com/ghosind/go-assert"
)

func testFormat(a *assert.Assertion, time *Time, format, layout string) {
	str := time.Format(format)
	expect := time.Time.Format(layout)

	a.Equal(str, expect)
}

func TestFormat(t *testing.T) {
	a := assert.New(t)

	tzNY, _ := time.LoadLocation("America/New_York")
	tzSH, _ := time.LoadLocation("Asia/Shanghai")

	for _, time := range []*Time{
		Date(1, 0, 0, 0, 0, 0, 0),                // zero
		Now(),                                    // now
		New(time.Time{}),                         // now with built-in Time
		Date(2023, 6, 05, 3, 6, 9, 0, tzNY),      // 1-digit
		Date(2023, 6, 10, 10, 20, 30, 999, tzSH), // 2-digits
	} {
		testFormat(a, time, "", "")
		testFormat(a, time, "YYYY-MM-DDTHH:mm:ss.SSS Z", "2006-01-02T15:04:05.000 -07:00")
		testFormat(a, time, "YY-MMMM-DDTHH:mm:ss.S ZZ", "06-January-02T15:04:05.0 -0700")
		testFormat(a, time, "YYYY-MMM-DD hhA mm:ss.SS", "2006-Jan-02 03PM 04:05.00")
		testFormat(a, time, "YY-M-D ha m:s \\Y", "06-1-2 3pm 4:5 Y")

		expect := strconv.Itoa(time.Hour())
		hour := time.Format("H")

		a.Equal(hour, expect)
	}
}
