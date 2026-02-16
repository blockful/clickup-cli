package api

import "fmt"

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

func (c *Client) ListFolders(spaceID string) (*FoldersResponse, error) {
	var resp FoldersResponse
	if err := c.Do("GET", fmt.Sprintf("/v2/space/%s/folder?archived=false", spaceID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetFolder(folderID string) (*Folder, error) {
	var resp Folder
	if err := c.Do("GET", fmt.Sprintf("/v2/folder/%s", folderID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateFolderRequest struct {
	Name string `json:"name"`
}

func (c *Client) CreateFolder(spaceID string, req *CreateFolderRequest) (*Folder, error) {
	var resp Folder
	if err := c.Do("POST", fmt.Sprintf("/v2/space/%s/folder", spaceID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
