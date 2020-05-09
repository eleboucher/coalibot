package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Method string

const (
	Get  = Method("GET")
	Post = Method("POST")
)

type Response struct {
	//	Public	//
	Status  Status
	Header  http.Header
	RawBody []byte
	Body    map[string]interface{}
}

func (c *ApiClient) requestBack(method Method, request string, data interface{}) (*Response, error) {
	url := fmt.Sprintf("%s/%s", c.ApiURL, request)
	preReq := newPreRequest(method, url, c.header)
	req, err := preReq.prepareRequest(data)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	rawBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &Response{
		Status:  processStatus(resp.StatusCode, resp.Status, request),
		Header:  resp.Header,
		RawBody: rawBody,
	}
	if json.Valid(ret.RawBody) {
		json.Unmarshal(ret.RawBody, &ret.Body)
	}
	return ret, nil
}

func (c *ApiClient) Request(method Method, request string, data interface{}) (*Response, error) {
	if err := c.wait(); err != nil {
		return nil, err
	}
	defer c.q.next()
	ret, err := c.requestBack(method, request, data)
	if err != nil {
		return nil, err
	}
	if ret.Status.Code == TooManyRequest {
		time.Sleep(time.Second * 2)
		return c.requestBack(method, request, data)
	}
	c.q.setLimits(ret.Header)
	return ret, nil
}

func (c *ApiClient) Post(request string, data interface{}) (*Response, error) {
	return c.Request(Post, request, data)
}

func (c *ApiClient) Get(request string, data interface{}) (*Response, error) {
	return c.Request(Get, request, data)
}
