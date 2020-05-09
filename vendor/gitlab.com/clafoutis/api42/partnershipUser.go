package api42

type PartnershipUser42 struct {
	ID            int     `json:"id"`
	PartnershipID int     `json:"partnership_id"`
	FinalMark     *int    `json:"final_mark"`
	User          *User42 `json:"user"`
	URL           string  `json:"url"`
	client        *Client42
}

func (c *Client42) backGetPartnershipUsers(params *RequestParameter, directFilter string, value interface{}) ([]PartnershipUser42, error) {
	var partnershipUserArray []PartnershipUser42
	url, err := buildUrl(partnershipUserRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &partnershipUserArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(partnershipUserArray); i++ {
		partnershipUserArray[i].client = c
	}
	return partnershipUserArray, nil
}

// Returns the PartnershipUser specified by the id.
func (c *Client42) GetPartnershipUser(id interface{}) (*PartnershipUser42, error) {
	var partnershipUser *PartnershipUser42
	url, err := buildUrl(partnershipUserRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &partnershipUser)
	if err != nil {
		return nil, err
	}
	partnershipUser.client = c
	return partnershipUser, nil
}

// Returns all the PartnershipUsers.
func (c *Client42) GetPartnershipUsers(params *RequestParameter) ([]PartnershipUser42, error) {
	return c.backGetPartnershipUsers(params, defaultRequest, nil)
}

// By Partnership

// Returns all the PartnershipUsers of the given Partnership.
func (c *Client42) GetPartnershipUsersByPartnership(id interface{}, params *RequestParameter) ([]PartnershipUser42, error) {
	return c.backGetPartnershipUsers(params, partnershipRequest, id)
}

// Returns all the PartnershipUsers of the given Partnership.
func (c *Partnership42) GetPartnershipUsers(params *RequestParameter) ([]PartnershipUser42, error) {
	return c.client.GetPartnershipUsersByPartnership(c.ID, params)
}
