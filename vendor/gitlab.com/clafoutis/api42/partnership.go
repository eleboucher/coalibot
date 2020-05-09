package api42

type Partnership42 struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	Slug                 string `json:"slug"`
	Tier                 int    `json:"tier"`
	URL                  string `json:"url"`
	PartnershipsUsersURL string `json:"partnerships_users_url"`
	client               *Client42
}

func (c *Client42) backGetPartnerships(params *RequestParameter, directFilter string, value interface{}) ([]Partnership42, error) {
	var partnershipArray []Partnership42
	url, err := buildUrl(partnershipRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &partnershipArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(partnershipArray); i++ {
		partnershipArray[i].client = c
	}
	return partnershipArray, nil
}

// Returns the Partnership specified by the id.
func (c *Client42) GetPartnership(id interface{}) (*Partnership42, error) {
	var partnership *Partnership42
	url, err := buildUrl(partnershipRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &partnership)
	if err != nil {
		return nil, err
	}
	partnership.client = c
	return partnership, nil
}

// Returns all the Partnerships.
func (c *Client42) GetPartnerships(params *RequestParameter) ([]Partnership42, error) {
	return c.backGetPartnerships(params, defaultRequest, nil)
}
