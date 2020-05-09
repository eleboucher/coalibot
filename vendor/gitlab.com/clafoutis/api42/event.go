package api42

import (
	"time"
)

type Event42 struct {
	ID             int        `json:"id"`
	MaxPeople      int        `json:"max_people"`
	NbrSubscribers int        `json:"nbr_subscribers"`
	CampusIds      []int      `json:"campus_ids"`
	CursusIds      []int      `json:"cursus_ids"`
	Name           string     `json:"name"`
	Description    string     `json:"description"`
	Location       string     `json:"location"`
	Kind           string     `json:"kind"`
	BeginAt        *time.Time `json:"begin_at"`
	EndAt          *time.Time `json:"end_at"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	client         *Client42
}

func (c *Client42) backGetEvents(params *RequestParameter, directFilter string, value interface{}) ([]Event42, error) {
	var eventArray []Event42
	url, err := buildUrl(eventRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &eventArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(eventArray); i++ {
		eventArray[i].client = c
	}
	return eventArray, nil
}

// Returns the Event specified by the id.
func (c *Client42) GetEvent(id interface{}) (*Event42, error) {
	var event *Event42
	url, err := buildUrl(eventRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &event)
	if err != nil {
		return nil, err
	}
	event.client = c
	return event, nil
}

// Returns all the Events.
func (c *Client42) GetEvents(params *RequestParameter) ([]Event42, error) {
	return c.backGetEvents(params, defaultRequest, nil)
}

// By Cursus

// Returns all the Events of the given Cursus.
func (c *Client42) GetEventsByCursus(id interface{}, params *RequestParameter) ([]Event42, error) {
	return c.backGetEvents(params, cursusRequest, id)
}

// Returns all the Events of the given Cursus.
func (c *Cursus42) GetEvents(params *RequestParameter) ([]Event42, error) {
	return c.client.GetEventsByCursus(c.ID, params)
}

// By Campus

// Returns all the Events of the given Campus.
func (c *Client42) GetEventsByCampus(id interface{}, params *RequestParameter) ([]Event42, error) {
	return c.backGetEvents(params, campusRequest, id)
}

// Returns all the Events of the given Campus.
func (c *Campus42) GetEvents(params *RequestParameter) ([]Event42, error) {
	return c.client.GetEventsByCampus(c.ID, params)
}

// By User

// Returns all the Events of the given User.
func (c *Client42) GetEventsByUser(id interface{}, params *RequestParameter) ([]Event42, error) {
	return c.backGetEvents(params, userRequest, id)
}

// Returns all the Events of the given User.
func (c *User42) GetEvents(params *RequestParameter) ([]Event42, error) {
	return c.client.GetEventsByUser(c.ID, params)
}
