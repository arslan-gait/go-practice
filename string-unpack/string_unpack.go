package myunpack

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

const (
	zeroASCII = 48
	backslash = 92
)

//UnpackString returns formatted string
func UnpackString(s string) (string, error) {
	var res strings.Builder
	isDigit := unicode.IsDigit
	for i := 0; i < len(s); i++ {
		if s[i] == backslash {
			i++
			res.WriteByte(s[i])
		} else if isDigit(rune(s[i])) {
			if i == 0 || (i > 0 && isDigit(rune(s[i-1])) && (i > 1 && s[i-2] != backslash)) {
				return "", ErrInvalidString
			}
			res.WriteString(strings.Repeat(string(s[i-1]), int(s[i]-zeroASCII-1)))
		} else {
			res.WriteByte(s[i])
		}
	}
	return res.String(), nil
}
