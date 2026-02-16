package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/blockful/clickup-cli/internal/api"
	"github.com/spf13/viper"
)

// runCommand executes a CLI command against a mock server and captures stdout.
func runCommand(t *testing.T, serverURL string, args ...string) (string, error) {
	t.Helper()

	// Set up client factory to use mock server
	oldFactory := clientFactory
	clientFactory = func() api.ClientInterface {
		c := api.NewClient("test-token")
		c.BaseURL = serverURL + "/api"
		c.MaxRetries = 0
		return c
	}
	defer func() { clientFactory = oldFactory }()

	// Set workspace in viper for commands that need it
	viper.Set("workspace", "12345678")
	defer viper.Reset()

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Reset and execute root command
	rootCmd.SetArgs(args)
	err := rootCmd.Execute()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)

	return buf.String(), err
}

// requestLog captures HTTP requests for verification.
type requestLog struct {
	mu      sync.Mutex
	Method  string
	Path    string
	Query   string
	Body    string
	Headers http.Header
}

func newMockServer(t *testing.T, handler http.HandlerFunc) (*httptest.Server, *requestLog) {
	t.Helper()
	log := &requestLog{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		log.mu.Lock()
		log.Method = r.Method
		log.Path = r.URL.Path
		log.Query = r.URL.RawQuery
		log.Body = string(body)
		log.Headers = r.Header.Clone()
		log.mu.Unlock()
		handler(w, r)
	}))
	t.Cleanup(server.Close)
	return server, log
}

func mustContainJSON(t *testing.T, output, key, expected string) {
	t.Helper()
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(output), &m); err != nil {
		// Try as array wrapper
		t.Fatalf("invalid JSON output: %v\n%s", err, output)
	}
	// Simple string check for nested values
	if !strings.Contains(output, expected) {
		t.Errorf("output missing expected value %q for key %q:\n%s", expected, key, output)
	}
}

// --- Workspace List ---

func TestWorkspaceList(t *testing.T) {
	server, log := newMockServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"teams": []map[string]interface{}{
				{
					"id":     "90901234567",
					"name":   "Blockful",
					"color":  "#536cfe",
					"avatar": "https://example.com/avatar.png",
					"members": []map[string]interface{}{
						{"user": map[string]interface{}{"id": 12345678, "username": "alice", "email": "alice@example.com"}},
					},
				},
				{
					"id":      "90907654321",
					"name":    "Personal",
					"color":   "#ff6900",
					"avatar":  "",
					"members": []map[string]interface{}{},
				},
			},
		})
	})

	out, err := runCommand(t, server.URL, "workspace", "list")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if log.Method != "GET" {
		t.Errorf("expected GET, got %s", log.Method)
	}
	if log.Path != "/api/v2/team" {
		t.Errorf("expected /api/v2/team, got %s", log.Path)
	}
	mustContainJSON(t, out, "teams[0].name", "Blockful")
	mustContainJSON(t, out, "teams[1].name", "Personal")
	mustContainJSON(t, out, "teams[0].id", "90901234567")
}

// --- Space List ---

func TestSpaceList(t *testing.T) {
	server, log := newMockServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"spaces": []map[string]interface{}{
				{
					"id":                 "98765432",
					"name":               "Engineering",
					"private":            false,
					"multiple_assignees": true,
					"features":           map[string]interface{}{"due_dates": map[string]interface{}{"enabled": true}},
				},
			},
		})
	})

	out, err := runCommand(t, server.URL, "space", "list", "--workspace", "90901234567")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if log.Method != "GET" {
		t.Errorf("expected GET, got %s", log.Method)
	}
	if !strings.Contains(log.Path, "/api/v2/team/90901234567/space") {
		t.Errorf("unexpected path: %s", log.Path)
	}
	mustContainJSON(t, out, "spaces[0].name", "Engineering")
}

// --- Task List ---

