package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCustomRoles(t *testing.T) {
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
			response:   `{"custom_roles":[{"id":1,"name":"Editor"},{"id":2,"name":"Viewer"}]}`,
			statusCode: 200,
			wantCount:  2,
		},
		{
			name:       "not found",
			response:   `{"err":"Team not found"}`,
			statusCode: 404,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/v2/team/t1/customroles" {
					t.Errorf("unexpected path %s", r.URL.Path)
				}
				w.WriteHeader(tt.statusCode)
				_, _ = w.Write([]byte(tt.response))
			}))
			defer server.Close()
			client := NewClient("pk_test")
			client.MaxRetries = 0
			client.BaseURL = server.URL
			resp, err := client.GetCustomRoles(ctx, "t1", false)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(resp.CustomRoles) != tt.wantCount {
				t.Errorf("expected %d roles, got %d", tt.wantCount, len(resp.CustomRoles))
			}
		})
	}
}
