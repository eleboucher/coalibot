package api42

import (
	"time"
)

type ScaleTeam42 struct {
	ID              int          `json:"id"`
	BeginAt         *time.Time   `json:"begin_at"`
	Comment         string       `json:"comment"`
	Correcteds      []User42     `json:"correcteds"`
	CorrectedString string       `json:"correcteds"`
	Corrector       *User42      `json:"corrector"`
	CorrectorString string       `json:"corrector"`
	CreatedAt       *time.Time   `json:"created_at"`
	Feedback        string       `json:"feedback"`
	Feedbacks       []Feedback42 `json:"feedbacks"`
	FilledAt        *time.Time   `json:"filled_at"`
	FinalMark       int          `json:"final_mark"`
	Flag            struct {
		ID        int        `json:"id"`
		CreatedAt *time.Time `json:"created_at"`
		Icon      string     `json:"icon"`
		Name      string     `json:"name"`
		Positive  bool       `json:"positive"`
		UpdatedAt *time.Time `json:"updated_at"`
	} `json:"flag"`
	Scale     *Scale42   `json:"scale"`
	ScaleID   int        `json:"scale_id"`
	Team      *Team42    `json:"team"`
	Truant    *User42    `json:"truant"`
	UpdatedAt *time.Time `json:"updated_at"`
	client    *Client42
}

func (c *Client42) backGetScaleTeams(params *RequestParameter, directFilter string, value interface{}, args ...interface{}) ([]ScaleTeam42, error) {
	var scaleTeamArray []ScaleTeam42
	url, err := buildUrl(scaleTeamRequest, directFilter, value, args...)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &scaleTeamArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(scaleTeamArray); i++ {
		scaleTeamArray[i].client = c
	}
	return scaleTeamArray, nil
}

// Returns the ScaleTeam specified by the id.
func (c *Client42) GetScaleTeam(id interface{}) (*ScaleTeam42, error) {
	var scaleTeam *ScaleTeam42
	url, err := buildUrl(scaleTeamRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &scaleTeam)
	if err != nil {
		return nil, err
	}
	scaleTeam.client = c
	return scaleTeam, nil
}

// Returns all the ScaleTeams.
func (c *Client42) GetScaleTeams(params *RequestParameter) ([]ScaleTeam42, error) {
	return c.backGetScaleTeams(params, defaultRequest, nil)
}

// By User

// Returns all the ScaleTeams of the given User.
func (c *Client42) GetScaleTeamsByUser(id interface{}, params *RequestParameter) ([]ScaleTeam42, error) {
	return c.backGetScaleTeams(params, userRequest, id)
}

// Returns all the ScaleTeams of the given User as corrector.
func (c *Client42) GetScaleTeamsByUserAsCorrector(id interface{}, params *RequestParameter) ([]ScaleTeam42, error) {
	return c.backGetScaleTeams(params, userRequest, id, defaultRequest, "as_corrector")
}

// Returns all the ScaleTeams of the given User as corrector.
func (c *User42) GetScaleTeamsAsCorrector(params *RequestParameter) ([]ScaleTeam42, error) {
	return c.client.GetScaleTeamsByUserAsCorrector(c.ID, params)
}

// Returns all the ScaleTeams of the given User as corrected.
func (c *Client42) GetScaleTeamsByUserAsCorrected(id interface{}, params *RequestParameter) ([]ScaleTeam42, error) {
	return c.backGetScaleTeams(params, userRequest, id, defaultRequest, "as_corrected")
}

// Returns all the ScaleTeams of the given User as corrected.
func (c *User42) GetScaleTeamsAsCorrected(params *RequestParameter) ([]ScaleTeam42, error) {
	return c.client.GetScaleTeamsByUserAsCorrected(c.ID, params)
}

// Returns all the ScaleTeams of the given User.
func (c *User42) GetScaleTeams(params *RequestParameter) ([]ScaleTeam42, error) {
	return c.client.GetScaleTeamsByUser(c.ID, params)
}

// By Project

// Returns all the ScaleTeams of the given Project.
func (c *Client42) GetScaleTeamsByProject(id interface{}, params *RequestParameter) ([]ScaleTeam42, error) {
	return c.backGetScaleTeams(params, projectRequest, id)
}

// Returns all the ScaleTeams of the given Project.
func (c *Project42) GetScaleTeams(params *RequestParameter) ([]ScaleTeam42, error) {
	return c.client.GetScaleTeamsByProject(c.ID, params)
}
