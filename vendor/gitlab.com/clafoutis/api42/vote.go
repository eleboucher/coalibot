package api42

import (
	"time"
)

const (
	upvote    = "upvotes"
	downvote  = "downvotes"
	trollvote = "trollvotes"
	problems  = "problems"
)

type Vote42 struct {
	ID        int        `json:"id"`
	Kind      string     `json:"kind"`
	CreatedAt *time.Time `json:"created_at"`
	User      *User42    `json:"user"`
	Message   *Message42 `json:"message"`
	client    *Client42
}

func (c *Client42) backGetVotes(params *RequestParameter, directFilter string, value interface{}, args ...interface{}) ([]Vote42, error) {
	var voteArray []Vote42
	url, err := buildUrl(voteRequest, directFilter, value, args...)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, params, &voteArray)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(voteArray); i++ {
		voteArray[i].client = c
	}
	return voteArray, nil
}

// Returns the Vote specified by the id.
func (c *Client42) GetVote(id interface{}) (*Vote42, error) {
	var vote *Vote42
	url, err := buildUrl(voteRequest, defaultRequest, id)
	if err != nil {
		return nil, err
	}
	_, err = c.Get(url, nil, &vote)
	if err != nil {
		return nil, err
	}
	vote.client = c
	return vote, nil
}

// Returns all the Votes.
func (c *Client42) GetVotes(params *RequestParameter) ([]Vote42, error) {
	return c.backGetVotes(params, defaultRequest, nil)
}

// Returns all the Upvotes.
func (c *Client42) GetUpvotes(params *RequestParameter) ([]Vote42, error) {
	return c.backGetVotes(params, defaultRequest, upvote)
}

// Returns all the Downvotes.
func (c *Client42) GetDownvotes(params *RequestParameter) ([]Vote42, error) {
	return c.backGetVotes(params, defaultRequest, downvote)
}

// Returns all the Trollvotes.
func (c *Client42) GetTrollvotes(params *RequestParameter) ([]Vote42, error) {
	return c.backGetVotes(params, defaultRequest, trollvote)
}

// Returns all the ProblemVotes.
func (c *Client42) GetProblemVotes(params *RequestParameter) ([]Vote42, error) {
	return c.backGetVotes(params, defaultRequest, problems)
}

// By Message

// Returns all the Votes of the given Message.
func (c *Client42) GetVotesByMessage(id interface{}, params *RequestParameter) ([]Vote42, error) {
	return c.backGetVotes(params, messageRequest, id)
}

// Returns all the Upvotes of the given Message.
func (c *Client42) GetUpvotesByMessage(id interface{}, params *RequestParameter) ([]Vote42, error) {
	return c.backGetVotes(params, messageRequest, id, defaultRequest, upvote)
}

// Returns all the Downvotes of the given Message.
func (c *Client42) GetDownvotesByMessage(id interface{}, params *RequestParameter) ([]Vote42, error) {
	return c.backGetVotes(params, messageRequest, id, defaultRequest, downvote)
}

// Returns all the Trollvotes of the given Message.
func (c *Client42) GetTrollvotesByMessage(id interface{}, params *RequestParameter) ([]Vote42, error) {
	return c.backGetVotes(params, messageRequest, id, defaultRequest, trollvote)
}

// Returns all the ProblemVotes of the given Message.
func (c *Client42) GetProblemVotesByMessage(id interface{}, params *RequestParameter) ([]Vote42, error) {
	return c.backGetVotes(params, messageRequest, id, defaultRequest, problems)
}

// Returns all the Votes of the given Message.
func (m *Message42) GetVotes(params *RequestParameter) ([]Vote42, error) {
	return m.client.GetVotesByMessage(m.ID, params)
}

// Returns all the Upvotes of the given Message.
func (m *Message42) GetUpvotes(params *RequestParameter) ([]Vote42, error) {
	return m.client.GetUpvotesByMessage(m.ID, params)
}

// Returns all the Downvotes of the given Message.
func (m *Message42) GetDownvotes(params *RequestParameter) ([]Vote42, error) {
	return m.client.GetDownvotesByMessage(m.ID, params)
}

// Returns all the Trollvotes of the given Message.
func (m *Message42) GetTrollvotes(params *RequestParameter) ([]Vote42, error) {
	return m.client.GetTrollvotesByMessage(m.ID, params)
}

// Returns all the ProblemVotes of the given Message.
func (m *Message42) GetProblemVotes(params *RequestParameter) ([]Vote42, error) {
	return m.client.GetProblemVotesByMessage(m.ID, params)
}

// By Topic

// Returns all the Votes of the given Topic.
func (c *Client42) GetVotesByTopic(id interface{}, params *RequestParameter) ([]Vote42, error) {
	return c.backGetVotes(params, topicRequest, id)
}

// Returns all the Upvotes of the given Topic.
func (c *Client42) GetUpvotesByTopic(id interface{}, params *RequestParameter) ([]Vote42, error) {
	return c.backGetVotes(params, topicRequest, id, defaultRequest, upvote)
}

// Returns all the Downvotes of the given Topic.
func (c *Client42) GetDownvotesByTopic(id interface{}, params *RequestParameter) ([]Vote42, error) {
	return c.backGetVotes(params, topicRequest, id, defaultRequest, downvote)
}

// Returns all the Trollvotes of the given Topic.
func (c *Client42) GetTrollvotesByTopic(id interface{}, params *RequestParameter) ([]Vote42, error) {
	return c.backGetVotes(params, topicRequest, id, defaultRequest, trollvote)
}

// Returns all the ProblemVotes of the given Topic.
func (c *Client42) GetProblemVotesByTopic(id interface{}, params *RequestParameter) ([]Vote42, error) {
	return c.backGetVotes(params, topicRequest, id, defaultRequest, problems)
}

// Returns all the Votes of the given Topic.
func (m *Topic42) GetVotes(params *RequestParameter) ([]Vote42, error) {
	return m.client.GetVotesByTopic(m.ID, params)
}

// Returns all the Upvotes of the given Topic.
func (m *Topic42) GetUpvotes(params *RequestParameter) ([]Vote42, error) {
	return m.client.GetUpvotesByTopic(m.ID, params)
}

// Returns all the Downvotes of the given Topic.
func (m *Topic42) GetDownvotes(params *RequestParameter) ([]Vote42, error) {
	return m.client.GetDownvotesByTopic(m.ID, params)
}

// Returns all the Trollvotes of the given Topic.
func (m *Topic42) GetTrollvotes(params *RequestParameter) ([]Vote42, error) {
	return m.client.GetTrollvotesByTopic(m.ID, params)
}

// Returns all the ProblemVotes of the given Topic.
func (m *Topic42) GetProblemVotes(params *RequestParameter) ([]Vote42, error) {
	return m.client.GetProblemVotesByTopic(m.ID, params)
}
