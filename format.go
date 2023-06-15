package date

import (
	"strings"
)

const (
	// layoutTokenEnd indicates the layout is empty.
	layoutTokenEnd int = iota
	// layoutTokenNone is not a token.
	layoutTokenNone
	// layoutTokenYear is the year.
	layoutTokenYear
	// layoutTokenYearTwo is the two-digits year.
	layoutTokenYearTwo
	// layoutTokenMonth is the month beginning at 1.
	layoutTokenMonth
	// layoutTokenMonthTwo is the two-digits month.
	layoutTokenMonthTwo
	// layoutTokenMonthAbbr is the abbreviated month name.
	layoutTokenMonthAbbr
	// layoutTokenMonthFull is the full month name.
	layoutTokenMonthFull
	// layoutTokenDay is the day beginning at 1.
	layoutTokenDay
	// layoutTokenDayTwo is the two-digits day.
	layoutTokenDayTwo
	// layoutTokenHour is the 24-hour clock hour that beginning at 1.
	layoutTokenHour
	// layoutTokenHourTwo is the two-digits, 24-hour clock hour.
	layoutTokenHourTwo
	// layoutTokenHour12 is the 12-hour clock hour that beginning at 1.
	layoutTokenHour12
	// layoutTokenHour12Two is the two-digits, 12-hour clock hour.
	layoutTokenHour12Two
	// layoutTokenMinute is the minute beginning at 1.
	layoutTokenMinute
	// layoutTokenMinuteTwo is the two-digits minute.
	layoutTokenMinuteTwo
	// layoutTokenSecond is the second beginning at 1.
	layoutTokenSecond
	// layoutTokenSecondTwo is the two-digits second.
	layoutTokenSecondTwo
	// layoutTokenMillisecond is the one-digit millisecond (hundreds of milliseconds).
	layoutTokenMillisecond
	// layoutTokenMillisecondTwo is the two-digits millisecond (tens of milliseconds).
	layoutTokenMillisecondTwo
	// layoutTokenMillisecondThree is the three-digits millisecond.
	layoutTokenMillisecondThree
	// layoutTokenPMUpper is post or ante meridiem in upper-case.
	layoutTokenPMUpper
	// layoutTokenPMLower is post or ante meridiem in upper-case.
	layoutTokenPMLower
	// layoutTokenTZ is the timezone offset from UTC.
	layoutTokenTZ
	// layoutTokenTZColon is the timezone offset from UTC that separate by colon.
	layoutTokenTZColon
)

