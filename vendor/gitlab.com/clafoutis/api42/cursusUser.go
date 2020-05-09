package api42

import (
	"time"
)

type CursusUser42 struct {
	ID           int        `json:"id"`
	BeginAt      *time.Time `json:"begin_at"`
	EndAt        *time.Time `json:"end_at"`
	Grade        string     `json:"grade"`
	Level        float64    `json:"level"`
	Skills       []Skill42  `json:"skills"`
	CursusID     int        `json:"cursus_id"`
	HasCoalition bool       `json:"has_coalition"`
	User         *User42    `json:"user"`
	Cursus       *Cursus42  `json:"cursus"`
	client       *Client42
}

func (c *Client42) backGetCursusUsers(params *RequestParameter, directFilter string, value interface{}) ([]CursusUser42, error) {
	var cursusUserArray []CursusUser42
	url, err := buildUrl(cursusUserRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &cursusUserArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(cursusUserArray); i++ {
		cursusUserArray[i].client = c
	}
	return cursusUserArray, nil
}

// Returns the CursusUser specified by the id.
func (c *Client42) GetCursusUser(id interface{}) (*CursusUser42, error) {
	var cursusUser *CursusUser42
	url, err := buildUrl(cursusUserRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &cursusUser)
	if err != nil {
		return nil, err
	}
	cursusUser.client = c
	return cursusUser, nil
}

// Returns all the CursusUsers.
func (c *Client42) GetCursusUsers(params *RequestParameter) ([]CursusUser42, error) {
	return c.backGetCursusUsers(params, defaultRequest, nil)
}

// By User

// Returns all the CursusUsers of the given User.
func (c *Client42) GetCursusUsersByUser(id interface{}, params *RequestParameter) ([]CursusUser42, error) {
	return c.backGetCursusUsers(params, userRequest, id)
}

// Returns all the CursusUsers of the given User.
func (u *User42) GetCursusUsers(params *RequestParameter) ([]CursusUser42, error) {
	return u.client.GetCursusUsersByUser(u.Login, params)
}

// By Cursus

// Returns all the CursusUsers of the given Cursus.
func (c *Client42) GetCursusUsersByCursus(id interface{}, params *RequestParameter) ([]CursusUser42, error) {
	return c.backGetCursusUsers(params, cursusRequest, id)
}

// Returns all the CursusUsers of the given Cursus.
func (u *Cursus42) GetCursusUsers(params *RequestParameter) ([]CursusUser42, error) {
	return u.client.GetCursusUsersByCursus(u.ID, params)
}
