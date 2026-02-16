package api

import (
	"context"
	"fmt"
)

type CustomRole struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	CustomType int    `json:"custom_type"`
	DateCreated string `json:"date_created,omitempty"`
}

type CustomRolesResponse struct {
	CustomRoles []CustomRole `json:"custom_roles"`
}

func (c *Client) GetCustomRoles(ctx context.Context, teamID string) (*CustomRolesResponse, error) {
	var resp CustomRolesResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/team/%s/customroles", teamID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
