package api42

type User42 struct {
	ID              int                 `json:"id"`
	Email           string              `json:"email"`
	Login           string              `json:"login"`
	FirstName       string              `json:"first_name"`
	LastName        string              `json:"last_name"`
	URL             string              `json:"url"`
	Phone           string              `json:"phone"`
	Displayname     string              `json:"displayname"`
	ImageURL        string              `json:"image_url"`
	Staff           bool                `json:"staff?"`
	CorrectionPoint int                 `json:"correction_point"`
	PoolMonth       string              `json:"pool_month"`
	PoolYear        string              `json:"pool_year"`
	Location        string              `json:"location"`
	Wallet          int                 `json:"wallet"`
	CursusUsers     []CursusUser42      `json:"cursus_users"`
	Projects        []ProjectUser42     `json:"projects_users"`
	Languages       []LanguageUser42    `json:"languages_users"`
	Achievements    []Achievement42     `json:"achievements"`
	Titles          []Title42           `json:"titles"`
	TitleUsers      []TitleUser42       `json:"titles_users"`
	Partnerships    []PartnershipUser42 `json:"partnerships"`
	Patroned        []User42            `json:"patroned"`
	Patroning       []User42            `json:"patroning"`
	ExpertisesUsers []ExpertiseUser42   `json:"expertises_users"`
	Campus          []Campus42          `json:"campus"`
	CampusUsers     []CampusUser42      `json:"campus_users"`
	client          *Client42
}

func (c *Client42) backGetUsers(params *RequestParameter, directFilter string, value interface{}) ([]User42, error) {
	var userArray []User42
	url, err := buildUrl(userRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &userArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(userArray); i++ {
		userArray[i].client = c
	}
	return userArray, nil
}

// Returns the User specified by the id.
func (c *Client42) GetUser(id interface{}) (*User42, error) {
	var user *User42
	url, err := buildUrl(userRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &user)
	if err != nil {
		return nil, err
	}
	user.client = c
	return user, nil
}

// Returns all the Users.
func (c *Client42) GetUsers(params *RequestParameter) ([]User42, error) {
	return c.backGetUsers(params, defaultRequest, nil)
}

// By Achievement

// Returns all the Users of the given Achievement.
func (c *Client42) GetUsersByAchievement(id interface{}, params *RequestParameter) ([]User42, error) {
	return c.backGetUsers(params, achievementRequest, id)
}

// Returns all the Users of the given Achievement.
func (c *Achievement42) GetUsers(params *RequestParameter) ([]User42, error) {
	return c.client.GetUsersByAchievement(c.ID, params)
}

// By Title

// Returns all the Users of the given Title.
func (c *Client42) GetUsersByTitle(id interface{}, params *RequestParameter) ([]User42, error) {
	return c.backGetUsers(params, titleRequest, id)
}

// Returns all the Users of the given Title.
func (c *Title42) GetUsers(params *RequestParameter) ([]User42, error) {
	return c.client.GetUsersByTitle(c.ID, params)
}

// By Campus

// Returns all the Users of the given Campus.
func (c *Client42) GetUsersByCampus(id interface{}, params *RequestParameter) ([]User42, error) {
	return c.backGetUsers(params, campusRequest, id)
}

// Returns all the Users of the given Campus.
func (c *Campus42) GetUsers(params *RequestParameter) ([]User42, error) {
	return c.client.GetUsersByCampus(c.ID, params)
}

// By Coalition

// Returns all the Users of the given Coalition.
func (c *Client42) GetUsersByCoalition(id interface{}, params *RequestParameter) ([]User42, error) {
	return c.backGetUsers(params, coalitionRequest, id)
}

// Returns all the Users of the given Coalition.
func (c *Coalition42) GetUsers(params *RequestParameter) ([]User42, error) {
	return c.client.GetUsersByCoalition(c.ID, params)
}

// By Cursus

// Returns all the Users of the given Cursus.
func (c *Client42) GetUsersByCursus(id interface{}, params *RequestParameter) ([]User42, error) {
	return c.backGetUsers(params, cursusRequest, id)
}

// Returns all the Users of the given Cursus.
func (c *Cursus42) GetUsers(params *RequestParameter) ([]User42, error) {
	return c.client.GetUsersByCursus(c.ID, params)
}

// By Event

// Returns all the Users of the given Event.
func (c *Client42) GetUsersByEvent(id interface{}, params *RequestParameter) ([]User42, error) {
	return c.backGetUsers(params, eventRequest, id)
}

// Returns all the Users of the given Event.
func (c *Event42) GetUsers(params *RequestParameter) ([]User42, error) {
	return c.client.GetUsersByEvent(c.ID, params)
}

// By Project

// Returns all the Users of the given Project.
func (c *Client42) GetUsersByProject(id interface{}, params *RequestParameter) ([]User42, error) {
	return c.backGetUsers(params, projectRequest, id)
}

// Returns all the Users of the given Project.
func (p *Project42) GetUsers(params *RequestParameter) ([]User42, error) {
	return p.client.GetUsersByProject(p.ID, params)
}
