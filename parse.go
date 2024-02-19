package date

import (
	"errors"
	"time"
)

var errParse error = errors.New("parse error") // a temporary error for parse

// Parse parses a formatted string with the layout and returns the time value it represents.
func Parse(layout, value string) (Time, error) {
	return parse(layout, value, time.Local)
}

// ParseInLocation parses a formatted string with the layout and the given location, and returns
// the time value it represents.
func ParseInLocation(layout, value string, loc *time.Location) (Time, error) {
	return parse(layout, value, loc)
}

// ParseInLocationName tries to load the location with the given name, parses a formatted string
// with the layout and the location, and returns the time value it represents.
func ParseInLocationName(layout, value, name string) (Time, error) {
	loc, err := time.LoadLocation(name)
	if err != nil {
		return Time{}, err
	}

	return parse(layout, value, loc)
}

func parse(layout, value string, loc *time.Location) (Time, error) {
	oLayout, oValue := layout, value
	am := false
	pm := false
	var str string
	var err error

	var (
		year     int
		month    int = 1
		day      int = 1
		hour     int
		min      int
		sec      int
		nsec     int
		tzOffset int = -1
	)

	for {
		token, s, suffix := nextLayoutToken(layout)
		if token == layoutTokenEnd {
			break
		}
		layout = suffix

		switch token {
		case layoutTokenYearLong:
			year, value, err = readNum(value, 4, true)
		case layoutTokenYear:
			year, value, err = readNum(value, 2, true)
			if err != nil {
				break
			}
			if year > 69 {
				year += 1900
			} else {
				year += 2000
			}
		case layoutTokenMonth:
			month, value, err = readNum(value, 2, false)
		case layoutTokenMonthLong:
			month, value, err = readNum(value, 2, true)
		case layoutTokenMonthAbbr:
			month, value, err = lookup(abbrMonthNames, value)
			if err != nil {
				break
			}
			month++
		case layoutTokenMonthFull:
			month, value, err = lookup(fullMonthNames, value)
			if err != nil {
				break
			}
			month++
		case layoutTokenDay:
			day, value, err = readNum(value, 2, false)
		case layoutTokenDayLong:
			day, value, err = readNum(value, 2, true)
		case layoutTokenHour:
			hour, value, err = readNum(value, 2, false)
		case layoutTokenHourLong:
			hour, value, err = readNum(value, 2, true)
		case layoutTokenHour12:
			hour, value, err = readNum(value, 2, false)
		case layoutTokenHour12Long:
			hour, value, err = readNum(value, 2, true)
		case layoutTokenMinute:
			min, value, err = readNum(value, 2, false)
		case layoutTokenMinuteLong:
			min, value, err = readNum(value, 2, true)
		case layoutTokenSecond:
			sec, value, err = readNum(value, 2, false)
		case layoutTokenSecondLong:
			sec, value, err = readNum(value, 2, true)
		case layoutTokenMillisecondHundred:
			nsec, value, err = readNum(value, 1, true)
			if err != nil {
				break
			}
			nsec *= 100
		case layoutTokenMillisecondTen:
			nsec, value, err = readNum(value, 2, true)
			if err != nil {
				break
			}
			nsec *= 10
		case layoutTokenMillisecond:
			nsec, value, err = readNum(value, 3, true)
		case layoutTokenPMUpper:
			if len(value) < 2 {
				return Time{}, newParseError(oLayout, oValue, s, value)
			}
			str, value = value[0:2], value[2:]
			switch str {
			case "AM":
				am = true
			case "PM":
				pm = true
			default:
				return Time{}, newParseError(oLayout, oValue, s, str)
			}
		case layoutTokenPMLower:
			if len(value) < 2 {
				return Time{}, newParseError(oLayout, oValue, s, value)
			}
			str, value = value[0:2], value[2:]
			switch str {
			case "am":
				am = true
			case "pm":
				pm = true
			default:
				return Time{}, newParseError(oLayout, oValue, s, str)
			}
		case layoutTokenTZ:
			if len(value) < 5 {
				return Time{}, newParseError(oLayout, oValue, s, value)
			}
			var tzHr, tzMm int
			tzHr, _, err = readNum(value[1:3], 2, true)
			if err == nil {
				tzMm, _, err = readNum(value[3:5], 2, true)
			}
			tzOffset = tzHr*60 + tzMm
			switch value[0] {
			case '+':
			case '-':
				tzOffset = -tzOffset
			default:
				return Time{}, newParseError(oLayout, oValue, s, value[0:5])
			}

			value = value[5:]
		case layoutTokenTZColon:
			if len(value) < 6 {
				return Time{}, newParseError(oLayout, oValue, s, value)
			}
			var tzHr, tzMm int
			tzHr, _, err = readNum(value[1:3], 2, true)
			if err == nil {
				tzMm, _, err = readNum(value[4:6], 2, true)
			}
			tzOffset = tzHr*60 + tzMm
			switch value[0] {
			case '+':
			case '-':
				tzOffset = -tzOffset
			default:
				return Time{}, newParseError(oLayout, oValue, s, value[0:6])
			}

			value = value[6:]
		case layoutTokenNone:
			if len(value) < len(s) {
				return Time{}, newParseError(oLayout, oValue, s, value)
			}
			str, value = value[0:len(s)], value[len(s):]
			if s != str {
				return Time{}, newParseError(oLayout, oValue, s, str)
			}
		}

		if err != nil {
			return Time{}, newParseError(oLayout, oValue, s, value)
		}
	}

	if pm && hour < 12 {
		hour += 12
	} else if am && hour == 12 {
		hour = 0
	}

	nsec *= int(time.Millisecond)

	if tzOffset == -1 {
		return Date(year, time.Month(month), day, hour, min, sec, nsec, loc), nil
	} else {
		tm := Date(year, time.Month(month), day, hour, min, sec, nsec, time.UTC)
		tm.Time = tm.Add(time.Duration(tzOffset) * time.Minute)

		return tm, nil
	}
}
