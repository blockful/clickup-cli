package api

import "fmt"

type Comment struct {
	ID          string        `json:"id"`
	CommentText string        `json:"comment_text"`
	User        User          `json:"user"`
	Date        string        `json:"date"`
	Comment     []CommentPart `json:"comment"`
}

type CommentPart struct {
	Text string `json:"text"`
}

type CommentsResponse struct {
	Comments []Comment `json:"comments"`
}

func (c *Client) ListComments(taskID string) (*CommentsResponse, error) {
	var resp CommentsResponse
	if err := c.Do("GET", fmt.Sprintf("/v2/task/%s/comment", taskID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListListComments(listID string) (*CommentsResponse, error) {
	var resp CommentsResponse
	if err := c.Do("GET", fmt.Sprintf("/v2/list/%s/comment", listID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateCommentRequest struct {
	CommentText string `json:"comment_text"`
	Assignee    *int   `json:"assignee,omitempty"`
	NotifyAll   bool   `json:"notify_all,omitempty"`
}

type CreateCommentResponse struct {
	ID     string `json:"id"`
	HistID string `json:"hist_id"`
	Date   int64  `json:"date"`
}

func (c *Client) CreateComment(taskID string, req *CreateCommentRequest) (*CreateCommentResponse, error) {
	var resp CreateCommentResponse
	if err := c.Do("POST", fmt.Sprintf("/v2/task/%s/comment", taskID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CreateListComment(listID string, req *CreateCommentRequest) (*CreateCommentResponse, error) {
	var resp CreateCommentResponse
	if err := c.Do("POST", fmt.Sprintf("/v2/list/%s/comment", listID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateCommentRequest struct {
	CommentText string `json:"comment_text"`
	Assignee    *int   `json:"assignee,omitempty"`
	Resolved    *bool  `json:"resolved,omitempty"`
}

func (c *Client) UpdateComment(commentID string, req *UpdateCommentRequest) error {
	return c.Do("PUT", fmt.Sprintf("/v2/comment/%s", commentID), req, nil)
}

func (c *Client) DeleteComment(commentID string) error {
	return c.Do("DELETE", fmt.Sprintf("/v2/comment/%s", commentID), nil, nil)
}
