package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListFolderlessLists(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/space/s1/list" {
			t.Errorf("path: %s", r.URL.Path)
		}
		_ = json.NewEncoder(w).Encode(ListsResponse{Lists: []List{{ID: "l1"}}})
	}))
	defer srv.Close()

	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resp, err := c.ListFolderlessLists(ctx, "s1")
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Lists) != 1 {
		t.Errorf("lists: %d", len(resp.Lists))
	}
}

func TestCreateListWithFields(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req CreateListRequest
		_ = json.NewDecoder(r.Body).Decode(&req)
		if req.Content != "desc" {
			t.Errorf("content: %s", req.Content)
		}
		if req.DueDate == nil || *req.DueDate != 999 {
			t.Error("due_date")
		}
		_ = json.NewEncoder(w).Encode(List{ID: "l1"})
	}))
	defer srv.Close()

	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	dd := int64(999)
	_, err := c.CreateList(ctx, "f1", &CreateListRequest{Name: "test", Content: "desc", DueDate: &dd})
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateList(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("method: %s", r.Method)
		}
		_ = json.NewEncoder(w).Encode(List{ID: "l1"})
	}))
	defer srv.Close()

	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	_, err := c.UpdateList(ctx, "l1", &UpdateListRequest{Name: "new"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteList(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("method: %s", r.Method)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()

	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	if err := c.DeleteList(ctx, "l1"); err != nil {
		t.Fatal(err)
	}
}
