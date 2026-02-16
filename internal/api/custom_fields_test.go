package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetListCustomFields(t *testing.T) {
	tests := []struct {
		name   string
		listID string
		path   string
	}{
		{"basic", "123", "/v2/list/123/field"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != tt.path {
					t.Errorf("path = %s, want %s", r.URL.Path, tt.path)
				}
				if r.Method != "GET" {
					t.Errorf("method = %s, want GET", r.Method)
				}
				json.NewEncoder(w).Encode(CustomFieldsResponse{Fields: []CustomField{{ID: "f1", Name: "Priority"}}})
			}))
			defer srv.Close()
			c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
			resp, err := c.GetListCustomFields(tt.listID)
			if err != nil {
				t.Fatal(err)
			}
			if len(resp.Fields) != 1 || resp.Fields[0].ID != "f1" {
				t.Errorf("unexpected fields: %+v", resp.Fields)
			}
		})
	}
}

func TestGetFolderCustomFields(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/folder/456/field" {
			t.Errorf("path: %s", r.URL.Path)
		}
		json.NewEncoder(w).Encode(CustomFieldsResponse{})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	_, err := c.GetFolderCustomFields("456")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetSpaceCustomFields(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/space/789/field" {
			t.Errorf("path: %s", r.URL.Path)
		}
		json.NewEncoder(w).Encode(CustomFieldsResponse{})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	_, err := c.GetSpaceCustomFields("789")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetWorkspaceCustomFields(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/team/111/field" {
			t.Errorf("path: %s", r.URL.Path)
		}
		json.NewEncoder(w).Encode(CustomFieldsResponse{})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	_, err := c.GetWorkspaceCustomFields("111")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetCustomFieldValue(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s", r.Method)
		}
		if r.URL.Path != "/v2/task/t1/field/f1" {
			t.Errorf("path = %s", r.URL.Path)
		}
		var req SetCustomFieldRequest
		json.NewDecoder(r.Body).Decode(&req)
		if req.Value != "hello" {
			t.Errorf("value = %v", req.Value)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	err := c.SetCustomFieldValue("t1", "f1", &SetCustomFieldRequest{Value: "hello"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveCustomFieldValue(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("method = %s", r.Method)
		}
		if r.URL.Path != "/v2/task/t1/field/f1" {
			t.Errorf("path = %s", r.URL.Path)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	err := c.RemoveCustomFieldValue("t1", "f1")
	if err != nil {
		t.Fatal(err)
	}
}
