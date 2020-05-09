package api42

import (
	"bytes"
	"fmt"
	"strconv"
)

func buildFilter(subject, filter string, buffer *bytes.Buffer) {
	switch filter {
	case defaultRequest:
		buffer.WriteString(subject)
	default:
		buffer.WriteString(filter)
	}
}

func buildValue(value interface{}, filter string, buffer *bytes.Buffer) error {
	switch val := value.(type) {
	case string:
		buffer.WriteString(val)
	case []byte:
		buffer.Write(val)
	case int:
		buffer.WriteString(strconv.Itoa(val))
	default:
		if filter != defaultRequest {
			return fmt.Errorf("Invalid id type '%t'", val)
		}
	}
	return nil
}

func buildUrl(subject, filter string, filterValue interface{}, args ...interface{}) (string, error) {
	var buffer bytes.Buffer
	buffer.WriteString(apiPrefix)
	buildFilter(subject, filter, &buffer)
	err := buildValue(filterValue, filter, &buffer)
	if err != nil {
		return "", err
	}
	for i, arg := range args {
		filter = defaultRequest
		switch a := arg.(type) {
		case string:
			switch i % 2 {
			case 0:
				filter = a
				buildFilter(subject, filter, &buffer)
				continue
			default:
				filterValue = a
				err = buildValue(filterValue, filter, &buffer)
				if err != nil {
					return "", err
				}
			}
		default:
			filterValue = a
			err = buildValue(filterValue, filter, &buffer)
			if err != nil {
				return "", err
			}
		}
		if filter == defaultRequest {
			break
		}
	}
	if filter != defaultRequest {
		buffer.WriteString(subject)
	}
	url := buffer.String()
	return url, nil
}
