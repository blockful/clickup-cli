package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCustomTaskTypes(t *testing.T) {
	ctx := context.Background()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/team/t1/custom_item" {
			t.Errorf("unexpected path %s", r.URL.Path)
		}
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"custom_items":[{"id":1,"name":"Bug"}]}`))
	}))
	defer server.Close()
	client := NewClient("pk_test")
	client.MaxRetries = 0
	client.BaseURL = server.URL
	resp, err := client.GetCustomTaskTypes(ctx, "t1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(resp.CustomItems) != 1 {
		t.Errorf("expected 1 item, got %d", len(resp.CustomItems))
	}
}
