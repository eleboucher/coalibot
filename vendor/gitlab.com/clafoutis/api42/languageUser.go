package api42

import (
	"time"
)

type LanguageUser42 struct {
	ID         int        `json:"id"`
	LanguageID int        `json:"language_id"`
	UserID     int        `json:"user_id"`
	Position   int        `json:"position"`
	CreatedAt  *time.Time `json:"created_at"`
	client     *Client42
}

func (c *Client42) backGetLanguageUsers(params *RequestParameter, directFilter string, value interface{}) ([]LanguageUser42, error) {
	var languageUserArray []LanguageUser42
	url, err := buildUrl(languageUserRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &languageUserArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(languageUserArray); i++ {
		languageUserArray[i].client = c
	}
	return languageUserArray, nil
}

func (c *Client42) backGetLanguageUser(id interface{}, directFilter string, value interface{}) (*LanguageUser42, error) {
	var languageUser *LanguageUser42
	var url string
	var err error
	switch directFilter {
	case defaultRequest:
		url, err = buildUrl(languageUserRequest, directFilter, id)
	default:
		url, err = buildUrl(languageUserRequest, directFilter, value, defaultRequest, id)
	}
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &languageUser)
	if err != nil {
		return nil, err
	}
	languageUser.client = c
	return languageUser, nil
}

// Returns all the LanguageUsers.
func (c *Client42) GetLanguageUsers(params *RequestParameter) ([]LanguageUser42, error) {
	return c.backGetLanguageUsers(params, defaultRequest, nil)
}

// Returns the LanguageUser specified by the id.
func (c *Client42) GetLanguageUser(id interface{}) (*LanguageUser42, error) {
	return c.backGetLanguageUser(id, defaultRequest, nil)
}

// By User

// Returns all the LanguageUsers of the given User.
func (c *Client42) GetLanguageUsersByUser(id interface{}, params *RequestParameter) ([]LanguageUser42, error) {
	return c.backGetLanguageUsers(params, achievementRequest, id)
}

// Returns the LanguageUser of the given id, associated with the given User.
func (c *Client42) GetLanguageUserByUser(id, userID interface{}) (*LanguageUser42, error) {
	return c.backGetLanguageUser(id, userRequest, userID)
}

// Returns all the LanguageUsers of the given User.
func (c *User42) GetLanguageUsers(params *RequestParameter) ([]LanguageUser42, error) {
	return c.client.GetLanguageUsersByUser(c.ID, params)
}

// Returns all the LanguageUser of the given User.
func (c *User42) GetLanguageUser(id interface{}) (*LanguageUser42, error) {
	return c.client.GetLanguageUserByUser(id, c.ID)
}
