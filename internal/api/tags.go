package api

import "fmt"

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

func (c *Client) GetSpaceTags(spaceID string) (*TagsResponse, error) {
	var resp TagsResponse
	if err := c.Do("GET", fmt.Sprintf("/v2/space/%s/tag", spaceID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CreateSpaceTag(spaceID string, req *CreateTagRequest) error {
	return c.Do("POST", fmt.Sprintf("/v2/space/%s/tag", spaceID), req, nil)
}

func (c *Client) UpdateSpaceTag(spaceID, tagName string, req *UpdateTagRequest) error {
	return c.Do("PUT", fmt.Sprintf("/v2/space/%s/tag/%s", spaceID, tagName), req, nil)
}

func (c *Client) DeleteSpaceTag(spaceID, tagName string) error {
	return c.Do("DELETE", fmt.Sprintf("/v2/space/%s/tag/%s", spaceID, tagName), nil, nil)
}

func (c *Client) AddTagToTask(taskID, tagName string) error {
	return c.Do("POST", fmt.Sprintf("/v2/task/%s/tag/%s", taskID, tagName), nil, nil)
}

func (c *Client) RemoveTagFromTask(taskID, tagName string) error {
	return c.Do("DELETE", fmt.Sprintf("/v2/task/%s/tag/%s", taskID, tagName), nil, nil)
}
