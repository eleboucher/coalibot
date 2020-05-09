package api42

import (
	"time"
)

type Subnotion42 struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Slug      string     `json:"slug"`
	CreatedAt *time.Time `json:"created_at"`
	client    *Client42
}

func (c *Client42) backGetSubnotions(params *RequestParameter, directFilter string, value interface{}) ([]Subnotion42, error) {
	var subnotionArray []Subnotion42
	url, err := buildUrl(subnotionRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &subnotionArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(subnotionArray); i++ {
		subnotionArray[i].client = c
	}
	return subnotionArray, nil
}

// Returns the Subnotion specified by the id.
func (c *Client42) GetSubnotion(id interface{}) (*Subnotion42, error) {
	var subnotion *Subnotion42
	url, err := buildUrl(subnotionRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &subnotion)
	if err != nil {
		return nil, err
	}
	subnotion.client = c
	return subnotion, nil
}

// Returns all the Subnotions.
func (c *Client42) GetSubnotions(params *RequestParameter) ([]Subnotion42, error) {
	return c.backGetSubnotions(params, defaultRequest, nil)
}

// By Notion

// Returns all the Subnotions of the given Notion.
func (c *Client42) GetSubnotionsByNotion(id interface{}, params *RequestParameter) ([]Subnotion42, error) {
	return c.backGetSubnotions(params, notionRequest, id)
}

// Returns all the Subnotions of the given Notion.
func (c *Notion42) GetSubnotions(params *RequestParameter) ([]Subnotion42, error) {
	return c.client.GetSubnotionsByNotion(c.ID, params)
}
