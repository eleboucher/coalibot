package api42

type TitleUser42 struct {
	ID       int  `json:"id"`
	Selected bool `json:"selected"`
	TitleID  int  `json:"title_id"`
	UserID   int  `json:"user_id"`
	client   *Client42
}

func (c *Client42) backGetTitleUsers(params *RequestParameter, directFilter string, value interface{}) ([]TitleUser42, error) {
	var titleUserArray []TitleUser42
	url, err := buildUrl(titleUserRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &titleUserArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(titleUserArray); i++ {
		titleUserArray[i].client = c
	}
	return titleUserArray, nil
}

// Returns the TitleUser specified by the id.
func (c *Client42) GetTitleUser(id interface{}) (*TitleUser42, error) {
	var titleUser *TitleUser42
	url, err := buildUrl(titleUserRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &titleUser)
	if err != nil {
		return nil, err
	}
	titleUser.client = c
	return titleUser, nil
}

// Returns all the TitleUsers.
func (c *Client42) GetTitleUsers(params *RequestParameter) ([]TitleUser42, error) {
	return c.backGetTitleUsers(params, defaultRequest, nil)
}

// By User

// Returns all the TitleUsers of the given User.
func (c *Client42) GetTitleUsersByUser(id interface{}, params *RequestParameter) ([]TitleUser42, error) {
	return c.backGetTitleUsers(params, userRequest, id)
}

// Returns all the TitleUsers of the given User.
func (u *User42) GetTitleUsers(params *RequestParameter) ([]TitleUser42, error) {
	return u.client.GetTitleUsersByUser(u.ID, params)
}

// By Title

// Returns all the TitleUsers of the given Title.
func (c *Client42) GetTitleUsersByTitle(id interface{}, params *RequestParameter) ([]TitleUser42, error) {
	return c.backGetTitleUsers(params, titleRequest, id)
}

// Returns all the TitleUsers of the given Title.
func (c *Title42) GetTitleUsers(params *RequestParameter) ([]TitleUser42, error) {
	return c.client.GetTitleUsersByTitle(c.ID, params)
}
