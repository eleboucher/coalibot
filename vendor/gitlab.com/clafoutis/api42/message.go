package api42

import (
	"time"
)

type Message42 struct {
	ID              int         `json:"id"`
	Author          *User42     `json:"author"`
	Content         string      `json:"content"`
	ContentHTML     string      `json:"content_html"`
	Replies         []Message42 `json:"replies"`
	CreatedAt       *time.Time  `json:"created_at"`
	UpdatedAt       *time.Time  `json:"updated_at"`
	ParentID        *int        `json:"parent_id"`
	IsRoot          bool        `json:"is_root"`
	MessageableID   int         `json:"messageable_id"`
	MessageableType string      `json:"messageable_type"`
	VotesCount      struct {
		Upvote    int `json:"upvote"`
		Downvote  int `json:"downvote"`
		Trollvote int `json:"trollvote"`
		Problem   int `json:"problem"`
	} `json:"votes_count"`
	UserVotes struct {
		Upvote    bool `json:"upvote"`
		Downvote  bool `json:"downvote"`
		Trollvote bool `json:"trollvote"`
		Problem   bool `json:"problem"`
	} `json:"user_votes"`
	Readings int `json:"readings"`
	client   *Client42
}

func (c *Client42) backGetMessages(params *RequestParameter, directFilter string, value interface{}, args ...interface{}) ([]Message42, error) {
	var messageArray []Message42
	url, err := buildUrl(messageRequest, directFilter, value, args...)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &messageArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(messageArray); i++ {
		messageArray[i].client = c
	}
	return messageArray, nil
}

// Returns the Message specified by the id.
func (c *Client42) GetMessage(id interface{}) (*Message42, error) {
	var message *Message42
	url, err := buildUrl(messageRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &message)
	if err != nil {
		return nil, err
	}
	message.client = c
	return message, nil
}

// Returns all the Messages.
func (c *Client42) GetMessages(params *RequestParameter) ([]Message42, error) {
	return c.backGetMessages(params, defaultRequest, nil)
}

// By Topic

// Returns all the Messages of the given Topic.
func (c *Client42) GetMessagesByTopic(id interface{}, params *RequestParameter) ([]Message42, error) {
	return c.backGetMessages(params, userRequest, id)
}

// Returns all the Messages of the given Message, associated with the given Topic.
func (c *Client42) GetMessagesByTopicByMessage(topicID, messageID interface{}, params *RequestParameter) ([]Message42, error) {
	return c.backGetMessages(params, topicRequest, topicID, messageRequest, messageID)
}

// Returns all the MessagesByMessage of the given Topic.
func (t *Topic42) GetMessagesByMessage(id interface{}, params *RequestParameter) ([]Message42, error) {
	return t.client.GetMessagesByTopicByMessage(t.ID, id, params)
}

// Returns all the MessagesByTopic of the given Message.
func (m *Message42) GetMessagesByTopic(id interface{}, params *RequestParameter) ([]Message42, error) {
	return m.client.GetMessagesByTopicByMessage(id, m.ID, params)
}

// Returns all the Messages of the given Topic.
func (t *Topic42) GetMessages(params *RequestParameter) ([]Message42, error) {
	return t.client.GetMessagesByTopic(t.ID, params)
}

// By User

// Returns all the Messages of the given User.
func (c *Client42) GetMessagesByUser(id interface{}, params *RequestParameter) ([]Message42, error) {
	return c.backGetMessages(params, projectRequest, id)
}

// Returns all the Messages of the given User.
func (u *User42) GetMessages(params *RequestParameter) ([]Message42, error) {
	return u.client.GetMessagesByUser(u.ID, params)
}
