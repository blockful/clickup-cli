//go:build integration

package api

import (
	"os"
	"testing"
)

// Integration tests require CLICKUP_TOKEN environment variable.
// Run with: go test -tags=integration ./internal/api/

func getIntegrationClient(t *testing.T) *Client {
	t.Helper()
	token := os.Getenv("CLICKUP_TOKEN")
	if token == "" {
		t.Skip("CLICKUP_TOKEN not set, skipping integration test")
	}
	return NewClient(token)
}

func TestIntegration_GetUser(t *testing.T) {
	client := getIntegrationClient(t)
	resp, err := client.GetUser()
	if err != nil {
		t.Fatalf("GetUser failed: %v", err)
	}
	if resp.User.Username == "" {
		t.Error("expected non-empty username")
	}
	t.Logf("Authenticated as: %s (%s)", resp.User.Username, resp.User.Email)
}

func TestIntegration_ListWorkspaces(t *testing.T) {
	client := getIntegrationClient(t)
	resp, err := client.ListWorkspaces()
	if err != nil {
		t.Fatalf("ListWorkspaces failed: %v", err)
	}
	if len(resp.Teams) == 0 {
		t.Error("expected at least one workspace")
	}
	for _, team := range resp.Teams {
		t.Logf("Workspace: %s (ID: %s)", team.Name, team.ID)
	}
}
