package api42

import (
	"time"
)

type TeamUpload42 struct {
	ID        int        `json:"id"`
	FinalMark int        `json:"final_mark"`
	Comment   string     `json:"comment"`
	CreatedAt *time.Time `json:"created_at"`
	UploadID  int        `json:"upload_id"`
	Upload    struct {
		ID           int        `json:"id"`
		EvaluationID int        `json:"evaluation_id"`
		Name         string     `json:"name"`
		Description  string     `json:"description"`
		CreatedAt    *time.Time `json:"created_at"`
		UpdatedAt    *time.Time `json:"updated_at"`
	} `json:"upload"`
	client *Client42
}

func (c *Client42) backGetTeamUploads(params *RequestParameter, directFilter string, value interface{}) ([]TeamUpload42, error) {
	var teamUploadArray []TeamUpload42
	url, err := buildUrl(teamUploadRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &teamUploadArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(teamUploadArray); i++ {
		teamUploadArray[i].client = c
	}
	return teamUploadArray, nil
}

// Returns the TeamUpload specified by the id.
func (c *Client42) GetTeamUpload(id interface{}) (*TeamUpload42, error) {
	var teamUpload *TeamUpload42
	url, err := buildUrl(teamUploadRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &teamUpload)
	if err != nil {
		return nil, err
	}
	teamUpload.client = c
	return teamUpload, nil
}

// Returns all the TeamUploads.
func (c *Client42) GetTeamUploads(params *RequestParameter) ([]TeamUpload42, error) {
	return c.backGetTeamUploads(params, defaultRequest, nil)
}

// By Team

// Returns all the TeamUploads of the given Team.
func (c *Client42) GetTeamUploadsByTeam(id interface{}, params *RequestParameter) ([]TeamUpload42, error) {
	return c.backGetTeamUploads(params, teamRequest, id)
}

// Returns all the TeamUploads of the given Team.
func (t *Team42) GetTeamUploads(params *RequestParameter) ([]TeamUpload42, error) {
	return t.client.GetTeamUploadsByTeam(t.ID, params)
}
