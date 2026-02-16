package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInviteUser(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name       string
		response   string
		statusCode int
		wantErr    bool
		errCode    string
	}{
		{
			name:       "success",
			response:   `{"team":{},"user":{"id":123}}`,
			statusCode: 200,
		},
		{
			name:       "unauthorized",
			response:   `{"err":"Token invalid","ECODE":"OAUTH_017"}`,
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
				if r.URL.Path != "/v2/team/123/user" {
					t.Errorf("expected /v2/team/123/user, got %s", r.URL.Path)
				}
				w.WriteHeader(tt.statusCode)
				_, _ = w.Write([]byte(tt.response))
			}))
			defer server.Close()
			client := NewClient("pk_test")
			client.MaxRetries = 0
			client.BaseURL = server.URL
			req := &InviteUserRequest{Email: "test@example.com"}
			_, err := client.InviteUser(ctx, "123", req)
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
		})
	}
}

func TestGetTeamUser(t *testing.T) {
	ctx := context.Background()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/v2/team/t1/user/u1" {
			t.Errorf("unexpected %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"user":{"id":1}}`))
	}))
	defer server.Close()
	client := NewClient("pk_test")
	client.MaxRetries = 0
	client.BaseURL = server.URL
	_, err := client.GetTeamUser(ctx, "t1", "u1", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestEditUser(t *testing.T) {
	ctx := context.Background()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" || r.URL.Path != "/v2/team/t1/user/u1" {
			t.Errorf("unexpected %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"user":{"id":1}}`))
	}))
	defer server.Close()
	client := NewClient("pk_test")
	client.MaxRetries = 0
	client.BaseURL = server.URL
	_, err := client.EditUser(ctx, "t1", "u1", &EditUserRequest{Username: "new"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRemoveUser(t *testing.T) {
	ctx := context.Background()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" || r.URL.Path != "/v2/team/t1/user/u1" {
			t.Errorf("unexpected %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(200)
	}))
	defer server.Close()
	client := NewClient("pk_test")
	client.MaxRetries = 0
	client.BaseURL = server.URL
	err := client.RemoveUser(ctx, "t1", "u1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
