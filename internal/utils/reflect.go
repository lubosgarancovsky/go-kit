package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// GetFieldName Returns variable name to look for in an .env file
func GetFieldName(field reflect.StructField) string {
	tag := field.Tag.Get("field")
	if tag == "" {
		return strings.ToLower(field.Name)
	}

	if strings.Contains(tag, ",") {
		return strings.Split(tag, ",")[0]
	}

	return tag
}

// SetFieldValue sets a value to target field with corresponding type
func SetFieldValue(field reflect.Value, strVal string) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(strVal)
		break
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(strVal, 10, field.Type().Bits())
		if err != nil {
			return err
		}
		field.SetInt(i)
		break
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u, err := strconv.ParseUint(strVal, 10, field.Type().Bits())
		if err != nil {
			return err
		}
		field.SetUint(u)
		break
	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(strVal, field.Type().Bits())
		if err != nil {
			return err
		}
		field.SetFloat(f)
	case reflect.Bool:
		b, err := strconv.ParseBool(strVal)
		if err != nil {
			return err
		}
		field.SetBool(b)
		break
	default:
		return fmt.Errorf("unsupported field type: %s", field.Kind())
	}
	return nil
}