func TestTaskList(t *testing.T) {
	server, log := newMockServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"tasks": []map[string]interface{}{
				{
					"id":            "abc123def",
					"name":          "Implement auth flow",
					"description":   "OAuth2 implementation",
					"status":        map[string]interface{}{"status": "in progress", "color": "#4194f6", "type": "custom"},
					"orderindex":    "1.00000",
					"date_created":  "1676000000000",
					"date_updated":  "1676100000000",
					"creator":       map[string]interface{}{"id": 12345678, "username": "alice", "email": "alice@example.com"},
					"assignees":     []map[string]interface{}{{"id": 12345678, "username": "alice"}},
					"tags":          []map[string]interface{}{{"name": "backend", "tag_fg": "#fff", "tag_bg": "#000"}},
					"parent":        nil,
					"priority":      map[string]interface{}{"id": "2", "priority": "high", "color": "#ffcc00"},
					"due_date":      "1677000000000",
					"points":        3.0,
					"time_estimate": 7200000,
					"custom_fields": []map[string]interface{}{
						{"id": "cf_uuid_1", "name": "Sprint", "type": "drop_down", "value": "Sprint 42"},
					},
					"url":    "https://app.clickup.com/t/abc123def",
					"list":   map[string]interface{}{"id": "901100200300", "name": "Sprint Backlog"},
					"folder": map[string]interface{}{"id": "800100200300", "name": "Product"},
					"space":  map[string]interface{}{"id": "98765432"},
				},
			},
		})
	})

	out, err := runCommand(t, server.URL, "task", "list", "--list", "901100200300")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if log.Method != "GET" {
		t.Errorf("expected GET, got %s", log.Method)
	}
	if !strings.Contains(log.Path, "/api/v2/list/901100200300/task") {
		t.Errorf("unexpected path: %s", log.Path)
	}
	mustContainJSON(t, out, "tasks[0].id", "abc123def")
	mustContainJSON(t, out, "tasks[0].name", "Implement auth flow")
	mustContainJSON(t, out, "priority", "high")
}

// --- Task Get ---

func TestTaskGet(t *testing.T) {
	server, log := newMockServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"id":                   "abc123def",
			"name":                 "Implement auth flow",
			"description":          "Full OAuth2 implementation with PKCE",
			"markdown_description": "# Auth Flow\n\nImplement OAuth2 with PKCE",
			"status":               map[string]interface{}{"status": "in progress", "color": "#4194f6", "type": "custom"},
			"orderindex":           "1.00000",
			"date_created":         "1676000000000",
			"date_updated":         "1676100000000",
			"creator":              map[string]interface{}{"id": 12345678, "username": "alice", "email": "alice@example.com"},
			"assignees":            []map[string]interface{}{{"id": 12345678, "username": "alice"}},
			"tags":                 []map[string]interface{}{},
			"parent":               nil,
			"priority":             map[string]interface{}{"id": "2", "priority": "high", "color": "#ffcc00"},
			"url":                  "https://app.clickup.com/t/abc123def",
			"list":                 map[string]interface{}{"id": "901100200300", "name": "Sprint Backlog"},
			"folder":               map[string]interface{}{"id": "800100200300", "name": "Product"},
			"space":                map[string]interface{}{"id": "98765432"},
		})
	})

	out, err := runCommand(t, server.URL, "task", "get", "--id", "abc123def")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if log.Method != "GET" {
		t.Errorf("expected GET, got %s", log.Method)
	}
	if !strings.Contains(log.Path, "/api/v2/task/abc123def") {
		t.Errorf("unexpected path: %s", log.Path)
	}
	mustContainJSON(t, out, "id", "abc123def")
	mustContainJSON(t, out, "markdown_description", "Auth Flow")
}

// --- Task Create ---

