package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSharedHierarchy(t *testing.T) {
	ctx := context.Background()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/team/t1/shared" {
			t.Errorf("unexpected path %s", r.URL.Path)
		}
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"shared":{"tasks":[],"lists":[],"folders":[]}}`))
	}))
	defer server.Close()
	client := NewClient("pk_test")
	client.MaxRetries = 0
	client.BaseURL = server.URL
	resp, err := client.GetSharedHierarchy(ctx, "t1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Shared == nil {
		t.Error("expected shared data")
	}
}