var abbrMonthNames = []string{
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

var fullMonthNames = []string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

// nextLayoutToken gets the next token in the layout, and return
func nextLayoutToken(layout string) (int, string, string) {
	switch layout[0] {
	case 'Y':
		if strings.HasPrefix(layout, "YYYY") {
			return layoutTokenYear, layout[0:4], layout[4:]
		} else if strings.HasPrefix(layout, "YY") {
			return layoutTokenYearTwo, layout[0:2], layout[2:]
		}
	case 'M':
		if strings.HasPrefix(layout, "MMMM") {
			return layoutTokenMonthFull, layout[0:4], layout[4:]
		} else if strings.HasPrefix(layout, "MMM") {
			return layoutTokenMonthAbbr, layout[0:3], layout[3:]
		} else if strings.HasPrefix(layout, "MM") {
			return layoutTokenMonthTwo, layout[0:2], layout[2:]
		} else {
			return layoutTokenMonth, layout[0:1], layout[1:]
		}
	case 'D':
		if strings.HasPrefix(layout, "DD") {
			return layoutTokenDayTwo, layout[0:2], layout[2:]
		} else {
			return layoutTokenDay, layout[0:1], layout[1:]
		}
	case 'H':
		if strings.HasPrefix(layout, "HH") {
			return layoutTokenHourTwo, layout[0:2], layout[2:]
		} else {
			return layoutTokenHour, layout[0:1], layout[1:]
		}
	case 'h':
		if strings.HasPrefix(layout, "hh") {
			return layoutTokenHour12Two, layout[0:2], layout[2:]
		} else {
			return layoutTokenHour12, layout[0:1], layout[1:]
		}
	case 'm':
		if strings.HasPrefix(layout, "mm") {
			return layoutTokenMinuteTwo, layout[0:2], layout[2:]
		} else {
			return layoutTokenMinute, layout[0:1], layout[1:]
		}
	case 's':
		if strings.HasPrefix(layout, "ss") {
			return layoutTokenSecondTwo, layout[0:2], layout[2:]
		} else {
			return layoutTokenSecond, layout[0:1], layout[1:]
		}
	case 'S':
		if strings.HasPrefix(layout, "SSS") {
			return layoutTokenMillisecondThree, layout[0:3], layout[3:]
		} else if strings.HasPrefix(layout, "SS") {
			return layoutTokenMillisecondTwo, layout[0:2], layout[2:]
		} else {
			return layoutTokenMillisecond, layout[0:1], layout[1:]
		}
	case 'A':
		return layoutTokenPMUpper, layout[0:1], layout[1:]
	case 'a':
		return layoutTokenPMLower, layout[0:1], layout[1:]
	case 'Z':
		if strings.HasPrefix(layout, "ZZ") {
			return layoutTokenTZ, layout[0:2], layout[2:]
		} else {
			return layoutTokenTZColon, layout[0:1], layout[1:]
		}
	case '\\': // Escape next character
		if len(layout) >= 2 {
			return layoutTokenNone, layout[1:2], layout[2:]
		}
	}

	return layoutTokenNone, layout[0:1], layout[1:]
}

// Format returns a string of the time formatted by the layout from the parameter.
func (t Time) Format(layout string) string {
	buf := make([]byte, 0, 64)
	buf = t.formatByLayout(layout, buf)

	return string(buf)
}

// formatByLayout appends the string of the time formatted by the layout into the buffer, and
// returns the reference of the buffer.
func (t Time) formatByLayout(layout string, buf []byte) []byte {
	for len(layout) > 0 {
		token, str, suffix := nextLayoutToken(layout)
		layout = suffix
		if token == layoutTokenEnd {
			break
		} else if token == layoutTokenNone {
			buf = append(buf, str...)
			continue
		}

		switch token {
		case layoutTokenYear:
			buf = appendIntToBuffer(buf, t.Year(), 4)
		case layoutTokenYearTwo:
			buf = appendIntToBuffer(buf, t.Year()%100, 2)
		case layoutTokenMonth:
			buf = appendIntToBuffer(buf, int(t.Month()), 0)
		case layoutTokenMonthTwo:
			buf = appendIntToBuffer(buf, int(t.Month()), 2)
		case layoutTokenMonthAbbr:
			abbr := abbrMonthNames[t.Month()-1]
			buf = append(buf, abbr...)
		case layoutTokenMonthFull:
			name := fullMonthNames[t.Month()-1]
			buf = append(buf, name...)
		case layoutTokenDay:
			buf = appendIntToBuffer(buf, t.Day(), 1)
		case layoutTokenDayTwo:
			buf = appendIntToBuffer(buf, t.Day(), 2)
		case layoutTokenHour:
			buf = appendIntToBuffer(buf, t.Hour(), 1)
		case layoutTokenHourTwo:
			buf = appendIntToBuffer(buf, t.Hour(), 2)
		case layoutTokenHour12:
			buf = appendIntToBuffer(buf, t.Hour12(), 1)
		case layoutTokenHour12Two:
			buf = appendIntToBuffer(buf, t.Hour12(), 2)
		case layoutTokenMinute:
			buf = appendIntToBuffer(buf, t.Minute(), 1)
		case layoutTokenMinuteTwo:
			buf = appendIntToBuffer(buf, t.Minute(), 2)
		case layoutTokenSecond:
			buf = appendIntToBuffer(buf, t.Second(), 1)
		case layoutTokenSecondTwo:
			buf = appendIntToBuffer(buf, t.Second(), 2)
		case layoutTokenMillisecond:
			buf = appendIntToBuffer(buf, t.Millisecond()/100, 1)
		case layoutTokenMillisecondTwo:
			buf = appendIntToBuffer(buf, t.Millisecond()/10, 2)
		case layoutTokenMillisecondThree:
			buf = appendIntToBuffer(buf, t.Millisecond(), 3)
		case layoutTokenPMUpper:
			hour := t.Hour()
			if hour > 12 {
				buf = append(buf, "PM"...)
			} else {
				buf = append(buf, "AM"...)
			}
		case layoutTokenPMLower:
			hour := t.Hour()
			if hour > 12 {
				buf = append(buf, "pm"...)
			} else {
				buf = append(buf, "am"...)
			}
		case layoutTokenTZ, layoutTokenTZColon:
			_, offset := t.Zone()
			zone := offset / 60
			if zone < 0 {
				buf = append(buf, '-')
				zone = -zone
			} else {
				buf = append(buf, '+')
			}

			buf = appendIntToBuffer(buf, zone/60, 2)
			if token == layoutTokenTZColon {
				buf = append(buf, ':')
			}
			buf = appendIntToBuffer(buf, zone%60, 2)
		}
	}

	return buf
}
