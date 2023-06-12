package date

import (
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
	for _, time := range []*Time{New(), New(time.Time{})} {
		testFormat(t, time, "", "")
		testFormat(t, time, "YYYY-MM-DDTHH:mm:ss.SSS", "2006-01-02T15:04:05.000")
		testFormat(t, time, "YYYY-MMM-DDThh:mm:ss.SS", "2006-Jan-02T03:04:05.00")
		testFormat(t, time, "YY-MMMM-DDTHH:mm:ss.S", "06-January-02T15:04:05.0")
		testFormat(t, time, "YY-M-D h:m:s \\Y", "06-1-2 3:4:5 Y")
	}
}
