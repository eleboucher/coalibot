package api42

import (
	"time"
)

type Expertise42 struct {
	ID                 int        `json:"id"`
	Name               string     `json:"name"`
	Slug               string     `json:"slug"`
	URL                string     `json:"url"`
	Kind               string     `json:"kind"`
	CreatedAt          *time.Time `json:"created_at"`
	ExpertisesUsersURL string     `json:"expertises_users_url"`
	client             *Client42
}

func (c *Client42) backGetExpertises(params *RequestParameter, directFilter string, value interface{}) ([]Expertise42, error) {
	var expertiseArray []Expertise42
	url, err := buildUrl(expertiseRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &expertiseArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(expertiseArray); i++ {
		expertiseArray[i].client = c
	}
	return expertiseArray, nil
}

// Returns the Expertise specified by the id.
func (c *Client42) GetExpertise(id interface{}) (*Expertise42, error) {
	var expertise *Expertise42
	url, err := buildUrl(expertiseRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &expertise)
	if err != nil {
		return nil, err
	}
	expertise.client = c
	return expertise, nil
}

// Returns all the Expertises.
func (c *Client42) GetExpertises(params *RequestParameter) ([]Expertise42, error) {
	return c.backGetExpertises(params, defaultRequest, nil)
}
