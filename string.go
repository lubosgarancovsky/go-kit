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
