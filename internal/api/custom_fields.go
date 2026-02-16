package api

import (
	"context"
	"fmt"
)

// CustomField is defined in tasks.go

type CustomFieldsResponse struct {
	Fields []CustomField `json:"fields"`
}

type SetCustomFieldRequest struct {
	Value interface{} `json:"value"`
}

func (c *Client) GetListCustomFields(ctx context.Context, listID string) (*CustomFieldsResponse, error) {
	var resp CustomFieldsResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/list/%s/field", listID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetFolderCustomFields(ctx context.Context, folderID string) (*CustomFieldsResponse, error) {
	var resp CustomFieldsResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/folder/%s/field", folderID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetSpaceCustomFields(ctx context.Context, spaceID string) (*CustomFieldsResponse, error) {
	var resp CustomFieldsResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/space/%s/field", spaceID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetWorkspaceCustomFields(ctx context.Context, teamID string) (*CustomFieldsResponse, error) {
	var resp CustomFieldsResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/team/%s/field", teamID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) SetCustomFieldValue(ctx context.Context, taskID, fieldID string, req *SetCustomFieldRequest, opts ...*TaskScopedOptions) error {
	var o *TaskScopedOptions
	if len(opts) > 0 { o = opts[0] }
	return c.Do(ctx, "POST", fmt.Sprintf("/v2/task/%s/field/%s", taskID, fieldID)+taskScopedQuery(o), req, nil)
}

func (c *Client) RemoveCustomFieldValue(ctx context.Context, taskID, fieldID string, opts ...*TaskScopedOptions) error {
	var o *TaskScopedOptions
	if len(opts) > 0 { o = opts[0] }
	return c.Do(ctx, "DELETE", fmt.Sprintf("/v2/task/%s/field/%s", taskID, fieldID)+taskScopedQuery(o), nil, nil)
}
