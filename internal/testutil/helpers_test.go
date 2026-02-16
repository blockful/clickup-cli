package testutil

import (
	"testing"
)

func TestAssertJSONEqual(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		wantFail bool
	}{
		{name: "equal objects", a: `{"a":1,"b":2}`, b: `{"b":2,"a":1}`},
		{name: "equal arrays", a: `[1,2,3]`, b: `[1,2,3]`},
		{name: "equal strings", a: `"hello"`, b: `"hello"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This doesn't fail for equal inputs
			AssertJSONEqual(t, tt.a, tt.b)
		})
	}
}

func TestNewTestServer(t *testing.T) {
	server := NewTestServer(t, 200, `{"ok":true}`)
	defer server.Close()

	if server.URL == "" {
		t.Error("server URL is empty")
	}
}

func TestNewTestServerWithValidation(t *testing.T) {
	server := NewTestServerWithValidation(t, "GET", "/v2/user", 200, `{"user":{}}`)
	defer server.Close()

	if server.URL == "" {
		t.Error("server URL is empty")
	}
}
