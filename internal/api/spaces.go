package api

import "fmt"

type Space struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Private  bool   `json:"private"`
	Status   struct {
		Status string `json:"status"`
		Color  string `json:"color"`
	} `json:"status"`
	Multiple bool `json:"multiple_assignees"`
	Features map[string]interface{} `json:"features"`
}

type SpacesResponse struct {
	Spaces []Space `json:"spaces"`
}

func (c *Client) ListSpaces(workspaceID string) (*SpacesResponse, error) {
	var resp SpacesResponse
	if err := c.Do("GET", fmt.Sprintf("/v2/team/%s/space?archived=false", workspaceID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetSpace(spaceID string) (*Space, error) {
	var resp Space
	if err := c.Do("GET", fmt.Sprintf("/v2/space/%s", spaceID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateSpaceRequest struct {
	Name             string `json:"name"`
	MultipleAssignees bool   `json:"multiple_assignees"`
	Features         map[string]interface{} `json:"features,omitempty"`
}

func (c *Client) CreateSpace(workspaceID string, req *CreateSpaceRequest) (*Space, error) {
	var resp Space
	if err := c.Do("POST", fmt.Sprintf("/v2/team/%s/space", workspaceID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
