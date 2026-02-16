package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListThreadedComments(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name       string
		commentID  string
		response   string
		statusCode int
		wantErr    bool
		errCode    string
		wantCount  int
	}{
		{
			name:       "success",
			commentID:  "c1",
			response:   `{"comments":[{"id":"r1","comment_text":"reply"}]}`,
			statusCode: 200,
			wantCount:  1,
		},
		{
			name:       "not found",
			commentID:  "bad",
			response:   `{"err":"Comment not found"}`,
			statusCode: 404,
			wantErr:    true,
			errCode:    "NOT_FOUND",
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

			resp, err := client.ListThreadedComments(ctx, tt.commentID)
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
			if len(resp.Comments) != tt.wantCount {
				t.Errorf("expected %d, got %d", tt.wantCount, len(resp.Comments))
			}
		})
	}
}

func TestCreateThreadedComment(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name       string
		commentID  string
		text       string
		response   string
		statusCode int
		wantErr    bool
		wantID     string
	}{
		{
			name:       "success",
			commentID:  "c1",
			text:       "reply text",
			response:   `{"id":"r1","hist_id":"h1","date":123}`,
			statusCode: 200,
			wantID:     "r1",
		},
		{
			name:       "unauthorized",
			commentID:  "c1",
			text:       "test",
			response:   `{"err":"unauthorized"}`,
			statusCode: 401,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "POST" {
					t.Errorf("expected POST, got %s", r.Method)
				}
				var body CreateCommentRequest
				_ = json.NewDecoder(r.Body).Decode(&body)
				if body.CommentText != tt.text {
					t.Errorf("expected text %q, got %q", tt.text, body.CommentText)
				}
				w.WriteHeader(tt.statusCode)
				_, _ = w.Write([]byte(tt.response))
			}))
			defer server.Close()

			client := NewClient("pk_test")
			client.MaxRetries = 0
			client.BaseURL = server.URL

			resp, err := client.CreateThreadedComment(ctx, tt.commentID, &CreateCommentRequest{CommentText: tt.text})
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if resp.ID != tt.wantID {
				t.Errorf("expected ID %q, got %q", tt.wantID, resp.ID)
			}
		})
	}
}

func TestListViewComments(t *testing.T) {
	ctx := context.Background()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"comments":[{"id":"vc1","comment_text":"view comment"}]}`))
	}))
	defer server.Close()

	client := NewClient("pk_test")
	client.MaxRetries = 0
	client.BaseURL = server.URL

	resp, err := client.ListViewComments(ctx, "v1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(resp.Comments) != 1 {
		t.Errorf("expected 1, got %d", len(resp.Comments))
	}
}

func TestCreateViewComment(t *testing.T) {
	ctx := context.Background()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"id":"vc1","hist_id":"h1","date":123}`))
	}))
	defer server.Close()

	client := NewClient("pk_test")
	client.MaxRetries = 0
	client.BaseURL = server.URL

	resp, err := client.CreateViewComment(ctx, "v1", &CreateCommentRequest{CommentText: "hello"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.ID != "vc1" {
		t.Errorf("expected vc1, got %s", resp.ID)
	}
}