func TestTaskCreate(t *testing.T) {
	server, log := newMockServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"id":           "xyz789abc",
			"name":         "New feature task",
			"description":  "Build the new feature",
			"status":       map[string]interface{}{"status": "to do", "color": "#d3d3d3", "type": "custom"},
			"orderindex":   "2.00000",
			"date_created": "1676200000000",
			"creator":      map[string]interface{}{"id": 12345678, "username": "alice"},
			"assignees":    []map[string]interface{}{{"id": 12345678, "username": "alice"}},
			"tags":         []map[string]interface{}{},
			"parent":       nil,
			"priority":     map[string]interface{}{"id": "3", "priority": "normal", "color": "#6fddff"},
			"url":          "https://app.clickup.com/t/xyz789abc",
			"list":         map[string]interface{}{"id": "901100200300", "name": "Sprint Backlog"},
			"folder":       map[string]interface{}{"id": "800100200300", "name": "Product"},
			"space":        map[string]interface{}{"id": "98765432"},
		})
	})

	out, err := runCommand(t, server.URL, "task", "create",
		"--list", "901100200300",
		"--name", "New feature task",
		"--description", "Build the new feature",
		"--priority", "3",
		"--assignee", "12345678",
		"--markdown-content", "# New Feature\n\nDetails here",
		"--custom-item-id", "42",
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if log.Method != "POST" {
		t.Errorf("expected POST, got %s", log.Method)
	}
	if !strings.Contains(log.Path, "/api/v2/list/901100200300/task") {
		t.Errorf("unexpected path: %s", log.Path)
	}

	// Verify request body
	var reqBody map[string]interface{}
	_ = json.Unmarshal([]byte(log.Body), &reqBody)
	if reqBody["name"] != "New feature task" {
		t.Errorf("expected name 'New feature task', got %v", reqBody["name"])
	}
	if reqBody["description"] != "Build the new feature" {
		t.Errorf("expected description, got %v", reqBody["description"])
	}
	if reqBody["markdown_description"] != "# New Feature\n\nDetails here" {
		t.Errorf("expected markdown_description, got %v", reqBody["markdown_description"])
	}
	if v, ok := reqBody["priority"].(float64); !ok || v != 3 {
		t.Errorf("expected priority 3, got %v", reqBody["priority"])
	}
	if v, ok := reqBody["custom_item_id"].(float64); !ok || v != 42 {
		t.Errorf("expected custom_item_id 42, got %v", reqBody["custom_item_id"])
	}

	mustContainJSON(t, out, "id", "xyz789abc")
}

// --- Task Update ---

func TestTaskUpdate(t *testing.T) {
	server, log := newMockServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"id":           "abc123def",
			"name":         "Updated task name",
			"status":       map[string]interface{}{"status": "complete", "color": "#6bc950", "type": "closed"},
			"orderindex":   "1.00000",
			"date_created": "1676000000000",
			"creator":      map[string]interface{}{"id": 12345678, "username": "alice"},
			"assignees":    []map[string]interface{}{},
			"tags":         []map[string]interface{}{},
			"parent":       nil,
			"priority":     map[string]interface{}{"id": "1", "priority": "urgent", "color": "#f50000"},
			"url":          "https://app.clickup.com/t/abc123def",
			"list":         map[string]interface{}{"id": "901100200300", "name": "Sprint Backlog"},
			"folder":       map[string]interface{}{"id": "800100200300", "name": "Product"},
			"space":        map[string]interface{}{"id": "98765432"},
		})
	})

	out, err := runCommand(t, server.URL, "task", "update",
		"--id", "abc123def",
		"--name", "Updated task name",
		"--status", "complete",
		"--priority", "1",
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if log.Method != "PUT" {
		t.Errorf("expected PUT, got %s", log.Method)
	}
	if !strings.Contains(log.Path, "/api/v2/task/abc123def") {
		t.Errorf("unexpected path: %s", log.Path)
	}

	var reqBody map[string]interface{}
	_ = json.Unmarshal([]byte(log.Body), &reqBody)
	if reqBody["name"] != "Updated task name" {
		t.Errorf("expected name, got %v", reqBody["name"])
	}
	if reqBody["status"] != "complete" {
		t.Errorf("expected status 'complete', got %v", reqBody["status"])
	}

	mustContainJSON(t, out, "name", "Updated task name")
}

// --- Comment Create (tests json.Number for numeric ID) ---

func TestCommentCreate(t *testing.T) {
	server, log := newMockServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// ClickUp returns numeric comment IDs
		_, _ = w.Write([]byte(`{"id":457898234,"hist_id":"abc123hist","date":1676300000000}`))
	})

	out, err := runCommand(t, server.URL, "comment", "create",
		"--task", "abc123def",
		"--text", "Great progress on this task!",
		"--assignee", "12345678",
		"--notify-all",
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if log.Method != "POST" {
		t.Errorf("expected POST, got %s", log.Method)
	}
	if !strings.Contains(log.Path, "/api/v2/task/abc123def/comment") {
		t.Errorf("unexpected path: %s", log.Path)
	}

	var reqBody map[string]interface{}
	_ = json.Unmarshal([]byte(log.Body), &reqBody)
	if reqBody["comment_text"] != "Great progress on this task!" {
		t.Errorf("expected comment_text, got %v", reqBody["comment_text"])
	}
	if v, ok := reqBody["assignee"].(float64); !ok || v != 12345678 {
		t.Errorf("expected assignee 12345678, got %v", reqBody["assignee"])
	}
	if reqBody["notify_all"] != true {
		t.Errorf("expected notify_all true, got %v", reqBody["notify_all"])
	}

	mustContainJSON(t, out, "id", "457898234")
	mustContainJSON(t, out, "hist_id", "abc123hist")
}

