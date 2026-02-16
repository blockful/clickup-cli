package api

import (
	"context"
	"fmt"
)

type Tag struct {
	Name  string `json:"name"`
	TagFg string `json:"tag_fg"`
	TagBg string `json:"tag_bg"`
}

type TagsResponse struct {
	Tags []Tag `json:"tags"`
}

type CreateTagRequest struct {
	Tag Tag `json:"tag"`
}

type UpdateTagRequest struct {
	Tag Tag `json:"tag"`
}

func (c *Client) GetSpaceTags(ctx context.Context, spaceID string) (*TagsResponse, error) {
	var resp TagsResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/space/%s/tag", spaceID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CreateSpaceTag(ctx context.Context, spaceID string, req *CreateTagRequest) error {
	return c.Do(ctx, "POST", fmt.Sprintf("/v2/space/%s/tag", spaceID), req, nil)
}

func (c *Client) UpdateSpaceTag(ctx context.Context, spaceID, tagName string, req *UpdateTagRequest) error {
	return c.Do(ctx, "PUT", fmt.Sprintf("/v2/space/%s/tag/%s", spaceID, tagName), req, nil)
}

func (c *Client) DeleteSpaceTag(ctx context.Context, spaceID, tagName string) error {
	return c.Do(ctx, "DELETE", fmt.Sprintf("/v2/space/%s/tag/%s", spaceID, tagName), nil, nil)
}

func (c *Client) AddTagToTask(ctx context.Context, taskID, tagName string) error {
	return c.Do(ctx, "POST", fmt.Sprintf("/v2/task/%s/tag/%s", taskID, tagName), nil, nil)
}

func (c *Client) RemoveTagFromTask(ctx context.Context, taskID, tagName string) error {
	return c.Do(ctx, "DELETE", fmt.Sprintf("/v2/task/%s/tag/%s", taskID, tagName), nil, nil)
}
