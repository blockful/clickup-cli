package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTaskTemplates(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name       string
		response   string
		statusCode int
		wantErr    bool
		wantCount  int
	}{
		{
			name:       "success",
			response:   `{"templates":[{"id":"t1","name":"Bug Report"},{"id":"t2","name":"Feature"}]}`,
			statusCode: 200,
			wantCount:  2,
		},
		{
			name:       "unauthorized",
			response:   `{"err":"unauthorized"}`,
			statusCode: 401,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
			resp, err := client.GetTaskTemplates(ctx, "team1", 0)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(resp.Templates) != tt.wantCount {
				t.Errorf("expected %d, got %d", tt.wantCount, len(resp.Templates))
			}
		})
	}
}

func TestCreateTaskFromTemplate(t *testing.T) {
	ctx := context.Background()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		var body CreateFromTemplateRequest
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body.Name != "My Task" {
			t.Errorf("expected name 'My Task', got %q", body.Name)
		}
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"id":"task1"}`))
	}))
	defer server.Close()
	client := NewClient("pk_test")
	client.MaxRetries = 0
	client.BaseURL = server.URL
	resp, err := client.CreateTaskFromTemplate(ctx, "list1", "tmpl1", &CreateFromTemplateRequest{Name: "My Task"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.ID != "task1" {
		t.Errorf("expected task1, got %s", resp.ID)
	}
}

func TestCreateFolderFromTemplate(t *testing.T) {
	ctx := context.Background()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"id":"folder1"}`))
	}))
	defer server.Close()
	client := NewClient("pk_test")
	client.MaxRetries = 0
	client.BaseURL = server.URL
	resp, err := client.CreateFolderFromTemplate(ctx, "space1", "tmpl1", &CreateFromTemplateRequest{Name: "My Folder"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.ID != "folder1" {
		t.Errorf("expected folder1, got %s", resp.ID)
	}
}

func TestCreateListFromFolderTemplate(t *testing.T) {
	ctx := context.Background()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"id":"list1"}`))
	}))
	defer server.Close()
	client := NewClient("pk_test")
	client.MaxRetries = 0
	client.BaseURL = server.URL
	resp, err := client.CreateListFromFolderTemplate(ctx, "folder1", "tmpl1", &CreateFromTemplateRequest{Name: "My List"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.ID != "list1" {
		t.Errorf("expected list1, got %s", resp.ID)
	}
}
