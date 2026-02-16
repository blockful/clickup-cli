package cmd

import (
	"context"

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
	Short: "List comments on a task, list, or view",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		taskID, _ := cmd.Flags().GetString("task")
		listID, _ := cmd.Flags().GetString("list")
		viewID, _ := cmd.Flags().GetString("view-id")
		if taskID == "" && listID == "" && viewID == "" {
			output.PrintError("VALIDATION_ERROR", "--task, --list, or --view-id is required")
			return &exitError{code: 1}
		}
		startID, _ := cmd.Flags().GetString("start-id")
		if viewID != "" {
			resp, err := client.ListViewComments(ctx, viewID, startID)
			if err != nil {
				return handleError(err)
			}
			output.JSON(resp)
			return nil
		}
		if listID != "" {
			resp, err := client.ListListComments(ctx, listID, startID)
			if err != nil {
				return handleError(err)
			}
			output.JSON(resp)
			return nil
		}
		resp, err := client.ListComments(ctx, taskID, startID, getTaskScopedOpts(cmd))
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var commentCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Add a comment to a task, list, or view",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		taskID, _ := cmd.Flags().GetString("task")
		listID, _ := cmd.Flags().GetString("list")
		viewID, _ := cmd.Flags().GetString("view-id")
		if taskID == "" && listID == "" && viewID == "" {
			output.PrintError("VALIDATION_ERROR", "--task, --list, or --view-id is required")
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
		if cmd.Flags().Changed("group-assignee") {
			v, _ := cmd.Flags().GetInt("group-assignee")
			req.GroupAssignee = api.IntPtr(v)
		}
		req.NotifyAll, _ = cmd.Flags().GetBool("notify-all")

		if viewID != "" {
			resp, err := client.CreateViewComment(ctx, viewID, req)
			if err != nil {
				return handleError(err)
			}
			output.JSON(resp)
			return nil
		}
		if listID != "" {
			resp, err := client.CreateListComment(ctx, listID, req)
			if err != nil {
				return handleError(err)
			}
			output.JSON(resp)
			return nil
		}
		resp, err := client.CreateComment(ctx, taskID, req, getTaskScopedOpts(cmd))
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
		ctx := context.Background()
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
		if cmd.Flags().Changed("group-assignee") {
			v, _ := cmd.Flags().GetInt("group-assignee")
			req.GroupAssignee = api.IntPtr(v)
		}
		if cmd.Flags().Changed("resolved") {
			v, _ := cmd.Flags().GetBool("resolved")
			req.Resolved = api.BoolPtr(v)
		}
		if err := client.UpdateComment(ctx, id, req); err != nil {
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
		ctx := context.Background()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		if err := client.DeleteComment(ctx, id); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"message": "comment deleted", "id": id})
		return nil
	},
}

var commentReplyCmd = &cobra.Command{
	Use:   "reply",
	Short: "Manage threaded comment replies",
}

var commentReplyListCmd = &cobra.Command{
	Use:   "list",
	Short: "List threaded comments on a comment",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		commentID, _ := cmd.Flags().GetString("comment-id")
		if commentID == "" {
			output.PrintError("VALIDATION_ERROR", "--comment-id is required")
			return &exitError{code: 1}
		}
		resp, err := client.ListThreadedComments(ctx, commentID)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var commentReplyCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a threaded comment reply",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		commentID, _ := cmd.Flags().GetString("comment-id")
		if commentID == "" {
			output.PrintError("VALIDATION_ERROR", "--comment-id is required")
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
		resp, err := client.CreateThreadedComment(ctx, commentID, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

func init() {
	commentListCmd.Flags().String("task", "", "Task ID")
	commentListCmd.Flags().String("list", "", "List ID")
	commentListCmd.Flags().String("view-id", "", "View ID (chat view comments)")
	commentListCmd.Flags().String("start-id", "", "Start comment ID for pagination")
	addTaskScopedFlags(commentListCmd)

	commentCreateCmd.Flags().String("task", "", "Task ID")
	commentCreateCmd.Flags().String("list", "", "List ID")
	commentCreateCmd.Flags().String("view-id", "", "View ID (chat view comment)")
	commentCreateCmd.Flags().String("text", "", "Comment text")
	commentCreateCmd.Flags().Int("assignee", 0, "Assignee user ID")
	commentCreateCmd.Flags().Int("group-assignee", 0, "Group assignee ID")
	commentCreateCmd.Flags().Bool("notify-all", false, "Notify all")
	addTaskScopedFlags(commentCreateCmd)

	commentUpdateCmd.Flags().String("id", "", "Comment ID")
	commentUpdateCmd.Flags().String("text", "", "Comment text")
	commentUpdateCmd.Flags().Int("assignee", 0, "Assignee user ID")
	commentUpdateCmd.Flags().Int("group-assignee", 0, "Group assignee ID")
	commentUpdateCmd.Flags().Bool("resolved", false, "Mark as resolved")

	commentDeleteCmd.Flags().String("id", "", "Comment ID")

	commentReplyListCmd.Flags().String("comment-id", "", "Comment ID")
	commentReplyCreateCmd.Flags().String("comment-id", "", "Comment ID")
	commentReplyCreateCmd.Flags().String("text", "", "Comment text")
	commentReplyCreateCmd.Flags().Int("assignee", 0, "Assignee user ID")
	commentReplyCreateCmd.Flags().Bool("notify-all", false, "Notify all")

	commentReplyCmd.AddCommand(commentReplyListCmd)
	commentReplyCmd.AddCommand(commentReplyCreateCmd)

	commentCmd.AddCommand(commentListCmd)
	commentCmd.AddCommand(commentCreateCmd)
	commentCmd.AddCommand(commentUpdateCmd)
	commentCmd.AddCommand(commentDeleteCmd)
	commentCmd.AddCommand(commentReplyCmd)
	rootCmd.AddCommand(commentCmd)
}
