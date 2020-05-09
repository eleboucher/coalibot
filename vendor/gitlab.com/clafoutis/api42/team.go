package api42

import (
	"time"
)

type Team42 struct {
	ID               int            `json:"id"`
	Name             string         `json:"name"`
	URL              string         `json:"url"`
	FinalMark        *int           `json:"final_mark"`
	ProjectID        int            `json:"project_id"`
	CreatedAt        *time.Time     `json:"created_at"`
	UpdatedAt        *time.Time     `json:"updated_at"`
	Status           string         `json:"status"`
	TerminatingAt    *time.Time     `json:"terminating_at"`
	Users            []User42       `json:"users"`
	Locked           bool           `json:"locked?"`
	Validated        bool           `json:"validated?"`
	Closed           bool           `json:"closed?"`
	RepoURL          string         `json:"repo_url"`
	RepoUUID         string         `json:"repo_uuid"`
	LockedAt         *time.Time     `json:"locked_at"`
	ClosedAt         *time.Time     `json:"closed_at"`
	ProjectSessionID int            `json:"project_session_id"`
	ScaleTeams       []ScaleTeam42  `json:"scale_teams"`
	TeamsUploads     []TeamUpload42 `json:"teams_uploads"`
	client           *Client42
}

func (c *Client42) backGetTeams(params *RequestParameter, directFilter string, value interface{}) ([]Team42, error) {
	var teamArray []Team42
	url, err := buildUrl(teamRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &teamArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(teamArray); i++ {
		teamArray[i].client = c
	}
	return teamArray, nil
}

// Returns the Team specified by the id.
func (c *Client42) GetTeam(id interface{}) (*Team42, error) {
	var team *Team42
	url, err := buildUrl(teamRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &team)
	if err != nil {
		return nil, err
	}
	team.client = c
	return team, nil
}

// Returns all the Teams.
func (c *Client42) GetTeams(params *RequestParameter) ([]Team42, error) {
	return c.backGetTeams(params, defaultRequest, nil)
}

// By Cursus

// Returns all the Teams of the given Cursus.
func (c *Client42) GetTeamsByCursus(id interface{}, params *RequestParameter) ([]Team42, error) {
	return c.backGetTeams(params, cursusRequest, id)
}

// Returns all the Teams of the given Cursus.
func (c *Cursus42) GetTeams(params *RequestParameter) ([]Team42, error) {
	return c.client.GetTeamsByCursus(c.ID, params)
}

// By User

// Returns all the Teams of the given User.
func (c *Client42) GetTeamsByUser(id interface{}, params *RequestParameter) ([]Team42, error) {
	return c.backGetTeams(params, userRequest, id)
}

// Returns all the Teams of the given User.
func (c *User42) GetTeams(params *RequestParameter) ([]Team42, error) {
	return c.client.GetTeamsByUser(c.ID, params)
}

// By Project

// Returns all the Teams of the given Project.
func (c *Client42) GetTeamsByProject(id interface{}, params *RequestParameter) ([]Team42, error) {
	return c.backGetTeams(params, projectRequest, id)
}

// Returns all the Teams of the given Project.
func (c *Project42) GetTeams(params *RequestParameter) ([]Team42, error) {
	return c.client.GetTeamsByProject(c.ID, params)
}

// By Project Session

// Returns all the Teams of the given ProjectSession.
func (c *Client42) GetTeamsByProjectSession(id interface{}, params *RequestParameter) ([]Team42, error) {
	return c.backGetTeams(params, projectSessionRequest, id)
}

// Returns all the Teams of the given ProjectSession.
func (c *ProjectSession42) GetTeams(params *RequestParameter) ([]Team42, error) {
	return c.client.GetTeamsByProjectSession(c.ID, params)
}
