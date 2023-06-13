package date

import (
	"strconv"
	"testing"
	"time"
)

func testFormat(t *testing.T, time *Time, format, layout string) {
	str := time.Format(format)
	expect := time.time.Format(layout)
	if str != expect {
		t.Errorf("time.Format(\"%s\") expect %s, actually %s", format, expect, str)
	}
}

func TestFormat(t *testing.T) {
	for _, time := range []*Time{
		New(),            // zero
		New(time.Time{}), // now
		New(time.Date(2023, 6, 05, 3, 6, 9, 0, time.UTC)),      // 1-digit
		New(time.Date(2023, 6, 10, 10, 20, 30, 999, time.UTC)), // 2-digits
	} {
		testFormat(t, time, "", "")
		testFormat(t, time, "YYYY-MM-DDTHH:mm:ss.SSS", "2006-01-02T15:04:05.000")
		testFormat(t, time, "YYYY-MMM-DDThh:mm:ss.SS", "2006-Jan-02T03:04:05.00")
		testFormat(t, time, "YY-MMMM-DDTHH:mm:ss.S", "06-January-02T15:04:05.0")
		testFormat(t, time, "YY-M-D h:m:s \\Y", "06-1-2 3:4:5 Y")

		expect := strconv.Itoa(time.time.Hour())
		hour := time.Format("H")
		if expect != hour {
			t.Errorf("time.Format(\"H\") expect %s, actually %s", expect, hour)
		}
	}
}
