package api42

import (
	"time"
)

type Topic42 struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Author      *User42     `json:"author"`
	Kind        string      `json:"kind"`
	CreatedAt   *time.Time  `json:"created_at"`
	UpdatedAt   *time.Time  `json:"updated_at"`
	PinnedAt    *time.Time  `json:"pinned_at"`
	LockedAt    *time.Time  `json:"locked_at"`
	Pinner      *User42     `json:"pinner"`
	Locker      *User42     `json:"locker"`
	Language    *Language42 `json:"language"`
	MessagesURL string      `json:"messages_url"`
	Message     *Message42  `json:"message"`
	Tags        []Tag42     `json:"tags"`
	client      *Client42
}

func (c *Client42) backGetTopics(params *RequestParameter, directFilter string, value interface{}) ([]Topic42, error) {
	var topicArray []Topic42
	url, err := buildUrl(topicRequest, directFilter, value)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &topicArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(topicArray); i++ {
		topicArray[i].client = c
	}
	return topicArray, nil
}

// Returns the Topic specified by the id.
func (c *Client42) GetTopic(id interface{}) (*Topic42, error) {
	var topic *Topic42
	url, err := buildUrl(topicRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &topic)
	if err != nil {
		return nil, err
	}
	topic.client = c
	return topic, nil
}

// Returns all the Topics.
func (c *Client42) GetTopics(params *RequestParameter) ([]Topic42, error) {
	return c.backGetTopics(params, defaultRequest, nil)
}

// Returns all the UnreadTopics.
func (c *Client42) GetUnreadTopics(params *RequestParameter) ([]Topic42, error) {
	return c.backGetTopics(params, defaultRequest, "unread")
}

// By Cursus

// Returns all the Topics of the given Cursus.
func (c *Client42) GetTopicsByCursus(id interface{}, params *RequestParameter) ([]Topic42, error) {
	return c.backGetTopics(params, cursusRequest, id)
}

// Returns all the Topics of the given Cursus.
func (c *Cursus42) GetTopics(params *RequestParameter) ([]Topic42, error) {
	return c.client.GetTopicsByCursus(c.ID, params)
}

// By Tag

// Returns all the Topics of the given Tag.
func (c *Client42) GetTopicsByTag(id interface{}, params *RequestParameter) ([]Topic42, error) {
	return c.backGetTopics(params, tagRequest, id)
}

// Returns all the Topics of the given Tag.
func (c *Tag42) GetTopics(params *RequestParameter) ([]Topic42, error) {
	return c.client.GetTopicsByTag(c.ID, params)
}

// By User

// Returns all the Topics of the given User.
func (c *Client42) GetTopicsByUser(id interface{}, params *RequestParameter) ([]Topic42, error) {
	return c.backGetTopics(params, userRequest, id)
}

// Returns all the Topics of the given User.
func (u *User42) GetTopics(params *RequestParameter) ([]Topic42, error) {
	return u.client.GetTopicsByUser(u.ID, params)
}
