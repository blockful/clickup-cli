package api

import (
	"context"
	"fmt"
)

type AddDependencyRequest struct {
	DependsOn    string `json:"depends_on,omitempty"`
	DependencyOf string `json:"dependency_of,omitempty"`
	Type         string `json:"type,omitempty"`
}

type DependencyResponse struct {
	Dependency interface{} `json:"dependency"`
}

type TaskLinkResponse struct {
	Link interface{} `json:"link"`
}

func (c *Client) AddDependency(ctx context.Context, taskID string, req *AddDependencyRequest) (*DependencyResponse, error) {
	var resp DependencyResponse
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/task/%s/dependency", taskID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteDependency(ctx context.Context, taskID, dependsOn, dependencyOf string) error {
	path := fmt.Sprintf("/v2/task/%s/dependency?", taskID)
	if dependsOn != "" {
		path += "depends_on=" + dependsOn
	}
	if dependencyOf != "" {
		if dependsOn != "" {
			path += "&"
		}
		path += "dependency_of=" + dependencyOf
	}
	return c.Do(ctx, "DELETE", path, nil, nil)
}

func (c *Client) AddTaskLink(ctx context.Context, taskID, linksTo string) (*TaskLinkResponse, error) {
	var resp TaskLinkResponse
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/task/%s/link/%s", taskID, linksTo), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteTaskLink(ctx context.Context, taskID, linksTo string) error {
	return c.Do(ctx, "DELETE", fmt.Sprintf("/v2/task/%s/link/%s", taskID, linksTo), nil, nil)
}
