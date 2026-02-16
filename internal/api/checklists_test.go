package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateChecklist(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/v2/task/t1/checklist" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		var req CreateChecklistRequest
		_ = json.NewDecoder(r.Body).Decode(&req)
		if req.Name != "My Checklist" {
			t.Errorf("name = %s", req.Name)
		}
		_ = json.NewEncoder(w).Encode(ChecklistResponse{Checklist: ChecklistDetailed{ID: "cl1", Name: "My Checklist"}})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resp, err := c.CreateChecklist(ctx, "t1", &CreateChecklistRequest{Name: "My Checklist"})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Checklist.ID != "cl1" {
		t.Errorf("id = %s", resp.Checklist.ID)
	}
}

func TestEditChecklist(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" || r.URL.Path != "/v2/checklist/cl1" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	if err := c.EditChecklist(ctx, "cl1", &EditChecklistRequest{Name: "Updated"}); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteChecklist(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" || r.URL.Path != "/v2/checklist/cl1" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	if err := c.DeleteChecklist(ctx, "cl1"); err != nil {
		t.Fatal(err)
	}
}

func TestCreateChecklistItem(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/v2/checklist/cl1/checklist_item" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		_ = json.NewEncoder(w).Encode(ChecklistResponse{Checklist: ChecklistDetailed{ID: "cl1"}})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	_, err := c.CreateChecklistItem(ctx, "cl1", &CreateChecklistItemRequest{Name: "Item 1"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestEditChecklistItem(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" || r.URL.Path != "/v2/checklist/cl1/checklist_item/i1" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		_ = json.NewEncoder(w).Encode(ChecklistResponse{Checklist: ChecklistDetailed{ID: "cl1"}})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resolved := true
	_, err := c.EditChecklistItem(ctx, "cl1", "i1", &EditChecklistItemRequest{Resolved: &resolved})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteChecklistItem(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" || r.URL.Path != "/v2/checklist/cl1/checklist_item/i1" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	if err := c.DeleteChecklistItem(ctx, "cl1", "i1"); err != nil {
		t.Fatal(err)
	}
}
