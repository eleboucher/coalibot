package api42

type Tag42 struct {
	ID         int           `json:"id"`
	Name       string        `json:"name"`
	Kind       string        `json:"kind"`
	Users      []User42      `json:"users"`
	Subnotions []Subnotion42 `json:"subnotions"`
	client     *Client42
}

func (c *Client42) backGetTags(params *RequestParameter, directFilter string, value interface{}) ([]Tag42, error) {
	var tagArray []Tag42
	url, err := buildUrl(tagRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &tagArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(tagArray); i++ {
		tagArray[i].client = c
	}
	return tagArray, nil
}

// Returns the Tag specified by the id.
func (c *Client42) GetTag(id interface{}) (*Tag42, error) {
	var tag *Tag42
	url, err := buildUrl(tagRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &tag)
	if err != nil {
		return nil, err
	}
	tag.client = c
	return tag, nil
}

// Returns all the Tags.
func (c *Client42) GetTags(params *RequestParameter) ([]Tag42, error) {
	return c.backGetTags(params, defaultRequest, nil)
}

// By Cursus

// Returns all the Tags of the given Cursus.
func (c *Client42) GetTagsByCursus(id interface{}, params *RequestParameter) ([]Tag42, error) {
	return c.backGetTags(params, cursusRequest, id)
}

// Returns all the Tags of the given Cursus.
func (c *Cursus42) GetTags(params *RequestParameter) ([]Tag42, error) {
	return c.client.GetTagsByCursus(c.ID, params)
}

// By Notion

// Returns all the Tags of the given Notion.
func (c *Client42) GetTagsByNotion(id interface{}, params *RequestParameter) ([]Tag42, error) {
	return c.backGetTags(params, notionRequest, id)
}

// Returns all the Tags of the given Notion.
func (c *Notion42) GetTags(params *RequestParameter) ([]Tag42, error) {
	return c.client.GetTagsByNotion(c.ID, params)
}

// By User

// Returns all the Tags of the given User.
func (c *Client42) GetTagsByUser(id interface{}, params *RequestParameter) ([]Tag42, error) {
	return c.backGetTags(params, userRequest, id)
}

// Returns all the Tags of the given User.
func (c *User42) GetTags(params *RequestParameter) ([]Tag42, error) {
	return c.client.GetTagsByUser(c.ID, params)
}

// By Project

// Returns all the Tags of the given Project.
func (c *Client42) GetTagsByProject(id interface{}, params *RequestParameter) ([]Tag42, error) {
	return c.backGetTags(params, projectRequest, id)
}

// Returns all the Tags of the given Project.
func (c *Project42) GetTags(params *RequestParameter) ([]Tag42, error) {
	return c.client.GetTagsByProject(c.ID, params)
}
