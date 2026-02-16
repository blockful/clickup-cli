package api

import (
	"context"
	"fmt"
)

type TaskTemplate struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type TaskTemplatesResponse struct {
	Templates []TaskTemplate `json:"templates"`
}

type CreateFromTemplateRequest struct {
	Name    string      `json:"name"`
	Options interface{} `json:"options,omitempty"`
}

type CreateFromTemplateResponse struct {
	ID string `json:"id,omitempty"`
}

func (c *Client) GetTaskTemplates(ctx context.Context, teamID string, page int) (*TaskTemplatesResponse, error) {
	var resp TaskTemplatesResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/team/%s/taskTemplate?page=%d", teamID, page), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CreateTaskFromTemplate(ctx context.Context, listID, templateID string, req *CreateFromTemplateRequest) (*CreateFromTemplateResponse, error) {
	var resp CreateFromTemplateResponse
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/list/%s/taskTemplate/%s", listID, templateID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CreateFolderFromTemplate(ctx context.Context, spaceID, templateID string, req *CreateFromTemplateRequest) (*CreateFromTemplateResponse, error) {
	var resp CreateFromTemplateResponse
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/space/%s/folder_template/%s", spaceID, templateID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CreateListFromFolderTemplate(ctx context.Context, folderID, templateID string, req *CreateFromTemplateRequest) (*CreateFromTemplateResponse, error) {
	var resp CreateFromTemplateResponse
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/folder/%s/list_template/%s", folderID, templateID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CreateListFromSpaceTemplate(ctx context.Context, spaceID, templateID string, req *CreateFromTemplateRequest) (*CreateFromTemplateResponse, error) {
	var resp CreateFromTemplateResponse
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/space/%s/list_template/%s", spaceID, templateID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
