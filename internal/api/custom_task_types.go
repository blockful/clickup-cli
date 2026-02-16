package api

import (
	"context"
	"fmt"
)

type CustomTaskType struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	Avatar      interface{} `json:"avatar,omitempty"`
}

type CustomTaskTypesResponse struct {
	CustomItems []CustomTaskType `json:"custom_items"`
}

func (c *Client) GetCustomTaskTypes(ctx context.Context, teamID string) (*CustomTaskTypesResponse, error) {
	var resp CustomTaskTypesResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/team/%s/custom_item", teamID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
