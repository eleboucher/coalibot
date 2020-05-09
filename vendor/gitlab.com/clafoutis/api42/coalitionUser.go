package api42

import (
	"time"
)

type CoalitionUser42 struct {
	ID          int        `json:"id"`
	CoalitionID int        `json:"coalition_id"`
	UserID      int        `json:"user_id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	client      *Client42
}

func (c *Client42) backGetCoalitionUsers(params *RequestParameter, directFilter string, value interface{}) ([]CoalitionUser42, error) {
	var coalitionUserArray []CoalitionUser42
	url, err := buildUrl(coalitionUserRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &coalitionUserArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(coalitionUserArray); i++ {
		coalitionUserArray[i].client = c
	}
	return coalitionUserArray, nil
}

// Returns the CoalitionUser specified by the id.
func (c *Client42) GetCoalitionUser(id interface{}) (*CoalitionUser42, error) {
	var coalitionUser *CoalitionUser42
	url, err := buildUrl(coalitionUserRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &coalitionUser)
	if err != nil {
		return nil, err
	}
	coalitionUser.client = c
	return coalitionUser, nil
}

// Returns all the CoalitionUsers.
func (c *Client42) GetCoalitionUsers(params *RequestParameter) ([]CoalitionUser42, error) {
	return c.backGetCoalitionUsers(params, defaultRequest, nil)
}

// By User

// Returns all the CoalitionUsers of the given User.
func (c *Client42) GetCoalitionUsersByUser(id interface{}, params *RequestParameter) ([]CoalitionUser42, error) {
	return c.backGetCoalitionUsers(params, userRequest, id)
}

// Returns all the CoalitionUsers of the given User.
func (u *User42) GetCoalitionUsers(params *RequestParameter) ([]CoalitionUser42, error) {
	return u.client.GetCoalitionUsersByUser(u.Login, params)
}

// By Coalition

// Returns all the CoalitionUsers of the given Coalition.
func (c *Client42) GetCoalitionUsersByCoalition(id interface{}, params *RequestParameter) ([]CoalitionUser42, error) {
	return c.backGetCoalitionUsers(params, coalitionRequest, id)
}

// Returns all the CoalitionUsers of the given Coalition.
func (u *Coalition42) GetCoalitionUsers(params *RequestParameter) ([]CoalitionUser42, error) {
	return u.client.GetCoalitionUsersByCoalition(u.ID, params)
}
