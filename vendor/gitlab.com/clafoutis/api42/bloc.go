package api42

import (
	"time"
)

type Bloc42 struct {
	ID         int           `json:"id"`
	CampusID   int           `json:"campus_id"`
	CursusID   int           `json:"cursus_id"`
	SquadSize  int           `json:"squad_size"`
	CreatedAt  *time.Time    `json:"created_at"`
	UpdatedAt  *time.Time    `json:"updated_at"`
	Coalitions []Coalition42 `json:"coalitions"`
	client     *Client42
}

func (c *Client42) backGetBlocs(params *RequestParameter, directFilter string, value interface{}) ([]Bloc42, error) {
	var blocArray []Bloc42
	url, err := buildUrl(blocRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &blocArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(blocArray); i++ {
		blocArray[i].client = c
	}
	return blocArray, nil
}

// Returns the Bloc specified by the id.
func (c *Client42) GetBloc(id interface{}) (*Bloc42, error) {
	var bloc *Bloc42
	url, err := buildUrl(blocRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &bloc)
	if err != nil {
		return nil, err
	}
	bloc.client = c
	return bloc, nil
}

// Returns all the Blocs.
func (c *Client42) GetBlocs(params *RequestParameter) ([]Bloc42, error) {
	return c.backGetBlocs(params, defaultRequest, nil)
}
