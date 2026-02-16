package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTimeEntryHistory(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/v2/team/123/time_entries/te1/history" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		_ = json.NewEncoder(w).Encode(TimeEntryHistoryResponse{Data: []interface{}{"h1"}})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resp, err := c.GetTimeEntryHistory(context.Background(), "123", "te1")
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data) != 1 {
		t.Errorf("count = %d", len(resp.Data))
	}
}

func TestAddTagsToTimeEntries(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/v2/team/123/time_entries/tags" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	err := c.AddTagsToTimeEntries(context.Background(), "123", &AddTagsToTimeEntriesRequest{
		TimeEntryIDs: []string{"te1"},
		Tags:         []Tag{{Name: "bug"}},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveTagsFromTimeEntries(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" || r.URL.Path != "/v2/team/123/time_entries/tags" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	err := c.RemoveTagsFromTimeEntries(context.Background(), "123", &RemoveTagsFromTimeEntriesRequest{
		TimeEntryIDs: []string{"te1"},
		Tags:         []Tag{{Name: "bug"}},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestChangeTagNames(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" || r.URL.Path != "/v2/team/123/time_entries/tags" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	err := c.ChangeTagNames(context.Background(), "123", &ChangeTagNameRequest{
		Name:    "old",
		NewName: "new",
		TagBg:   "#000",
		TagFg:   "#fff",
	})
	if err != nil {
		t.Fatal(err)
	}
}
