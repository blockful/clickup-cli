package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientDo_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "pk_test123" {
			t.Errorf("expected Authorization header 'pk_test123', got '%s'", r.Header.Get("Authorization"))
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected Content-Type 'application/json', got '%s'", r.Header.Get("Content-Type"))
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}))
	defer server.Close()

	client := NewClient("pk_test123")
	client.BaseURL = server.URL

	var result map[string]string
	err := client.Do("GET", "/test", nil, &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result["status"] != "ok" {
		t.Errorf("expected status 'ok', got '%s'", result["status"])
	}
}

func TestClientDo_Unauthorized(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(map[string]string{"err": "Token invalid", "ECODE": "OAUTH_017"})
	}))
	defer server.Close()

	client := NewClient("bad_token")
	client.BaseURL = server.URL

	err := client.Do("GET", "/test", nil, nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	clientErr, ok := err.(*ClientError)
	if !ok {
		t.Fatalf("expected *ClientError, got %T", err)
	}
	if clientErr.Code != "UNAUTHORIZED" {
		t.Errorf("expected code 'UNAUTHORIZED', got '%s'", clientErr.Code)
	}
	if clientErr.Message != "Token invalid" {
		t.Errorf("expected message 'Token invalid', got '%s'", clientErr.Message)
	}
}

func TestClientDo_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(map[string]string{"err": "Resource not found"})
	}))
	defer server.Close()

	client := NewClient("pk_test")
	client.BaseURL = server.URL

	err := client.Do("GET", "/missing", nil, nil)
	if err == nil {
		t.Fatal("expected error")
	}
	clientErr := err.(*ClientError)
	if clientErr.Code != "NOT_FOUND" {
		t.Errorf("expected NOT_FOUND, got %s", clientErr.Code)
	}
}

func TestClientDo_PostWithBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		var body map[string]string
		json.NewDecoder(r.Body).Decode(&body)
		if body["name"] != "Test Task" {
			t.Errorf("expected name 'Test Task', got '%s'", body["name"])
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]string{"id": "abc123", "name": "Test Task"})
	}))
	defer server.Close()

	client := NewClient("pk_test")
	client.BaseURL = server.URL

	reqBody := map[string]string{"name": "Test Task"}
	var result map[string]string
	err := client.Do("POST", "/task", reqBody, &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result["id"] != "abc123" {
		t.Errorf("expected id 'abc123', got '%s'", result["id"])
	}
}

func TestClientDo_RateLimited(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(429)
		json.NewEncoder(w).Encode(map[string]string{"err": "Rate limit exceeded"})
	}))
	defer server.Close()

	client := NewClient("pk_test")
	client.BaseURL = server.URL

	err := client.Do("GET", "/test", nil, nil)
	clientErr := err.(*ClientError)
	if clientErr.Code != "RATE_LIMITED" {
		t.Errorf("expected RATE_LIMITED, got %s", clientErr.Code)
	}
}
