package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListSpaces(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name        string
		workspaceID string
		response    string
		statusCode  int
		wantErr     bool
		errCode     string
		wantCount   int
	}{
		{
			name:        "success",
			workspaceID: "ws1",
			response:    `{"spaces":[{"id":"s1","name":"Space 1"},{"id":"s2","name":"Space 2"}]}`,
			statusCode:  200,
			wantCount:   2,
		},
		{
			name:        "empty",
			workspaceID: "ws2",
			response:    `{"spaces":[]}`,
			statusCode:  200,
			wantCount:   0,
		},
		{
			name:        "unauthorized",
			workspaceID: "ws1",
			response:    `{"err":"unauthorized"}`,
			statusCode:  401,
			wantErr:     true,
			errCode:     "UNAUTHORIZED",
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

			resp, err := client.ListSpaces(ctx, tt.workspaceID)
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
			if len(resp.Spaces) != tt.wantCount {
				t.Errorf("expected %d spaces, got %d", tt.wantCount, len(resp.Spaces))
			}
		})
	}
}

func TestGetSpace(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name       string
		spaceID    string
		response   string
		statusCode int
		wantErr    bool
		errCode    string
		wantName   string
	}{
		{
			name:       "success",
			spaceID:    "s1",
			response:   `{"id":"s1","name":"My Space"}`,
			statusCode: 200,
			wantName:   "My Space",
		},
		{
			name:       "not found",
			spaceID:    "bad",
			response:   `{"err":"Space not found"}`,
			statusCode: 404,
			wantErr:    true,
			errCode:    "NOT_FOUND",
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

			space, err := client.GetSpace(ctx, tt.spaceID)
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
			if space.Name != tt.wantName {
				t.Errorf("expected %q, got %q", tt.wantName, space.Name)
			}
		})
	}
}

func TestCreateSpace(t *testing.T) {
	ctx := context.Background()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		w.WriteHeader(200)
		w.Write([]byte(`{"id":"new1","name":"New Space"}`))
	}))
	defer server.Close()

	client := NewClient("pk_test")
	client.MaxRetries = 0
	client.BaseURL = server.URL

	space, err := client.CreateSpace(ctx, "ws1", &CreateSpaceRequest{Name: "New Space"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if space.ID != "new1" {
		t.Errorf("expected ID 'new1', got %q", space.ID)
	}
}
