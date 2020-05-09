package api42

import (
	"time"
)

type Scale42 struct {
	ID                 int            `json:"id"`
	EvaluationID       int            `json:"evaluation_id"`
	Name               string         `json:"name"`
	IsPrimary          bool           `json:"is_primary"`
	Comment            string         `json:"comment"`
	IntroductionMd     string         `json:"introduction_md"`
	DisclaimerMd       string         `json:"disclaimer_md"`
	GuidelinesMd       string         `json:"guidelines_md"`
	CreatedAt          *time.Time     `json:"created_at"`
	CorrectionNumber   int            `json:"correction_number"`
	Duration           *time.Duration `json:"duration"`
	ManualSubscription bool           `json:"manual_subscription"`
	// Languages          []Language42 `json:"languages"`
	// Sections []interface{} `json:"sections"`
	// Evaluation []Evaluation42 `json:"evaluation"`
	client *Client42
}

func (c *Client42) backGetScales(params *RequestParameter, directFilter string, value interface{}) ([]Scale42, error) {
	var scaleArray []Scale42
	url, err := buildUrl(scaleRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &scaleArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(scaleArray); i++ {
		scaleArray[i].client = c
	}
	return scaleArray, nil
}

// Returns the Scale specified by the id.
func (c *Client42) GetScale(id interface{}) (*Scale42, error) {
	var scale *Scale42
	url, err := buildUrl(scaleRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &scale)
	if err != nil {
		return nil, err
	}
	scale.client = c
	return scale, nil
}

// Returns all the Scales.
func (c *Client42) GetScales(params *RequestParameter) ([]Scale42, error) {
	return c.backGetScales(params, defaultRequest, nil)
}

// By User

// Returns all the Scales of the given User.
func (c *Client42) GetScalesByUser(id interface{}, params *RequestParameter) ([]Scale42, error) {
	return c.backGetScales(params, userRequest, id)
}

// Returns all the Scales of the given User.
func (c *User42) GetScales(params *RequestParameter) ([]Scale42, error) {
	return c.client.GetScalesByUser(c.ID, params)
}

// By Project

// Returns all the Scales of the given Project.
func (c *Client42) GetScalesByProject(id interface{}, params *RequestParameter) ([]Scale42, error) {
	return c.backGetScales(params, projectRequest, id)
}

// Returns all the Scales of the given Project.
func (c *Project42) GetScales(params *RequestParameter) ([]Scale42, error) {
	return c.client.GetScalesByProject(c.ID, params)
}
