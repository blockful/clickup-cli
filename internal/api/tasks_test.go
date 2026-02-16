package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTask(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name       string
		taskID     string
		response   string
		statusCode int
		wantErr    bool
		errCode    string
		wantName   string
	}{
		{
			name:       "success",
			taskID:     "abc123",
			response:   `{"id":"abc123","name":"Test Task","status":{"status":"open","color":"#000"}}`,
			statusCode: 200,
			wantName:   "Test Task",
		},
		{
			name:       "not found",
			taskID:     "missing",
			response:   `{"err":"Task not found"}`,
			statusCode: 404,
			wantErr:    true,
			errCode:    "NOT_FOUND",
		},
		{
			name:       "unauthorized",
			taskID:     "abc123",
			response:   `{"err":"Token invalid"}`,
			statusCode: 401,
			wantErr:    true,
			errCode:    "UNAUTHORIZED",
		},
		{
			name:       "rate limited",
			taskID:     "abc123",
			response:   `{"err":"Rate limit exceeded"}`,
			statusCode: 429,
			wantErr:    true,
			errCode:    "RATE_LIMITED",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				wantPath := "/v2/task/" + tt.taskID
				if r.URL.Path != wantPath {
					t.Errorf("expected path %s, got %s", wantPath, r.URL.Path)
				}
				if r.Method != "GET" {
					t.Errorf("expected GET, got %s", r.Method)
				}
				w.WriteHeader(tt.statusCode)
				_, _ = w.Write([]byte(tt.response))
			}))
			defer server.Close()

			client := NewClient("pk_test")
			client.MaxRetries = 0
			client.BaseURL = server.URL

			task, err := client.GetTask(ctx, tt.taskID)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				if ce, ok := err.(*ClientError); ok && ce.Code != tt.errCode {
					t.Errorf("expected error code %q, got %q", tt.errCode, ce.Code)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if task.Name != tt.wantName {
				t.Errorf("expected name %q, got %q", tt.wantName, task.Name)
			}
		})
	}
}

func TestCreateTask(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name       string
		listID     string
		req        *CreateTaskRequest
		response   string
		statusCode int
		wantErr    bool
		errCode    string
		wantID     string
	}{
		{
			name:       "success",
			listID:     "list1",
			req:        &CreateTaskRequest{Name: "New Task", Description: "desc"},
			response:   `{"id":"task1","name":"New Task"}`,
			statusCode: 200,
			wantID:     "task1",
		},
		{
			name:       "unauthorized",
			listID:     "list1",
			req:        &CreateTaskRequest{Name: "New"},
			response:   `{"err":"unauthorized"}`,
			statusCode: 401,
			wantErr:    true,
			errCode:    "UNAUTHORIZED",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "POST" {
					t.Errorf("expected POST, got %s", r.Method)
				}
				var body CreateTaskRequest
				_ = json.NewDecoder(r.Body).Decode(&body)
				if body.Name != tt.req.Name {
					t.Errorf("expected name %q in body, got %q", tt.req.Name, body.Name)
				}
				w.WriteHeader(tt.statusCode)
				_, _ = w.Write([]byte(tt.response))
			}))
			defer server.Close()

			client := NewClient("pk_test")
			client.MaxRetries = 0
			client.BaseURL = server.URL

			task, err := client.CreateTask(ctx, tt.listID, tt.req)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				if ce, ok := err.(*ClientError); ok && ce.Code != tt.errCode {
					t.Errorf("expected %q, got %q", tt.errCode, ce.Code)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if task.ID != tt.wantID {
				t.Errorf("expected ID %q, got %q", tt.wantID, task.ID)
			}
		})
	}
}

