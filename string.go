package go_kit

import (
	"regexp"
	"strings"
	"unicode"
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

	// Step 1: split into words
	words := splitWords(s)
	if len(words) == 0 {
		return ""
	}

	// Step 2: normalize casing
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}

	// Step 3: build camelCase
	for i := 1; i < len(words); i++ {
		words[i] = capitalize(words[i])
	}

	return strings.Join(words, "")
}

func splitWords(s string) []string {
	var words []string
	var current []rune

	for i, r := range s {
		switch {
		case r == '_' || r == '-' || r == ' ':
			if len(current) > 0 {
				words = append(words, string(current))
				current = nil
			}
		case unicode.IsUpper(r):
			if i > 0 && len(current) > 0 {
				words = append(words, string(current))
				current = nil
			}
			current = append(current, r)
		default:
			current = append(current, r)
		}
	}

	if len(current) > 0 {
		words = append(words, string(current))
	}

	return words
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
