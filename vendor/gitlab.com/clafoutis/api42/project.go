package api42

import (
	"time"
)

type Project42 struct {
	ID          int                      `json:"id"`
	Tier        int                      `json:"tier"`
	Name        string                   `json:"name"`
	Slug        string                   `json:"slug"`
	Description string                   `json:"description"`
	Exam        bool                     `json:"exam"`
	Children    []map[string]interface{} `json:'children'`
	CreatedAt   *time.Time               `json:"created_at"`
	UpdatedAt   *time.Time               `json:"updated_at"`
	Cursus      []Cursus42               `json:"cursus"`
	Tags        []Tag42                  `json:"tags"`
	Parent      *Project42               `json:"parent"`
	client      *Client42
}

func (c *Client42) backGetProjects(params *RequestParameter, directFilter string, value interface{}) ([]Project42, error) {
	var projectArray []Project42
	url, err := buildUrl(projectRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &projectArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(projectArray); i++ {
		projectArray[i].client = c
	}
	return projectArray, nil
}

// Returns the Project specified by the id.
func (c *Client42) GetProject(id interface{}) (*Project42, error) {
	var project *Project42
	url, err := buildUrl(projectRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &project)
	if err != nil {
		return nil, err
	}
	project.client = c
	return project, nil
}

// Returns all the Projects.
func (c *Client42) GetProjects(params *RequestParameter) ([]Project42, error) {
	return c.backGetProjects(params, defaultRequest, nil)
}

// By Cursus

// Returns all the ProjectsByCursus of the given Cursus.
func (c *Client42) GetProjectsByCursus(id interface{}, params *RequestParameter) ([]Project42, error) {
	return c.backGetProjects(params, cursusRequest, id)
}

// Returns all the Projects of the given Cursus.
func (c *Cursus42) GetProjects(params *RequestParameter) ([]Project42, error) {
	return c.client.GetProjectsByCursus(c.ID, params)
}

// By Project

// Returns all the SubprojectsByProjects of the given Project.
func (c *Client42) GetSubprojectsByProjects(id interface{}, params *RequestParameter) ([]Project42, error) {
	return c.backGetProjects(params, projectRequest, id)
}

// Returns all the Subprojects of the given Project.
func (c *Project42) GetSubprojects(params *RequestParameter) ([]Project42, error) {
	return c.client.GetSubprojectsByProjects(c.ID, params)
}
