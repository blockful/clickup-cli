package api

import (
	"context"
	"fmt"
)

type Workspace struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Color   string `json:"color"`
	Avatar  string `json:"avatar"`
	Members []struct {
		User struct {
			ID       int    `json:"id"`
			Username string `json:"username"`
			Email    string `json:"email"`
		} `json:"user"`
	} `json:"members"`
}

type WorkspacesResponse struct {
	Teams []Workspace `json:"teams"`
}

type SeatsResponse struct {
	Members interface{} `json:"members,omitempty"`
	Seats   interface{} `json:"seats,omitempty"`
}

type PlanResponse struct {
	Plan interface{} `json:"plan,omitempty"`
}

func (c *Client) GetWorkspaceSeats(ctx context.Context, teamID string) (*SeatsResponse, error) {
	var resp SeatsResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/team/%s/seats", teamID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetWorkspacePlan(ctx context.Context, teamID string) (*PlanResponse, error) {
	var resp PlanResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/team/%s/plan", teamID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListWorkspaces(ctx context.Context) (*WorkspacesResponse, error) {
	var resp WorkspacesResponse
	if err := c.Do(ctx, "GET", "/v2/team", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
