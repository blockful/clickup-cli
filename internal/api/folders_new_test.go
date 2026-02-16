package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateFolder(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("method: %s", r.Method)
		}
		json.NewEncoder(w).Encode(Folder{ID: "f1", Name: "updated"})
	}))
	defer srv.Close()

	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	f, err := c.UpdateFolder(ctx, "f1", &UpdateFolderRequest{Name: "updated"})
	if err != nil {
		t.Fatal(err)
	}
	if f.Name != "updated" {
		t.Errorf("name: %s", f.Name)
	}
}

func TestDeleteFolder(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("method: %s", r.Method)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()

	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	if err := c.DeleteFolder(ctx, "f1"); err != nil {
		t.Fatal(err)
	}
}
