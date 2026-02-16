package api

import (
	"context"
	"fmt"
)

type List struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	OrderIndex int    `json:"orderindex"`
	Content    string `json:"content,omitempty"`
	Status     struct {
		Status string `json:"status"`
		Color  string `json:"color"`
	} `json:"status,omitempty"`
	Priority struct {
		Priority string `json:"priority"`
		Color    string `json:"color"`
	} `json:"priority,omitempty"`
	Assignee  interface{} `json:"assignee"`
	DueDate   string      `json:"due_date,omitempty"`
	TaskCount int         `json:"task_count"`
	Folder    struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Hidden bool   `json:"hidden"`
	} `json:"folder"`
	Space struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"space"`
}

type ListsResponse struct {
	Lists []List `json:"lists"`
}

func (c *Client) ListLists(ctx context.Context, folderID string) (*ListsResponse, error) {
	var resp ListsResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/folder/%s/list?archived=false", folderID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListFolderlessLists(ctx context.Context, spaceID string) (*ListsResponse, error) {
	var resp ListsResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/space/%s/list?archived=false", spaceID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetList(ctx context.Context, listID string) (*List, error) {
	var resp List
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/list/%s", listID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateListRequest struct {
	Name            string `json:"name"`
	Content         string `json:"content,omitempty"`
	MarkdownContent string `json:"markdown_content,omitempty"`
	DueDate         *int64 `json:"due_date,omitempty"`
	DueDateTime     *bool  `json:"due_date_time,omitempty"`
	Priority        *int   `json:"priority,omitempty"`
	Assignee        *int   `json:"assignee,omitempty"`
	Status          string `json:"status,omitempty"`
}

func (c *Client) CreateList(ctx context.Context, folderID string, req *CreateListRequest) (*List, error) {
	var resp List
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/folder/%s/list", folderID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CreateFolderlessList(ctx context.Context, spaceID string, req *CreateListRequest) (*List, error) {
	var resp List
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/space/%s/list", spaceID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateListRequest struct {
	Name            string `json:"name,omitempty"`
	Content         string `json:"content,omitempty"`
	MarkdownContent string `json:"markdown_content,omitempty"`
	DueDate         *int64 `json:"due_date,omitempty"`
	DueDateTime     *bool  `json:"due_date_time,omitempty"`
	Priority        *int   `json:"priority,omitempty"`
	Assignee        *int   `json:"assignee,omitempty"`
	Status          string `json:"status,omitempty"`
	UnsetStatus     bool   `json:"unset_status,omitempty"`
}

func (c *Client) UpdateList(ctx context.Context, listID string, req *UpdateListRequest) (*List, error) {
	var resp List
	if err := c.Do(ctx, "PUT", fmt.Sprintf("/v2/list/%s", listID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteList(ctx context.Context, listID string) error {
	return c.Do(ctx, "DELETE", fmt.Sprintf("/v2/list/%s", listID), nil, nil)
}
