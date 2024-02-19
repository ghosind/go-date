package date

import (
	"strconv"
	"strings"
	"time"
)

// getTime tries to get a Time instance, and panics if the object is not a Time.
func getTime(t any) time.Time {
	if ut, ok := t.(Time); ok {
		return ut.Time
	} else if ut, ok := t.(time.Time); ok {
		return ut
	} else {
		panic(ErrNotTime)
	}
}

// appendIntToBuffer converts the integer value to a textual representation string, and padding
// with '0' if the length is less than the minimum width requirement.
func appendIntToBuffer(buf []byte, val int, width int) []byte {
	start := len(buf)

	// add negative sign
	if val < 0 {
		buf = append(buf, '-')
		start = start + 1
		width--
		val = -val
	}

	for val > 0 {
		n := val % 10
		buf = append(buf, byte(n)+'0')

		width--
		val /= 10
	}

	// padding with '0'
	for width > 0 {
		buf = append(buf, '0')
		width--
	}

	// reverse
	for i := len(buf); i > start; i-- {
		buf[i-1], buf[start] = buf[start], buf[i-1]
		start++
	}

	return buf
}

// lookup tries to find the index in the list that the element is the prefix of the provided
// string, and it is case-insensitive.
func lookup(list []string, value string) (int, string, error) {
	for i, v := range list {
		if len(value) < len(v) {
			continue
		}
		s := value[0:len(v)]
		if v == s || strings.EqualFold(v, s) {
			return i, value[len(v):], nil
		}
	}

	return -1, value, errParse
}

// readNum tries to read a width-length string for a fixed-length format, or a string that length
// greater than 0 and less or equal to width if it is not a fixed-length format. After getting the
// string, it will try to convert the string to an integer number.
func readNum(value string, width int, fixed bool) (int, string, error) {
	if len(value) == 0 || (fixed && len(value) < width) {
		return -1, value, errParse
	}

	var s, t string
	if fixed {
		s, t = value[0:width], value[width:]
	} else {
		i := 0
		for ; i < width && i < len(value); i++ {
			if value[i] < '0' || value[i] > '9' {
				break
			}
		}
		s, t = value[0:i], value[i:]
	}

	num, err := strconv.Atoi(s)
	if err != nil {
		return -1, value, errParse
	}

	value = t

	return num, value, nil
}