// --- Comment List ---

func TestCommentList(t *testing.T) {
	server, log := newMockServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"comments": []map[string]interface{}{
				{
					"id":           "90901010101",
					"comment_text": "Initial review complete",
					"user":         map[string]interface{}{"id": 12345678, "username": "alice", "email": "alice@example.com"},
					"date":         "1676400000000",
					"comment":      []map[string]interface{}{{"text": "Initial review complete"}},
				},
				{
					"id":           "90901010102",
					"comment_text": "LGTM, merging",
					"user":         map[string]interface{}{"id": 87654321, "username": "bob", "email": "bob@example.com"},
					"date":         "1676500000000",
					"comment":      []map[string]interface{}{{"text": "LGTM, merging"}},
				},
			},
		})
	})

	out, err := runCommand(t, server.URL, "comment", "list", "--task", "abc123def")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if log.Method != "GET" {
		t.Errorf("expected GET, got %s", log.Method)
	}
	mustContainJSON(t, out, "comments[0].comment_text", "Initial review complete")
	mustContainJSON(t, out, "comments[1].comment_text", "LGTM, merging")
}

// --- Custom Field List ---

func TestCustomFieldList(t *testing.T) {
	server, log := newMockServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"fields": []map[string]interface{}{
				{
					"id":               "cf_uuid_sprint",
					"name":             "Sprint",
					"type":             "drop_down",
					"type_config":      map[string]interface{}{"options": []map[string]interface{}{{"id": "opt1", "name": "Sprint 42"}}},
					"date_created":     "1670000000000",
					"hide_from_guests": false,
					"required":         true,
				},
				{
					"id":               "cf_uuid_points",
					"name":             "Story Points",
					"type":             "number",
					"date_created":     "1670000000000",
					"hide_from_guests": false,
					"required":         false,
				},
			},
		})
	})

	out, err := runCommand(t, server.URL, "custom-field", "list", "--list", "901100200300")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if log.Method != "GET" {
		t.Errorf("expected GET, got %s", log.Method)
	}
	if !strings.Contains(log.Path, "/api/v2/list/901100200300/field") {
		t.Errorf("unexpected path: %s", log.Path)
	}
	mustContainJSON(t, out, "fields[0].name", "Sprint")
	mustContainJSON(t, out, "fields[1].name", "Story Points")
}

// --- Custom Field Set ---

func TestCustomFieldSet(t *testing.T) {
	server, log := newMockServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	out, err := runCommand(t, server.URL, "custom-field", "set",
		"--task", "abc123def",
		"--field", "cf_uuid_sprint",
		"--value", `"Sprint 43"`,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if log.Method != "POST" {
		t.Errorf("expected POST, got %s", log.Method)
	}
	if !strings.Contains(log.Path, "/api/v2/task/abc123def/field/cf_uuid_sprint") {
		t.Errorf("unexpected path: %s", log.Path)
	}

	var reqBody map[string]interface{}
	_ = json.Unmarshal([]byte(log.Body), &reqBody)
	if reqBody["value"] != "Sprint 43" {
		t.Errorf("expected value 'Sprint 43', got %v", reqBody["value"])
	}

	mustContainJSON(t, out, "status", "ok")
}

// --- Doc List (tests json.Number for date_created) ---

func TestDocList(t *testing.T) {
	server, log := newMockServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// Use numeric date_created to test json.Number handling
		_, _ = w.Write([]byte(`{"docs":[{"id":"doc_abc123","name":"API Design Doc","workspace_id":"12345678","date_created":1676000000000,"deleted":false,"visibility":"public"},{"id":"doc_def456","name":"Architecture Notes","workspace_id":"12345678","date_created":1676100000000,"deleted":false,"visibility":"private"}]}`))
	})

	out, err := runCommand(t, server.URL, "doc", "list", "--workspace", "12345678")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if log.Method != "GET" {
		t.Errorf("expected GET, got %s", log.Method)
	}
	if !strings.Contains(log.Path, "/api/v3/workspaces/12345678/docs") {
		t.Errorf("unexpected path: %s", log.Path)
	}
	mustContainJSON(t, out, "docs[0].name", "API Design Doc")
	mustContainJSON(t, out, "docs[1].name", "Architecture Notes")
	// Verify numeric date_created survives round-trip
	mustContainJSON(t, out, "date_created", "1676000000000")
}

