package api

import (
	"context"
	"fmt"
)

type SharedHierarchyResponse struct {
	Shared interface{} `json:"shared"`
}

func (c *Client) GetSharedHierarchy(ctx context.Context, teamID string) (*SharedHierarchyResponse, error) {
	var resp SharedHierarchyResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/team/%s/shared", teamID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
