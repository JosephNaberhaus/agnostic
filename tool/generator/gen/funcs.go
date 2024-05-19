package gen

import (
	"regexp"
	"strings"
)

func makePointer(str string) string {
	if strings.HasPrefix(str, "[]") {
		// Slices are already pointers
		return str
	}

	return "*" + str
}

func removeOptional(str string) string {
	optionalRegex := regexp.MustCompile(`Optional\[(.*)\]`)
	matches := optionalRegex.FindStringSubmatch(str)
	if matches == nil {
		// Already not an optional. Just return the original str.
		return str
	}

	return matches[1]
}

func title(str string) string {
	return strings.ToUpper(str[:1]) + str[1:]
}
