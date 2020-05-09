package api

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func sliceToString(slice interface{}) (string, error) {
	var builder strings.Builder
	sV := reflect.ValueOf(slice)
	wrote := false
	for i := 0; i < sV.Len(); i++ {
		if wrote {
			builder.WriteString(",")
		}
		wrote = true
		vStr, err := toString(sV.Index(i).Interface())
		if err != nil {
			return "", err
		}
		builder.WriteString(vStr)
	}
	return builder.String(), nil
}

func toString(value interface{}) (string, error) {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice, reflect.Array:
		return sliceToString(value)
	default:
	}
	switch val := value.(type) {
	case string:
		return val, nil
	case int:
		return strconv.Itoa(val), nil
	case int64:
		return strconv.FormatInt(val, 10), nil
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64), nil
	case bool:
		return strconv.FormatBool(val), nil
	case []byte:
		return string(val), nil
	case rune:
		return string(val), nil
	case byte:
		return string(val), nil
	case time.Time:
		return val.Format("2006-01-02T15:04:05.999999999Z"), nil
	default:
	}
	return "", fmt.Errorf("Type %v is not supported.", reflect.TypeOf(value))
}
