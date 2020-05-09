package api42

type ProjectUser42 struct {
	ID            int        `json:"id"`
	Occurrence    int        `json:"occurrence"`
	FinalMark     *int       `json:"final_mark"`
	Status        string     `json:"status"`
	Validated     bool       `json:"validated?"`
	CurrentTeamID int        `json:"current_team_id"`
	Project       *Project42 `json:"project"`
	CursusIds     []int      `json:"cursus_ids"`
	User          *User42    `json:"user"`
	Teams         []Team42   `json:"teams"`
	client        *Client42
}

func (c *Client42) backGetProjectUsers(params *RequestParameter, directFilter string, value interface{}) ([]ProjectUser42, error) {
	var projectUserArray []ProjectUser42
	url, err := buildUrl(projectUserRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &projectUserArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(projectUserArray); i++ {
		projectUserArray[i].client = c
	}
	return projectUserArray, nil
}

// Returns the ProjectUser specified by the id.
func (c *Client42) GetProjectUser(id interface{}) (*ProjectUser42, error) {
	var projectUser *ProjectUser42
	url, err := buildUrl(projectUserRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &projectUser)
	if err != nil {
		return nil, err
	}
	projectUser.client = c
	return projectUser, nil
}

// Returns all the ProjectUsers.
func (c *Client42) GetProjectUsers(params *RequestParameter) ([]ProjectUser42, error) {
	return c.backGetProjectUsers(params, defaultRequest, nil)
}

// By User

// Returns all the ProjectUsers of the given User.
func (c *Client42) GetProjectUsersByUser(id interface{}, params *RequestParameter) ([]ProjectUser42, error) {
	return c.backGetProjectUsers(params, userRequest, id)
}

// Returns all the ProjectUsers of the given User.
func (c *User42) GetProjectUsers(params *RequestParameter) ([]ProjectUser42, error) {
	return c.client.GetProjectUsersByUser(c.ID, params)
}

// By Project

// Returns all the ProjectUsers of the given Project.
func (c *Client42) GetProjectUsersByProject(id interface{}, params *RequestParameter) ([]ProjectUser42, error) {
	return c.backGetProjectUsers(params, projectRequest, id)
}

// Returns all the ProjectUsers of the given Project.
func (c *Project42) GetProjectUsers(params *RequestParameter) ([]ProjectUser42, error) {
	return c.client.GetProjectUsersByProject(c.ID, params)
}
