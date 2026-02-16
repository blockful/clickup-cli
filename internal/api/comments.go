package api

import (
	"context"
	"fmt"
	"net/url"
)

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

func (c *Client) ListComments(ctx context.Context, taskID, startID string, opts ...*TaskScopedOptions) (*CommentsResponse, error) {
	params := url.Values{}
	if startID != "" {
		params.Set("start_id", startID)
	}
	if len(opts) > 0 && opts[0] != nil {
		if opts[0].CustomTaskIDs {
			params.Set("custom_task_ids", "true")
		}
		if opts[0].TeamID != "" {
			params.Set("team_id", opts[0].TeamID)
		}
	}
	path := fmt.Sprintf("/v2/task/%s/comment", taskID)
	if q := params.Encode(); q != "" {
		path += "?" + q
	}
	var resp CommentsResponse
	if err := c.Do(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListListComments(ctx context.Context, listID, startID string) (*CommentsResponse, error) {
	path := fmt.Sprintf("/v2/list/%s/comment", listID)
	if startID != "" {
		path += "?start_id=" + startID
	}
	var resp CommentsResponse
	if err := c.Do(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateCommentRequest struct {
	CommentText   string `json:"comment_text"`
	Assignee      *int   `json:"assignee,omitempty"`
	GroupAssignee *int   `json:"group_assignee,omitempty"`
	NotifyAll     bool   `json:"notify_all,omitempty"`
}

type CreateCommentResponse struct {
	ID     string `json:"id"`
	HistID string `json:"hist_id"`
	Date   int64  `json:"date"`
}

func (c *Client) CreateComment(ctx context.Context, taskID string, req *CreateCommentRequest, opts ...*TaskScopedOptions) (*CreateCommentResponse, error) {
	var o *TaskScopedOptions
	if len(opts) > 0 {
		o = opts[0]
	}
	var resp CreateCommentResponse
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/task/%s/comment", taskID)+taskScopedQuery(o), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CreateListComment(ctx context.Context, listID string, req *CreateCommentRequest) (*CreateCommentResponse, error) {
	var resp CreateCommentResponse
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/list/%s/comment", listID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateCommentRequest struct {
	CommentText   string `json:"comment_text"`
	Assignee      *int   `json:"assignee,omitempty"`
	GroupAssignee *int   `json:"group_assignee,omitempty"`
	Resolved      *bool  `json:"resolved,omitempty"`
}

func (c *Client) UpdateComment(ctx context.Context, commentID string, req *UpdateCommentRequest) error {
	return c.Do(ctx, "PUT", fmt.Sprintf("/v2/comment/%s", commentID), req, nil)
}

func (c *Client) DeleteComment(ctx context.Context, commentID string) error {
	return c.Do(ctx, "DELETE", fmt.Sprintf("/v2/comment/%s", commentID), nil, nil)
}

// Threaded Comments

func (c *Client) ListThreadedComments(ctx context.Context, commentID string) (*CommentsResponse, error) {
	var resp CommentsResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/comment/%s/reply", commentID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CreateThreadedComment(ctx context.Context, commentID string, req *CreateCommentRequest) (*CreateCommentResponse, error) {
	var resp CreateCommentResponse
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/comment/%s/reply", commentID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Chat View Comments

func (c *Client) ListViewComments(ctx context.Context, viewID, startID string) (*CommentsResponse, error) {
	path := fmt.Sprintf("/v2/view/%s/comment", viewID)
	if startID != "" {
		path += "?start_id=" + startID
	}
	var resp CommentsResponse
	if err := c.Do(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CreateViewComment(ctx context.Context, viewID string, req *CreateCommentRequest) (*CreateCommentResponse, error) {
	var resp CreateCommentResponse
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/view/%s/comment", viewID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
