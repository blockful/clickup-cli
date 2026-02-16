package api

import (
	"context"
	"fmt"
	"net/url"
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

func (c *Client) AddDependency(ctx context.Context, taskID string, req *AddDependencyRequest, opts ...*TaskScopedOptions) (*DependencyResponse, error) {
	var o *TaskScopedOptions
	if len(opts) > 0 { o = opts[0] }
	var resp DependencyResponse
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/task/%s/dependency", taskID)+taskScopedQuery(o), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteDependency(ctx context.Context, taskID, dependsOn, dependencyOf string, opts ...*TaskScopedOptions) error {
	params := url.Values{}
	if dependsOn != "" {
		params.Set("depends_on", dependsOn)
	}
	if dependencyOf != "" {
		params.Set("dependency_of", dependencyOf)
	}
	if len(opts) > 0 && opts[0] != nil {
		if opts[0].CustomTaskIDs {
			params.Set("custom_task_ids", "true")
		}
		if opts[0].TeamID != "" {
			params.Set("team_id", opts[0].TeamID)
		}
	}
	path := fmt.Sprintf("/v2/task/%s/dependency", taskID)
	if q := params.Encode(); q != "" {
		path += "?" + q
	}
	return c.Do(ctx, "DELETE", path, nil, nil)
}

func (c *Client) AddTaskLink(ctx context.Context, taskID, linksTo string, opts ...*TaskScopedOptions) (*TaskLinkResponse, error) {
	var o *TaskScopedOptions
	if len(opts) > 0 { o = opts[0] }
	var resp TaskLinkResponse
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/task/%s/link/%s", taskID, linksTo)+taskScopedQuery(o), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteTaskLink(ctx context.Context, taskID, linksTo string, opts ...*TaskScopedOptions) error {
	var o *TaskScopedOptions
	if len(opts) > 0 { o = opts[0] }
	return c.Do(ctx, "DELETE", fmt.Sprintf("/v2/task/%s/link/%s", taskID, linksTo)+taskScopedQuery(o), nil, nil)
}
