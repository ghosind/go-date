package date

import (
	"strings"
)

const (
	// layoutTokenEnd indicates the layout is empty.
	layoutTokenEnd int = iota
	// layoutTokenNone is not a token.
	layoutTokenNone
	// layoutTokenYearLong is the year.
	layoutTokenYearLong
	// layoutTokenYear is the two-digits year.
	layoutTokenYear
	// layoutTokenMonth is the month beginning at 1.
	layoutTokenMonth
	// layoutTokenMonthLong is the two-digits month.
	layoutTokenMonthLong
	// layoutTokenMonthAbbr is the abbreviated month name.
	layoutTokenMonthAbbr
	// layoutTokenMonthFull is the full month name.
	layoutTokenMonthFull
	// layoutTokenDay is the day beginning at 1.
	layoutTokenDay
	// layoutTokenDayLong is the two-digits day.
	layoutTokenDayLong
	// layoutTokenDayOfWeek is the day of week that beginning at 0 (Sunday).
	layoutTokenDayOfWeek
	// layoutTokenDayOfWeek is the abbreviated name of the day of week.
	layoutTokenDayOfWeekAbbr
	// layoutTokenDayOfWeekFull is the name of the day of week.
	layoutTokenDayOfWeekFull
	// layoutTokenHour is the 24-hour clock hour that beginning at 1.
	layoutTokenHour
	// layoutTokenHourLong is the two-digits, 24-hour clock hour.
	layoutTokenHourLong
	// layoutTokenHour12 is the 12-hour clock hour that beginning at 1.
	layoutTokenHour12
	// layoutTokenHour12Long is the two-digits, 12-hour clock hour.
	layoutTokenHour12Long
	// layoutTokenMinute is the minute beginning at 1.
	layoutTokenMinute
	// layoutTokenMinuteLong is the two-digits minute.
	layoutTokenMinuteLong
	// layoutTokenSecond is the second beginning at 1.
	layoutTokenSecond
	// layoutTokenSecondLong is the two-digits second.
	layoutTokenSecondLong
	// layoutTokenMillisecondHundred is the one-digit millisecond (hundreds of milliseconds).
	layoutTokenMillisecondHundred
	// layoutTokenMillisecondTen is the two-digits millisecond (tens of milliseconds).
	layoutTokenMillisecondTen
	// layoutTokenMillisecondThree is the three-digits millisecond.
	layoutTokenMillisecond
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

var abbrWeekdayNames = []string{
	"Sun",
	"Mon",
	"Tue",
	"Wed",
	"Thu",
	"Fri",
	"Sat",
}

var fullWeekdayNames = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

// nextLayoutToken gets the next token in the layout, and return
func nextLayoutToken(layout string) (int, string, string) {
	if len(layout) == 0 {
		return layoutTokenEnd, "", ""
	}

	switch layout[0] {
	case 'Y':
		if strings.HasPrefix(layout, "YYYY") {
			return layoutTokenYearLong, layout[0:4], layout[4:]
		} else if strings.HasPrefix(layout, "YY") {
			return layoutTokenYear, layout[0:2], layout[2:]
		}
	case 'M':
		if strings.HasPrefix(layout, "Monday") {
			return layoutTokenDayOfWeekFull, layout[0:6], layout[6:]
		} else if strings.HasPrefix(layout, "Mon") {
			return layoutTokenDayOfWeekAbbr, layout[0:3], layout[3:]
		}

		if strings.HasPrefix(layout, "MMMM") {
			return layoutTokenMonthFull, layout[0:4], layout[4:]
		} else if strings.HasPrefix(layout, "MMM") {
			return layoutTokenMonthAbbr, layout[0:3], layout[3:]
		} else if strings.HasPrefix(layout, "MM") {
			return layoutTokenMonthLong, layout[0:2], layout[2:]
		} else {
			return layoutTokenMonth, layout[0:1], layout[1:]
		}
	case 'J':
		if strings.HasPrefix(layout, "Jan") {
			if strings.HasPrefix(layout, "January") {
				return layoutTokenMonthFull, layout[0:7], layout[7:]
			} else {
				return layoutTokenMonthAbbr, layout[0:3], layout[3:]
			}
		}
	case 'D':
		if strings.HasPrefix(layout, "DD") {
			return layoutTokenDayLong, layout[0:2], layout[2:]
		} else {
			return layoutTokenDay, layout[0:1], layout[1:]
		}
	case 'd':
		if strings.HasPrefix(layout, "dddd") {
			return layoutTokenDayOfWeekFull, layout[0:4], layout[4:]
		} else if strings.HasPrefix(layout, "ddd") {
			return layoutTokenDayOfWeekAbbr, layout[0:3], layout[3:]
		} else {
			return layoutTokenDayOfWeek, layout[0:1], layout[1:]
		}
	case 'H':
		if strings.HasPrefix(layout, "HH") {
			return layoutTokenHourLong, layout[0:2], layout[2:]
		} else {
			return layoutTokenHour, layout[0:1], layout[1:]
		}
	case 'h':
		if strings.HasPrefix(layout, "hh") {
			return layoutTokenHour12Long, layout[0:2], layout[2:]
		} else {
			return layoutTokenHour12, layout[0:1], layout[1:]
		}
	case 'm':
		if strings.HasPrefix(layout, "mm") {
			return layoutTokenMinuteLong, layout[0:2], layout[2:]
		} else {
			return layoutTokenMinute, layout[0:1], layout[1:]
		}
	case 's':
		if strings.HasPrefix(layout, "ss") {
			return layoutTokenSecondLong, layout[0:2], layout[2:]
		} else {
			return layoutTokenSecond, layout[0:1], layout[1:]
		}
	case 'S':
		if strings.HasPrefix(layout, "SSS") {
			return layoutTokenMillisecond, layout[0:3], layout[3:]
		} else if strings.HasPrefix(layout, "SS") {
			return layoutTokenMillisecondTen, layout[0:2], layout[2:]
		} else {
			return layoutTokenMillisecondHundred, layout[0:1], layout[1:]
		}
	case 'A':
		return layoutTokenPMUpper, layout[0:1], layout[1:]
	case 'a':
		return layoutTokenPMLower, layout[0:1], layout[1:]
	case 'P':
		if strings.HasPrefix(layout, "PM") {
			return layoutTokenPMUpper, layout[0:2], layout[2:]
		}
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
	case '0':
		if len(layout) < 2 {
			break
		}
		token := layoutTokenNone
		switch layout[1] {
		case '1':
			token = layoutTokenMonthLong
		case '2':
			token = layoutTokenDayLong
		case '3':
			token = layoutTokenHour12Long
		case '4':
			token = layoutTokenMinuteLong
		case '5':
			token = layoutTokenSecondLong
		case '6':
			token = layoutTokenYear
		}
		if token != layoutTokenNone {
			return token, layout[0:2], layout[2:]
		}
	case '1':
		if len(layout) >= 2 && layout[1] == '5' {
			return layoutTokenHourLong, layout[0:2], layout[2:]
		}
		return layoutTokenMonth, layout[0:1], layout[1:]
	case '2':
		if strings.HasPrefix(layout, "2006") {
			return layoutTokenYearLong, layout[0:4], layout[4:]
		}
		return layoutTokenDay, layout[0:1], layout[1:]
	case '3':
		return layoutTokenHour12, layout[0:1], layout[1:]
	case '4':
		return layoutTokenMinute, layout[0:1], layout[1:]
	case '5':
		return layoutTokenSecond, layout[0:1], layout[1:]
	}

	return layoutTokenNone, layout[0:1], layout[1:]
}

// AppendFormat is like Format but appends the textual representation to b and returns the extended
// buffer.
func (t Time) AppendFormat(b []byte, layout string) []byte {
	buf := t.formatByLayout(layout, b)
	return buf
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
	for {
		token, str, suffix := nextLayoutToken(layout)
		layout = suffix
		if token == layoutTokenEnd {
			break
		} else if token == layoutTokenNone {
			buf = append(buf, str...)
			continue
		}

		switch token {
		case layoutTokenYearLong:
			buf = appendIntToBuffer(buf, t.Year(), 4)
		case layoutTokenYear:
			buf = appendIntToBuffer(buf, t.Year()%100, 2)
		case layoutTokenMonth:
			buf = appendIntToBuffer(buf, int(t.Month()), 0)
		case layoutTokenMonthLong:
			buf = appendIntToBuffer(buf, int(t.Month()), 2)
		case layoutTokenMonthAbbr:
			abbr := abbrMonthNames[t.Month()-1]
			buf = append(buf, abbr...)
		case layoutTokenMonthFull:
			name := fullMonthNames[t.Month()-1]
			buf = append(buf, name...)
		case layoutTokenDay:
			buf = appendIntToBuffer(buf, t.Day(), 1)
		case layoutTokenDayLong:
			buf = appendIntToBuffer(buf, t.Day(), 2)
		case layoutTokenDayOfWeek:
			buf = appendIntToBuffer(buf, int(t.Weekday()), 1)
		case layoutTokenDayOfWeekAbbr:
			abbr := abbrWeekdayNames[t.Weekday()]
			buf = append(buf, abbr...)
		case layoutTokenDayOfWeekFull:
			name := fullWeekdayNames[t.Weekday()]
			buf = append(buf, name...)
		case layoutTokenHour:
			buf = appendIntToBuffer(buf, t.Hour(), 1)
		case layoutTokenHourLong:
			buf = appendIntToBuffer(buf, t.Hour(), 2)
		case layoutTokenHour12:
			buf = appendIntToBuffer(buf, t.Hour12(), 1)
		case layoutTokenHour12Long:
			buf = appendIntToBuffer(buf, t.Hour12(), 2)
		case layoutTokenMinute:
			buf = appendIntToBuffer(buf, t.Minute(), 1)
		case layoutTokenMinuteLong:
			buf = appendIntToBuffer(buf, t.Minute(), 2)
		case layoutTokenSecond:
			buf = appendIntToBuffer(buf, t.Second(), 1)
		case layoutTokenSecondLong:
			buf = appendIntToBuffer(buf, t.Second(), 2)
		case layoutTokenMillisecondHundred:
			buf = appendIntToBuffer(buf, t.Millisecond()/100, 1)
		case layoutTokenMillisecondTen:
			buf = appendIntToBuffer(buf, t.Millisecond()/10, 2)
		case layoutTokenMillisecond:
			buf = appendIntToBuffer(buf, t.Millisecond(), 3)
		case layoutTokenPMUpper:
			hour := t.Hour()
			if hour >= 12 {
				buf = append(buf, "PM"...)
			} else {
				buf = append(buf, "AM"...)
			}
		case layoutTokenPMLower:
			hour := t.Hour()
			if hour >= 12 {
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
