package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/blockful/clickup-cli/internal/config"
)

// newChangesServer mocks the task search and docs endpoints.
func newChangesServer(t *testing.T, docsJSON string) (*httptest.Server, *requestLog) {
	t.Helper()
	log := &requestLog{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.mu.Lock()
		log.Method = r.Method
		log.Path = r.URL.Path
		log.Query = r.URL.RawQuery
		log.mu.Unlock()
		switch {
		case strings.Contains(r.URL.Path, "/v2/team/"):
			fmt.Fprint(w, `{"tasks":[{"id":"task1","name":"Updated Task"}]}`)
		case strings.Contains(r.URL.Path, "/v3/workspaces/") && strings.HasSuffix(r.URL.Path, "/docs"):
			fmt.Fprint(w, docsJSON)
		default:
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	t.Cleanup(server.Close)
	return server, log
}

func TestChangesCommand(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	docsJSON := `{"docs":[
		{"id":"doc-new","name":"Fresh","date_updated":9000000000000},
		{"id":"doc-old","name":"Stale","date_updated":1000},
		{"id":"doc-created-only","name":"NoUpdatedField","date_created":9000000000000}
	]}`
	server, _ := newChangesServer(t, docsJSON)

	out, err := runCommand(t, server.URL, "changes", "--since", "5000")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var resp struct {
		Since     int64 `json:"since"`
		TaskCount int   `json:"task_count"`
		DocCount  int   `json:"doc_count"`
		Tasks     []struct {
			ID string `json:"id"`
		} `json:"tasks"`
		Docs []struct {
			ID string `json:"id"`
		} `json:"docs"`
	}
	if err := json.Unmarshal([]byte(out), &resp); err != nil {
		t.Fatalf("invalid JSON output: %v\n%s", err, out)
	}
	if resp.Since != 5000 {
		t.Errorf("expected since=5000, got %d", resp.Since)
	}
	if resp.TaskCount != 1 || resp.Tasks[0].ID != "task1" {
		t.Errorf("expected 1 task task1, got %+v", resp.Tasks)
	}
	if resp.DocCount != 2 {
		t.Fatalf("expected 2 docs after filtering, got %d: %+v", resp.DocCount, resp.Docs)
	}
	if resp.Docs[0].ID != "doc-new" || resp.Docs[1].ID != "doc-created-only" {
		t.Errorf("unexpected docs after filtering: %+v", resp.Docs)
	}
}

func TestChangesCompactOutput(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	docsJSON := `{"docs":[{"id":"doc-new","name":"Fresh","date_updated":9000000000000,"creator":{"id":1},"visibility":"private"}]}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/v2/team/") {
			fmt.Fprint(w, `{"tasks":[{"id":"task1","name":"Updated Task","status":{"status":"in progress","color":"#fff","type":"custom"},"date_updated":"9000000000000","url":"https://app.clickup.com/t/task1","list":{"id":"list1","name":"Sprint"},"folder":{"id":"folder1","name":"Product"},"space":{"id":"space1"},"assignees":[{"id":1,"username":"Agent"}],"custom_fields":[{"id":"field1","name":"Huge field"}]}]}`)
			return
		}
		fmt.Fprint(w, docsJSON)
	}))
	t.Cleanup(server.Close)

	out, err := runCommand(t, server.URL, "changes", "--since", "5000", "--compact")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.Contains(out, "custom_fields") || strings.Contains(out, "assignees") || strings.Contains(out, "creator") {
		t.Fatalf("compact output should omit token-heavy fields, got: %s", out)
	}

	var resp struct {
		TaskCount    int `json:"task_count"`
		TaskReturned int `json:"task_returned"`
		DocCount     int `json:"doc_count"`
		DocReturned  int `json:"doc_returned"`
		Tasks        []struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			Status      string `json:"status"`
			URL         string `json:"url"`
			ListID      string `json:"list_id"`
			ListName    string `json:"list_name"`
			FolderName  string `json:"folder_name"`
			SpaceID     string `json:"space_id"`
			DateUpdated string `json:"date_updated"`
		} `json:"tasks"`
		Docs []struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			DateUpdated int64  `json:"date_updated"`
		} `json:"docs"`
	}
	if err := json.Unmarshal([]byte(out), &resp); err != nil {
		t.Fatalf("invalid JSON output: %v\n%s", err, out)
	}
	if resp.TaskCount != 1 || resp.TaskReturned != 1 || resp.DocCount != 1 || resp.DocReturned != 1 {
		t.Fatalf("unexpected counts: %+v", resp)
	}
	task := resp.Tasks[0]
	if task.ID != "task1" || task.Status != "in progress" || task.ListName != "Sprint" || task.FolderName != "Product" || task.SpaceID != "space1" || task.URL == "" || task.DateUpdated == "" {
		t.Errorf("unexpected compact task: %+v", task)
	}
	if resp.Docs[0].ID != "doc-new" || resp.Docs[0].Name != "Fresh" || resp.Docs[0].DateUpdated == 0 {
		t.Errorf("unexpected compact doc: %+v", resp.Docs[0])
	}
}

func TestChangesLimitCompactsReturnedItemsWithoutChangingCounts(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/v2/team/") {
			fmt.Fprint(w, `{"tasks":[{"id":"task1","name":"One"},{"id":"task2","name":"Two"},{"id":"task3","name":"Three"}]}`)
			return
		}
		fmt.Fprint(w, `{"docs":[{"id":"doc1","name":"One","date_updated":9000000000000},{"id":"doc2","name":"Two","date_updated":9000000000000}]}`)
	}))
	t.Cleanup(server.Close)

	out, err := runCommand(t, server.URL, "changes", "--since", "5000", "--compact", "--limit", "1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var resp struct {
		TaskCount    int                   `json:"task_count"`
		TaskReturned int                   `json:"task_returned"`
		DocCount     int                   `json:"doc_count"`
		DocReturned  int                   `json:"doc_returned"`
		Truncated    bool                  `json:"truncated"`
		Tasks        []struct{ ID string } `json:"tasks"`
		Docs         []struct{ ID string } `json:"docs"`
	}
	if err := json.Unmarshal([]byte(out), &resp); err != nil {
		t.Fatalf("invalid JSON output: %v\n%s", err, out)
	}
	if resp.TaskCount != 3 || resp.TaskReturned != 1 || resp.DocCount != 2 || resp.DocReturned != 1 || !resp.Truncated {
		t.Fatalf("expected full counts but limited returned items, got %+v", resp)
	}
	if resp.Tasks[0].ID != "task1" || resp.Docs[0].ID != "doc1" {
		t.Errorf("limit should preserve updated ordering, got tasks=%+v docs=%+v", resp.Tasks, resp.Docs)
	}
}

func TestChangesTaskQueryParams(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	var taskQuery string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/v2/team/") {
			taskQuery = r.URL.RawQuery
			fmt.Fprint(w, `{"tasks":[]}`)
			return
		}
		fmt.Fprint(w, `{"docs":[]}`)
	}))
	t.Cleanup(server.Close)

	_, err := runCommand(t, server.URL, "changes", "--since", "5000", "--space-ids", "sp1", "--list-ids", "l1,l2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for _, want := range []string{
		"date_updated_gt=5000",
		"include_closed=true",
		"subtasks=true",
		"order_by=updated",
		"space_ids%5B%5D=sp1",
		"list_ids%5B%5D=l1",
		"list_ids%5B%5D=l2",
	} {
		if !strings.Contains(taskQuery, want) {
			t.Errorf("expected task query to contain %q, got %q", want, taskQuery)
		}
	}
}

func TestChangesSkipDocs(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/docs") {
			t.Errorf("docs endpoint should not be called with --skip-docs")
		}
		fmt.Fprint(w, `{"tasks":[]}`)
	}))
	t.Cleanup(server.Close)

	out, err := runCommand(t, server.URL, "changes", "--since", "5000", "--skip-docs")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, `"doc_count": 0`) {
		t.Errorf("expected doc_count 0, got: %s", out)
	}
}

func TestChangesSinceLast(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	server, _ := newChangesServer(t, `{"docs":[]}`)

	// First run: no saved state → defaults to a 24h window and records now.
	before := time.Now().UnixMilli()
	out, err := runCommand(t, server.URL, "changes", "--since", "last")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var resp struct {
		Since    int64 `json:"since"`
		FirstRun bool  `json:"first_run"`
	}
	if err := json.Unmarshal([]byte(out), &resp); err != nil {
		t.Fatalf("invalid JSON output: %v\n%s", err, out)
	}
	if !resp.FirstRun {
		t.Errorf("expected first_run=true on first run")
	}
	if got, want := resp.Since, before-24*time.Hour.Milliseconds(); got < want-5000 || got > want+5000 {
		t.Errorf("expected since≈now-24h (%d), got %d", want, got)
	}

	saved, err := config.GetLastChangesCheck("12345678")
	if err != nil {
		t.Fatalf("failed reading state: %v", err)
	}
	if saved < before {
		t.Errorf("expected saved timestamp >= %d, got %d", before, saved)
	}

	// Second run: uses the saved timestamp.
	out, err = runCommand(t, server.URL, "changes", "--since", "last")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	resp.FirstRun = false // Unmarshal leaves absent (omitempty) fields untouched
	if err := json.Unmarshal([]byte(out), &resp); err != nil {
		t.Fatalf("invalid JSON output: %v\n%s", err, out)
	}
	if resp.FirstRun {
		t.Errorf("expected first_run=false on second run")
	}
	if resp.Since != saved {
		t.Errorf("expected since=%d (saved), got %d", saved, resp.Since)
	}
}

func TestChangesNoSave(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	server, _ := newChangesServer(t, `{"docs":[]}`)

	_, err := runCommand(t, server.URL, "changes", "--since", "last", "--no-save")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, err := os.Stat(filepath.Join(home, config.StateFileName)); !os.IsNotExist(err) {
		t.Errorf("expected no state file with --no-save, stat err: %v", err)
	}
}

func TestChangesPagination(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	pages := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/docs") {
			fmt.Fprint(w, `{"docs":[]}`)
			return
		}
		pages++
		if pages == 1 {
			tasks := make([]string, 100)
			for i := range tasks {
				tasks[i] = fmt.Sprintf(`{"id":"t%d"}`, i)
			}
			fmt.Fprintf(w, `{"tasks":[%s]}`, strings.Join(tasks, ","))
			return
		}
		fmt.Fprint(w, `{"tasks":[{"id":"t100"}]}`)
	}))
	t.Cleanup(server.Close)

	out, err := runCommand(t, server.URL, "changes", "--since", "5000")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pages != 2 {
		t.Errorf("expected 2 task pages fetched, got %d", pages)
	}
	if !strings.Contains(out, `"task_count": 101`) {
		t.Errorf("expected task_count 101, got: %s", out)
	}
}

func TestChangesInvalidSince(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	server, _ := newChangesServer(t, `{"docs":[]}`)

	_, err := runCommand(t, server.URL, "changes", "--since", "yesterday-ish")
	if err == nil {
		t.Fatal("expected error for invalid --since value")
	}
}

func TestParseSince(t *testing.T) {
	now := int64(10_000_000_000_000)
	cases := []struct {
		in      string
		want    int64
		wantErr bool
	}{
		{"1765432100000", 1765432100000, false},
		{"24h", now - 24*time.Hour.Milliseconds(), false},
		{"30m", now - 30*time.Minute.Milliseconds(), false},
		{"7d", now - 7*24*time.Hour.Milliseconds(), false},
		{"2w", now - 14*24*time.Hour.Milliseconds(), false},
		{"2026-06-09T12:00:00Z", time.Date(2026, 6, 9, 12, 0, 0, 0, time.UTC).UnixMilli(), false},
		{"", 0, true},
		{"banana", 0, true},
	}
	for _, c := range cases {
		got, err := parseSince(c.in, now)
		if c.wantErr {
			if err == nil {
				t.Errorf("parseSince(%q): expected error, got %d", c.in, got)
			}
			continue
		}
		if err != nil {
			t.Errorf("parseSince(%q): unexpected error: %v", c.in, err)
			continue
		}
		if got != c.want {
			t.Errorf("parseSince(%q) = %d, want %d", c.in, got, c.want)
		}
	}
	// Date-only values parse in local time; just check it round-trips to the right day.
	got, err := parseSince("2026-06-09", now)
	if err != nil {
		t.Fatalf("parseSince(2026-06-09): %v", err)
	}
	if d := time.UnixMilli(got).Format("2006-01-02"); d != "2026-06-09" {
		t.Errorf("parseSince(2026-06-09) = %d (%s), want same day", got, d)
	}
}
