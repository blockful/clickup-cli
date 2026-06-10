package cmd

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/blockful/clickup-cli/internal/api"
	"github.com/blockful/clickup-cli/internal/config"
	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
)

// defaultChangesWindow is used for --since last when no previous check is
// recorded for the workspace.
const defaultChangesWindow = 24 * time.Hour

// maxChangesPages caps task pagination as a safety limit (100 tasks/page).
const maxChangesPages = 50

type changesResponse struct {
	WorkspaceID string     `json:"workspace_id"`
	Since       int64      `json:"since"`
	Until       int64      `json:"until"`
	FirstRun    bool       `json:"first_run,omitempty"`
	TaskCount   int        `json:"task_count"`
	DocCount    int        `json:"doc_count"`
	Tasks       []api.Task `json:"tasks"`
	Docs        []api.Doc  `json:"docs"`
}

var changesCmd = &cobra.Command{
	Use:   "changes",
	Short: "List tasks and docs updated since a point in time",
	Long: `List everything updated in a workspace since a point in time.

Tasks are filtered server-side (date_updated_gt). Docs do not support a
server-side updated filter, so they are fetched and filtered client-side
by their date_updated field.

--since accepts:
  last                 changes since the previous 'changes --since last' run
                       (tracked per workspace in ~/` + config.StateFileName + `;
                       defaults to 24h on first run)
  a duration           e.g. 30m, 24h, 7d, 2w
  a date or timestamp  e.g. 2026-06-09 or RFC3339
  Unix milliseconds    e.g. 1765432100000`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		workspaceID := getWorkspaceID(cmd)

		now := time.Now().UnixMilli()
		sinceFlag, _ := cmd.Flags().GetString("since")
		trackLast := sinceFlag == "last"

		var since int64
		firstRun := false
		if trackLast {
			last, err := config.GetLastChangesCheck(workspaceID)
			if err != nil {
				output.PrintError("STATE_ERROR", err.Error())
				return &exitError{code: 1}
			}
			since = last
			if since == 0 {
				since = now - defaultChangesWindow.Milliseconds()
				firstRun = true
			}
		} else {
			var err error
			since, err = parseSince(sinceFlag, now)
			if err != nil {
				output.PrintError("VALIDATION_ERROR", err.Error())
				return &exitError{code: 1}
			}
		}

		opts := &api.SearchTasksOptions{
			DateUpdatedGt: since,
			IncludeClosed: true,
			Subtasks:      true,
			OrderBy:       "updated",
		}
		opts.SpaceIDs, _ = cmd.Flags().GetStringSlice("space-ids")
		opts.ListIDs, _ = cmd.Flags().GetStringSlice("list-ids")
		opts.FolderIDs, _ = cmd.Flags().GetStringSlice("folder-ids")

		tasks := []api.Task{}
		for page := 0; page < maxChangesPages; page++ {
			opts.Page = page
			resp, err := client.SearchTasks(ctx, workspaceID, opts)
			if err != nil {
				return handleError(err)
			}
			tasks = append(tasks, resp.Tasks...)
			if len(resp.Tasks) < 100 {
				break
			}
		}

		docs := []api.Doc{}
		if skipDocs, _ := cmd.Flags().GetBool("skip-docs"); !skipDocs {
			resp, err := client.SearchDocs(ctx, workspaceID)
			if err != nil {
				return handleError(err)
			}
			for _, doc := range resp.Docs {
				if docUpdatedAt(doc) > since {
					docs = append(docs, doc)
				}
			}
		}

		if noSave, _ := cmd.Flags().GetBool("no-save"); trackLast && !noSave {
			if err := config.SetLastChangesCheck(workspaceID, now); err != nil {
				output.PrintError("STATE_ERROR", err.Error())
				return &exitError{code: 1}
			}
		}

		output.JSON(changesResponse{
			WorkspaceID: workspaceID,
			Since:       since,
			Until:       now,
			FirstRun:    firstRun,
			TaskCount:   len(tasks),
			DocCount:    len(docs),
			Tasks:       tasks,
			Docs:        docs,
		})
		return nil
	},
}

// docUpdatedAt returns the doc's date_updated in Unix ms, falling back to
// date_created when the API omits date_updated.
func docUpdatedAt(doc api.Doc) int64 {
	if v, err := doc.DateUpdated.Int64(); err == nil && v > 0 {
		return v
	}
	if v, err := doc.DateCreated.Int64(); err == nil {
		return v
	}
	return 0
}

// parseSince converts a --since value into a Unix ms timestamp. It accepts
// raw Unix ms, durations (30m, 24h, 7d, 2w), dates (2006-01-02), and
// RFC3339 timestamps.
func parseSince(s string, now int64) (int64, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, fmt.Errorf("--since requires a value (e.g. last, 24h, 7d, 2026-06-09, or Unix ms)")
	}

	if ms, err := strconv.ParseInt(s, 10, 64); err == nil {
		return ms, nil
	}

	if d, err := parseDuration(s); err == nil {
		return now - d.Milliseconds(), nil
	}

	if t, err := time.Parse("2006-01-02", s); err == nil {
		return t.UnixMilli(), nil
	}
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		return t.UnixMilli(), nil
	}

	return 0, fmt.Errorf("invalid --since value %q: use last, a duration (24h, 7d), a date (2026-06-09), an RFC3339 timestamp, or Unix ms", s)
}

// parseDuration extends time.ParseDuration with d (days) and w (weeks) units.
func parseDuration(s string) (time.Duration, error) {
	if n, err := strconv.ParseFloat(strings.TrimSuffix(s, "d"), 64); err == nil && strings.HasSuffix(s, "d") {
		return time.Duration(n * 24 * float64(time.Hour)), nil
	}
	if n, err := strconv.ParseFloat(strings.TrimSuffix(s, "w"), 64); err == nil && strings.HasSuffix(s, "w") {
		return time.Duration(n * 7 * 24 * float64(time.Hour)), nil
	}
	return time.ParseDuration(s)
}

func init() {
	changesCmd.Flags().String("workspace", "", "Workspace/Team ID")
	changesCmd.Flags().String("since", "last", "Point in time: last, duration (24h, 7d), date, RFC3339, or Unix ms")
	changesCmd.Flags().Bool("skip-docs", false, "Skip checking docs for updates")
	changesCmd.Flags().Bool("no-save", false, "Don't record this check as the new 'last' timestamp")
	changesCmd.Flags().StringSlice("space-ids", nil, "Limit task changes to space IDs")
	changesCmd.Flags().StringSlice("folder-ids", nil, "Limit task changes to folder IDs")
	changesCmd.Flags().StringSlice("list-ids", nil, "Limit task changes to list IDs")
	rootCmd.AddCommand(changesCmd)
}