// --- Time Entry List ---

func TestTimeEntryList(t *testing.T) {
	server, log := newMockServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"data": []map[string]interface{}{
				{
					"id":          "3600000123456",
					"task":        map[string]interface{}{"id": "abc123def", "name": "Implement auth flow"},
					"wid":         "12345678",
					"user":        map[string]interface{}{"id": 12345678, "username": "alice"},
					"billable":    true,
					"start":       "1676000000000",
					"end":         "1676003600000",
					"duration":    "3600000",
					"description": "Working on OAuth2 flow",
					"tags":        []map[string]interface{}{{"name": "development"}},
					"source":      "clickup",
					"at":          "1676003600000",
				},
			},
		})
	})

	out, err := runCommand(t, server.URL, "time-entry", "list", "--workspace", "12345678")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if log.Method != "GET" {
		t.Errorf("expected GET, got %s", log.Method)
	}
	if !strings.Contains(log.Path, "/api/v2/team/12345678/time_entries") {
		t.Errorf("unexpected path: %s", log.Path)
	}
	mustContainJSON(t, out, "data[0].id", "3600000123456")
	mustContainJSON(t, out, "data[0].description", "Working on OAuth2 flow")
}

// --- Task Search ---

func TestTaskSearch(t *testing.T) {
	server, log := newMockServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"tasks": []map[string]interface{}{
				{
					"id":           "abc123def",
					"name":         "Implement auth flow",
					"status":       map[string]interface{}{"status": "in progress", "color": "#4194f6", "type": "custom"},
					"orderindex":   "1.00000",
					"date_created": "1676000000000",
					"creator":      map[string]interface{}{"id": 12345678, "username": "alice"},
					"assignees":    []map[string]interface{}{},
					"tags":         []map[string]interface{}{},
					"parent":       nil,
					"priority":     nil,
					"url":          "https://app.clickup.com/t/abc123def",
					"list":         map[string]interface{}{"id": "901100200300", "name": "Sprint Backlog"},
					"folder":       map[string]interface{}{"id": "800100200300", "name": "Product"},
					"space":        map[string]interface{}{"id": "98765432"},
				},
			},
		})
	})

	out, err := runCommand(t, server.URL, "task", "search",
		"--workspace", "12345678",
		"--status", "in progress",
		"--include-closed",
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if log.Method != "GET" {
		t.Errorf("expected GET, got %s", log.Method)
	}
	if !strings.Contains(log.Path, "/api/v2/team/12345678/task") {
		t.Errorf("unexpected path: %s", log.Path)
	}
	if !strings.Contains(log.Query, "statuses") {
		t.Errorf("expected statuses in query, got: %s", log.Query)
	}
	if !strings.Contains(log.Query, "include_closed=true") {
		t.Errorf("expected include_closed=true in query, got: %s", log.Query)
	}
	mustContainJSON(t, out, "tasks[0].id", "abc123def")
}

// --- Webhook Create ---

func TestWebhookCreate(t *testing.T) {
	server, log := newMockServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"id": "wh_abc123def456",
			"webhook": map[string]interface{}{
				"id":        "wh_abc123def456",
				"userid":    12345678,
				"team_id":   90901234567,
				"endpoint":  "https://example.com/webhook",
				"client_id": "client_xyz",
				"events":    []string{"taskCreated", "taskUpdated"},
				"health":    map[string]interface{}{"status": "active", "fail_count": 0},
				"secret":    "whsec_randomsecret123",
			},
		})
	})

	out, err := runCommand(t, server.URL, "webhook", "create",
		"--workspace", "12345678",
		"--endpoint", "https://example.com/webhook",
		"--events", "taskCreated,taskUpdated",
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if log.Method != "POST" {
		t.Errorf("expected POST, got %s", log.Method)
	}
	if !strings.Contains(log.Path, "/api/v2/team/12345678/webhook") {
		t.Errorf("unexpected path: %s", log.Path)
	}

	var reqBody map[string]interface{}
	_ = json.Unmarshal([]byte(log.Body), &reqBody)
	if reqBody["endpoint"] != "https://example.com/webhook" {
		t.Errorf("expected endpoint, got %v", reqBody["endpoint"])
	}

	mustContainJSON(t, out, "id", "wh_abc123def456")
	mustContainJSON(t, out, "secret", "whsec_randomsecret123")
}

