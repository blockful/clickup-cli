package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type CustomField struct {
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	Type           string      `json:"type"`
	TypeConfig     interface{} `json:"type_config,omitempty"`
	DateCreated    string      `json:"date_created,omitempty"`
	HideFromGuests bool        `json:"hide_from_guests,omitempty"`
	Value          interface{} `json:"value,omitempty"`
	Required       bool        `json:"required,omitempty"`
}

type Checklist struct {
	ID         string      `json:"id"`
	TaskID     string      `json:"task_id"`
	Name       string      `json:"name"`
	OrderIndex int         `json:"orderindex"`
	Resolved   int         `json:"resolved"`
	Unresolved int         `json:"unresolved"`
	Items      interface{} `json:"items,omitempty"`
}

type LinkedTask struct {
	TaskID      string `json:"task_id"`
	LinkID      string `json:"link_id"`
	DateCreated string `json:"date_created"`
	Userid      string `json:"userid"`
}

type Dependency struct {
	TaskID      string `json:"task_id"`
	DependsOn   string `json:"depends_on"`
	Type        int    `json:"type"`
	DateCreated string `json:"date_created"`
	Userid      string `json:"userid"`
}

type Task struct {
	ID                  string        `json:"id"`
	CustomID            string        `json:"custom_id,omitempty"`
	Name                string        `json:"name"`
	Description         string        `json:"description,omitempty"`
	MarkdownDescription string        `json:"markdown_description,omitempty"`
	Status              TaskStatus    `json:"status"`
	OrderIndex          string        `json:"orderindex"`
	DateCreated         string        `json:"date_created"`
	DateUpdated         string        `json:"date_updated,omitempty"`
	DateClosed          string        `json:"date_closed,omitempty"`
	Creator             User          `json:"creator"`
	Assignees           []User        `json:"assignees"`
	Watchers            []User        `json:"watchers,omitempty"`
	Tags                []TaskTag     `json:"tags"`
	Parent              interface{}   `json:"parent"`
	Priority            *TaskPriority `json:"priority"`
	DueDate             string        `json:"due_date,omitempty"`
	StartDate           string        `json:"start_date,omitempty"`
	Points              interface{}   `json:"points"`
	TimeEstimate        interface{}   `json:"time_estimate"`
	TimeSpent           interface{}   `json:"time_spent,omitempty"`
	CustomFields        []CustomField `json:"custom_fields,omitempty"`
	Checklists          []Checklist   `json:"checklists,omitempty"`
	LinkedTasks         []LinkedTask  `json:"linked_tasks,omitempty"`
	Dependencies        []Dependency  `json:"dependencies,omitempty"`
	Attachments         []Attachment  `json:"attachments,omitempty"`
	URL                 string        `json:"url"`
	List                struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"list"`
	Folder struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"folder"`
	Space struct {
		ID string `json:"id"`
	} `json:"space"`
}

type TaskStatus struct {
	Status string `json:"status"`
	Color  string `json:"color"`
	Type   string `json:"type"`
}

type TaskTag struct {
	Name    string `json:"name"`
	TagFg   string `json:"tag_fg"`
	TagBg   string `json:"tag_bg"`
	Creator int    `json:"creator"`
}

type TaskPriority struct {
	ID       string `json:"id"`
	Priority string `json:"priority"`
	Color    string `json:"color"`
}

type TasksResponse struct {
	Tasks []Task `json:"tasks"`
}

type ListTasksOptions struct {
	Statuses        []string
	Assignees       []string
	Tags            []string
	Watchers        []string
	Page            int
	OrderBy         string
	Reverse         bool
	Subtasks        bool
	IncludeClosed   bool
	Archived        bool
	IncludeMarkdown bool
	IncludeTiml     bool
	DueDateGt       int64
	DueDateLt       int64
	DateCreatedGt   int64
	DateCreatedLt   int64
	DateUpdatedGt   int64
	DateUpdatedLt   int64
	DateDoneGt      int64
	DateDoneLt      int64
	CustomFields    string
	CustomItems     []int
}

