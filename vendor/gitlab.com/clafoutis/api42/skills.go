package api42

import (
	"time"
)

type Skill42 struct {
	CreatedAt time.Time `json:"created_at"`
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	client    *Client42
}

func (c *Client42) backGetSkills(params *RequestParameter, directFilter string, value interface{}) ([]Skill42, error) {
	var skillArray []Skill42
	url, err := buildUrl(skillRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &skillArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(skillArray); i++ {
		skillArray[i].client = c
	}
	return skillArray, nil
}

// Returns the Skill specified by the id.
func (c *Client42) GetSkill(id interface{}) (*Skill42, error) {
	var skill *Skill42
	url, err := buildUrl(skillRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &skill)
	if err != nil {
		return nil, err
	}
	skill.client = c
	return skill, nil
}

// Returns all the Skills.
func (c *Client42) GetSkills(params *RequestParameter) ([]Skill42, error) {
	return c.backGetSkills(params, defaultRequest, nil)
}

// By Cursus

// Returns all the Skills of the given Cursus.
func (c *Client42) GetSkillsByCursus(id interface{}, params *RequestParameter) ([]Skill42, error) {
	return c.backGetSkills(params, cursusRequest, id)
}

// Returns all the Skills of the given Cursus.
func (c *Cursus42) GetSkills(params *RequestParameter) ([]Skill42, error) {
	return c.client.GetSkillsByCursus(c.ID, params)
}

// By Project

// Returns all the Skills of the given Project.
func (c *Client42) GetSkillsByProject(id interface{}, params *RequestParameter) ([]Skill42, error) {
	return c.backGetSkills(params, projectRequest, id)
}

// Returns all the Skills of the given Project.
func (c *Project42) GetSkills(params *RequestParameter) ([]Skill42, error) {
	return c.client.GetSkillsByProject(c.ID, params)
}
