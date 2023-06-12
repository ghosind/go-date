package date

import (
	"time"
)

type Time struct {
	time time.Time
}

// New creates and returns a new Time instance. It'll initialize by the parameter, or set the time
// to now.
func New(t ...time.Time) *Time {
	ins := new(Time)

	if len(t) > 0 {
		ins.time = t[0]
	} else {
		ins.time = time.Now()
	}

	return ins
}
