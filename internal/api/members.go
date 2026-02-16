package api

import (
	"context"
	"fmt"
	"net/url"
)

type Member struct {
	ID             int         `json:"id"`
	Username       string      `json:"username"`
	Email          string      `json:"email"`
	Color          string      `json:"color,omitempty"`
	Initials       string      `json:"initials,omitempty"`
	ProfilePicture interface{} `json:"profilePicture,omitempty"`
	ProfileInfo    interface{} `json:"profileInfo,omitempty"`
}

type MembersResponse struct {
	Members []Member `json:"members"`
}

type Group struct {
	ID          string      `json:"id"`
	TeamID      string      `json:"team_id"`
	UserID      int         `json:"userid"`
	Name        string      `json:"name"`
	Handle      string      `json:"handle"`
	DateCreated string      `json:"date_created"`
	Initials    string      `json:"initials"`
	Members     interface{} `json:"members"`
	Avatar      interface{} `json:"avatar"`
}

type GroupsResponse struct {
	Groups []Group `json:"groups"`
}

type CreateGroupRequest struct {
	Name    string `json:"name"`
	Handle  string `json:"handle,omitempty"`
	Members []int  `json:"members"`
}

type UpdateGroupRequest struct {
	Name    string `json:"name,omitempty"`
	Handle  string `json:"handle,omitempty"`
	Members *struct {
		Add []int `json:"add,omitempty"`
		Rem []int `json:"rem,omitempty"`
	} `json:"members,omitempty"`
}

type Guest struct {
	User interface{} `json:"user,omitempty"`
}

type GuestResponse struct {
	Guest Guest `json:"guest"`
}

type InviteGuestRequest struct {
	Email                 string `json:"email"`
	CanEditTags           *bool  `json:"can_edit_tags,omitempty"`
	CanSeeTimeSpent       *bool  `json:"can_see_time_spent,omitempty"`
	CanSeeTimeEstimated   *bool  `json:"can_see_time_estimated,omitempty"`
	CanCreateViews        *bool  `json:"can_create_views,omitempty"`
	CanSeePointsEstimated *bool  `json:"can_see_points_estimated,omitempty"`
	CustomRoleID          *int   `json:"custom_role_id,omitempty"`
}

type EditGuestRequest struct {
	CanSeePointsEstimated *bool `json:"can_see_points_estimated,omitempty"`
	CanEditTags           *bool `json:"can_edit_tags,omitempty"`
	CanSeeTimeSpent       *bool `json:"can_see_time_spent,omitempty"`
	CanSeeTimeEstimated   *bool `json:"can_see_time_estimated,omitempty"`
	CanCreateViews        *bool `json:"can_create_views,omitempty"`
	CustomRoleID          *int  `json:"custom_role_id,omitempty"`
}

