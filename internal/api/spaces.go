package api

import (
	"context"
	"fmt"
)

type Space struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Private bool   `json:"private"`
	Status  struct {
		Status string `json:"status"`
		Color  string `json:"color"`
	} `json:"status"`
	Multiple bool                   `json:"multiple_assignees"`
	Features map[string]interface{} `json:"features"`
}

type SpacesResponse struct {
	Spaces []Space `json:"spaces"`
}

func (c *Client) ListSpaces(ctx context.Context, workspaceID string) (*SpacesResponse, error) {
	var resp SpacesResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/team/%s/space?archived=false", workspaceID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetSpace(ctx context.Context, spaceID string) (*Space, error) {
	var resp Space
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/space/%s", spaceID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateSpaceRequest struct {
	Name              string                 `json:"name"`
	MultipleAssignees bool                   `json:"multiple_assignees"`
	Features          map[string]interface{} `json:"features,omitempty"`
}

func (c *Client) CreateSpace(ctx context.Context, workspaceID string, req *CreateSpaceRequest) (*Space, error) {
	var resp Space
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/team/%s/space", workspaceID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateSpaceRequest struct {
	Name              string                 `json:"name,omitempty"`
	MultipleAssignees *bool                  `json:"multiple_assignees,omitempty"`
	Private           *bool                  `json:"private,omitempty"`
	AdminCanManage    *bool                  `json:"admin_can_manage,omitempty"`
	Features          map[string]interface{} `json:"features,omitempty"`
}

func (c *Client) UpdateSpace(ctx context.Context, spaceID string, req *UpdateSpaceRequest) (*Space, error) {
	var resp Space
	if err := c.Do(ctx, "PUT", fmt.Sprintf("/v2/space/%s", spaceID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteSpace(ctx context.Context, spaceID string) error {
	return c.Do(ctx, "DELETE", fmt.Sprintf("/v2/space/%s", spaceID), nil, nil)
}
