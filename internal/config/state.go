package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const StateFileName = ".clickup-cli-state.json"

// state holds CLI-managed bookkeeping that is not user configuration.
// It lives in a separate file from the config so writing it never
// touches the token.
type state struct {
	// LastChangesCheck maps workspace ID to the Unix ms timestamp of the
	// last successful `clickup changes --since last` run.
	LastChangesCheck map[string]int64 `json:"last_changes_check,omitempty"`
}

func StateFilePath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, StateFileName)
}

func readState() (*state, error) {
	s := &state{}
	data, err := os.ReadFile(StateFilePath())
	if err != nil {
		if os.IsNotExist(err) {
			return s, nil
		}
		return nil, fmt.Errorf("failed to read state file: %w", err)
	}
	if err := json.Unmarshal(data, s); err != nil {
		return nil, fmt.Errorf("failed to parse state file %s: %w", StateFilePath(), err)
	}
	return s, nil
}

// GetLastChangesCheck returns the timestamp (Unix ms) of the last recorded
// changes check for a workspace, or 0 if none is recorded.
func GetLastChangesCheck(workspaceID string) (int64, error) {
	s, err := readState()
	if err != nil {
		return 0, err
	}
	return s.LastChangesCheck[workspaceID], nil
}

// SetLastChangesCheck records the timestamp (Unix ms) of a changes check
// for a workspace.
func SetLastChangesCheck(workspaceID string, ts int64) error {
	s, err := readState()
	if err != nil {
		return err
	}
	if s.LastChangesCheck == nil {
		s.LastChangesCheck = map[string]int64{}
	}
	s.LastChangesCheck[workspaceID] = ts

	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal state: %w", err)
	}
	if err := os.WriteFile(StateFilePath(), data, 0o600); err != nil {
		return fmt.Errorf("failed to write state file: %w", err)
	}
	return nil
}
