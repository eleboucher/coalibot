package api42

type RequestFilters map[string]interface{}
type RequestCustoms map[string]interface{}
type RequestRanges map[string][2]interface{}

type RequestParameter struct {
	Filters       RequestFilters
	Ranges        RequestRanges
	Customs       RequestCustoms
	Sort          string
	Page, PerPage int
}

func NewParameter(args ...interface{}) *RequestParameter {
	ret := RequestParameter{make(RequestFilters), make(RequestRanges), make(RequestCustoms), "", 0, 0}
	intCount := 0
	msCount := 0
	for _, arg := range args {
		switch a := arg.(type) {
		case map[string]interface{}:
			switch msCount % 2 {
			case 0:
				ret.AddFilters(RequestFilters(a))
			default:
				ret.AddCustoms(RequestCustoms(a))
			}
			msCount += 1
		case RequestFilters:
			ret.AddFilters(a)
		case RequestCustoms:
			ret.AddCustoms(a)
		case map[string][2]interface{}:
			ret.AddRanges(RequestRanges(a))
		case RequestRanges:
			ret.AddRanges(a)
		case string:
			ret.Sort = a
		case int:
			switch intCount % 2 {
			case 0:
				ret.Page = a
			default:
				ret.PerPage = a
			}
			intCount += 1
		default:
		}
	}
	return &ret
}

func (r *RequestParameter) DeleteFilter(key string) {
	delete(r.Filters, key)
}

func (r *RequestParameter) AddFilter(key string, value interface{}) {
	r.Filters[key] = value
}

func (r *RequestParameter) AddFilters(filters RequestFilters) {
	for key, value := range filters {
		r.AddFilter(key, value)
	}
}

func (r *RequestParameter) DeleteRange(key string) {
	delete(r.Ranges, key)
}

func (r *RequestParameter) AddRange(key string, min, max interface{}) {
	array := [2]interface{}{min, max}
	r.Ranges[key] = array
}

func (r *RequestParameter) AddRanges(ranges RequestRanges) {
	for key, value := range ranges {
		r.AddRange(key, value[0], value[1])
	}
}

func (r *RequestParameter) DeleteCustom(key string) {
	delete(r.Customs, key)
}

func (r *RequestParameter) AddCustom(key string, value interface{}) {
	r.Customs[key] = value
}

func (r *RequestParameter) AddCustoms(ranges RequestCustoms) {
	for key, value := range ranges {
		r.AddCustom(key, value)
	}
}
