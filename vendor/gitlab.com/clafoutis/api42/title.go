package api42

type Title42 struct {
	ID     int
	Name   string
	client *Client42
}

func (c *Client42) backGetTitles(params *RequestParameter, directFilter string, value interface{}) ([]Title42, error) {
	var titleArray []Title42
	url, err := buildUrl(titleRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &titleArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(titleArray); i++ {
		titleArray[i].client = c
	}
	return titleArray, nil
}

// Returns all the Titles.
func (c *Client42) GetTitles(params *RequestParameter) ([]Title42, error) {
	return c.backGetTitles(params, defaultRequest, nil)
}

// Returns the Title specified by the id.
func (c *Client42) GetTitle(id interface{}) (*Title42, error) {
	var title *Title42
	url, err := buildUrl(titleRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &title)
	if err != nil {
		return nil, err
	}
	title.client = c
	return title, nil
}

// By User

// Returns all the Titles of the given User.
func (c *Client42) GetTitlesByUser(id interface{}, params *RequestParameter) ([]Title42, error) {
	return c.backGetTitles(params, userRequest, id)
}

// Returns all the Titles of the given User.
func (u *User42) GetTitles(params *RequestParameter) ([]Title42, error) {
	return u.client.GetTitlesByUser(u.ID, params)
}