func (c *Client) ListTasks(ctx context.Context, listID string, opts *ListTasksOptions) (*TasksResponse, error) {
	params := url.Values{}
	if opts != nil {
		for _, s := range opts.Statuses {
			params.Add("statuses[]", s)
		}
		for _, a := range opts.Assignees {
			params.Add("assignees[]", a)
		}
		for _, t := range opts.Tags {
			params.Add("tags[]", t)
		}
		for _, w := range opts.Watchers {
			params.Add("watchers[]", w)
		}
		if opts.Page > 0 {
			params.Set("page", strconv.Itoa(opts.Page))
		}
		if opts.OrderBy != "" {
			params.Set("order_by", opts.OrderBy)
		}
		if opts.Reverse {
			params.Set("reverse", "true")
		}
		if opts.Subtasks {
			params.Set("subtasks", "true")
		}
		if opts.IncludeClosed {
			params.Set("include_closed", "true")
		}
		if opts.Archived {
			params.Set("archived", "true")
		}
		if opts.IncludeMarkdown {
			params.Set("include_markdown_description", "true")
		}
		if opts.IncludeTiml {
			params.Set("include_task_in_multiple_lists", "true")
		}
		if opts.DueDateGt > 0 {
			params.Set("due_date_gt", strconv.FormatInt(opts.DueDateGt, 10))
		}
		if opts.DueDateLt > 0 {
			params.Set("due_date_lt", strconv.FormatInt(opts.DueDateLt, 10))
		}
		if opts.DateCreatedGt > 0 {
			params.Set("date_created_gt", strconv.FormatInt(opts.DateCreatedGt, 10))
		}
		if opts.DateCreatedLt > 0 {
			params.Set("date_created_lt", strconv.FormatInt(opts.DateCreatedLt, 10))
		}
		if opts.DateUpdatedGt > 0 {
			params.Set("date_updated_gt", strconv.FormatInt(opts.DateUpdatedGt, 10))
		}
		if opts.DateUpdatedLt > 0 {
			params.Set("date_updated_lt", strconv.FormatInt(opts.DateUpdatedLt, 10))
		}
		if opts.DateDoneGt > 0 {
			params.Set("date_done_gt", strconv.FormatInt(opts.DateDoneGt, 10))
		}
		if opts.DateDoneLt > 0 {
			params.Set("date_done_lt", strconv.FormatInt(opts.DateDoneLt, 10))
		}
		if opts.CustomFields != "" {
			params.Set("custom_fields", opts.CustomFields)
		}
		for _, ci := range opts.CustomItems {
			params.Add("custom_items[]", strconv.Itoa(ci))
		}
	}

	path := fmt.Sprintf("/v2/list/%s/task", listID)
	q := params.Encode()
	if q != "" {
		path += "?" + q
	}

	var resp TasksResponse
	if err := c.Do(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetTaskOptions struct {
	CustomTaskIDs   bool
	TeamID          string
	IncludeSubtasks bool
	IncludeMarkdown bool
}

func (c *Client) GetTask(ctx context.Context, taskID string, opts ...GetTaskOptions) (*Task, error) {
	params := url.Values{}
	if len(opts) > 0 {
		o := opts[0]
		if o.CustomTaskIDs {
			params.Set("custom_task_ids", "true")
		}
		if o.TeamID != "" {
			params.Set("team_id", o.TeamID)
		}
		if o.IncludeSubtasks {
			params.Set("include_subtasks", "true")
		}
		if o.IncludeMarkdown {
			params.Set("include_markdown_description", "true")
		}
	}
	path := fmt.Sprintf("/v2/task/%s", taskID)
	q := params.Encode()
	if q != "" {
		path += "?" + q
	}
	var resp Task
	if err := c.Do(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CustomFieldValue struct {
	ID    string      `json:"id"`
	Value interface{} `json:"value"`
}

type CreateTaskRequest struct {
	Name                      string             `json:"name"`
	Description               string             `json:"description,omitempty"`
	MarkdownDescription       string             `json:"markdown_description,omitempty"`
	Assignees                 []int              `json:"assignees,omitempty"`
	GroupAssignees            []string           `json:"group_assignees,omitempty"`
	Tags                      []string           `json:"tags,omitempty"`
	Status                    string             `json:"status,omitempty"`
	Priority                  *int               `json:"priority,omitempty"`
	DueDate                   *int64             `json:"due_date,omitempty"`
	DueDateTime               *bool              `json:"due_date_time,omitempty"`
	StartDate                 *int64             `json:"start_date,omitempty"`
	StartDateTime             *bool              `json:"start_date_time,omitempty"`
	TimeEstimate              *int64             `json:"time_estimate,omitempty"`
	Points                    *float64           `json:"points,omitempty"`
	NotifyAll                 bool               `json:"notify_all,omitempty"`
	Parent                    string             `json:"parent,omitempty"`
	LinksTo                   string             `json:"links_to,omitempty"`
	CustomFields              []CustomFieldValue `json:"custom_fields,omitempty"`
	CustomItemID              *int               `json:"custom_item_id,omitempty"`
	CheckRequiredCustomFields *bool              `json:"check_required_custom_fields,omitempty"`
}

func (c *Client) CreateTask(ctx context.Context, listID string, req *CreateTaskRequest) (*Task, error) {
	var resp Task
	if err := c.Do(ctx, "POST", fmt.Sprintf("/v2/list/%s/task", listID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateTaskAssignees struct {
	Add []int `json:"add,omitempty"`
	Rem []int `json:"rem,omitempty"`
}

type UpdateTaskGroupAssignees struct {
	Add []string `json:"add,omitempty"`
	Rem []string `json:"rem,omitempty"`
}

type UpdateTaskRequest struct {
	Name                *string                   `json:"name,omitempty"`
	Description         *string                   `json:"description,omitempty"`
	MarkdownDescription *string                   `json:"markdown_description,omitempty"`
	Status              *string                   `json:"status,omitempty"`
	Priority            *int                      `json:"priority,omitempty"`
	Assignees           *UpdateTaskAssignees       `json:"assignees,omitempty"`
	GroupAssignees      *UpdateTaskGroupAssignees  `json:"group_assignees,omitempty"`
	DueDate             *int64                     `json:"due_date,omitempty"`
	DueDateTime         *bool                      `json:"due_date_time,omitempty"`
	StartDate           *int64                     `json:"start_date,omitempty"`
	StartDateTime       *bool                      `json:"start_date_time,omitempty"`
	TimeEstimate        *int64                     `json:"time_estimate,omitempty"`
	Points              *float64                   `json:"points,omitempty"`
	Archived            *bool                      `json:"archived,omitempty"`
	Parent              *string                    `json:"parent,omitempty"`
	CustomItemID        *int                       `json:"custom_item_id,omitempty"`
}

type UpdateTaskOptions struct {
	CustomTaskIDs bool
	TeamID        string
}

func (c *Client) UpdateTask(ctx context.Context, taskID string, req *UpdateTaskRequest, opts ...UpdateTaskOptions) (*Task, error) {
	params := url.Values{}
	if len(opts) > 0 {
		o := opts[0]
		if o.CustomTaskIDs {
			params.Set("custom_task_ids", "true")
		}
		if o.TeamID != "" {
			params.Set("team_id", o.TeamID)
		}
	}
	path := fmt.Sprintf("/v2/task/%s", taskID)
	q := params.Encode()
	if q != "" {
		path += "?" + q
	}
	var resp Task
	if err := c.Do(ctx, "PUT", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteTask(ctx context.Context, taskID string) error {
	return c.Do(ctx, "DELETE", fmt.Sprintf("/v2/task/%s", taskID), nil, nil)
}

// SearchTasks searches tasks across a workspace (GET /v2/team/{team_id}/task).
type SearchTasksOptions struct {
	Page            int
	OrderBy         string
	Reverse         bool
	Subtasks        bool
	Statuses        []string
	Assignees       []string
	Tags            []string
	DueDateGt       int64
	DueDateLt       int64
	DateCreatedGt   int64
	DateCreatedLt   int64
	DateUpdatedGt   int64
	DateUpdatedLt   int64
	DateDoneGt      int64
	DateDoneLt      int64
	IncludeClosed   bool
	CustomFields    string
	CustomItems     []int
	ListIDs         []string
	ProjectIDs      []string
	SpaceIDs        []string
	FolderIDs       []string
	IncludeMarkdown bool
}

func (c *Client) SearchTasks(ctx context.Context, teamID string, opts *SearchTasksOptions) (*TasksResponse, error) {
	params := url.Values{}
	if opts != nil {
		if opts.Page > 0 {
			params.Set("page", strconv.Itoa(opts.Page))
		}
		if opts.OrderBy != "" {
			params.Set("order_by", opts.OrderBy)
		}
		if opts.Reverse {
			params.Set("reverse", "true")
		}
		if opts.Subtasks {
			params.Set("subtasks", "true")
		}
		if opts.IncludeClosed {
			params.Set("include_closed", "true")
		}
		for _, s := range opts.Statuses {
			params.Add("statuses[]", s)
		}
		for _, a := range opts.Assignees {
			params.Add("assignees[]", a)
		}
		for _, t := range opts.Tags {
			params.Add("tags[]", t)
		}
		if opts.DueDateGt > 0 {
			params.Set("due_date_gt", strconv.FormatInt(opts.DueDateGt, 10))
		}
		if opts.DueDateLt > 0 {
			params.Set("due_date_lt", strconv.FormatInt(opts.DueDateLt, 10))
		}
		if opts.DateCreatedGt > 0 {
			params.Set("date_created_gt", strconv.FormatInt(opts.DateCreatedGt, 10))
		}
		if opts.DateCreatedLt > 0 {
			params.Set("date_created_lt", strconv.FormatInt(opts.DateCreatedLt, 10))
		}
		if opts.DateUpdatedGt > 0 {
			params.Set("date_updated_gt", strconv.FormatInt(opts.DateUpdatedGt, 10))
		}
		if opts.DateUpdatedLt > 0 {
			params.Set("date_updated_lt", strconv.FormatInt(opts.DateUpdatedLt, 10))
		}
		if opts.DateDoneGt > 0 {
			params.Set("date_done_gt", strconv.FormatInt(opts.DateDoneGt, 10))
		}
		if opts.DateDoneLt > 0 {
			params.Set("date_done_lt", strconv.FormatInt(opts.DateDoneLt, 10))
		}
		if opts.CustomFields != "" {
			params.Set("custom_fields", opts.CustomFields)
		}
		for _, ci := range opts.CustomItems {
			params.Add("custom_items[]", strconv.Itoa(ci))
		}
		for _, id := range opts.ListIDs {
			params.Add("list_ids[]", id)
		}
		for _, id := range opts.ProjectIDs {
			params.Add("project_ids[]", id)
		}
		for _, id := range opts.SpaceIDs {
			params.Add("space_ids[]", id)
		}
		for _, id := range opts.FolderIDs {
			params.Add("folder_ids[]", id)
		}
		if opts.IncludeMarkdown {
			params.Set("include_markdown_description", "true")
		}
	}
	path := fmt.Sprintf("/v2/team/%s/task", teamID)
	q := params.Encode()
	if q != "" {
		path += "?" + q
	}
	var resp TasksResponse
	if err := c.Do(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ParseCustomFields parses a JSON string into a slice of CustomFieldValue.
func ParseCustomFields(s string) ([]CustomFieldValue, error) {
	var fields []CustomFieldValue
	if err := json.Unmarshal([]byte(s), &fields); err != nil {
		return nil, fmt.Errorf("invalid custom-fields JSON: %v", err)
	}
	return fields, nil
}

// MergeTasksRequest is the request body for merging tasks.
type MergeTasksRequest struct {
	MergeWith []string `json:"merge_with"`
}

// TimeInStatusResponse is the response for time in status.
type TimeInStatusResponse struct {
	CurrentStatus    interface{}   `json:"current_status"`
	StatusHistory    []interface{} `json:"status_history"`
}

// BulkTimeInStatusResponse is the response for bulk time in status.
type BulkTimeInStatusResponse map[string]interface{}

// TaskScopedOptions holds custom_task_ids and team_id query params for task-scoped endpoints.
type TaskScopedOptions struct {
	CustomTaskIDs bool
	TeamID        string
}

func taskScopedQuery(opts *TaskScopedOptions) string {
	if opts == nil {
		return ""
	}
	params := url.Values{}
	if opts.CustomTaskIDs {
		params.Set("custom_task_ids", "true")
	}
	if opts.TeamID != "" {
		params.Set("team_id", opts.TeamID)
	}
	q := params.Encode()
	if q != "" {
		return "?" + q
	}
	return ""
}

func (c *Client) MergeTasks(ctx context.Context, taskID string, req *MergeTasksRequest, opts ...*TaskScopedOptions) error {
	var o *TaskScopedOptions
	if len(opts) > 0 { o = opts[0] }
	return c.Do(ctx, "POST", fmt.Sprintf("/v2/task/%s/merge", taskID)+taskScopedQuery(o), req, nil)
}

func (c *Client) GetTimeInStatus(ctx context.Context, taskID string, opts ...*TaskScopedOptions) (*TimeInStatusResponse, error) {
	var o *TaskScopedOptions
	if len(opts) > 0 { o = opts[0] }
	var resp TimeInStatusResponse
	if err := c.Do(ctx, "GET", fmt.Sprintf("/v2/task/%s/time_in_status", taskID)+taskScopedQuery(o), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetBulkTimeInStatus(ctx context.Context, taskIDs []string) (*BulkTimeInStatusResponse, error) {
	params := url.Values{}
	for _, id := range taskIDs {
		params.Add("task_ids", id)
	}
	var resp BulkTimeInStatusResponse
	if err := c.Do(ctx, "GET", "/v2/task/bulk_time_in_status/task_ids?"+params.Encode(), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) AddTaskToList(ctx context.Context, listID, taskID string) error {
	return c.Do(ctx, "POST", fmt.Sprintf("/v2/list/%s/task/%s", listID, taskID), nil, nil)
}

func (c *Client) RemoveTaskFromList(ctx context.Context, listID, taskID string) error {
	return c.Do(ctx, "DELETE", fmt.Sprintf("/v2/list/%s/task/%s", listID, taskID), nil, nil)
}
