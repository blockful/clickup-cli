package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListTasksOptions_AllParams(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name   string
		opts   *ListTasksOptions
		expect map[string]string // key -> expected query param value
	}{
		{
			name: "archived and markdown",
			opts: &ListTasksOptions{Archived: true, IncludeMarkdown: true},
			expect: map[string]string{
				"archived":                     "true",
				"include_markdown_description": "true",
			},
		},
		{
			name: "date filters",
			opts: &ListTasksOptions{DueDateGt: 1000, DueDateLt: 2000, DateCreatedGt: 3000},
			expect: map[string]string{
				"due_date_gt":     "1000",
				"due_date_lt":     "2000",
				"date_created_gt": "3000",
			},
		},
		{
			name: "include timl",
			opts: &ListTasksOptions{IncludeTiml: true},
			expect: map[string]string{
				"include_task_in_multiple_lists": "true",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				for k, v := range tc.expect {
					got := r.URL.Query().Get(k)
					if got != v {
						t.Errorf("param %s: got %q, want %q", k, got, v)
					}
				}
				json.NewEncoder(w).Encode(TasksResponse{})
			}))
			defer srv.Close()

			c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
			_, err := c.ListTasks(ctx, "list1", tc.opts)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestGetTask_WithOptions(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("include_subtasks") != "true" {
			t.Error("expected include_subtasks=true")
		}
		if r.URL.Query().Get("include_markdown_description") != "true" {
			t.Error("expected include_markdown_description=true")
		}
		json.NewEncoder(w).Encode(Task{ID: "t1"})
	}))
	defer srv.Close()

	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	task, err := c.GetTask(ctx, "t1", GetTaskOptions{IncludeSubtasks: true, IncludeMarkdown: true})
	if err != nil {
		t.Fatal(err)
	}
	if task.ID != "t1" {
		t.Errorf("got ID %s, want t1", task.ID)
	}
}

func TestCreateTask_AllFields(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req CreateTaskRequest
		json.NewDecoder(r.Body).Decode(&req)
		if req.Name != "test" {
			t.Errorf("name: got %s", req.Name)
		}
		if req.Parent != "parent1" {
			t.Errorf("parent: got %s", req.Parent)
		}
		if req.DueDate == nil || *req.DueDate != 12345 {
			t.Error("due_date not set correctly")
		}
		if req.NotifyAll != true {
			t.Error("notify_all should be true")
		}
		json.NewEncoder(w).Encode(Task{ID: "new"})
	}))
	defer srv.Close()

	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	dd := int64(12345)
	_, err := c.CreateTask(ctx, "list1", &CreateTaskRequest{
		Name:      "test",
		Parent:    "parent1",
		DueDate:   &dd,
		NotifyAll: true,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateTask_WithAssignees(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req UpdateTaskRequest
		json.NewDecoder(r.Body).Decode(&req)
		if req.Assignees == nil {
			t.Fatal("assignees nil")
		}
		if len(req.Assignees.Add) != 1 || req.Assignees.Add[0] != 42 {
			t.Errorf("add: %v", req.Assignees.Add)
		}
		if r.URL.Query().Get("custom_task_ids") != "true" {
			t.Error("expected custom_task_ids query param")
		}
		json.NewEncoder(w).Encode(Task{ID: "t1"})
	}))
	defer srv.Close()

	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	_, err := c.UpdateTask(ctx, "t1", &UpdateTaskRequest{
		Assignees: &UpdateTaskAssignees{Add: []int{42}},
	}, UpdateTaskOptions{CustomTaskIDs: true})
	if err != nil {
		t.Fatal(err)
	}
}

func TestSearchTasks(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/team/team1/task" {
			t.Errorf("path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("include_closed") != "true" {
			t.Error("expected include_closed")
		}
		json.NewEncoder(w).Encode(TasksResponse{})
	}))
	defer srv.Close()

	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	_, err := c.SearchTasks(ctx, "team1", &SearchTasksOptions{IncludeClosed: true})
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseCustomFields(t *testing.T) {
	input := `[{"id":"abc","value":"hello"}]`
	fields, err := ParseCustomFields(input)
	if err != nil {
		t.Fatal(err)
	}
	if len(fields) != 1 || fields[0].ID != "abc" {
		t.Errorf("unexpected: %+v", fields)
	}

	_, err = ParseCustomFields("not json")
	if err == nil {
		t.Error("expected error for invalid JSON")
	}
}
