package cmd

import (
	"context"
	"strings"

	"github.com/blockful/clickup-cli/internal/api"
	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
)

var legacyCmd = &cobra.Command{
	Use:   "legacy",
	Short: "Task-level time tracking (legacy endpoints)",
}

var legacyListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get tracked time for a task",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		taskID, _ := cmd.Flags().GetString("task-id")
		if taskID == "" {
			output.PrintError("VALIDATION_ERROR", "--task-id is required")
			return &exitError{code: 1}
		}
		subcategoryID, _ := cmd.Flags().GetString("subcategory-id")
		resp, err := client.GetLegacyTrackedTime(ctx, taskID, subcategoryID, getTaskScopedOpts(cmd))
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var legacyCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Track time on a task",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		taskID, _ := cmd.Flags().GetString("task-id")
		if taskID == "" {
			output.PrintError("VALIDATION_ERROR", "--task-id is required")
			return &exitError{code: 1}
		}
		time_, _ := cmd.Flags().GetInt64("time")
		if time_ == 0 {
			output.PrintError("VALIDATION_ERROR", "--time is required")
			return &exitError{code: 1}
		}
		start, _ := cmd.Flags().GetInt64("start")
		end, _ := cmd.Flags().GetInt64("end")
		tagsStr, _ := cmd.Flags().GetString("tags")

		req := &api.LegacyTrackTimeRequest{
			Time:  time_,
			Start: start,
			End:   end,
		}
		if tagsStr != "" {
			for _, t := range strings.Split(tagsStr, ",") {
				req.Tags = append(req.Tags, api.Tag{Name: strings.TrimSpace(t)})
			}
		}

		resp, err := client.TrackLegacyTime(ctx, taskID, req, getTaskScopedOpts(cmd))
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var legacyUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Edit tracked time on a task",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		taskID, _ := cmd.Flags().GetString("task-id")
		intervalID, _ := cmd.Flags().GetString("interval-id")
		if taskID == "" || intervalID == "" {
			output.PrintError("VALIDATION_ERROR", "--task-id and --interval-id are required")
			return &exitError{code: 1}
		}

		req := &api.LegacyEditTimeRequest{}
		if cmd.Flags().Changed("time") {
			req.Time, _ = cmd.Flags().GetInt64("time")
		}
		if cmd.Flags().Changed("start") {
			req.Start, _ = cmd.Flags().GetInt64("start")
		}
		if cmd.Flags().Changed("end") {
			req.End, _ = cmd.Flags().GetInt64("end")
		}
		if cmd.Flags().Changed("tag-action") {
			req.TagAction, _ = cmd.Flags().GetString("tag-action")
		}
		if cmd.Flags().Changed("tags") {
			tagsStr, _ := cmd.Flags().GetString("tags")
			for _, t := range strings.Split(tagsStr, ",") {
				req.Tags = append(req.Tags, api.Tag{Name: strings.TrimSpace(t)})
			}
		}

		if err := client.EditLegacyTime(ctx, taskID, intervalID, req, getTaskScopedOpts(cmd)); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

var legacyDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete tracked time on a task",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		taskID, _ := cmd.Flags().GetString("task-id")
		intervalID, _ := cmd.Flags().GetString("interval-id")
		if taskID == "" || intervalID == "" {
			output.PrintError("VALIDATION_ERROR", "--task-id and --interval-id are required")
			return &exitError{code: 1}
		}
		if err := client.DeleteLegacyTime(ctx, taskID, intervalID, getTaskScopedOpts(cmd)); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

func init() {
	timeEntryCmd.AddCommand(legacyCmd)
	legacyCmd.AddCommand(legacyListCmd, legacyCreateCmd, legacyUpdateCmd, legacyDeleteCmd)

	legacyListCmd.Flags().String("task-id", "", "Task ID (required)")
	legacyListCmd.Flags().String("subcategory-id", "", "Subcategory ID filter")
	addTaskScopedFlags(legacyListCmd)

	legacyCreateCmd.Flags().String("task-id", "", "Task ID (required)")
	legacyCreateCmd.Flags().Int64("time", 0, "Time in milliseconds (required)")
	legacyCreateCmd.Flags().Int64("start", 0, "Start date (unix ms)")
	legacyCreateCmd.Flags().Int64("end", 0, "End date (unix ms)")
	legacyCreateCmd.Flags().String("tags", "", "Comma-separated tag names")
	addTaskScopedFlags(legacyCreateCmd)

	legacyUpdateCmd.Flags().String("task-id", "", "Task ID (required)")
	legacyUpdateCmd.Flags().String("interval-id", "", "Interval ID (required)")
	legacyUpdateCmd.Flags().Int64("time", 0, "Time in milliseconds")
	legacyUpdateCmd.Flags().Int64("start", 0, "Start date (unix ms)")
	legacyUpdateCmd.Flags().Int64("end", 0, "End date (unix ms)")
	legacyUpdateCmd.Flags().String("tag-action", "", "Tag action (add/replace)")
	legacyUpdateCmd.Flags().String("tags", "", "Comma-separated tag names")
	addTaskScopedFlags(legacyUpdateCmd)

	legacyDeleteCmd.Flags().String("task-id", "", "Task ID (required)")
	legacyDeleteCmd.Flags().String("interval-id", "", "Interval ID (required)")
	addTaskScopedFlags(legacyDeleteCmd)
}
