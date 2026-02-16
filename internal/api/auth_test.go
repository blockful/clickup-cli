package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUser(t *testing.T) {
	tests := []struct {
		name       string
		response   string
		statusCode int
		wantErr    bool
		errCode    string
		wantUser   string
	}{
		{
			name:       "success",
			response:   `{"user":{"id":123,"username":"testuser","email":"test@example.com"}}`,
			statusCode: 200,
			wantUser:   "testuser",
		},
		{
			name:       "invalid token",
			response:   `{"err":"Token invalid","ECODE":"OAUTH_017"}`,
			statusCode: 401,
			wantErr:    true,
			errCode:    "UNAUTHORIZED",
		},
		{
			name:       "rate limited",
			response:   `{"err":"Rate limit exceeded"}`,
			statusCode: 429,
			wantErr:    true,
			errCode:    "RATE_LIMITED",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/v2/user" {
					t.Errorf("expected path /v2/user, got %s", r.URL.Path)
				}
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.response))
			}))
			defer server.Close()

			client := NewClient("pk_test")
			client.BaseURL = server.URL

			resp, err := client.GetUser()
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
			if resp.User.Username != tt.wantUser {
				t.Errorf("expected %q, got %q", tt.wantUser, resp.User.Username)
			}
		})
	}
}
