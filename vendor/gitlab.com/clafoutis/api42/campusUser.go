package api42

type CampusUser42 struct {
	CampusID  int  `json:"campus_id"`
	ID        int  `json:"id"`
	IsPrimary bool `json:"is_primary"`
	UserID    int  `json:"user_id"`
	client    *Client42
}

func (c *Client42) backGetCampusUsers(params *RequestParameter, directFilter string, value interface{}) ([]CampusUser42, error) {
	var campusUserArray []CampusUser42
	url, err := buildUrl(campusUserRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &campusUserArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(campusUserArray); i++ {
		campusUserArray[i].client = c
	}
	return campusUserArray, nil
}

// Returns the CampusUser specified by the id.
func (c *Client42) GetCampusUser(id interface{}) (*CampusUser42, error) {
	var campusUser *CampusUser42
	url, err := buildUrl(campusUserRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &campusUser)
	if err != nil {
		return nil, err
	}
	campusUser.client = c
	return campusUser, nil
}

// Returns all the CampusUsers.
func (c *Client42) GetCampusUsers(params *RequestParameter) ([]CampusUser42, error) {
	return c.backGetCampusUsers(params, defaultRequest, nil)
}

// By User

// Returns all the CampusUsers of the given User.
func (c *Client42) GetCampusUsersByUser(id interface{}, params *RequestParameter) ([]CampusUser42, error) {
	return c.backGetCampusUsers(params, userRequest, id)
}

// Returns all the CampusUsers of the given User.
func (c *User42) GetCampusUsers(params *RequestParameter) ([]CampusUser42, error) {
	return c.client.GetCampusUsersByUser(c.ID, params)
}