// --- Error Responses ---

func TestErrorUnauthorized(t *testing.T) {
	server, _ := newMockServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(401)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"err":   "Token invalid",
			"ECODE": "OAUTH_025",
		})
	})

	_, err := runCommand(t, server.URL, "workspace", "list")
	if err == nil {
		t.Fatal("expected error for 401 response")
	}
}

func TestErrorNotFound(t *testing.T) {
	server, _ := newMockServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"err":   "Task not found",
			"ECODE": "ITEM_015",
		})
	})

	_, err := runCommand(t, server.URL, "task", "get", "--id", "nonexistent")
	if err == nil {
		t.Fatal("expected error for 404 response")
	}
}

func TestErrorRateLimit(t *testing.T) {
	server, _ := newMockServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(429)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"err":   "Rate limit exceeded",
			"ECODE": "RATE_001",
		})
	})

	_, err := runCommand(t, server.URL, "workspace", "list")
	if err == nil {
		t.Fatal("expected error for 429 response")
	}
}

// --- Empty Lists ---

func TestWorkspaceListEmpty(t *testing.T) {
	server, _ := newMockServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"teams": []interface{}{}})
	})

	out, err := runCommand(t, server.URL, "workspace", "list")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "teams") {
		t.Errorf("expected teams key in output: %s", out)
	}
}

func TestTaskListEmpty(t *testing.T) {
	server, _ := newMockServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"tasks": []interface{}{}})
	})

	out, err := runCommand(t, server.URL, "task", "list", "--list", "901100200300")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "tasks") {
		t.Errorf("expected tasks key in output: %s", out)
	}
}

// --- Authorization Header ---

func TestAuthorizationHeader(t *testing.T) {
	server, log := newMockServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"teams": []interface{}{}})
	})

	_, err := runCommand(t, server.URL, "workspace", "list")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if log.Headers.Get("Authorization") != "test-token" {
		t.Errorf("expected Authorization header 'test-token', got %q", log.Headers.Get("Authorization"))
	}
}

// --- Task List with Query Params ---

func TestTaskListWithFilters(t *testing.T) {
	server, log := newMockServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"tasks": []interface{}{}})
	})

	_, err := runCommand(t, server.URL, "task", "list",
		"--list", "901100200300",
		"--status", "in progress",
		"--include-closed",
		"--subtasks",
		"--order-by", "due_date",
		"--page", "2",
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(log.Query, "statuses") {
		t.Errorf("expected statuses in query: %s", log.Query)
	}
	if !strings.Contains(log.Query, "include_closed=true") {
		t.Errorf("expected include_closed: %s", log.Query)
	}
	if !strings.Contains(log.Query, "subtasks=true") {
		t.Errorf("expected subtasks: %s", log.Query)
	}
	if !strings.Contains(log.Query, "order_by=due_date") {
		t.Errorf("expected order_by: %s", log.Query)
	}
	if !strings.Contains(log.Query, "page=2") {
		t.Errorf("expected page=2: %s", log.Query)
	}
}

// --- Valid JSON output ---

func TestOutputIsValidJSON(t *testing.T) {
	tests := []struct {
		name string
		args []string
		resp string
	}{
		{
			name: "workspace list",
			args: []string{"workspace", "list"},
			resp: `{"teams":[{"id":"123","name":"Test","color":"","avatar":"","members":[]}]}`,
		},
		{
			name: "task get",
			args: []string{"task", "get", "--id", "abc"},
			resp: `{"id":"abc","name":"Test","status":{"status":"open","color":"#000","type":"open"},"orderindex":"0","date_created":"0","creator":{"id":0,"username":""},"assignees":[],"tags":[],"parent":null,"priority":null,"url":"","list":{"id":"","name":""},"folder":{"id":"","name":""},"space":{"id":""}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, _ := newMockServer(t, func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(tt.resp))
			})

			out, err := runCommand(t, server.URL, tt.args...)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !json.Valid([]byte(strings.TrimSpace(out))) {
				t.Errorf("output is not valid JSON:\n%s", out)
			}
		})
	}
}
