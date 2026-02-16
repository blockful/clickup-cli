package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateTaskAttachment(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/v2/task/t1/attachment" {
			t.Errorf("unexpected: %s %s", r.Method, r.URL.Path)
		}
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			t.Fatal(err)
		}
		file, _, err := r.FormFile("attachment")
		if err != nil {
			t.Fatal(err)
		}
		file.Close()
		_ = json.NewEncoder(w).Encode(Attachment{ID: "att1", Title: "test.txt"})
	}))
	defer srv.Close()

	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	if err := os.WriteFile(tmpFile, []byte("hello"), 0644); err != nil {
		t.Fatal(err)
	}

	c := &Client{BaseURL: srv.URL, Token: "test", HTTPClient: srv.Client()}
	att, err := c.CreateTaskAttachment(ctx, "t1", tmpFile)
	if err != nil {
		t.Fatal(err)
	}
	if att.ID != "att1" {
		t.Errorf("id = %s", att.ID)
	}
}
