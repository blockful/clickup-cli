package testutil

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

// NewTestServer creates an httptest.Server that responds with the given status and body.
func NewTestServer(t *testing.T, statusCode int, responseBody string) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		w.Write([]byte(responseBody))
	}))
}

// NewTestServerWithValidation creates a server that validates method/path and responds.
func NewTestServerWithValidation(t *testing.T, wantMethod, wantPath string, statusCode int, responseBody string) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != wantMethod {
			t.Errorf("expected method %s, got %s", wantMethod, r.Method)
		}
		if r.URL.Path != wantPath {
			t.Errorf("expected path %s, got %s", wantPath, r.URL.Path)
		}
		if r.Header.Get("Authorization") == "" {
			t.Error("missing Authorization header")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		w.Write([]byte(responseBody))
	}))
}

// AssertJSONEqual compares two JSON strings for semantic equality.
func AssertJSONEqual(t *testing.T, expected, actual string) {
	t.Helper()
	var e, a interface{}
	if err := json.Unmarshal([]byte(expected), &e); err != nil {
		t.Fatalf("invalid expected JSON: %v", err)
	}
	if err := json.Unmarshal([]byte(actual), &a); err != nil {
		t.Fatalf("invalid actual JSON: %v", err)
	}
	eb, _ := json.Marshal(e)
	ab, _ := json.Marshal(a)
	if string(eb) != string(ab) {
		t.Errorf("JSON mismatch:\nexpected: %s\nactual:   %s", eb, ab)
	}
}

// GoldenFile compares output against a golden file. If the -update flag or
// UPDATE_GOLDEN env var is set, updates the golden file instead.
func GoldenFile(t *testing.T, name string, actual []byte) {
	t.Helper()
	golden := filepath.Join("testdata", name+".golden")

	if os.Getenv("UPDATE_GOLDEN") != "" {
		os.MkdirAll(filepath.Dir(golden), 0o755)
		if err := os.WriteFile(golden, actual, 0o644); err != nil {
			t.Fatalf("failed to update golden file: %v", err)
		}
		return
	}

	expected, err := os.ReadFile(golden)
	if err != nil {
		t.Fatalf("golden file %s not found (run with UPDATE_GOLDEN=1 to create): %v", golden, err)
	}

	if string(expected) != string(actual) {
		t.Errorf("golden file mismatch for %s:\nexpected:\n%s\nactual:\n%s", name, expected, actual)
	}
}
