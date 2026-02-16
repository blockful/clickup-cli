package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSpaceTags(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/space/512/tag" {
			t.Errorf("path: %s", r.URL.Path)
		}
		json.NewEncoder(w).Encode(TagsResponse{Tags: []Tag{{Name: "bug", TagFg: "#fff", TagBg: "#f00"}}})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resp, err := c.GetSpaceTags(ctx, "512")
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Tags) != 1 || resp.Tags[0].Name != "bug" {
		t.Errorf("unexpected: %+v", resp.Tags)
	}
}

func TestCreateSpaceTag(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/v2/space/512/tag" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		var req CreateTagRequest
		json.NewDecoder(r.Body).Decode(&req)
		if req.Tag.Name != "feature" {
			t.Errorf("name = %s", req.Tag.Name)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	err := c.CreateSpaceTag(ctx, "512", &CreateTagRequest{Tag: Tag{Name: "feature", TagFg: "#000", TagBg: "#0f0"}})
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSpaceTag(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" || r.URL.Path != "/v2/space/512/tag/old" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	err := c.UpdateSpaceTag(ctx, "512", "old", &UpdateTagRequest{Tag: Tag{Name: "new"}})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteSpaceTag(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" || r.URL.Path != "/v2/space/512/tag/old" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	err := c.DeleteSpaceTag(ctx, "512", "old")
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddTagToTask(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/v2/task/t1/tag/bug" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	if err := c.AddTagToTask(ctx, "t1", "bug"); err != nil {
		t.Fatal(err)
	}
}

func TestRemoveTagFromTask(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" || r.URL.Path != "/v2/task/t1/tag/bug" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	if err := c.RemoveTagFromTask(ctx, "t1", "bug"); err != nil {
		t.Fatal(err)
	}
}
