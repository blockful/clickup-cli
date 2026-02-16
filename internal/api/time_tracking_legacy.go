package api

import (
	"context"
	"fmt"
)

type LegacyTimeInterval struct {
	ID       string      `json:"id"`
	Start    string      `json:"start"`
	End      string      `json:"end"`
	Time     string      `json:"time"`
	Source   string      `json:"source,omitempty"`
	DateAdded string     `json:"date_added,omitempty"`
	Tags     []Tag       `json:"tags,omitempty"`
	Taskid   string      `json:"taskid,omitempty"`
	User     interface{} `json:"user,omitempty"`
}

type LegacyTimeResponse struct {
	Data []LegacyTimeInterval `json:"data"`
}

type LegacyTrackTimeRequest struct {
	Time  int64  `json:"time"`
	Start int64  `json:"start,omitempty"`
	End   int64  `json:"end,omitempty"`
	Tags  []Tag  `json:"tags,omitempty"`
}

type LegacyEditTimeRequest struct {
	Time      int64  `json:"time,omitempty"`
	Start     int64  `json:"start,omitempty"`
	End       int64  `json:"end,omitempty"`
	TagAction string `json:"tag_action,omitempty"`
	Tags      []Tag  `json:"tags,omitempty"`
}

func (c *Client) GetLegacyTrackedTime(ctx context.Context, taskID string, subcategoryID string) (*LegacyTimeResponse, error) {
	path := fmt.Sprintf("/v2/task/%s/time", taskID)
	if subcategoryID != "" {
		path += "?subcategory_id=" + subcategoryID
	}
	var resp LegacyTimeResponse
	if err := c.Do(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) TrackLegacyTime(ctx context.Context, taskID string, req *LegacyTrackTimeRequest) (*LegacyTimeResponse, error) {
	var resp LegacyTimeResponse
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/task/%s/time", taskID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) EditLegacyTime(ctx context.Context, taskID, intervalID string, req *LegacyEditTimeRequest) error {
	return c.Do(ctx, "PUT", fmt.Sprintf("/v2/task/%s/time/%s", taskID, intervalID), req, nil)
}

func (c *Client) DeleteLegacyTime(ctx context.Context, taskID, intervalID string) error {
	return c.Do(ctx, "DELETE", fmt.Sprintf("/v2/task/%s/time/%s", taskID, intervalID), nil, nil)
}
