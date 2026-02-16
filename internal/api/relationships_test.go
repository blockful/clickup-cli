package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddDependency(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/v2/task/t1/dependency" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		var req AddDependencyRequest
		_ = json.NewDecoder(r.Body).Decode(&req)
		if req.DependsOn != "t2" {
			t.Errorf("depends_on = %s", req.DependsOn)
		}
		_ = json.NewEncoder(w).Encode(DependencyResponse{})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	_, err := c.AddDependency(ctx, "t1", &AddDependencyRequest{DependsOn: "t2"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteDependency(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("method = %s", r.Method)
		}
		if r.URL.Query().Get("depends_on") != "t2" {
			t.Errorf("depends_on = %s", r.URL.Query().Get("depends_on"))
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	if err := c.DeleteDependency(ctx, "t1", "t2", ""); err != nil {
		t.Fatal(err)
	}
}

func TestAddTaskLink(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/v2/task/t1/link/t2" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		_ = json.NewEncoder(w).Encode(TaskLinkResponse{})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	_, err := c.AddTaskLink(ctx, "t1", "t2")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteTaskLink(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" || r.URL.Path != "/v2/task/t1/link/t2" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	if err := c.DeleteTaskLink(ctx, "t1", "t2"); err != nil {
		t.Fatal(err)
	}
}
