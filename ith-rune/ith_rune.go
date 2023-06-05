package ithrune

import "unicode/utf8"

// IthRune returns the i-th rune in s, i starting at 0.
// ok is false if s is not valid UTF-8, or if i is out of bounds
func IthRuneInString(s string, i int) (r rune, ok bool) {
	if i < 0 {
		return utf8.RuneError, false
	}

	for {
		r, size := utf8.DecodeRuneInString(s)
		if r == utf8.RuneError {
			return utf8.RuneError, false
		}
		if i == 0 {
			return r, true
		}
		i--
		s = s[size:]
	}
}

// IthRune returns the i-th rune in buf, i starting at 0.
// ok is false if s is not valid UTF-8, or if i is out of bounds
func IthRuneInBytes(buf []byte, i int) (r rune, ok bool) {
	if i < 0 {
		return utf8.RuneError, false
	}

	for {
		r, size := utf8.DecodeRune(buf)
		if r == utf8.RuneError {
			return utf8.RuneError, false
		}
		if i == 0 {
			return r, true
		}
		i--
		buf = buf[size:]
	}
}
