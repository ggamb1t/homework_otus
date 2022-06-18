package hw02unpackstring

import (
	"fmt"
	"strconv"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = fmt.Errorf("invalid string")

func Unpack(str string) (string, error) {
	if str == "" {
		return "", nil
	}

	if ok := utf8.ValidString(str); !ok {
		return "", ErrInvalidString
	}

	var (
		prevRune rune
		lastRune rune
		resBuf   = make([]byte, 0, 2*utf8.RuneCountInString(str))
	)

	for i, currRune := range str {

		if i == 0 { // first rune special check
			if unicode.IsDigit(currRune) {
				return "", ErrInvalidString
			}
			if utf8.RuneCountInString(str) == 1 {
				resBuf = utf8.AppendRune(resBuf, currRune) //nolint:typecheck
				return string(resBuf), nil
			}
			prevRune = currRune
			continue
		}

		if unicode.IsDigit(currRune) && unicode.IsDigit(prevRune) {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(currRune) {
			count, _ := strconv.Atoi(string(currRune))
			for i := 0; i < count; i++ {
				resBuf = utf8.AppendRune(resBuf, prevRune) //nolint:typecheck
			}
		} else if !unicode.IsDigit(prevRune) {
			{
				resBuf = utf8.AppendRune(resBuf, prevRune) //nolint:typecheck
			}
		}
		prevRune, lastRune = currRune, currRune
	}
	if !unicode.IsDigit(lastRune) {
		resBuf = utf8.AppendRune(resBuf, lastRune) //nolint:typecheck
	}
	return string(resBuf), nil
}
