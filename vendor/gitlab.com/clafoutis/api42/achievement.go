package api42

type Achievement42 struct {
	ID           int             `json:"id"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	Tier         string          `json:"tier"`
	Kind         string          `json:"kind"`
	Visible      bool            `json:"visible"`
	Image        string          `json:"image"`
	NbrOfSuccess int             `json:"nbr_of_success"`
	UsersURL     string          `json:"users_url"`
	Achievements []Achievement42 `json:"achievements"`
	Parent       *Achievement42  `json:"parent"`
	Title        *Title42        `json:"title"`
	client       *Client42
}

func (c *Client42) backGetAchievements(params *RequestParameter, directFilter string, value interface{}) ([]Achievement42, error) {
	var achievementArray []Achievement42
	url, err := buildUrl(achievementRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &achievementArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(achievementArray); i++ {
		achievementArray[i].client = c
	}
	return achievementArray, nil
}

// Returns the Achievement specified by the id.
func (c *Client42) GetAchievement(id interface{}) (*Achievement42, error) {
	var achievement *Achievement42
	url, err := buildUrl(achievementRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &achievement)
	if err != nil {
		return nil, err
	}
	achievement.client = c
	return achievement, nil
}

// Returns all the Achievements.
func (c *Client42) GetAchievements(params *RequestParameter) ([]Achievement42, error) {
	return c.backGetAchievements(params, defaultRequest, nil)
}

// By Cursus

// Returns all the Achievements of the given Cursus.
func (c *Client42) GetAchievementsByCursus(id interface{}, params *RequestParameter) ([]Achievement42, error) {
	return c.backGetAchievements(params, cursusRequest, id)
}

// Returns all the Achievements of the given Cursus.
func (u *Cursus42) GetAchievements(params *RequestParameter) ([]Achievement42, error) {
	return u.client.GetAchievementsByCursus(u.ID, params)
}

// By Campus

// Returns all the Achievements of the given Campus.
func (c *Client42) GetAchievementsByCampus(id interface{}, params *RequestParameter) ([]Achievement42, error) {
	return c.backGetAchievements(params, campusRequest, id)
}

// Returns all the Achievements of the given Campus.
func (u *Campus42) GetAchievements(params *RequestParameter) ([]Achievement42, error) {
	return u.client.GetAchievementsByCampus(u.ID, params)
}

// By Title

// Returns all the Achievements of the given Title.
func (c *Client42) GetAchievementsByTitle(id interface{}) ([]Achievement42, error) {
	return c.backGetAchievements(nil, titleRequest, id)
}

// Returns all the Achievement of the given Title.
func (u *Title42) GetAchievement() ([]Achievement42, error) {
	return u.client.GetAchievementsByTitle(u.ID)
}
