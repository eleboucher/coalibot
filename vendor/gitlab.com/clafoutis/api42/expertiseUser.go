package api42

import (
	"time"
)

type ExpertiseUser42 struct {
	ID          int          `json:"id"`
	ExpertiseID int          `json:"expertise_id"`
	Interested  bool         `json:"interested"`
	Value       int          `json:"value"`
	ContactMe   bool         `json:"contact_me"`
	CreatedAt   *time.Time   `json:"created_at"`
	UserID      int          `json:"user_id"`
	Expertise   *Expertise42 `json:"expertise"`
	User        *User42      `json:"user"`
	client      *Client42
}

func (c *Client42) backGetExpertiseUsers(params *RequestParameter, directFilter string, value interface{}) ([]ExpertiseUser42, error) {
	var expertiseUserArray []ExpertiseUser42
	url, err := buildUrl(expertiseUserRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &expertiseUserArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(expertiseUserArray); i++ {
		expertiseUserArray[i].client = c
	}
	return expertiseUserArray, nil
}

// Returns the ExpertiseUser specified by the id.
func (c *Client42) GetExpertiseUser(id interface{}) (*Expertise42, error) {
	var expertiseUser *Expertise42
	url, err := buildUrl(expertiseUserRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &expertiseUser)
	if err != nil {
		return nil, err
	}
	expertiseUser.client = c
	return expertiseUser, nil
}

// Returns all the ExpertiseUsers.
func (c *Client42) GetExpertiseUsers(params *RequestParameter) ([]ExpertiseUser42, error) {
	return c.backGetExpertiseUsers(params, defaultRequest, nil)
}

// By Expertise

// Returns all the ExpertiseUsers of the given Expertise.
func (c *Client42) GetExpertiseUsersByExpertise(id interface{}, params *RequestParameter) ([]ExpertiseUser42, error) {
	return c.backGetExpertiseUsers(params, expertiseRequest, id)
}

// Returns all the ExpertiseUsers of the given Expertise.
func (c *Expertise42) GetExpertiseUsers(params *RequestParameter) ([]ExpertiseUser42, error) {
	return c.client.GetExpertiseUsersByExpertise(c.ID, params)
}

// By User

// Returns all the ExpertiseUsers of the given User.
func (c *Client42) GetExpertiseUsersByUser(id interface{}, params *RequestParameter) ([]ExpertiseUser42, error) {
	return c.backGetExpertiseUsers(params, userRequest, id)
}

// Returns all the ExpertiseUsers of the given User.
func (c *User42) GetExpertiseUsers(params *RequestParameter) ([]ExpertiseUser42, error) {
	return c.client.GetExpertiseUsersByUser(c.ID, params)
}
