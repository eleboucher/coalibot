package api

import (
	"encoding/json"
	"fmt"
	"time"
)

type Token struct {
	createdAt time.Time
	c         chan bool
	expiresIn time.Duration

	Client      *ApiClient
	Token, Type string
}

func (t *Token) Init() error {
	data := map[string]interface{}{
		"grant_type":    "client_credentials",
		"client_id":     t.Client.Uid,
		"client_secret": t.Client.Secret,
	}
	resp, err := t.Client.requestBack(Post, "/oauth/token", data)
	if err != nil {
		return err
	}
	if resp.Status.Code != OkStatus {
		return resp.Status.Error()
	}
	if !isToken(resp) {
		return fmt.Errorf("Could not read content to find token")
	}
	t.Token = resp.Body["access_token"].(string)
	t.Type = resp.Body["token_type"].(string)
	t.Client.header = make(map[string]string)
	t.Client.header["Authorization"] = fmt.Sprintf("%s %s", t.Type, t.Token)
	return nil
}

func (t *Token) SetTokenInfo() {
	resp, err := t.Client.requestBack(Get, "/oauth/token/info", nil)
	if err != nil {
		return
	}
	if !json.Valid(resp.RawBody) {
		return
	}
	if resp.Body["created_at"] == nil || resp.Body["expires_in_seconds"] == nil {
		return
	}
	t.createdAt = time.Unix(int64(resp.Body["created_at"].(float64)), 0)
	t.expiresIn, _ = time.ParseDuration(fmt.Sprintf("%ds", int(resp.Body["expires_in_seconds"].(float64))))
}

func (t *Token) Reinit() (*Token, error) {
	t.Kill()
	if err := t.Init(); err != nil {
		return nil, err
	}
	t.Client.Token = t
	t.SetTokenInfo()
	go t.startTicker()
	return t, nil
}

func NewToken(client *ApiClient) (*Token, error) {
	ret := Token{
		createdAt: time.Now(),
		expiresIn: 30 * time.Minute,
		c:         make(chan bool),
		Client:    client,
	}
	if err := ret.Init(); err != nil {
		return nil, err
	}
	return &ret, nil
}

func isToken(resp *Response) bool {
	if !json.Valid(resp.RawBody) {
		return false
	}
	tokType, token := false, false
	for key, _ := range resp.Body {
		switch key {
		case "token_type":
			tokType = true
		case "access_token":
			token = true
		default:
		}
	}
	return tokType && token
}
