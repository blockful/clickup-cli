package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListComments(t *testing.T) {
	tests := []struct {
		name       string
		taskID     string
		response   string
		statusCode int
		wantErr    bool
		errCode    string
		wantCount  int
	}{
		{
			name:       "success",
			taskID:     "t1",
			response:   `{"comments":[{"id":"c1","comment_text":"Hello"},{"id":"c2","comment_text":"World"}]}`,
			statusCode: 200,
			wantCount:  2,
		},
		{
			name:       "task not found",
			taskID:     "bad",
			response:   `{"err":"Task not found"}`,
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
			client.BaseURL = server.URL

			resp, err := client.ListComments(tt.taskID)
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

func TestCreateComment(t *testing.T) {
	tests := []struct {
		name       string
		taskID     string
		text       string
		response   string
		statusCode int
		wantErr    bool
		wantID     string
	}{
		{
			name:       "success",
			taskID:     "t1",
			text:       "Great work!",
			response:   `{"id":"c1","hist_id":"h1","date":1234567890}`,
			statusCode: 200,
			wantID:     "c1",
		},
		{
			name:       "unauthorized",
			taskID:     "t1",
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
				json.NewDecoder(r.Body).Decode(&body)
				if body.CommentText != tt.text {
					t.Errorf("expected text %q, got %q", tt.text, body.CommentText)
				}
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.response))
			}))
			defer server.Close()

			client := NewClient("pk_test")
			client.BaseURL = server.URL

			resp, err := client.CreateComment(tt.taskID, &CreateCommentRequest{CommentText: tt.text})
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
