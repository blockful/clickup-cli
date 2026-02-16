package cmd

import (
	"context"
	"github.com/blockful/clickup-cli/internal/api"
	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Manage lists",
}

var listListCmd = &cobra.Command{
	Use:   "list",
	Short: "List lists in a folder or space",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		folderID, _ := cmd.Flags().GetString("folder")
		spaceID, _ := cmd.Flags().GetString("space")
		if folderID == "" && spaceID == "" {
			output.PrintError("VALIDATION_ERROR", "--folder or --space is required")
			return &exitError{code: 1}
		}
		if spaceID != "" {
			resp, err := client.ListFolderlessLists(ctx, spaceID)
			if err != nil {
				return handleError(err)
			}
			output.JSON(resp)
			return nil
		}
		resp, err := client.ListLists(ctx, folderID)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var listGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a list by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		resp, err := client.GetList(ctx, id)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var listCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new list",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		folderID, _ := cmd.Flags().GetString("folder")
		spaceID, _ := cmd.Flags().GetString("space")
		if folderID == "" && spaceID == "" {
			output.PrintError("VALIDATION_ERROR", "--folder or --space is required")
			return &exitError{code: 1}
		}
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			output.PrintError("VALIDATION_ERROR", "--name is required")
			return &exitError{code: 1}
		}
		req := &api.CreateListRequest{Name: name}
		req.Content, _ = cmd.Flags().GetString("content")
		req.Status, _ = cmd.Flags().GetString("status")
		if cmd.Flags().Changed("due-date") {
			v, _ := cmd.Flags().GetInt64("due-date")
			req.DueDate = api.Int64Ptr(v)
		}
		if cmd.Flags().Changed("priority") {
			v, _ := cmd.Flags().GetInt("priority")
			req.Priority = api.IntPtr(v)
		}
		if cmd.Flags().Changed("assignee") {
			v, _ := cmd.Flags().GetInt("assignee")
			req.Assignee = api.IntPtr(v)
		}

		if spaceID != "" {
			resp, err := client.CreateFolderlessList(ctx, spaceID, req)
			if err != nil {
				return handleError(err)
			}
			output.JSON(resp)
			return nil
		}
		resp, err := client.CreateList(ctx, folderID, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var listUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a list",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		req := &api.UpdateListRequest{}
		if cmd.Flags().Changed("name") {
			req.Name, _ = cmd.Flags().GetString("name")
		}
		if cmd.Flags().Changed("content") {
			req.Content, _ = cmd.Flags().GetString("content")
		}
		if cmd.Flags().Changed("status") {
			req.Status, _ = cmd.Flags().GetString("status")
		}
		if cmd.Flags().Changed("due-date") {
			v, _ := cmd.Flags().GetInt64("due-date")
			req.DueDate = api.Int64Ptr(v)
		}
		if cmd.Flags().Changed("priority") {
			v, _ := cmd.Flags().GetInt("priority")
			req.Priority = api.IntPtr(v)
		}
		if cmd.Flags().Changed("assignee") {
			v, _ := cmd.Flags().GetInt("assignee")
			req.Assignee = api.IntPtr(v)
		}
		req.UnsetStatus, _ = cmd.Flags().GetBool("unset-status")

		resp, err := client.UpdateList(ctx, id, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var listDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a list",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		if err := client.DeleteList(ctx, id); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"message": "list deleted", "id": id})
		return nil
	},
}

func init() {
	listListCmd.Flags().String("folder", "", "Folder ID")
	listListCmd.Flags().String("space", "", "Space ID (for folderless lists)")

	listGetCmd.Flags().String("id", "", "List ID")

	listCreateCmd.Flags().String("folder", "", "Folder ID")
	listCreateCmd.Flags().String("space", "", "Space ID (for folderless list)")
	listCreateCmd.Flags().String("name", "", "List name")
	listCreateCmd.Flags().String("content", "", "List description/content")
	listCreateCmd.Flags().Int64("due-date", 0, "Due date (Unix ms)")
	listCreateCmd.Flags().Int("priority", 0, "Priority (1=urgent, 2=high, 3=normal, 4=low)")
	listCreateCmd.Flags().Int("assignee", 0, "Assignee user ID")
	listCreateCmd.Flags().String("status", "", "List status")

	listUpdateCmd.Flags().String("id", "", "List ID")
	listUpdateCmd.Flags().String("name", "", "List name")
	listUpdateCmd.Flags().String("content", "", "List description/content")
	listUpdateCmd.Flags().Int64("due-date", 0, "Due date (Unix ms)")
	listUpdateCmd.Flags().Int("priority", 0, "Priority")
	listUpdateCmd.Flags().Int("assignee", 0, "Assignee user ID")
	listUpdateCmd.Flags().String("status", "", "List status")
	listUpdateCmd.Flags().Bool("unset-status", false, "Remove list status")

	listDeleteCmd.Flags().String("id", "", "List ID")

	listCmd.AddCommand(listListCmd)
	listCmd.AddCommand(listGetCmd)
	listCmd.AddCommand(listCreateCmd)
	listCmd.AddCommand(listUpdateCmd)
	listCmd.AddCommand(listDeleteCmd)
	rootCmd.AddCommand(listCmd)
}
