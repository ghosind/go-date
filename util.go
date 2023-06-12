package date

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
