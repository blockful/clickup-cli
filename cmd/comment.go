package cmd

import (
	"github.com/blockful/clickup-cli/internal/api"
	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
)

var commentCmd = &cobra.Command{
	Use:   "comment",
	Short: "Manage comments",
}

var commentListCmd = &cobra.Command{
	Use:   "list",
	Short: "List comments on a task",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		taskID, _ := cmd.Flags().GetString("task")
		if taskID == "" {
			output.PrintError("VALIDATION_ERROR", "--task is required")
			return &exitError{code: 1}
		}
		resp, err := client.ListComments(taskID)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var commentCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Add a comment to a task",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		taskID, _ := cmd.Flags().GetString("task")
		if taskID == "" {
			output.PrintError("VALIDATION_ERROR", "--task is required")
			return &exitError{code: 1}
		}
		text, _ := cmd.Flags().GetString("text")
		if text == "" {
			output.PrintError("VALIDATION_ERROR", "--text is required")
			return &exitError{code: 1}
		}
		resp, err := client.CreateComment(taskID, &api.CreateCommentRequest{CommentText: text})
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

func init() {
	commentListCmd.Flags().String("task", "", "Task ID")
	commentCreateCmd.Flags().String("task", "", "Task ID")
	commentCreateCmd.Flags().String("text", "", "Comment text")

	commentCmd.AddCommand(commentListCmd)
	commentCmd.AddCommand(commentCreateCmd)
	rootCmd.AddCommand(commentCmd)
}
