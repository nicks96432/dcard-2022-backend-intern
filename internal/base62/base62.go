package base62

import (
	"errors"
	"strconv"
)

const (
	codes  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	length = len(codes)
)

// Encode the number into a base62 string.
//
// Since an uint64 is not bigger than 62 ^ 11,
// so a string of 11 characters is enough.
func Encode(number uint64) string {
	s := make([]rune, 11)

	for i := len(s) - 1; i >= 0; i, number = i-1, number/uint64(length) {
		s[i] = rune(codes[number%uint64(length)])
	}

	return string(s)
}

// Decode the base62 string into a number.
func Decode(str string) (uint64, error) {
	if len(str) > 11 {
		return 0, errors.New("string too long")
	}

	var ret uint64 = 0

	for i, a := range str {
		ret *= 62
		if '0' <= a && a <= '9' {
			ret += uint64(a) - '0'
		} else if 'a' <= a && a <= 'z' {
			ret += uint64(a) - 'a' + 10
		} else if 'A' <= a && a <= 'Z' {
			ret += uint64(a) - 'A' + 36
		} else {
			return 0, errors.New("invalid character at position " + strconv.Itoa(i) + " :" + string(a))
		}
	}

	return ret, nil
}
