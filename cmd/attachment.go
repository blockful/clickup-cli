package cmd

import (
	"context"

	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
)

var attachmentCmd = &cobra.Command{
	Use:   "attachment",
	Short: "Manage attachments",
}

var attachmentCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a task attachment",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		taskID, _ := cmd.Flags().GetString("task-id")
		filePath, _ := cmd.Flags().GetString("file")

		if taskID == "" || filePath == "" {
			output.PrintError("VALIDATION_ERROR", "--task-id and --file are required")
			return &exitError{code: 1}
		}

		resp, err := client.CreateTaskAttachment(ctx, taskID, filePath, getTaskScopedOpts(cmd))
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(attachmentCmd)
	attachmentCmd.AddCommand(attachmentCreateCmd)

	attachmentCreateCmd.Flags().String("task-id", "", "Task ID (required)")
	attachmentCreateCmd.Flags().String("file", "", "Path to file (required)")
	addTaskScopedFlags(attachmentCreateCmd)
}
