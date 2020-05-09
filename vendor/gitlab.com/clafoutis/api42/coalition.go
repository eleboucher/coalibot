package api42

type Coalition42 struct {
	ID       int    `json:"id"`
	Score    int    `json:"score"`
	UserID   int    `json:"user_id"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	ImageUrl string `json:"image_url"`
	Color    string `json:"color"`
	client   *Client42
}

func (c *Client42) backGetCoalitions(params *RequestParameter, directFilter string, value interface{}) ([]Coalition42, error) {
	var coaArray []Coalition42
	url, err := buildUrl(coalitionRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &coaArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(coaArray); i++ {
		coaArray[i].client = c
	}
	return coaArray, nil
}

// Returns the Coalition specified by the id.
func (c *Client42) GetCoalition(id interface{}) (*Coalition42, error) {
	var coalition *Coalition42
	url, err := buildUrl(coalitionRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &coalition)
	if err != nil {
		return nil, err
	}
	coalition.client = c
	return coalition, nil
}

// Returns all the Coalitions.
func (c *Client42) GetCoalitions(params *RequestParameter) ([]Coalition42, error) {
	return c.backGetCoalitions(params, defaultRequest, nil)
}

// Returns all the Coalitions of the given Bloc.
func (c *Client42) GetCoalitionsByBloc(id interface{}, params *RequestParameter) ([]Coalition42, error) {
	if params == nil {
		params = NewParameter()
	}
	params.AddCustom("bloc_id", id)
	return c.GetCoalitions(params)
}

// By User

// Returns all the Coalitions of the given User.
func (c *Client42) GetCoalitionsByUser(id interface{}, params *RequestParameter) ([]Coalition42, error) {
	return c.backGetCoalitions(params, userRequest, id)
}

// Returns all the Coalitions of the given User.
func (u *User42) GetCoalitions(params *RequestParameter) ([]Coalition42, error) {
	return u.client.GetCoalitionsByUser(u.ID, params)
}
