package api42

import (
	"time"
)

type Feedback42 struct {
	ID              int        `json:"id"`
	Comment         string     `json:"comment"`
	CreatedAt       *time.Time `json:"created_at"`
	FeedbackDetails []struct {
		ID   int    `json:"id"`
		Kind string `json:"kind"`
		Rate int    `json:"rate"`
	} `json:"feedback_details"`
	FeedbackableID   int     `json:"feedbackable_id"`
	FeedbackableType string  `json:"feedbackable_type"`
	Rating           *int    `json:"rating"`
	User             *User42 `json:"user"`
	client           *Client42
}

func (c *Client42) backGetFeedbacks(params *RequestParameter, directFilter string, value interface{}) ([]Feedback42, error) {
	var feedbackArray []Feedback42
	url, err := buildUrl(feedbackRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &feedbackArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(feedbackArray); i++ {
		feedbackArray[i].client = c
	}
	return feedbackArray, nil
}

// Returns the Feedback specified by the id.
func (c *Client42) GetFeedback(id interface{}) (*Feedback42, error) {
	var feedback *Feedback42
	url, err := buildUrl(feedbackRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &feedback)
	if err != nil {
		return nil, err
	}
	feedback.client = c
	return feedback, nil
}

// Returns all the Feedbacks.
func (c *Client42) GetFeedbacks(params *RequestParameter) ([]Feedback42, error) {
	return c.backGetFeedbacks(params, defaultRequest, nil)
}

// By Event

// Returns all the Feedbacks of the given Event.
func (c *Client42) GetFeedbacksByEvent(id interface{}, params *RequestParameter) ([]Feedback42, error) {
	return c.backGetFeedbacks(params, eventRequest, id)
}

// Returns all the Feedbacks of the given Event.
func (c *Event42) GetFeedbacks(params *RequestParameter) ([]Feedback42, error) {
	return c.client.GetFeedbacksByEvent(c.ID, params)
}

// By ScaleTeam

// Returns all the Feedbacks of the given ScaleTeam.
func (c *Client42) GetFeedbacksByScaleTeam(id interface{}, params *RequestParameter) ([]Feedback42, error) {
	return c.backGetFeedbacks(params, scaleTeamRequest, id)
}

// Returns all the Feedbacks of the given ScaleTeam.
func (c *ScaleTeam42) GetFeedbacks(params *RequestParameter) ([]Feedback42, error) {
	return c.client.GetFeedbacksByScaleTeam(c.ID, params)
}
