package common

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

type validationRule struct {
	Required bool
	Regexp   *regexp.Regexp
	Min      int64
	Max      int64
}

func ValidateStruct(s interface{}) error {
	valueOf := reflect.ValueOf(s)

	if valueOf.Kind() != reflect.Struct {
		return errors.New("not a struct")
	}

	for i := 0; i < valueOf.NumField(); i++ {
		field := valueOf.Type().Field(i)

		// Get validation rule from struct tag
		rule, err := parseTag(field.Tag.Get("validate"))
		if err != nil {
			return fmt.Errorf("%s: %v", field.Name, err)
		}

		// Get field value and perform validation
		fieldValue := valueOf.Field(i).Interface()
		if err := validateField(fieldValue, rule, field.Name); err != nil {
			return err
		}
	}

	return nil
}

func ContainsString(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func parseTag(tag string) (validationRule, error) {
	rule := validationRule{}

	for _, s := range strings.Split(tag, ",") {
		switch {
		case s == "required":
			rule.Required = true
		case strings.HasPrefix(s, "regexp="):
			r, err := regexp.Compile(strings.TrimPrefix(s, "regexp="))
			if err != nil {
				return rule, fmt.Errorf("invalid regexp: %v", err)
			}
			rule.Regexp = r
		case strings.HasPrefix(s, "min="):
			i, err := strconv.ParseInt(strings.TrimPrefix(s, "min="), 10, 64)
			if err != nil {
				return rule, fmt.Errorf("invalid min: %v", err)
			}
			rule.Min = i
		case strings.HasPrefix(s, "max="):
			i, err := strconv.ParseInt(strings.TrimPrefix(s, "max="), 10, 64)
			if err != nil {
				return rule, fmt.Errorf("invalid max: %v", err)
			}
			rule.Max = i
		default:
			return rule, nil
		}
	}

	return rule, nil
}

func validateField(value interface{}, rule validationRule, fieldName string) error {
	if rule.Required && isEmpty(value) {
		return fmt.Errorf("%s: required", fieldName)
	}

	if rule.Regexp != nil {
		strValue, ok := value.(string)
		if ok && !rule.Regexp.MatchString(strValue) {
			return fmt.Errorf("%s: invalid format", fieldName)
		}
	}

	intValue, ok := toInt(value)
	if ok && rule.Min != 0 && rule.Max != 0 {
		if intValue < rule.Min || intValue > rule.Max {
			return fmt.Errorf("%s: must be between %d and %d", fieldName, rule.Min, rule.Max)
		}
	} else if ok && rule.Min != 0 {
		if intValue < rule.Min {
			return fmt.Errorf("%s: minimum value is %d", fieldName, rule.Min)
		}
	} else if ok && rule.Max != 0 {
		if intValue > rule.Max {
			return fmt.Errorf("%s: maximum value is %d", fieldName, rule.Max)
		}
	}

	return nil
}
func isEmpty(value interface{}) bool {
	if value == nil {
		return true
	}

	switch v := value.(type) {
	case string:
		return len(v) == 0 || utf8.RuneCountInString(v) == 0
	case []byte:
		return len(v) == 0 || utf8.RuneCountInString(string(v)) == 0
	default:
		rv := reflect.ValueOf(value)
		switch rv.Kind() {
		case reflect.String:
			strVal := rv.String()
			return len(strVal) == 0 || utf8.RuneCountInString(strVal) == 0
		case reflect.Slice:
			if rv.Type().Elem().Kind() == reflect.Uint8 {
				strVal := string(rv.Bytes())
				return len(strVal) == 0 || utf8.RuneCountInString(strVal) == 0
			}
		}

		switch v := value.(type) {
		case map[interface{}]interface{}:
			return len(v) == 0
		case map[string]interface{}:
			return len(v) == 0
		case []interface{}:
			return len(v) == 0
		case []string:
			return len(v) == 0
		case []int:
			return len(v) == 0
		case []int64:
			return len(v) == 0
		case []float64:
			return len(v) == 0
		case *string:
			if v == nil {
				return true
			}
			return len([]rune(*v)) == 0
		default:
			return false
		}
	}
}

func toInt(value interface{}) (int64, bool) {
	switch v := value.(type) {
	case int:
		return int64(v), true
	case int64:
		return v, true
	case float64:
		return int64(v), true
	case string:
		i, err := strconv.ParseInt(v, 10, 64)
		return i, err == nil
	case *string:
		if v == nil {
			return 0, false
		}
		return toInt(*v)
	default:
		return 0, false
	}
}
