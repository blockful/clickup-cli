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
	io.Copy(&buf, r)
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
	io.Copy(&buf, r)
	return buf.String()
}

func TestJSON_ValidOutput(t *testing.T) {
	out := captureStdout(func() {
		JSON(map[string]string{"id": "123", "name": "test"})
	})

	var result map[string]string
	if err := json.Unmarshal([]byte(out), &result); err != nil {
		t.Fatalf("output is not valid JSON: %v\nOutput: %s", err, out)
	}
	if result["id"] != "123" {
		t.Errorf("expected id '123', got '%s'", result["id"])
	}
}

func TestJSON_Array(t *testing.T) {
	out := captureStdout(func() {
		JSON([]string{"a", "b", "c"})
	})

	var result []string
	if err := json.Unmarshal([]byte(out), &result); err != nil {
		t.Fatalf("output is not valid JSON: %v", err)
	}
	if len(result) != 3 {
		t.Errorf("expected 3 items, got %d", len(result))
	}
}

func TestPrintError_ValidJSON(t *testing.T) {
	out := captureStderr(func() {
		PrintError("NOT_FOUND", "resource not found")
	})

	var result ErrorResponse
	if err := json.Unmarshal([]byte(out), &result); err != nil {
		t.Fatalf("error output is not valid JSON: %v\nOutput: %s", err, out)
	}
	if result.Code != "NOT_FOUND" {
		t.Errorf("expected code 'NOT_FOUND', got '%s'", result.Code)
	}
	if result.Error != "resource not found" {
		t.Errorf("expected error 'resource not found', got '%s'", result.Error)
	}
}
