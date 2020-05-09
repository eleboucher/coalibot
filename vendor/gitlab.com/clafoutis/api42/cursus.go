package api42

import (
	"time"
)

type Cursus42 struct {
	ID        int        `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	Name      string     `json:"name"`
	Slug      string     `json:"slug"`
	client    *Client42
}

func (c *Client42) backGetCursus(params *RequestParameter, directFilter string, value interface{}) ([]Cursus42, error) {
	var cursusArray []Cursus42
	url, err := buildUrl(cursusRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &cursusArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(cursusArray); i++ {
		cursusArray[i].client = c
	}
	return cursusArray, nil
}

// Returns the Cursus specified by the id.
func (c *Client42) GetCursusByID(id interface{}) (*Cursus42, error) {
	var cursus *Cursus42
	url, err := buildUrl(cursusRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &cursus)
	if err != nil {
		return nil, err
	}
	cursus.client = c
	return cursus, nil
}

// Returns all the Cursus.
func (c *Client42) GetCursus(params *RequestParameter) ([]Cursus42, error) {
	return c.backGetCursus(params, defaultRequest, nil)
}
