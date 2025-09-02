package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	runes := map[rune]bool{}
	repeatingRunes := strings.FieldsFunc(word, func(r rune) bool {
		switch r {
		case '-':
			return true
		case ' ':
			return true
		default:
			upperRune := unicode.ToUpper(r)
			_, exists := runes[upperRune]
			runes[upperRune] = true
			return !exists
		}
	})
	return len(repeatingRunes) == 0
}
