package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddGuestToTask(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name       string
		statusCode int
		response   string
		wantErr    bool
	}{
		{"success", 200, `{"guest":{}}`, false},
		{"not found", 404, `{"err":"not found"}`, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "POST" {
					t.Errorf("expected POST, got %s", r.Method)
				}
				w.WriteHeader(tt.statusCode)
				_, _ = w.Write([]byte(tt.response))
			}))
			defer server.Close()
			client := NewClient("pk_test")
			client.MaxRetries = 0
			client.BaseURL = server.URL
			_, err := client.AddGuestToTask(ctx, "t1", 123, &GuestPermissionRequest{PermissionLevel: "read"})
			if tt.wantErr && err == nil {
				t.Fatal("expected error")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestRemoveGuestFromTask(t *testing.T) {
	ctx := context.Background()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		w.WriteHeader(200)
	}))
	defer server.Close()
	client := NewClient("pk_test")
	client.MaxRetries = 0
	client.BaseURL = server.URL
	if err := client.RemoveGuestFromTask(ctx, "t1", 123); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAddGuestToList(t *testing.T) {
	ctx := context.Background()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"guest":{}}`))
	}))
	defer server.Close()
	client := NewClient("pk_test")
	client.MaxRetries = 0
	client.BaseURL = server.URL
	_, err := client.AddGuestToList(ctx, "l1", 123, &GuestPermissionRequest{PermissionLevel: "read"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRemoveGuestFromFolder(t *testing.T) {
	ctx := context.Background()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		w.WriteHeader(200)
	}))
	defer server.Close()
	client := NewClient("pk_test")
	client.MaxRetries = 0
	client.BaseURL = server.URL
	if err := client.RemoveGuestFromFolder(ctx, "f1", 456); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
