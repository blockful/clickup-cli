package api

import (
	"context"
	"fmt"
)

type TimeEntry struct {
	ID           string      `json:"id"`
	Task         interface{} `json:"task,omitempty"`
	Wid          string      `json:"wid,omitempty"`
	User         interface{} `json:"user,omitempty"`
	Billable     bool        `json:"billable"`
	Start        string      `json:"start"`
	End          string      `json:"end,omitempty"`
	Duration     string      `json:"duration"`
	Description  string      `json:"description"`
	Tags         []Tag       `json:"tags,omitempty"`
	Source       string      `json:"source,omitempty"`
	At           string      `json:"at,omitempty"`
	TaskLocation interface{} `json:"task_location,omitempty"`
	TaskTags     interface{} `json:"task_tags,omitempty"`
	TaskURL      string      `json:"task_url,omitempty"`
}

type TimeEntriesResponse struct {
	Data []TimeEntry `json:"data"`
}

type SingleTimeEntryResponse struct {
	Data TimeEntry `json:"data"`
}

type ListTimeEntriesOptions struct {
	StartDate            string
	EndDate              string
	Assignee             string
	IncludeTaskTags      bool
	IncludeLocationNames bool
	SpaceID              string
	FolderID             string
	ListID               string
	TaskID               string
	CustomTaskIDs        bool
	TeamID               string
	IsBillable           *bool
}

type CreateTimeEntryRequest struct {
	Description string `json:"description,omitempty"`
	Tags        []Tag  `json:"tags,omitempty"`
	Start       int64  `json:"start"`
	Billable    bool   `json:"billable,omitempty"`
	Duration    int64  `json:"duration"`
	Assignee    *int   `json:"assignee,omitempty"`
	Tid         string `json:"tid,omitempty"`
}

type UpdateTimeEntryRequest struct {
	Description string `json:"description,omitempty"`
	Tags        []Tag  `json:"tags,omitempty"`
	TagAction   string `json:"tag_action,omitempty"`
	Start       *int64 `json:"start,omitempty"`
	End         *int64 `json:"end,omitempty"`
	Tid         string `json:"tid,omitempty"`
	Billable    *bool  `json:"billable,omitempty"`
	Duration    *int64 `json:"duration,omitempty"`
}

type StartTimerRequest struct {
	Tid         string `json:"tid,omitempty"`
	Description string `json:"description,omitempty"`
	Tags        []Tag  `json:"tags,omitempty"`
	Billable    bool   `json:"billable,omitempty"`
}

type TimeEntryTagsResponse struct {
	Data []Tag `json:"data"`
}

func (c *Client) GetTimeEntries(ctx context.Context, teamID string, opts *ListTimeEntriesOptions) (*TimeEntriesResponse, error) {
	path := fmt.Sprintf("/v2/team/%s/time_entries", teamID)
	if opts != nil {
		q := ""
		sep := "?"
		add := func(k, v string) {
			if v != "" {
				q += sep + k + "=" + v
				sep = "&"
			}
		}
		add("start_date", opts.StartDate)
		add("end_date", opts.EndDate)
		add("assignee", opts.Assignee)
		add("space_id", opts.SpaceID)
		add("folder_id", opts.FolderID)
		add("list_id", opts.ListID)
		add("task_id", opts.TaskID)
		if opts.IncludeTaskTags {
			q += sep + "include_task_tags=true"
			sep = "&"
		}
		if opts.IncludeLocationNames {
			q += sep + "include_location_names=true"
			sep = "&"
		}
		if opts.IsBillable != nil {
			if *opts.IsBillable {
				q += sep + "is_billable=true"
			} else {
				q += sep + "is_billable=false"
			}
		}
		path += q
	}
	var resp TimeEntriesResponse
	if err := c.Do(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CreateTimeEntry(ctx context.Context, teamID string, req *CreateTimeEntryRequest) (*TimeEntry, error) {
	var resp TimeEntry
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/team/%s/time_entries", teamID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetTimeEntry(ctx context.Context, teamID, timerID string) (*SingleTimeEntryResponse, error) {
	var resp SingleTimeEntryResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/team/%s/time_entries/%s", teamID, timerID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateTimeEntry(ctx context.Context, teamID, timerID string, req *UpdateTimeEntryRequest) error {
	return c.Do(ctx, "PUT", fmt.Sprintf("/v2/team/%s/time_entries/%s", teamID, timerID), req, nil)
}

func (c *Client) DeleteTimeEntry(ctx context.Context, teamID, timerID string) error {
	return c.Do(ctx, "DELETE", fmt.Sprintf("/v2/team/%s/time_entries/%s", teamID, timerID), nil, nil)
}

func (c *Client) StartTimer(ctx context.Context, teamID string, req *StartTimerRequest) (*SingleTimeEntryResponse, error) {
	var resp SingleTimeEntryResponse
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/team/%s/time_entries/start", teamID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) StopTimer(ctx context.Context, teamID string) (*SingleTimeEntryResponse, error) {
	var resp SingleTimeEntryResponse
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/team/%s/time_entries/stop", teamID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetRunningTimer(ctx context.Context, teamID string, assignee string) (*SingleTimeEntryResponse, error) {
	path := fmt.Sprintf("/v2/team/%s/time_entries/current", teamID)
	if assignee != "" {
		path += "?assignee=" + assignee
	}
	var resp SingleTimeEntryResponse
	if err := c.Do(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetTimeEntryTags(ctx context.Context, teamID string) (*TimeEntryTagsResponse, error) {
	var resp TimeEntryTagsResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/team/%s/time_entries/tags", teamID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
