package api42

import (
	"encoding/json"
	"fmt"
	"gitlab.com/clafoutis/api"
	"reflect"
	"strconv"
)

const (
	// 42 API URL
	api42Url  = "https://api.intra.42.fr/"
	apiPrefix = "/v2"

	// Supported Requests
	achievementRequest     = "/achievements/"
	blocRequest            = "/blocs/"
	campusRequest          = "/campus/"
	campusUserRequest      = "/campus_users/"
	coalitionRequest       = "/coalitions/"
	coalitionUserRequest   = "/coalitions_users/"
	cursusRequest          = "/cursus/"
	cursusUserRequest      = "/cursus_users/"
	defaultRequest         = "/"
	eventRequest           = "/events/"
	expertiseRequest       = "/expertises/"
	expertiseUserRequest   = "/expertises_users/"
	feedbackRequest        = "/feedbacks/"
	languageRequest        = "/languages/"
	languageUserRequest    = "/languages_users/"
	locationRequest        = "/locations/"
	messageRequest         = "/messages/"
	notionRequest          = "/notions/"
	partnershipRequest     = "/partnerships/"
	partnershipUserRequest = "/partnerships_users/"
	projectRequest         = "/projects/"
	projectSessionRequest  = "/project_sessions/"
	projectUserRequest     = "/projects_users/"
	scaleRequest           = "/scales/"
	scaleTeamRequest       = "/scale_teams/"
	skillRequest           = "/skills/"
	subnotionRequest       = "/subnotions/"
	tagRequest             = "/tags/"
	teamRequest            = "/teams/"
	teamUploadRequest      = "/teams_uploads/"
	topicRequest           = "/topics/"
	titleRequest           = "/titles/"
	titleUserRequest       = "/titles_users/"
	userRequest            = "/users/"
	voteRequest            = "/votes/"
)

type Client42 struct {
	client *api.ApiClient
}

func NewAPI(uid, secret string) (*Client42, error) {
	client, err := api.NewAPI(api42Url, uid, secret)
	if err != nil {
		return nil, err
	}
	return &Client42{client}, err
}

// Request main

func concatenateSlice(dst, src interface{}) {
	dv := reflect.ValueOf(dst)
	sv := reflect.ValueOf(src)
	if sv.Type() != dv.Type() {
		return
	}
	tmp := reflect.AppendSlice(dv.Elem(), sv.Elem())
	dv.Elem().Set(tmp)
}

func getPager(page *int) bool {
	if *page != -1 {
		return false
	}
	*page = 1
	return true
}

func (c *Client42) Request(method api.Method, url string, params *RequestParameter, respHost interface{}) ([]byte, error) {
	var err error
	if params == nil {
		params = NewParameter()
	}
	pager := getPager(&params.Page)
	resp, err := c.client.Request(method, url, prepareData(params))
	if err != nil {
		return nil, err
	}
	ret, err := handleResp(resp, respHost)
	if err != nil {
		return nil, err
	}
	if pager {
		perPage, _ := strconv.Atoi(resp.Header["X-Per-Page"][0])
		total, _ := strconv.Atoi(resp.Header["X-Total"][0])
		pages := total/perPage + 1
		if total%perPage > 0 {
			pages += 1
		}
		for i := 1; i < pages; i++ {
			newHost := reflect.New(reflect.ValueOf(respHost).Elem().Type()).Interface()
			params.Page = i + 1
			_, err := c.Request(method, url, params, newHost)
			if err != nil {
				return nil, err
			}
			concatenateSlice(respHost, newHost)
		}
		return nil, nil
	}
	return ret, nil
}

func handleResp(resp *api.Response, respHost interface{}) ([]byte, error) {
	var err error
	if resp.Status.Code != api.OkStatus {
		return nil, resp.Status.Error()
	}
	if !json.Valid(resp.RawBody) {
		return nil, fmt.Errorf("Could not parse response")
	}
	if respHost != nil {
		err = json.Unmarshal(resp.RawBody, respHost)
	}
	return resp.RawBody, err
}

// Requests Front

func (c *Client42) Get(url string, params *RequestParameter, respHost interface{}) ([]byte, error) {
	return c.Request(api.Get, url, params, respHost)
}

func (c *Client42) Post(url string, params *RequestParameter, respHost interface{}) ([]byte, error) {
	return c.Request(api.Post, url, params, respHost)
}
