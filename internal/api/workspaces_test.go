package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetWorkspaceSeats(t *testing.T) {
	ctx := context.Background()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/team/t1/seats" {
			t.Errorf("unexpected path %s", r.URL.Path)
		}
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"seats":{"filled":5,"total":10}}`))
	}))
	defer server.Close()
	client := NewClient("pk_test")
	client.MaxRetries = 0
	client.BaseURL = server.URL
	resp, err := client.GetWorkspaceSeats(ctx, "t1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Seats == nil {
		t.Error("expected seats data")
	}
}

func TestGetWorkspacePlan(t *testing.T) {
	ctx := context.Background()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/team/t1/plan" {
			t.Errorf("unexpected path %s", r.URL.Path)
		}
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"plan":{"name":"Business"}}`))
	}))
	defer server.Close()
	client := NewClient("pk_test")
	client.MaxRetries = 0
	client.BaseURL = server.URL
	resp, err := client.GetWorkspacePlan(ctx, "t1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Plan == nil {
		t.Error("expected plan data")
	}
}
