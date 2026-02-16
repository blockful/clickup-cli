package api

import "context"

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

func (c *Client) ListWorkspaces(ctx context.Context) (*WorkspacesResponse, error) {
	var resp WorkspacesResponse
	if err := c.Do(ctx, "GET", "/v2/team", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
