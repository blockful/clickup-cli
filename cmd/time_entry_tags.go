package cmd

import (
	"context"
	"strings"

	"github.com/blockful/clickup-cli/internal/api"
	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
)

var timeEntryTagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Manage time entry tags",
}

var timeEntryHistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "Get time entry history",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		resp, err := client.GetTimeEntryHistory(ctx, wid, id)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var timeEntryTagAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add tags to time entries",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		idsStr, _ := cmd.Flags().GetString("time-entry-ids")
		tagsStr, _ := cmd.Flags().GetString("tags")
		if idsStr == "" || tagsStr == "" {
			output.PrintError("VALIDATION_ERROR", "--time-entry-ids and --tags are required")
			return &exitError{code: 1}
		}
		var tags []api.Tag
		for _, t := range strings.Split(tagsStr, ",") {
			tags = append(tags, api.Tag{Name: strings.TrimSpace(t)})
		}
		req := &api.AddTagsToTimeEntriesRequest{
			TimeEntryIDs: strings.Split(idsStr, ","),
			Tags:         tags,
		}
		if err := client.AddTagsToTimeEntries(ctx, wid, req); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

var timeEntryTagRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove tags from time entries",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		idsStr, _ := cmd.Flags().GetString("time-entry-ids")
		tagsStr, _ := cmd.Flags().GetString("tags")
		if idsStr == "" || tagsStr == "" {
			output.PrintError("VALIDATION_ERROR", "--time-entry-ids and --tags are required")
			return &exitError{code: 1}
		}
		var tags []api.Tag
		for _, t := range strings.Split(tagsStr, ",") {
			tags = append(tags, api.Tag{Name: strings.TrimSpace(t)})
		}
		req := &api.RemoveTagsFromTimeEntriesRequest{
			TimeEntryIDs: strings.Split(idsStr, ","),
			Tags:         tags,
		}
		if err := client.RemoveTagsFromTimeEntries(ctx, wid, req); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

var timeEntryTagUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Change tag names on time entries",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		name, _ := cmd.Flags().GetString("name")
		newName, _ := cmd.Flags().GetString("new-name")
		if name == "" || newName == "" {
			output.PrintError("VALIDATION_ERROR", "--name and --new-name are required")
			return &exitError{code: 1}
		}
		tagBg, _ := cmd.Flags().GetString("tag-bg")
		tagFg, _ := cmd.Flags().GetString("tag-fg")
		req := &api.ChangeTagNameRequest{
			Name:    name,
			NewName: newName,
			TagBg:   tagBg,
			TagFg:   tagFg,
		}
		if err := client.ChangeTagNames(ctx, wid, req); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

func init() {
	timeEntryCmd.AddCommand(timeEntryHistoryCmd, timeEntryTagCmd)
	timeEntryTagCmd.AddCommand(timeEntryTagAddCmd, timeEntryTagRemoveCmd, timeEntryTagUpdateCmd)

	timeEntryHistoryCmd.Flags().String("id", "", "Time entry ID (required)")

	timeEntryTagAddCmd.Flags().String("time-entry-ids", "", "Comma-separated time entry IDs (required)")
	timeEntryTagAddCmd.Flags().String("tags", "", "Comma-separated tag names (required)")

	timeEntryTagRemoveCmd.Flags().String("time-entry-ids", "", "Comma-separated time entry IDs (required)")
	timeEntryTagRemoveCmd.Flags().String("tags", "", "Comma-separated tag names (required)")

	timeEntryTagUpdateCmd.Flags().String("name", "", "Current tag name (required)")
	timeEntryTagUpdateCmd.Flags().String("new-name", "", "New tag name (required)")
	timeEntryTagUpdateCmd.Flags().String("tag-bg", "", "Tag background color")
	timeEntryTagUpdateCmd.Flags().String("tag-fg", "", "Tag foreground color")
}
