package api42

type Campus42 struct {
	ID          int        `json:"id"`
	UsersCount  int        `json:"users_count"`
	VogsphereID int        `json:"vogsphere_id"`
	Name        string     `json:"name"`
	TimeZone    string     `json:"time_zone"`
	Language    Language42 `json:"language"`
	client      *Client42
}

func (c *Client42) backGetCampus(params *RequestParameter, directFilter string, value interface{}) ([]Campus42, error) {
	var campusArray []Campus42
	url, err := buildUrl(campusRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &campusArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(campusArray); i++ {
		campusArray[i].client = c
	}
	return campusArray, nil
}

// Returns the Campus specified by the id.
func (c *Client42) GetCampusByID(id interface{}) (*Campus42, error) {
	var campus *Campus42
	url, err := buildUrl(campusRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &campus)
	if err != nil {
		return nil, err
	}
	campus.client = c
	return campus, nil
}

// Returns all the Campus.
func (c *Client42) GetCampus(params *RequestParameter) ([]Campus42, error) {
	return c.backGetCampus(params, defaultRequest, nil)
}
