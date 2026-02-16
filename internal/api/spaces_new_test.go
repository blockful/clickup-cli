package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateSpace(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("method: %s", r.Method)
		}
		var req UpdateSpaceRequest
		json.NewDecoder(r.Body).Decode(&req)
		if req.Name != "newname" {
			t.Errorf("name: %s", req.Name)
		}
		json.NewEncoder(w).Encode(Space{ID: "s1", Name: "newname"})
	}))
	defer srv.Close()

	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	s, err := c.UpdateSpace("s1", &UpdateSpaceRequest{Name: "newname"})
	if err != nil {
		t.Fatal(err)
	}
	if s.Name != "newname" {
		t.Errorf("name: %s", s.Name)
	}
}

func TestDeleteSpace(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("method: %s", r.Method)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()

	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	if err := c.DeleteSpace("s1"); err != nil {
		t.Fatal(err)
	}
}
