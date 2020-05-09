package api42

import (
	"time"
)

type ProjectSession42 struct {
	ID               int            `json:"id"`
	Solo             bool           `json:"solo"`
	BeginAt          *time.Time     `json:"begin_at"`
	EndAt            *time.Time     `json:"end_at"`
	EstimateTime     *time.Duration `json:"estimate_time"`
	DurationDays     *time.Duration `json:"duration_days"`
	TerminatingAfter *time.Duration `json:"terminating_after"`
	ProjectID        int            `json:"project_id"`
	CampusID         *int           `json:"campus_id"`
	CursusID         *int           `json:"cursus_id"`
	CreatedAt        *time.Time     `json:"created_at"`
	UpdatedAt        *time.Time     `json:"updated_at"`
	MaxPeople        *int           `json:"max_people"`
	IsSubscriptable  bool           `json:"is_subscriptable"`
	Scales           []Scale42      `json:"scales"`
	Uploads          []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"uploads"`
	TeamBehaviour string     `json:"team_behaviour"`
	Project       *Project42 `json:"project"`
	Campus        *Campus42  `json:"campus"`
	Cursus        *Cursus42  `json:"cursus"`
	// Evaluations			[]Evaluation42  `json:"evaluations"`
	client *Client42
}

func (c *Client42) backGetProjectSessions(params *RequestParameter, directFilter string, value interface{}) ([]ProjectSession42, error) {
	var projectSessionArray []ProjectSession42
	url, err := buildUrl(projectSessionRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &projectSessionArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(projectSessionArray); i++ {
		projectSessionArray[i].client = c
	}
	return projectSessionArray, nil
}

// Returns the ProjectSession specified by the id.
func (c *Client42) GetProjectSession(id interface{}) (*ProjectSession42, error) {
	var projectSession *ProjectSession42
	url, err := buildUrl(projectSessionRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &projectSession)
	if err != nil {
		return nil, err
	}
	projectSession.client = c
	return projectSession, nil
}

// Returns all the ProjectSessions.
func (c *Client42) GetProjectSessions(params *RequestParameter) ([]ProjectSession42, error) {
	return c.backGetProjectSessions(params, defaultRequest, nil)
}

// By Project

// Returns all the ProjectSessions of the given Project.
func (c *Client42) GetProjectSessionsByProject(id interface{}, params *RequestParameter) ([]ProjectSession42, error) {
	return c.backGetProjectSessions(params, projectRequest, id)
}

// Returns all the ProjectSessions of the given Project.
func (c *Project42) GetProjectSessions(params *RequestParameter) ([]ProjectSession42, error) {
	return c.client.GetProjectSessionsByProject(c.ID, params)
}
