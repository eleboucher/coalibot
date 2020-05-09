package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type preRequest struct {
	Method     Method
	URL        string
	DataReader io.Reader
	Header     map[string]string
}

func (r *preRequest) prepareRequest(data interface{}) (*http.Request, error) {
	var req *http.Request
	err := r.treatData(data)
	if err != nil {
		return nil, err
	}
	req, err = http.NewRequest(string(r.Method), r.URL, r.DataReader)
	if err != nil {
		return nil, err
	}
	for key, value := range r.Header {
		req.Header.Add(key, value)
	}
	return req, nil
}

func (r *preRequest) mapStringToR(data map[string]interface{}) error {
	var buffer bytes.Buffer
	wrote := false
	for key, value := range data {
		if wrote {
			buffer.WriteString("&")
		}
		wrote = true
		buffer.WriteString(key)
		buffer.WriteString("=")
		vStr, err := toString(value)
		if err != nil {
			return err
		}
		buffer.WriteString(vStr)
	}
	switch r.Method {
	case Get:
		r.URL = fmt.Sprintf("%s?%s", r.URL, buffer.String())
	default:
		r.DataReader = bytes.NewReader(buffer.Bytes())
	}
	return nil
}

func (r *preRequest) sToR(data interface{}) {
	var b []byte
	switch dt := data.(type) {
	case string:
		b = []byte(dt)
	default:
		b = dt.([]byte)
	}
	switch r.Method {
	case Get:
		r.URL = fmt.Sprintf("%s?%s", r.URL, string(b))
	default:
		r.DataReader = bytes.NewReader(b)
	}
}

func (r *preRequest) treatData(data interface{}) error {
	switch dt := data.(type) {
	case map[string]interface{}:
		if err := r.mapStringToR(dt); err != nil {
			return err
		}
	case string, []byte:
		r.sToR(dt)
	default:
		fmt.Errorf("Type of date '%t' is not supported.", data)
	}
	return nil
}

func newPreRequest(method Method, url string, header map[string]string) *preRequest {
	return &preRequest{
		DataReader: nil,
		Method:     method,
		URL:        url,
		Header:     header,
	}
}
