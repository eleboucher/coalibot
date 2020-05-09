package api42

import (
	"fmt"
)

func applyFilter(filterName, key string, value interface{}, data *map[string]interface{}) {
	builtKey := fmt.Sprintf("%s[%s]", filterName, key)
	(*data)[builtKey] = value
}

func customFilter(filterName string, filter map[string]interface{}, data *map[string]interface{}) {
	for key, value := range filter {
		applyFilter(filterName, key, value, data)
	}
}

func prepareData(params *RequestParameter) map[string]interface{} {
	ret := make(map[string]interface{})
	for key, value := range params.Filters {
		applyFilter("filter", key, value, &ret)
	}
	for key, value := range params.Customs {
		switch val := value.(type) {
		case map[string]interface{}:
			customFilter(key, val, &ret)
		default:
			ret[key] = value
		}
	}
	for key, value := range params.Ranges {
		applyFilter("range", key, value, &ret)
	}
	if params.Sort != "" {
		ret["sort"] = params.Sort
	}
	if params.Page != 0 {
		ret["page"] = params.Page
	}
	if params.PerPage > 0 {
		ret["per_page"] = params.PerPage
	}
	return ret
}
