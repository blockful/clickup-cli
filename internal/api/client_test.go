package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientDo(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name       string
		method     string
		path       string
		body       interface{}
		statusCode int
		response   string
		wantErr    bool
		errCode    string
		checkAuth  bool
	}{
		{
			name:       "GET success",
			method:     "GET",
			path:       "/v2/user",
			statusCode: 200,
			response:   `{"user":{"id":1,"username":"test"}}`,
			checkAuth:  true,
		},
		{
			name:       "POST with body",
			method:     "POST",
			path:       "/v2/list/123/task",
			body:       map[string]string{"name": "Test Task"},
			statusCode: 200,
			response:   `{"id":"abc123","name":"Test Task"}`,
		},
		{
			name:       "401 unauthorized",
			method:     "GET",
			path:       "/v2/user",
			statusCode: 401,
			response:   `{"err":"Token invalid","ECODE":"OAUTH_017"}`,
			wantErr:    true,
			errCode:    "UNAUTHORIZED",
		},
		{
			name:       "403 forbidden",
			method:     "GET",
			path:       "/v2/space/1",
			statusCode: 403,
			response:   `{"err":"Forbidden"}`,
			wantErr:    true,
			errCode:    "FORBIDDEN",
		},
		{
			name:       "404 not found",
			method:     "GET",
			path:       "/v2/task/missing",
			statusCode: 404,
			response:   `{"err":"Resource not found"}`,
			wantErr:    true,
			errCode:    "NOT_FOUND",
		},
		{
			name:       "429 rate limited",
			method:     "GET",
			path:       "/v2/team",
			statusCode: 429,
			response:   `{"err":"Rate limit exceeded"}`,
			wantErr:    true,
			errCode:    "RATE_LIMITED",
		},
		{
			name:       "500 server error",
			method:     "GET",
			path:       "/v2/team",
			statusCode: 500,
			response:   `{"err":"Internal server error"}`,
			wantErr:    true,
			errCode:    "API_ERROR",
		},
		{
			name:       "error response without err field",
			method:     "GET",
			path:       "/v2/team",
			statusCode: 400,
			response:   `{"message":"Bad request"}`,
			wantErr:    true,
			errCode:    "API_ERROR",
		},
		{
			name:       "error response with no JSON body",
			method:     "GET",
			path:       "/v2/team",
			statusCode: 502,
			response:   `Bad Gateway`,
			wantErr:    true,
			errCode:    "API_ERROR",
		},
		{
			name:       "empty success response",
			method:     "DELETE",
			path:       "/v2/task/123",
			statusCode: 200,
			response:   ``,
		},
		{
			name:       "malformed JSON in success response",
			method:     "GET",
			path:       "/v2/user",
			statusCode: 200,
			response:   `{not json`,
			wantErr:    true,
			errCode:    "UNMARSHAL_ERROR",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if tt.checkAuth && r.Header.Get("Authorization") != "pk_test_token" {
					t.Errorf("expected Authorization 'pk_test_token', got '%s'", r.Header.Get("Authorization"))
				}
				if r.Header.Get("Content-Type") != "application/json" {
					t.Errorf("missing Content-Type header")
				}
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.response))
			}))
			defer server.Close()

			client := NewClient("pk_test_token")
			client.MaxRetries = 0
			client.BaseURL = server.URL

			var result map[string]interface{}
			var resultPtr interface{}
			// Only pass result pointer if we expect to unmarshal (non-empty success)
			if !tt.wantErr && tt.response != "" {
				resultPtr = &result
			}
			// For the malformed JSON test, we need a result pointer to trigger unmarshal
			if tt.name == "malformed JSON in success response" {
				resultPtr = &result
			}

			err := client.Do(ctx, tt.method, tt.path, tt.body, resultPtr)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				clientErr, ok := err.(*ClientError)
				if !ok {
					t.Fatalf("expected *ClientError, got %T", err)
				}
				if clientErr.Code != tt.errCode {
					t.Errorf("expected code %q, got %q", tt.errCode, clientErr.Code)
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}
		})
	}
}

func TestClientDo_NetworkError(t *testing.T) {
	ctx := context.Background()
	client := NewClient("pk_test")
	client.MaxRetries = 0
	client.BaseURL = "http://127.0.0.1:1" // unreachable

	err := client.Do(ctx, "GET", "/v2/user", nil, nil)
	if err == nil {
		t.Fatal("expected network error")
	}
	clientErr, ok := err.(*ClientError)
	if !ok {
		t.Fatalf("expected *ClientError, got %T", err)
	}
	if clientErr.Code != "NETWORK_ERROR" {
		t.Errorf("expected NETWORK_ERROR, got %s", clientErr.Code)
	}
}

func TestClientDo_PostBodyValidation(t *testing.T) {
	ctx := context.Background()
	var receivedBody map[string]string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		json.NewDecoder(r.Body).Decode(&receivedBody)
		w.WriteHeader(200)
		w.Write([]byte(`{"id":"123"}`))
	}))
	defer server.Close()

	client := NewClient("pk_test")
	client.MaxRetries = 0
	client.BaseURL = server.URL

	body := map[string]string{"name": "My Task", "description": "A test"}
	var result map[string]string
	err := client.Do(ctx, "POST", "/v2/list/1/task", body, &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if receivedBody["name"] != "My Task" {
		t.Errorf("expected name 'My Task', got '%s'", receivedBody["name"])
	}
	if receivedBody["description"] != "A test" {
		t.Errorf("expected description 'A test', got '%s'", receivedBody["description"])
	}
}

func TestErrorCodeFromStatus(t *testing.T) {
	tests := []struct {
		status int
		want   string
	}{
		{401, "UNAUTHORIZED"},
		{403, "FORBIDDEN"},
		{404, "NOT_FOUND"},
		{429, "RATE_LIMITED"},
		{500, "API_ERROR"},
		{502, "API_ERROR"},
		{400, "API_ERROR"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := errorCodeFromStatus(tt.status)
			if got != tt.want {
				t.Errorf("errorCodeFromStatus(%d) = %q, want %q", tt.status, got, tt.want)
			}
		})
	}
}
