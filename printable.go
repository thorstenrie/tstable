package lpstr

import (
	"errors"
	"strings"
	"unicode"
)

func Printable(a string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, a)
}

func IsPrintable(a []string) (bool, error) {
	if (a == nil) || (len(a) == 0) {
		return false, errors.New("empty array")
	}
	for _, s := range a {
		if s != Printable(s) {
			return false, nil
		}
	}
	return true, nil
}
