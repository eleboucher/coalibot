package api42

type Language42 struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
	client     *Client42
}

func (c *Client42) backGetLanguages(params *RequestParameter, directFilter string, value interface{}) ([]Language42, error) {
	var languageArray []Language42
	url, err := buildUrl(languageRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &languageArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(languageArray); i++ {
		languageArray[i].client = c
	}
	return languageArray, nil
}

// Returns the Language specified by the id.
func (c *Client42) GetLanguage(id interface{}) (*Language42, error) {
	var language *Language42
	url, err := buildUrl(languageRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &language)
	if err != nil {
		return nil, err
	}
	language.client = c
	return language, nil
}

// Returns all the Languages.
func (c *Client42) GetLanguages(params *RequestParameter) ([]Language42, error) {
	return c.backGetLanguages(params, defaultRequest, nil)
}
