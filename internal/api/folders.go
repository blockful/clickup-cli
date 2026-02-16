package api

import (
	"context"
	"fmt"
)

type Folder struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	OrderIndex       int    `json:"orderindex"`
	OverrideStatuses bool   `json:"override_statuses"`
	Hidden           bool   `json:"hidden"`
	Space            struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"space"`
	TaskCount string `json:"task_count"`
	Lists     []List `json:"lists"`
}

type FoldersResponse struct {
	Folders []Folder `json:"folders"`
}

func (c *Client) ListFolders(ctx context.Context, spaceID string) (*FoldersResponse, error) {
	var resp FoldersResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/space/%s/folder?archived=false", spaceID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetFolder(ctx context.Context, folderID string) (*Folder, error) {
	var resp Folder
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/folder/%s", folderID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateFolderRequest struct {
	Name string `json:"name"`
}

func (c *Client) CreateFolder(ctx context.Context, spaceID string, req *CreateFolderRequest) (*Folder, error) {
	var resp Folder
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/space/%s/folder", spaceID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateFolderRequest struct {
	Name string `json:"name"`
}

func (c *Client) UpdateFolder(ctx context.Context, folderID string, req *UpdateFolderRequest) (*Folder, error) {
	var resp Folder
	if err := c.Do(ctx, "PUT", fmt.Sprintf("/v2/folder/%s", folderID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteFolder(ctx context.Context, folderID string) error {
	return c.Do(ctx, "DELETE", fmt.Sprintf("/v2/folder/%s", folderID), nil, nil)
}
