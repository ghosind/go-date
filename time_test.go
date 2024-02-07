package date

import (
	"testing"
	"time"

	"github.com/ghosind/go-assert"
)

func TestNew(t *testing.T) {
	a := assert.New(t)

	a.LtNow(time.Since(New().Time), time.Microsecond)
	tm := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	a.EqualNow(New(tm).Time, tm)
}

func TestNow(t *testing.T) {
	a := assert.New(t)

	now := Now()

	a.LtNow(time.Since(now.Time), time.Microsecond)
}

func TestDate(t *testing.T) {
	a := assert.New(t)

	a.EqualNow(Date(2024, 1, 1, 0, 0, 0, 0).Time, time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC))
	a.EqualNow(Date(2024, 1, 1, 0, 0, 0, 0, nil).Time, time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC))

	tzLA, _ := time.LoadLocation("America/Los_Angeles")
	a.EqualNow(Date(2024, 1, 1, 0, 0, 0, 0, tzLA).Time, time.Date(2024, 1, 1, 0, 0, 0, 0, tzLA))
}

func TestTimeHour12(t *testing.T) {
	a := assert.New(t)

	a.EqualNow(Date(2024, 1, 1, 0, 0, 0, 0).Hour12(), 12)
	a.EqualNow(Date(2024, 1, 1, 1, 0, 0, 0).Hour12(), 1)
	a.EqualNow(Date(2024, 1, 1, 12, 0, 0, 0).Hour12(), 12)
	a.EqualNow(Date(2024, 1, 1, 13, 0, 0, 0).Hour12(), 1)
}

func TestMicrosecond(t *testing.T) {
	a := assert.New(t)

	a.EqualNow(Date(2024, 1, 1, 0, 0, 0, 900).Microsecond(), 0)
	a.EqualNow(Date(2024, 1, 1, 0, 0, 0, 9000).Microsecond(), 9)
	a.EqualNow(Date(2024, 1, 1, 0, 0, 0, 99000).Microsecond(), 99)
	a.EqualNow(Date(2024, 1, 1, 0, 0, 0, 999000).Microsecond(), 999)
	a.EqualNow(Date(2024, 1, 1, 0, 0, 0, 9999000).Microsecond(), 9999)
	a.EqualNow(Date(2024, 1, 1, 0, 0, 0, 99999000).Microsecond(), 99999)
	a.EqualNow(Date(2024, 1, 1, 0, 0, 0, 999999000).Microsecond(), 999999)
	a.EqualNow(Date(2024, 1, 1, 0, 0, 0, 9999999000).Microsecond(), 999999)
}

func TestMillisecond(t *testing.T) {
	a := assert.New(t)

	a.EqualNow(Date(2024, 1, 1, 0, 0, 0, 900).Millisecond(), 0)
	a.EqualNow(Date(2024, 1, 1, 0, 0, 0, 900000).Millisecond(), 0)
	a.EqualNow(Date(2024, 1, 1, 0, 0, 0, 9000000).Millisecond(), 9)
	a.EqualNow(Date(2024, 1, 1, 0, 0, 0, 99000000).Millisecond(), 99)
	a.EqualNow(Date(2024, 1, 1, 0, 0, 0, 999000000).Millisecond(), 999)
	a.EqualNow(Date(2024, 1, 1, 0, 0, 0, 9999000000).Millisecond(), 999)
}