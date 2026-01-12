package go_kit

import (
	"regexp"
	"strings"
)

func ToSnakeCase(str string) string {
	// Step 0: convert dashes to underscores
	str = strings.ReplaceAll(str, "-", "_")

	// Step 1: handle camelCase / PascalCase boundaries
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

	// Step 2: lowercase everything
	return strings.ToLower(snake)
}

func ToCamelCase(s string) string {
	if s == "" {
		return s
	}

	lowerUpper := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	acronymWord := regexp.MustCompile(`([A-Z]+)([A-Z][a-z])`)

	// Normalize separators
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.ReplaceAll(s, "-", " ")

	// Insert word boundaries
	s = acronymWord.ReplaceAllString(s, `$1 $2`)
	s = lowerUpper.ReplaceAllString(s, `$1 $2`)

	words := strings.Fields(s)
	if len(words) == 0 {
		return ""
	}

	for i := range words {
		words[i] = strings.ToLower(words[i])
	}

	for i := 1; i < len(words); i++ {
		words[i] = capitalize(words[i])
	}

	return strings.Join(words, "")
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
