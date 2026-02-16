package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetLegacyTrackedTime(t *testing.T) {
	tests := []struct {
		name          string
		taskID        string
		subcategoryID string
		wantPath      string
	}{
		{"without subcategory", "abc", "", "/v2/task/abc/time"},
		{"with subcategory", "abc", "sub1", "/v2/task/abc/time"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "GET" {
					t.Errorf("method = %s", r.Method)
				}
				if r.URL.Path != tt.wantPath {
					t.Errorf("path = %s, want %s", r.URL.Path, tt.wantPath)
				}
				if tt.subcategoryID != "" && r.URL.Query().Get("subcategory_id") != tt.subcategoryID {
					t.Errorf("subcategory_id = %s", r.URL.Query().Get("subcategory_id"))
				}
				_ = json.NewEncoder(w).Encode(LegacyTimeResponse{Data: []LegacyTimeInterval{{ID: "i1"}}})
			}))
			defer srv.Close()
			c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
			resp, err := c.GetLegacyTrackedTime(context.Background(), tt.taskID, tt.subcategoryID)
			if err != nil {
				t.Fatal(err)
			}
			if len(resp.Data) != 1 || resp.Data[0].ID != "i1" {
				t.Errorf("unexpected response: %+v", resp)
			}
		})
	}
}

func TestTrackLegacyTime(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/v2/task/abc/time" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		_ = json.NewEncoder(w).Encode(LegacyTimeResponse{Data: []LegacyTimeInterval{{ID: "i2"}}})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resp, err := c.TrackLegacyTime(context.Background(), "abc", &LegacyTrackTimeRequest{Time: 3600000})
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data) != 1 {
		t.Errorf("count = %d", len(resp.Data))
	}
}

func TestEditLegacyTime(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" || r.URL.Path != "/v2/task/abc/time/i1" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	if err := c.EditLegacyTime(context.Background(), "abc", "i1", &LegacyEditTimeRequest{Time: 1000}); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteLegacyTime(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" || r.URL.Path != "/v2/task/abc/time/i1" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	if err := c.DeleteLegacyTime(context.Background(), "abc", "i1"); err != nil {
		t.Fatal(err)
	}
}
