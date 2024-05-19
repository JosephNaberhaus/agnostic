package gen

import (
	"regexp"
	"strings"
)

func removeTypePrefix(str string) string {
	return strings.ReplaceAll(str, "~", "")
}

func makePointer(str string) string {
	if strings.HasPrefix(str, "[]") {
		// For slices we might need to make the underlying type a pointer.
		return "[]" + makePointer(strings.TrimPrefix(str, "[]"))
	}

	if strings.HasPrefix(str, "~") {
		// Types are implemented by interfaces which are already pointers.
		return str
	}

	if strings.ToLower(str[:1]) == str[:1] {
		// Primitive types don't need to be pointers.
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
