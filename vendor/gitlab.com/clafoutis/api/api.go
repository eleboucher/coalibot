package api

import (
	"net/http"
)

type ApiClient struct {
	//	Private	//
	client *http.Client
	header map[string]string
	q      *queue
	//	Public	//
	ApiURL, Uid, Secret string
	Token               *Token
}

func NewAPI(url, uid, secret string) (*ApiClient, error) {
	ret := &ApiClient{
		client: &http.Client{},
		Token:  nil,
		ApiURL: url,
		Uid:    uid,
		Secret: secret,
		q:      newQueue(),
	}
	if err := ret.NewToken(); err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *ApiClient) NewToken() error {
	var err error
	c.Token, err = NewToken(c)
	if err != nil {
		return err
	}
	c.Token.SetTokenInfo()
	go c.Token.startTicker()
	return nil
}

func (c *ApiClient) wait() error {
	c.q.waitTurn()
	for c.Token == nil {
	}
	if err := c.Token.Expired(); err != nil {
		return err
	}
	return nil
}
