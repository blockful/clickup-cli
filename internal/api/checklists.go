package api

import (
	"context"
	"fmt"
)

type ChecklistItem struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	OrderIndex  interface{}   `json:"orderindex"`
	Assignee    interface{}   `json:"assignee"`
	Resolved    bool          `json:"resolved"`
	Parent      interface{}   `json:"parent"`
	DateCreated string        `json:"date_created,omitempty"`
	Children    []interface{} `json:"children,omitempty"`
}

// Checklist is defined in tasks.go

type ChecklistDetailed struct {
	ID          string          `json:"id"`
	TaskID      string          `json:"task_id"`
	Name        string          `json:"name"`
	DateCreated string          `json:"date_created,omitempty"`
	OrderIndex  interface{}     `json:"orderindex"`
	Resolved    int             `json:"resolved"`
	Unresolved  int             `json:"unresolved"`
	Items       []ChecklistItem `json:"items"`
}

type ChecklistResponse struct {
	Checklist ChecklistDetailed `json:"checklist"`
}

type CreateChecklistRequest struct {
	Name string `json:"name"`
}

type EditChecklistRequest struct {
	Name     string `json:"name,omitempty"`
	Position *int   `json:"position,omitempty"`
}

type CreateChecklistItemRequest struct {
	Name     string `json:"name,omitempty"`
	Assignee *int   `json:"assignee,omitempty"`
}

type EditChecklistItemRequest struct {
	Name     string  `json:"name,omitempty"`
	Assignee *string `json:"assignee,omitempty"`
	Resolved *bool   `json:"resolved,omitempty"`
	Parent   *string `json:"parent,omitempty"`
}

func (c *Client) CreateChecklist(ctx context.Context, taskID string, req *CreateChecklistRequest, opts ...*TaskScopedOptions) (*ChecklistResponse, error) {
	var o *TaskScopedOptions
	if len(opts) > 0 { o = opts[0] }
	var resp ChecklistResponse
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/task/%s/checklist", taskID)+taskScopedQuery(o), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) EditChecklist(ctx context.Context, checklistID string, req *EditChecklistRequest) error {
	return c.Do(ctx, "PUT", fmt.Sprintf("/v2/checklist/%s", checklistID), req, nil)
}

func (c *Client) DeleteChecklist(ctx context.Context, checklistID string) error {
	return c.Do(ctx, "DELETE", fmt.Sprintf("/v2/checklist/%s", checklistID), nil, nil)
}

func (c *Client) CreateChecklistItem(ctx context.Context, checklistID string, req *CreateChecklistItemRequest) (*ChecklistResponse, error) {
	var resp ChecklistResponse
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/checklist/%s/checklist_item", checklistID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) EditChecklistItem(ctx context.Context, checklistID, checklistItemID string, req *EditChecklistItemRequest) (*ChecklistResponse, error) {
	var resp ChecklistResponse
	if err := c.Do(ctx, "PUT", fmt.Sprintf("/v2/checklist/%s/checklist_item/%s", checklistID, checklistItemID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteChecklistItem(ctx context.Context, checklistID, checklistItemID string) error {
	return c.Do(ctx, "DELETE", fmt.Sprintf("/v2/checklist/%s/checklist_item/%s", checklistID, checklistItemID), nil, nil)
}
