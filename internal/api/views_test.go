package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTeamViews(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/team/123/view" {
			t.Errorf("path = %s", r.URL.Path)
		}
		json.NewEncoder(w).Encode(ViewsResponse{Views: []View{{ID: "v1"}}})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resp, err := c.GetTeamViews("123")
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Views) != 1 {
		t.Errorf("count = %d", len(resp.Views))
	}
}

func TestGetView(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/view/v1" {
			t.Errorf("path = %s", r.URL.Path)
		}
		json.NewEncoder(w).Encode(ViewResponse{View: View{ID: "v1", Name: "Test"}})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resp, err := c.GetView("v1")
	if err != nil {
		t.Fatal(err)
	}
	if resp.View.Name != "Test" {
		t.Errorf("name = %s", resp.View.Name)
	}
}

func TestDeleteView(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" || r.URL.Path != "/v2/view/v1" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	if err := c.DeleteView("v1"); err != nil {
		t.Fatal(err)
	}
}

func TestGetViewTasks(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/view/v1/task" {
			t.Errorf("path = %s", r.URL.Path)
		}
		json.NewEncoder(w).Encode(ViewTasksResponse{LastPage: true})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resp, err := c.GetViewTasks("v1", 0)
	if err != nil {
		t.Fatal(err)
	}
	if !resp.LastPage {
		t.Error("expected last_page=true")
	}
}
