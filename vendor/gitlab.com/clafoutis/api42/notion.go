package api42

import (
	"time"
)

type Notion42 struct {
	ID         int           `json:"id"`
	Name       string        `json:"name"`
	Slug       string        `json:"slug"`
	CreatedAt  *time.Time    `json:"created_at"`
	Subnotions []Subnotion42 `json:"subnotions"`
	Tags       []Tag42       `json:"tags"`
	Cursus     []Cursus42    `json:"cursus"`
	client     *Client42
}

func (c *Client42) backGetNotions(params *RequestParameter, directFilter string, value interface{}) ([]Notion42, error) {
	var notionArray []Notion42
	url, err := buildUrl(notionRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &notionArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(notionArray); i++ {
		notionArray[i].client = c
	}
	return notionArray, nil
}

// Returns the Notion specified by the id.
func (c *Client42) GetNotion(id interface{}) (*Notion42, error) {
	var notion *Notion42
	url, err := buildUrl(notionRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &notion)
	if err != nil {
		return nil, err
	}
	notion.client = c
	return notion, nil
}

// Returns all the Notions.
func (c *Client42) GetNotions(params *RequestParameter) ([]Notion42, error) {
	return c.backGetNotions(params, defaultRequest, nil)
}

// By Cursus

// Returns all the Notions of the given Cursus.
func (c *Client42) GetNotionsByCursus(id interface{}, params *RequestParameter) ([]Notion42, error) {
	return c.backGetNotions(params, cursusRequest, id)
}

// Returns all the Notions of the given Cursus.
func (c *Cursus42) GetNotions(params *RequestParameter) ([]Notion42, error) {
	return c.client.GetNotionsByCursus(c.ID, params)
}

// By Tag

// Returns all the Notions of the given Tag.
func (c *Client42) GetNotionsByTag(id interface{}, params *RequestParameter) ([]Notion42, error) {
	return c.backGetNotions(params, tagRequest, id)
}

// Returns all the Notions of the given Tag.
func (c *Tag42) GetNotions(params *RequestParameter) ([]Notion42, error) {
	return c.client.GetNotionsByTag(c.ID, params)
}
