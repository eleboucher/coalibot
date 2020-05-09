package api

import (
	"errors"
	"fmt"
	"strings"
)

type StatusCode int

type Status struct {
	Code StatusCode
	Str  string
}

const (
	OkStatus            = StatusCode(200)
	BadRequest          = StatusCode(400)
	Unauthorized        = StatusCode(401)
	Forbidden           = StatusCode(403)
	NotFound            = StatusCode(404)
	UnprocessableEntity = StatusCode(422)
	TooManyRequest      = StatusCode(429)
	ServerError         = StatusCode(500)
)

var statusFormats = map[int]string{
	200: "OK",
	400: "Bad Request",
	401: "Unauthorized: {request}",
	403: "Forbidden: {request}",
	404: "Page or ressource not found: {request}",
	422: "Unprocessable Entity",
	429: "Too Many Request",
	500: "Server Error",
}

func processStatus(statusCode int, status, request string) Status {
	statusFormat := statusFormats[statusCode]
	if statusFormat == "" {
		statusFormats[statusCode] = status
		statusFormat = status
	}
	return Status{
		Code: StatusCode(statusCode),
		Str:  strings.Replace(statusFormat, "{request}", request, -1),
	}
}

func (s *Status) String() string {
	return fmt.Sprintf("%d: %s", s.Code, s.Str)
}

func (s *Status) Error() error {
	return errors.New(s.String())
}
