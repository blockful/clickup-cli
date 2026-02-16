package api

import (
	"context"
	"fmt"
)

type InviteUserRequest struct {
	Email        string `json:"email"`
	Admin        *bool  `json:"admin,omitempty"`
	CustomRoleID *int   `json:"custom_role_id,omitempty"`
	MemberGroups []int  `json:"member_groups,omitempty"`
}

type EditUserRequest struct {
	Username     string `json:"username,omitempty"`
	Admin        *bool  `json:"admin,omitempty"`
	CustomRoleID *int   `json:"custom_role_id,omitempty"`
}

type TeamUser struct {
	User   interface{} `json:"user,omitempty"`
	Invite *bool       `json:"invite,omitempty"`
}

type TeamUserResponse struct {
	Team interface{} `json:"team,omitempty"`
	User interface{} `json:"user,omitempty"`
}

func (c *Client) InviteUser(ctx context.Context, teamID string, req *InviteUserRequest) (*TeamUserResponse, error) {
	var resp TeamUserResponse
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/team/%s/user", teamID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetTeamUser(ctx context.Context, teamID, userID string) (*TeamUserResponse, error) {
	var resp TeamUserResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/team/%s/user/%s", teamID, userID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) EditUser(ctx context.Context, teamID, userID string, req *EditUserRequest) (*TeamUserResponse, error) {
	var resp TeamUserResponse
	if err := c.Do(ctx, "PUT", fmt.Sprintf("/v2/team/%s/user/%s", teamID, userID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) RemoveUser(ctx context.Context, teamID, userID string) error {
	return c.Do(ctx, "DELETE", fmt.Sprintf("/v2/team/%s/user/%s", teamID, userID), nil, nil)
}
