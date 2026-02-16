package api

import "fmt"

type Comment struct {
	ID          string      `json:"id"`
	CommentText string     `json:"comment_text"`
	User        User        `json:"user"`
	Date        string      `json:"date"`
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

type CreateCommentRequest struct {
	CommentText string `json:"comment_text"`
}

type CreateCommentResponse struct {
	ID      string `json:"id"`
	HistID  string `json:"hist_id"`
	Date    int64  `json:"date"`
}

func (c *Client) CreateComment(taskID string, req *CreateCommentRequest) (*CreateCommentResponse, error) {
	var resp CreateCommentResponse
	if err := c.Do("POST", fmt.Sprintf("/v2/task/%s/comment", taskID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
