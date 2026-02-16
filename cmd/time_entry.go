package cmd

import (
	"context"

	"github.com/blockful/clickup-cli/internal/api"
	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
)

var timeEntryCmd = &cobra.Command{
	Use:   "time-entry",
	Short: "Manage time tracking entries",
}

var timeEntryListCmd = &cobra.Command{
	Use:   "list",
	Short: "List time entries",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)

		opts := &api.ListTimeEntriesOptions{}
		opts.StartDate, _ = cmd.Flags().GetString("start-date")
		opts.EndDate, _ = cmd.Flags().GetString("end-date")
		opts.Assignee, _ = cmd.Flags().GetString("assignee")
		opts.SpaceID, _ = cmd.Flags().GetString("space")
		opts.FolderID, _ = cmd.Flags().GetString("folder")
		opts.ListID, _ = cmd.Flags().GetString("list")
		opts.TaskID, _ = cmd.Flags().GetString("task")
		opts.IncludeTaskTags, _ = cmd.Flags().GetBool("include-task-tags")
		opts.IncludeLocationNames, _ = cmd.Flags().GetBool("include-location-names")

		resp, err := client.GetTimeEntries(ctx, wid, opts)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var timeEntryGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a time entry",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		resp, err := client.GetTimeEntry(ctx, wid, id)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var timeEntryCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a time entry",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)

		start, _ := cmd.Flags().GetInt64("start")
		duration, _ := cmd.Flags().GetInt64("duration")
		description, _ := cmd.Flags().GetString("description")
		tid, _ := cmd.Flags().GetString("task")
		billable, _ := cmd.Flags().GetBool("billable")

		if start == 0 || duration == 0 {
			output.PrintError("VALIDATION_ERROR", "--start and --duration are required")
			return &exitError{code: 1}
		}

		req := &api.CreateTimeEntryRequest{
			Start:       start,
			Duration:    duration,
			Description: description,
			Tid:         tid,
			Billable:    billable,
		}

		resp, err := client.CreateTimeEntry(ctx, wid, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var timeEntryUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a time entry",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}

		req := &api.UpdateTimeEntryRequest{}
		if cmd.Flags().Changed("description") {
			req.Description, _ = cmd.Flags().GetString("description")
		}
		if cmd.Flags().Changed("task") {
			req.Tid, _ = cmd.Flags().GetString("task")
		}
		if cmd.Flags().Changed("tag-action") {
			req.TagAction, _ = cmd.Flags().GetString("tag-action")
		}

		if err := client.UpdateTimeEntry(ctx, wid, id, req); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

var timeEntryDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a time entry",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		if err := client.DeleteTimeEntry(ctx, wid, id); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

var timeEntryStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a timer",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		tid, _ := cmd.Flags().GetString("task")
		description, _ := cmd.Flags().GetString("description")
		billable, _ := cmd.Flags().GetBool("billable")

		req := &api.StartTimerRequest{
			Tid:         tid,
			Description: description,
			Billable:    billable,
		}

		resp, err := client.StartTimer(ctx, wid, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var timeEntryStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the running timer",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		resp, err := client.StopTimer(ctx, wid)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var timeEntryCurrentCmd = &cobra.Command{
	Use:   "current",
	Short: "Get the running timer",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		assignee, _ := cmd.Flags().GetString("assignee")
		resp, err := client.GetRunningTimer(ctx, wid, assignee)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(timeEntryCmd)
	timeEntryCmd.AddCommand(timeEntryListCmd, timeEntryGetCmd, timeEntryCreateCmd, timeEntryUpdateCmd, timeEntryDeleteCmd, timeEntryStartCmd, timeEntryStopCmd, timeEntryCurrentCmd)

	timeEntryListCmd.Flags().String("start-date", "", "Start date (unix ms)")
	timeEntryListCmd.Flags().String("end-date", "", "End date (unix ms)")
	timeEntryListCmd.Flags().String("assignee", "", "Assignee user ID")
	timeEntryListCmd.Flags().String("space", "", "Space ID filter")
	timeEntryListCmd.Flags().String("folder", "", "Folder ID filter")
	timeEntryListCmd.Flags().String("list", "", "List ID filter")
	timeEntryListCmd.Flags().String("task", "", "Task ID filter")
	timeEntryListCmd.Flags().Bool("include-task-tags", false, "Include task tags")
	timeEntryListCmd.Flags().Bool("include-location-names", false, "Include location names")

	timeEntryGetCmd.Flags().String("id", "", "Time entry ID (required)")

	timeEntryCreateCmd.Flags().Int64("start", 0, "Start time (unix ms, required)")
	timeEntryCreateCmd.Flags().Int64("duration", 0, "Duration (ms, required)")
	timeEntryCreateCmd.Flags().String("description", "", "Description")
	timeEntryCreateCmd.Flags().String("task", "", "Task ID")
	timeEntryCreateCmd.Flags().Bool("billable", false, "Billable")

	timeEntryUpdateCmd.Flags().String("id", "", "Time entry ID (required)")
	timeEntryUpdateCmd.Flags().String("description", "", "Description")
	timeEntryUpdateCmd.Flags().String("task", "", "Task ID")
	timeEntryUpdateCmd.Flags().String("tag-action", "", "Tag action (add/replace)")

	timeEntryDeleteCmd.Flags().String("id", "", "Time entry ID (required)")

	timeEntryStartCmd.Flags().String("task", "", "Task ID")
	timeEntryStartCmd.Flags().String("description", "", "Description")
	timeEntryStartCmd.Flags().Bool("billable", false, "Billable")

	timeEntryCurrentCmd.Flags().String("assignee", "", "Assignee user ID")
}
