package cfg

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"reflect"
	"strings"
)

func LoadEnv(target any, paths ...string) error {
	v := reflect.ValueOf(target)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("target must be pointer to struct")
	}

	v = v.Elem()  // struct value
	t := v.Type() // struct type (string, uint64...)

	cfgMap := LoadEnvFiles(paths)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// Skip unexported fields
		if !value.CanSet() {
			continue
		}

		defaultValue := field.Tag.Get("default")
		fieldName := GetFieldName(field)

		envValue, ok := cfgMap[fieldName]
		if !ok {
			envValue = defaultValue
		}

		if envValue != "" {
			if err := SetFieldValue(value, envValue); err != nil {
				return fmt.Errorf("failed to set field %s: %w", fieldName, err)
			}
		}

	}
	return nil
}

func LoadEnvFiles(paths []string) map[string]string {
	if len(paths) == 0 {
		return LoadEnvFile(".env")
	}

	result := map[string]string{}

	for _, p := range paths {
		scannedMap := LoadEnvFile(p)
		result = MergeMaps(result, scannedMap)
	}
	return result
}

func LoadEnvFile(path string) map[string]string {
	file, err := os.Open(path)
	if err != nil {
		return map[string]string{}
	}
	defer file.Close()
	return ScanEnvFile(file)
}

func ScanEnvFile(file fs.File) map[string]string {
	result := map[string]string{}
	scanner := bufio.NewScanner(file)
	lineNum := 1

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "#") {
			lineNum++
			continue
		}

		parts := strings.SplitN(line, "=", 2)

		if len(parts) == 1 {
			result[parts[0]] = ""
		}

		if len(parts) >= 2 {
			result[parts[0]] = parts[1]
		}

		lineNum++
	}

	return result
}