func TestUpdateTask(t *testing.T) {
	ctx := context.Background()
	name := "Updated"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"id":"t1","name":"Updated"}`))
	}))
	defer server.Close()

	client := NewClient("pk_test")
	client.MaxRetries = 0
	client.BaseURL = server.URL

	task, err := client.UpdateTask(ctx, "t1", &UpdateTaskRequest{Name: &name})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if task.Name != "Updated" {
		t.Errorf("expected name 'Updated', got %q", task.Name)
	}
}

func TestDeleteTask(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name       string
		taskID     string
		statusCode int
		response   string
		wantErr    bool
		errCode    string
	}{
		{name: "success", taskID: "t1", statusCode: 200, response: ""},
		{name: "not found", taskID: "bad", statusCode: 404, response: `{"err":"not found"}`, wantErr: true, errCode: "NOT_FOUND"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "DELETE" {
					t.Errorf("expected DELETE, got %s", r.Method)
				}
				w.WriteHeader(tt.statusCode)
				_, _ = w.Write([]byte(tt.response))
			}))
			defer server.Close()

			client := NewClient("pk_test")
			client.MaxRetries = 0
			client.BaseURL = server.URL

			err := client.DeleteTask(ctx, tt.taskID)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				if ce, ok := err.(*ClientError); ok && ce.Code != tt.errCode {
					t.Errorf("expected %q, got %q", tt.errCode, ce.Code)
				}
			} else if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestListTasks(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name      string
		listID    string
		opts      *ListTasksOptions
		wantQuery map[string]string
	}{
		{
			name:   "no options",
			listID: "list1",
			opts:   nil,
		},
		{
			name:      "with pagination",
			listID:    "list1",
			opts:      &ListTasksOptions{Page: 2},
			wantQuery: map[string]string{"page": "2"},
		},
		{
			name:      "with filters",
			listID:    "list1",
			opts:      &ListTasksOptions{Reverse: true, Subtasks: true, IncludeClosed: true, OrderBy: "created"},
			wantQuery: map[string]string{"reverse": "true", "subtasks": "true", "include_closed": "true", "order_by": "created"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				for k, v := range tt.wantQuery {
					if got := r.URL.Query().Get(k); got != v {
						t.Errorf("expected query %s=%s, got %s", k, v, got)
					}
				}
				w.WriteHeader(200)
				_, _ = w.Write([]byte(`{"tasks":[]}`))
			}))
			defer server.Close()

			client := NewClient("pk_test")
			client.MaxRetries = 0
			client.BaseURL = server.URL

			resp, err := client.ListTasks(ctx, tt.listID, tt.opts)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if resp == nil {
				t.Fatal("expected response, got nil")
			}
		})
	}
}

func TestMergeTasks(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/v2/task/t1/merge" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		var req MergeTasksRequest
		_ = json.NewDecoder(r.Body).Decode(&req)
		if len(req.SourceTaskIDs) != 2 {
			t.Errorf("merge_with len = %d", len(req.SourceTaskIDs))
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	err := c.MergeTasks(ctx, "t1", &MergeTasksRequest{SourceTaskIDs: []string{"t2", "t3"}})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetTimeInStatus(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/v2/task/t1/time_in_status" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		_ = json.NewEncoder(w).Encode(TimeInStatusResponse{CurrentStatus: map[string]string{"status": "open"}})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resp, err := c.GetTimeInStatus(ctx, "t1")
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("expected response")
	}
}

func TestGetBulkTimeInStatus(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method = %s", r.Method)
		}
		ids := r.URL.Query()["task_ids"]
		if len(ids) != 2 {
			t.Errorf("task_ids len = %d", len(ids))
		}
		_ = json.NewEncoder(w).Encode(BulkTimeInStatusResponse{"t1": "data", "t2": "data"})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resp, err := c.GetBulkTimeInStatus(ctx, []string{"t1", "t2"})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("expected response")
	}
}

func TestAddTaskToList(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/v2/list/l1/task/t1" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	if err := c.AddTaskToList(ctx, "l1", "t1"); err != nil {
		t.Fatal(err)
	}
}

func TestRemoveTaskFromList(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" || r.URL.Path != "/v2/list/l1/task/t1" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	if err := c.RemoveTaskFromList(ctx, "l1", "t1"); err != nil {
		t.Fatal(err)
	}
}
