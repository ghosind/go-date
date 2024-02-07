package date

import (
	"testing"

	"github.com/ghosind/go-assert"
)

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
