package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTimeEntries(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method = %s", r.Method)
		}
		if r.URL.Path != "/v2/team/123/time_entries" {
			t.Errorf("path = %s", r.URL.Path)
		}
		json.NewEncoder(w).Encode(TimeEntriesResponse{Data: []TimeEntry{{ID: "te1"}}})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resp, err := c.GetTimeEntries(ctx, "123", nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data) != 1 {
		t.Errorf("count = %d", len(resp.Data))
	}
}

func TestGetTimeEntriesWithOpts(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("assignee") != "5" {
			t.Errorf("assignee = %s", r.URL.Query().Get("assignee"))
		}
		json.NewEncoder(w).Encode(TimeEntriesResponse{})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	_, err := c.GetTimeEntries(ctx, "123", &ListTimeEntriesOptions{Assignee: "5"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateTimeEntry(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/v2/team/123/time_entries" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		json.NewEncoder(w).Encode(TimeEntry{ID: "te1"})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resp, err := c.CreateTimeEntry(ctx, "123", &CreateTimeEntryRequest{Start: 1000, Duration: 3600000})
	if err != nil {
		t.Fatal(err)
	}
	if resp.ID != "te1" {
		t.Errorf("id = %s", resp.ID)
	}
}

func TestGetTimeEntry(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/team/123/time_entries/te1" {
			t.Errorf("path = %s", r.URL.Path)
		}
		json.NewEncoder(w).Encode(SingleTimeEntryResponse{Data: TimeEntry{ID: "te1"}})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resp, err := c.GetTimeEntry(ctx, "123", "te1")
	if err != nil {
		t.Fatal(err)
	}
	if resp.Data.ID != "te1" {
		t.Errorf("id = %s", resp.Data.ID)
	}
}

func TestDeleteTimeEntry(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" || r.URL.Path != "/v2/team/123/time_entries/te1" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	if err := c.DeleteTimeEntry(ctx, "123", "te1"); err != nil {
		t.Fatal(err)
	}
}

func TestStartTimer(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/v2/team/123/time_entries/start" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		json.NewEncoder(w).Encode(SingleTimeEntryResponse{Data: TimeEntry{ID: "te2"}})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resp, err := c.StartTimer(ctx, "123", &StartTimerRequest{Tid: "t1"})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Data.ID != "te2" {
		t.Errorf("id = %s", resp.Data.ID)
	}
}

func TestStopTimer(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/v2/team/123/time_entries/stop" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		json.NewEncoder(w).Encode(SingleTimeEntryResponse{Data: TimeEntry{ID: "te1"}})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	_, err := c.StopTimer(ctx, "123")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetRunningTimer(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/team/123/time_entries/current" {
			t.Errorf("path = %s", r.URL.Path)
		}
		json.NewEncoder(w).Encode(SingleTimeEntryResponse{Data: TimeEntry{ID: "te1"}})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	_, err := c.GetRunningTimer(ctx, "123", "")
	if err != nil {
		t.Fatal(err)
	}
}
