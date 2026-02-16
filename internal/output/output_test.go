package output

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"testing"
)

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	return buf.String()
}

func captureStderr(f func()) string {
	old := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w
	f()
	w.Close()
	os.Stderr = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	return buf.String()
}

func TestJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		wantKeys []string
	}{
		{
			name:     "map output",
			input:    map[string]string{"id": "123", "name": "test"},
			wantKeys: []string{"id", "name"},
		},
		{
			name:  "array output",
			input: []string{"a", "b", "c"},
		},
		{
			name:  "nested struct",
			input: map[string]interface{}{"task": map[string]string{"id": "1", "name": "nested"}},
		},
		{
			name:  "empty map",
			input: map[string]string{},
		},
		{
			name:  "nil value",
			input: (*string)(nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := captureStdout(func() {
				JSON(tt.input)
			})

			// Verify it's valid JSON
			var parsed interface{}
			if err := json.Unmarshal([]byte(out), &parsed); err != nil {
				t.Fatalf("output is not valid JSON: %v\nOutput: %s", err, out)
			}

			// For map outputs, verify expected keys exist
			if tt.wantKeys != nil {
				m, ok := parsed.(map[string]interface{})
				if !ok {
					t.Fatalf("expected map output, got %T", parsed)
				}
				for _, k := range tt.wantKeys {
					if _, exists := m[k]; !exists {
						t.Errorf("missing key %q in output", k)
					}
				}
			}
		})
	}
}

func TestPrintError(t *testing.T) {
	tests := []struct {
		name    string
		code    string
		message string
	}{
		{name: "not found", code: "NOT_FOUND", message: "resource not found"},
		{name: "unauthorized", code: "UNAUTHORIZED", message: "invalid token"},
		{name: "validation", code: "VALIDATION_ERROR", message: "--id is required"},
		{name: "empty message", code: "ERROR", message: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := captureStderr(func() {
				PrintError(tt.code, tt.message)
			})

			var result ErrorResponse
			if err := json.Unmarshal([]byte(out), &result); err != nil {
				t.Fatalf("error output is not valid JSON: %v\nOutput: %s", err, out)
			}
			if result.Code != tt.code {
				t.Errorf("expected code %q, got %q", tt.code, result.Code)
			}
			if result.Error != tt.message {
				t.Errorf("expected error %q, got %q", tt.message, result.Error)
			}
		})
	}
}

func TestJSON_GoldenFile(t *testing.T) {
	input := map[string]interface{}{
		"id":   "task_123",
		"name": "Test Task",
		"status": map[string]string{
			"status": "open",
			"color":  "#000000",
		},
	}

	out := captureStdout(func() {
		JSON(input)
	})

	// Verify structure matches expected golden output
	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(out), &parsed); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if parsed["id"] != "task_123" {
		t.Errorf("expected id 'task_123', got %v", parsed["id"])
	}
	status, ok := parsed["status"].(map[string]interface{})
	if !ok {
		t.Fatal("expected status to be a map")
	}
	if status["status"] != "open" {
		t.Errorf("expected status 'open', got %v", status["status"])
	}
}
