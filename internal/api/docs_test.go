package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateDoc(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/v3/workspaces/w1/docs" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		var req CreateDocRequest
		json.NewDecoder(r.Body).Decode(&req)
		if req.Name != "My Doc" {
			t.Errorf("name = %s", req.Name)
		}
		json.NewEncoder(w).Encode(Doc{ID: "d1", Name: "My Doc"})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resp, err := c.CreateDoc("w1", &CreateDocRequest{Name: "My Doc"})
	if err != nil {
		t.Fatal(err)
	}
	if resp.ID != "d1" {
		t.Errorf("id = %s", resp.ID)
	}
}

func TestSearchDocs(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v3/workspaces/w1/docs" {
			t.Errorf("path: %s", r.URL.Path)
		}
		json.NewEncoder(w).Encode(DocsResponse{Docs: []Doc{{ID: "d1"}}})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resp, err := c.SearchDocs("w1")
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Docs) != 1 {
		t.Errorf("docs count = %d", len(resp.Docs))
	}
}

func TestGetDoc(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v3/workspaces/w1/docs/d1" {
			t.Errorf("path: %s", r.URL.Path)
		}
		json.NewEncoder(w).Encode(Doc{ID: "d1"})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resp, err := c.GetDoc("w1", "d1")
	if err != nil {
		t.Fatal(err)
	}
	if resp.ID != "d1" {
		t.Errorf("id = %s", resp.ID)
	}
}

func TestCreatePage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/v3/workspaces/w1/docs/d1/pages" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		json.NewEncoder(w).Encode(DocPage{ID: "p1", Name: "Page 1"})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resp, err := c.CreatePage("w1", "d1", &CreatePageRequest{Name: "Page 1"})
	if err != nil {
		t.Fatal(err)
	}
	if resp.ID != "p1" {
		t.Errorf("id = %s", resp.ID)
	}
}

func TestGetPage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v3/workspaces/w1/docs/d1/pages/p1" {
			t.Errorf("path: %s", r.URL.Path)
		}
		json.NewEncoder(w).Encode(DocPage{ID: "p1"})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resp, err := c.GetPage("w1", "d1", "p1")
	if err != nil {
		t.Fatal(err)
	}
	if resp.ID != "p1" {
		t.Errorf("id = %s", resp.ID)
	}
}

func TestEditPage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" || r.URL.Path != "/v3/workspaces/w1/docs/d1/pages/p1" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		json.NewEncoder(w).Encode(DocPage{ID: "p1", Name: "Updated"})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resp, err := c.EditPage("w1", "d1", "p1", &EditPageRequest{Name: "Updated"})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Name != "Updated" {
		t.Errorf("name = %s", resp.Name)
	}
}

func TestGetDocPageListing(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v3/workspaces/w1/docs/d1/page_listing" {
			t.Errorf("path: %s", r.URL.Path)
		}
		json.NewEncoder(w).Encode(DocPagesResponse{Pages: []DocPage{{ID: "p1"}, {ID: "p2"}}})
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	resp, err := c.GetDocPageListing("w1", "d1")
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Pages) != 2 {
		t.Errorf("pages count = %d", len(resp.Pages))
	}
}
