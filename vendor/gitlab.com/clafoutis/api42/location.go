package api42

import (
	"time"
)

type Location42 struct {
	ID       int        `json:"id"`
	CampusID int        `json:"campus_id"`
	Primary  bool       `json:"primary"`
	Host     string     `json:"host"`
	User     *User42    `json:"user"`
	EndAt    *time.Time `json:"end_at"`
	BeginAt  *time.Time `json:"begin_at"`
	client   *Client42
}

func (c *Client42) backGetLocations(params *RequestParameter, directFilter string, value interface{}) ([]Location42, error) {
	var locationArray []Location42
	url, err := buildUrl(locationRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &locationArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(locationArray); i++ {
		locationArray[i].client = c
	}
	return locationArray, nil
}

// Returns all the Locations.
func (c *Client42) GetLocations(params *RequestParameter) ([]Location42, error) {
	return c.backGetLocations(params, defaultRequest, nil)
}

// Returns the Location specified by the id.
func (c *Client42) GetLocation(id interface{}) (*Location42, error) {
	var location *Location42
	url, err := buildUrl(locationRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &location)
	if err != nil {
		return nil, err
	}
	location.client = c
	return location, nil
}

// Returns all the ActiveLocations.
func (c *Client42) GetActiveLocations(params *RequestParameter) ([]Location42, error) {
	if params == nil {
		params = NewParameter()
	}
	params.AddFilter("active", "true")
	return c.GetLocations(params)
}

// Returns all the InactiveLocations.
func (c *Client42) GetInactiveLocations(params *RequestParameter) ([]Location42, error) {
	if params == nil {
		params = NewParameter()
	}
	params.AddFilter("inactive", "true")
	return c.GetLocations(params)
}

// By Users

// Returns all the UserLocations.
func (c *Client42) GetUserLocations(id interface{}, params *RequestParameter) ([]Location42, error) {
	return c.backGetLocations(params, userRequest, id)
}

// Returns all the UserCurrentLocations.
func (c *Client42) GetUserCurrentLocations(id interface{}, params *RequestParameter) ([]Location42, error) {
	if params == nil {
		params = NewParameter()
	}
	params.AddFilter("active", "true")
	return c.backGetLocations(params, userRequest, id)
}

// Returns all the UserOldLocations.
func (c *Client42) GetUserOldLocations(id interface{}, params *RequestParameter) ([]Location42, error) {
	if params == nil {
		params = NewParameter()
	}
	params.AddFilter("inactive", "true")
	return c.backGetLocations(params, userRequest, id)
}

// Returns all the Locations of the given User.
func (u *User42) GetLocations(params *RequestParameter) ([]Location42, error) {
	return u.client.GetUserLocations(u.ID, params)
}

// Returns all the CurrentLocations of the given User.
func (u *User42) GetCurrentLocations(params *RequestParameter) ([]Location42, error) {
	return u.client.GetUserCurrentLocations(u.ID, params)
}

// Returns all the OldLocations of the given User.
func (u *User42) GetOldLocations(params *RequestParameter) ([]Location42, error) {
	return u.client.GetUserOldLocations(u.ID, params)
}

// By Campus

// Returns all the Locations of the given Campus.
func (c *Client42) GetLocationsByCampus(id interface{}, params *RequestParameter) ([]Location42, error) {
	return c.backGetLocations(params, campusRequest, id)
}

// Returns all the ActiveLocations of the given Campus.
func (c *Client42) GetActiveLocationsByCampus(id interface{}, params *RequestParameter) ([]Location42, error) {
	if params == nil {
		params = NewParameter()
	}
	params.AddFilter("active", "true")
	return c.GetLocationsByCampus(id, params)
}

// Returns all the InactiveLocations of the given Campus.
func (c *Client42) GetInactiveLocationsByCampus(id interface{}, params *RequestParameter) ([]Location42, error) {
	if params == nil {
		params = NewParameter()
	}
	params.AddFilter("inactive", "true")
	return c.GetLocationsByCampus(id, params)
}

// Returns all the Locations of the given Campus.
func (c *Campus42) GetLocations(params *RequestParameter) ([]Location42, error) {
	return c.client.GetLocationsByCampus(c.ID, params)
}

// Returns all the ActiveLocations of the given Campus.
func (c *Campus42) GetActiveLocations(params *RequestParameter) ([]Location42, error) {
	return c.client.GetActiveLocationsByCampus(c.ID, params)
}

// Returns all the InactiveLocations of the given Campus.
func (c *Campus42) GetInactiveLocations(params *RequestParameter) ([]Location42, error) {
	return c.client.GetInactiveLocationsByCampus(c.ID, params)
}
