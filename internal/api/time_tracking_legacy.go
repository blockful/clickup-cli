package api

import (
	"context"
	"fmt"
	"net/url"
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

func (c *Client) GetLegacyTrackedTime(ctx context.Context, taskID string, subcategoryID string, opts ...*TaskScopedOptions) (*LegacyTimeResponse, error) {
	params := url.Values{}
	if subcategoryID != "" {
		params.Set("subcategory_id", subcategoryID)
	}
	if len(opts) > 0 && opts[0] != nil {
		if opts[0].CustomTaskIDs {
			params.Set("custom_task_ids", "true")
		}
		if opts[0].TeamID != "" {
			params.Set("team_id", opts[0].TeamID)
		}
	}
	path := fmt.Sprintf("/v2/task/%s/time", taskID)
	if q := params.Encode(); q != "" {
		path += "?" + q
	}
	var resp LegacyTimeResponse
	if err := c.Do(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) TrackLegacyTime(ctx context.Context, taskID string, req *LegacyTrackTimeRequest, opts ...*TaskScopedOptions) (*LegacyTimeResponse, error) {
	var o *TaskScopedOptions
	if len(opts) > 0 { o = opts[0] }
	var resp LegacyTimeResponse
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/task/%s/time", taskID)+taskScopedQuery(o), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) EditLegacyTime(ctx context.Context, taskID, intervalID string, req *LegacyEditTimeRequest, opts ...*TaskScopedOptions) error {
	var o *TaskScopedOptions
	if len(opts) > 0 { o = opts[0] }
	return c.Do(ctx, "PUT", fmt.Sprintf("/v2/task/%s/time/%s", taskID, intervalID)+taskScopedQuery(o), req, nil)
}

func (c *Client) DeleteLegacyTime(ctx context.Context, taskID, intervalID string, opts ...*TaskScopedOptions) error {
	var o *TaskScopedOptions
	if len(opts) > 0 { o = opts[0] }
	return c.Do(ctx, "DELETE", fmt.Sprintf("/v2/task/%s/time/%s", taskID, intervalID)+taskScopedQuery(o), nil, nil)
}