func (c *Client) GetListMembers(ctx context.Context, listID string) (*MembersResponse, error) {
	var resp MembersResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/list/%s/member", listID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetTaskMembers(ctx context.Context, taskID string) (*MembersResponse, error) {
	var resp MembersResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/task/%s/member", taskID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetGroups(ctx context.Context, teamID string, groupIDs []string) (*GroupsResponse, error) {
	path := "/v2/group"
	sep := "?"
	if teamID != "" {
		path += sep + "team_id=" + teamID
		sep = "&"
	}
	for _, gid := range groupIDs {
		path += sep + "group_ids=" + gid
		sep = "&"
	}
	var resp GroupsResponse
	if err := c.Do(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CreateGroup(ctx context.Context, teamID string, req *CreateGroupRequest) (*Group, error) {
	var resp Group
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/team/%s/group", teamID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateGroup(ctx context.Context, groupID string, req *UpdateGroupRequest) (*Group, error) {
	var resp Group
	if err := c.Do(ctx, "PUT", fmt.Sprintf("/v2/group/%s", groupID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteGroup(ctx context.Context, groupID string) error {
	return c.Do(ctx, "DELETE", fmt.Sprintf("/v2/group/%s", groupID), nil, nil)
}

func (c *Client) InviteGuest(ctx context.Context, teamID string, req *InviteGuestRequest) error {
	return c.Do(ctx, "POST", fmt.Sprintf("/v2/team/%s/guest", teamID), req, nil)
}

func (c *Client) GetGuest(ctx context.Context, teamID, guestID string) (*GuestResponse, error) {
	var resp GuestResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/team/%s/guest/%s", teamID, guestID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) EditGuest(ctx context.Context, teamID, guestID string, req *EditGuestRequest) (*GuestResponse, error) {
	var resp GuestResponse
	if err := c.Do(ctx, "PUT", fmt.Sprintf("/v2/team/%s/guest/%s", teamID, guestID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) RemoveGuest(ctx context.Context, teamID, guestID string) error {
	return c.Do(ctx, "DELETE", fmt.Sprintf("/v2/team/%s/guest/%s", teamID, guestID), nil, nil)
}

// Guest assignments to task/list/folder

type GuestPermissionRequest struct {
	PermissionLevel string `json:"permission_level"`
}

func includeSharedQuery(includeShared bool) string {
	if includeShared {
		return "?include_shared=true"
	}
	return ""
}

func (c *Client) AddGuestToTask(ctx context.Context, taskID string, guestID int, req *GuestPermissionRequest, includeShared bool, opts ...*TaskScopedOptions) (*GuestResponse, error) {
	params := url.Values{}
	if includeShared {
		params.Set("include_shared", "true")
	}
	if len(opts) > 0 && opts[0] != nil {
		if opts[0].CustomTaskIDs { params.Set("custom_task_ids", "true") }
		if opts[0].TeamID != "" { params.Set("team_id", opts[0].TeamID) }
	}
	path := fmt.Sprintf("/v2/task/%s/guest/%d", taskID, guestID)
	if q := params.Encode(); q != "" { path += "?" + q }
	var resp GuestResponse
	if err := c.Do(ctx, "POST", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) RemoveGuestFromTask(ctx context.Context, taskID string, guestID int, includeShared bool, opts ...*TaskScopedOptions) error {
	params := url.Values{}
	if includeShared {
		params.Set("include_shared", "true")
	}
	if len(opts) > 0 && opts[0] != nil {
		if opts[0].CustomTaskIDs { params.Set("custom_task_ids", "true") }
		if opts[0].TeamID != "" { params.Set("team_id", opts[0].TeamID) }
	}
	path := fmt.Sprintf("/v2/task/%s/guest/%d", taskID, guestID)
	if q := params.Encode(); q != "" { path += "?" + q }
	return c.Do(ctx, "DELETE", path, nil, nil)
}

func (c *Client) AddGuestToList(ctx context.Context, listID string, guestID int, req *GuestPermissionRequest, includeShared bool) (*GuestResponse, error) {
	var resp GuestResponse
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/list/%s/guest/%d%s", listID, guestID, includeSharedQuery(includeShared)), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) RemoveGuestFromList(ctx context.Context, listID string, guestID int, includeShared bool) error {
	return c.Do(ctx, "DELETE", fmt.Sprintf("/v2/list/%s/guest/%d%s", listID, guestID, includeSharedQuery(includeShared)), nil, nil)
}

func (c *Client) AddGuestToFolder(ctx context.Context, folderID string, guestID int, req *GuestPermissionRequest, includeShared bool) (*GuestResponse, error) {
	var resp GuestResponse
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/folder/%s/guest/%d%s", folderID, guestID, includeSharedQuery(includeShared)), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) RemoveGuestFromFolder(ctx context.Context, folderID string, guestID int, includeShared bool) error {
	return c.Do(ctx, "DELETE", fmt.Sprintf("/v2/folder/%s/guest/%d%s", folderID, guestID, includeSharedQuery(includeShared)), nil, nil)
}
