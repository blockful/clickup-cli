package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListFolders(t *testing.T) {
	tests := []struct {
		name       string
		spaceID    string
		response   string
		statusCode int
		wantErr    bool
		errCode    string
		wantCount  int
	}{
		{
			name:       "success",
			spaceID:    "s1",
			response:   `{"folders":[{"id":"f1","name":"Folder 1"}]}`,
			statusCode: 200,
			wantCount:  1,
		},
		{
			name:       "not found",
			spaceID:    "bad",
			response:   `{"err":"Space not found"}`,
			statusCode: 404,
			wantErr:    true,
			errCode:    "NOT_FOUND",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.response))
			}))
			defer server.Close()

			client := NewClient("pk_test")
			client.BaseURL = server.URL

			resp, err := client.ListFolders(tt.spaceID)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				if ce, ok := err.(*ClientError); ok && ce.Code != tt.errCode {
					t.Errorf("expected %q, got %q", tt.errCode, ce.Code)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(resp.Folders) != tt.wantCount {
				t.Errorf("expected %d, got %d", tt.wantCount, len(resp.Folders))
			}
		})
	}
}

func TestGetFolder(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(`{"id":"f1","name":"Test Folder"}`))
	}))
	defer server.Close()

	client := NewClient("pk_test")
	client.BaseURL = server.URL

	folder, err := client.GetFolder("f1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if folder.Name != "Test Folder" {
		t.Errorf("expected 'Test Folder', got %q", folder.Name)
	}
}

func TestCreateFolder(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		w.WriteHeader(200)
		w.Write([]byte(`{"id":"f2","name":"New Folder"}`))
	}))
	defer server.Close()

	client := NewClient("pk_test")
	client.BaseURL = server.URL

	folder, err := client.CreateFolder("s1", &CreateFolderRequest{Name: "New Folder"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if folder.ID != "f2" {
		t.Errorf("expected ID 'f2', got %q", folder.ID)
	}
}
