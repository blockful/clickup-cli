package api

import "fmt"

type List struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	OrderIndex int    `json:"orderindex"`
	Status     struct {
		Status string `json:"status"`
		Color  string `json:"color"`
	} `json:"status,omitempty"`
	Priority struct {
		Priority string `json:"priority"`
		Color    string `json:"color"`
	} `json:"priority,omitempty"`
	Assignee  interface{} `json:"assignee"`
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

func (c *Client) ListLists(folderID string) (*ListsResponse, error) {
	var resp ListsResponse
	if err := c.Do("GET", fmt.Sprintf("/v2/folder/%s/list?archived=false", folderID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetList(listID string) (*List, error) {
	var resp List
	if err := c.Do("GET", fmt.Sprintf("/v2/list/%s", listID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateListRequest struct {
	Name string `json:"name"`
}

func (c *Client) CreateList(folderID string, req *CreateListRequest) (*List, error) {
	var resp List
	if err := c.Do("POST", fmt.Sprintf("/v2/folder/%s/list", folderID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
