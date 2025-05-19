package util

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func StartsWithLowerCase(s string) bool {
	if len(s) == 0 {
		return false
	}
	firstChar := rune(s[0])
	return unicode.IsLower(firstChar)
}

func IsTimeString(layout string, str ...string) bool {
	if len(str) == 0 {
		return false
	}

	for _, s := range str {
		if _, err := time.Parse(layout, s); err != nil {
			return false
		}
	}

	return true
}

func AbbreviateTxHash(txHash string, fromStartIndex, fromEndIndex int) string {
	var abbreviation string
	if len(txHash) < fromStartIndex+fromEndIndex {
		abbreviation = txHash
	} else {
		firstChars := txHash[:fromStartIndex]
		lastChars := txHash[len(txHash)-fromEndIndex:]
		abbreviation = fmt.Sprintf("%s...%s", firstChars, lastChars)
	}
	return abbreviation
}
