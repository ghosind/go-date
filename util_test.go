package date

import "testing"

func testAppendIntToBuffer(t *testing.T, val, width int, expect string) {
	buf := make([]byte, 0, len(expect))
	buf = appendIntToBuffer(buf, val, width)
	if string(buf) != expect {
		t.Errorf("appendIntToBuffer(buf, %d, 0) expect returns \"%s\", actually %s", val, expect, string(buf))
	}
}

func TestAppendIntToBuffer(t *testing.T) {
	testAppendIntToBuffer(t, 123, 0, "123")
	testAppendIntToBuffer(t, 123, 4, "0123")
	testAppendIntToBuffer(t, -123, 0, "-123")
	testAppendIntToBuffer(t, -123, 5, "-0123")
}
