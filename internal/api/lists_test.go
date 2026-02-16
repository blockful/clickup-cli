package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListLists(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name       string
		folderID   string
		response   string
		statusCode int
		wantErr    bool
		wantCount  int
	}{
		{
			name:       "success",
			folderID:   "f1",
			response:   `{"lists":[{"id":"l1","name":"List 1"},{"id":"l2","name":"List 2"}]}`,
			statusCode: 200,
			wantCount:  2,
		},
		{
			name:       "empty",
			folderID:   "f2",
			response:   `{"lists":[]}`,
			statusCode: 200,
			wantCount:  0,
		},
		{
			name:       "forbidden",
			folderID:   "f3",
			response:   `{"err":"Forbidden"}`,
			statusCode: 403,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.response))
			}))
			defer server.Close()

			client := NewClient("pk_test")
			client.MaxRetries = 0
			client.BaseURL = server.URL

			resp, err := client.ListLists(ctx, tt.folderID)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(resp.Lists) != tt.wantCount {
				t.Errorf("expected %d, got %d", tt.wantCount, len(resp.Lists))
			}
		})
	}
}

func TestGetList(t *testing.T) {
	ctx := context.Background()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(`{"id":"l1","name":"My List","task_count":5}`))
	}))
	defer server.Close()

	client := NewClient("pk_test")
	client.MaxRetries = 0
	client.BaseURL = server.URL

	list, err := client.GetList(ctx, "l1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if list.Name != "My List" {
		t.Errorf("expected 'My List', got %q", list.Name)
	}
	if list.TaskCount != 5 {
		t.Errorf("expected task_count 5, got %d", list.TaskCount)
	}
}

func TestCreateList(t *testing.T) {
	ctx := context.Background()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		w.WriteHeader(200)
		w.Write([]byte(`{"id":"l2","name":"New List"}`))
	}))
	defer server.Close()

	client := NewClient("pk_test")
	client.MaxRetries = 0
	client.BaseURL = server.URL

	list, err := client.CreateList(ctx, "f1", &CreateListRequest{Name: "New List"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if list.ID != "l2" {
		t.Errorf("expected 'l2', got %q", list.ID)
	}
}
