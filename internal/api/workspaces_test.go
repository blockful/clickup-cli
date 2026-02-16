package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListWorkspaces(t *testing.T) {
	tests := []struct {
		name       string
		response   string
		statusCode int
		wantErr    bool
		errCode    string
		wantCount  int
	}{
		{
			name:       "success with teams",
			response:   `{"teams":[{"id":"1","name":"Team A"},{"id":"2","name":"Team B"}]}`,
			statusCode: 200,
			wantCount:  2,
		},
		{
			name:       "empty",
			response:   `{"teams":[]}`,
			statusCode: 200,
			wantCount:  0,
		},
		{
			name:       "unauthorized",
			response:   `{"err":"unauthorized"}`,
			statusCode: 401,
			wantErr:    true,
			errCode:    "UNAUTHORIZED",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "GET" {
					t.Errorf("expected GET, got %s", r.Method)
				}
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.response))
			}))
			defer server.Close()

			client := NewClient("pk_test")
			client.BaseURL = server.URL

			resp, err := client.ListWorkspaces()
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
			if len(resp.Teams) != tt.wantCount {
				t.Errorf("expected %d teams, got %d", tt.wantCount, len(resp.Teams))
			}
		})
	}
}
