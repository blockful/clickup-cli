package api

import (
	"fmt"
	"net/url"
	"strconv"
)

type Task struct {
	ID          string      `json:"id"`
	CustomID    string      `json:"custom_id,omitempty"`
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	Status      TaskStatus  `json:"status"`
	OrderIndex  string      `json:"orderindex"`
	DateCreated string      `json:"date_created"`
	DateUpdated string      `json:"date_updated,omitempty"`
	DateClosed  string      `json:"date_closed,omitempty"`
	Creator     User        `json:"creator"`
	Assignees   []User      `json:"assignees"`
	Tags        []TaskTag   `json:"tags"`
	Parent      interface{} `json:"parent"`
	Priority    *TaskPriority `json:"priority"`
	DueDate     string      `json:"due_date,omitempty"`
	StartDate   string      `json:"start_date,omitempty"`
	Points      interface{} `json:"points"`
	TimeEstimate interface{} `json:"time_estimate"`
	URL         string      `json:"url"`
	List        struct {
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
	Statuses     []string
	Assignees    []string
	Tags         []string
	Page         int
	OrderBy      string
	Reverse      bool
	Subtasks     bool
	IncludeClosed bool
}

func (c *Client) ListTasks(listID string, opts *ListTasksOptions) (*TasksResponse, error) {
	params := url.Values{}
	if opts != nil {
		if len(opts.Statuses) > 0 {
			for _, s := range opts.Statuses {
				params.Add("statuses[]", s)
			}
		}
		if len(opts.Assignees) > 0 {
			for _, a := range opts.Assignees {
				params.Add("assignees[]", a)
			}
		}
		if len(opts.Tags) > 0 {
			for _, t := range opts.Tags {
				params.Add("tags[]", t)
			}
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
	}

	path := fmt.Sprintf("/v2/list/%s/task", listID)
	q := params.Encode()
	if q != "" {
		path += "?" + q
	}

	var resp TasksResponse
	if err := c.Do("GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetTask(taskID string) (*Task, error) {
	var resp Task
	if err := c.Do("GET", fmt.Sprintf("/v2/task/%s", taskID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateTaskRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Assignees   []int    `json:"assignees,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Status      string   `json:"status,omitempty"`
	Priority    *int     `json:"priority,omitempty"`
	DueDate     int64    `json:"due_date,omitempty"`
	StartDate   int64    `json:"start_date,omitempty"`
}

func (c *Client) CreateTask(listID string, req *CreateTaskRequest) (*Task, error) {
	var resp Task
	if err := c.Do("POST", fmt.Sprintf("/v2/list/%s/task", listID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateTaskRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Status      *string `json:"status,omitempty"`
	Priority    *int    `json:"priority,omitempty"`
}

func (c *Client) UpdateTask(taskID string, req *UpdateTaskRequest) (*Task, error) {
	var resp Task
	if err := c.Do("PUT", fmt.Sprintf("/v2/task/%s", taskID), req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteTask(taskID string) error {
	return c.Do("DELETE", fmt.Sprintf("/v2/task/%s", taskID), nil, nil)
}

// StringPtr returns a pointer to the given string.
func StringPtr(s string) *string {
	return &s
}

// IntPtr returns a pointer to the given int.
func IntPtr(i int) *int {
	return &i
}

