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
	Short: "List comments on a task or list",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		taskID, _ := cmd.Flags().GetString("task")
		listID, _ := cmd.Flags().GetString("list")
		if taskID == "" && listID == "" {
			output.PrintError("VALIDATION_ERROR", "--task or --list is required")
			return &exitError{code: 1}
		}
		if listID != "" {
			resp, err := client.ListListComments(listID)
			if err != nil {
				return handleError(err)
			}
			output.JSON(resp)
			return nil
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
		listID, _ := cmd.Flags().GetString("list")
		if taskID == "" && listID == "" {
			output.PrintError("VALIDATION_ERROR", "--task or --list is required")
			return &exitError{code: 1}
		}
		text, _ := cmd.Flags().GetString("text")
		if text == "" {
			output.PrintError("VALIDATION_ERROR", "--text is required")
			return &exitError{code: 1}
		}
		req := &api.CreateCommentRequest{CommentText: text}
		if cmd.Flags().Changed("assignee") {
			v, _ := cmd.Flags().GetInt("assignee")
			req.Assignee = api.IntPtr(v)
		}
		req.NotifyAll, _ = cmd.Flags().GetBool("notify-all")

		if listID != "" {
			resp, err := client.CreateListComment(listID, req)
			if err != nil {
				return handleError(err)
			}
			output.JSON(resp)
			return nil
		}
		resp, err := client.CreateComment(taskID, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var commentUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a comment",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		text, _ := cmd.Flags().GetString("text")
		if text == "" {
			output.PrintError("VALIDATION_ERROR", "--text is required")
			return &exitError{code: 1}
		}
		req := &api.UpdateCommentRequest{CommentText: text}
		if cmd.Flags().Changed("assignee") {
			v, _ := cmd.Flags().GetInt("assignee")
			req.Assignee = api.IntPtr(v)
		}
		if cmd.Flags().Changed("resolved") {
			v, _ := cmd.Flags().GetBool("resolved")
			req.Resolved = api.BoolPtr(v)
		}
		if err := client.UpdateComment(id, req); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"message": "comment updated", "id": id})
		return nil
	},
}

var commentDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a comment",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		if err := client.DeleteComment(id); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"message": "comment deleted", "id": id})
		return nil
	},
}

func init() {
	commentListCmd.Flags().String("task", "", "Task ID")
	commentListCmd.Flags().String("list", "", "List ID")

	commentCreateCmd.Flags().String("task", "", "Task ID")
	commentCreateCmd.Flags().String("list", "", "List ID")
	commentCreateCmd.Flags().String("text", "", "Comment text")
	commentCreateCmd.Flags().Int("assignee", 0, "Assignee user ID")
	commentCreateCmd.Flags().Bool("notify-all", false, "Notify all")

	commentUpdateCmd.Flags().String("id", "", "Comment ID")
	commentUpdateCmd.Flags().String("text", "", "Comment text")
	commentUpdateCmd.Flags().Int("assignee", 0, "Assignee user ID")
	commentUpdateCmd.Flags().Bool("resolved", false, "Mark as resolved")

	commentDeleteCmd.Flags().String("id", "", "Comment ID")

	commentCmd.AddCommand(commentListCmd)
	commentCmd.AddCommand(commentCreateCmd)
	commentCmd.AddCommand(commentUpdateCmd)
	commentCmd.AddCommand(commentDeleteCmd)
	rootCmd.AddCommand(commentCmd)
}
