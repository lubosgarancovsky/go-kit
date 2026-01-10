package go_kit

import (
	"fmt"
	"strings"
)

type Sort struct {
	Field     string
	Direction string
}

func BuildSort(sort string, fieldMap map[string]string) ([]Sort, error) {
	if sort == "" {
		return nil, nil
	}

	var fields []Sort
	parts := strings.Split(sort, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		var rawField, dir string
		switch {
		case strings.HasSuffix(part, ";"):
			rawField = strings.TrimSuffix(part, ";")
			dir = "DESC"
		case strings.HasSuffix(part, ":"):
			rawField = strings.TrimSuffix(part, ":")
			dir = "ASC"
		default:
			return nil, fmt.Errorf("invalid sort format: %s", part)
		}

		dbField, ok := fieldMap[rawField]
		if !ok {
			continue
		}

		fields = append(fields, Sort{
			Field:     dbField,
			Direction: dir,
		})
	}

	return fields, nil
}
