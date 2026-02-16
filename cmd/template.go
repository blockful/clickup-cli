package cmd

import (
	"context"
	"encoding/json"

	"github.com/blockful/clickup-cli/internal/api"
	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
)

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Manage templates",
}

var templateListCmd = &cobra.Command{
	Use:   "list",
	Short: "List task templates",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		page, _ := cmd.Flags().GetInt("page")
		resp, err := client.GetTaskTemplates(ctx, wid, page)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var templateCreateTaskCmd = &cobra.Command{
	Use:   "create-task",
	Short: "Create a task from a template",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		listID, _ := cmd.Flags().GetString("list")
		templateID, _ := cmd.Flags().GetString("template-id")
		name, _ := cmd.Flags().GetString("name")
		if listID == "" {
			output.PrintError("VALIDATION_ERROR", "--list is required")
			return &exitError{code: 1}
		}
		if templateID == "" {
			output.PrintError("VALIDATION_ERROR", "--template-id is required")
			return &exitError{code: 1}
		}
		if name == "" {
			output.PrintError("VALIDATION_ERROR", "--name is required")
			return &exitError{code: 1}
		}
		resp, err := client.CreateTaskFromTemplate(ctx, listID, templateID, &api.CreateFromTemplateRequest{Name: name})
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var templateCreateFolderCmd = &cobra.Command{
	Use:   "create-folder",
	Short: "Create a folder from a template",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		spaceID, _ := cmd.Flags().GetString("space")
		templateID, _ := cmd.Flags().GetString("template-id")
		name, _ := cmd.Flags().GetString("name")
		if spaceID == "" {
			output.PrintError("VALIDATION_ERROR", "--space is required")
			return &exitError{code: 1}
		}
		if templateID == "" {
			output.PrintError("VALIDATION_ERROR", "--template-id is required")
			return &exitError{code: 1}
		}
		if name == "" {
			output.PrintError("VALIDATION_ERROR", "--name is required")
			return &exitError{code: 1}
		}
		req := &api.CreateFromTemplateRequest{Name: name}
		if optStr, _ := cmd.Flags().GetString("options"); optStr != "" {
			var v interface{}
			if err := json.Unmarshal([]byte(optStr), &v); err == nil {
				req.Options = v
			}
		}
		resp, err := client.CreateFolderFromTemplate(ctx, spaceID, templateID, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var templateCreateListCmd = &cobra.Command{
	Use:   "create-list",
	Short: "Create a list from a template",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		folderID, _ := cmd.Flags().GetString("folder")
		spaceID, _ := cmd.Flags().GetString("space")
		templateID, _ := cmd.Flags().GetString("template-id")
		name, _ := cmd.Flags().GetString("name")
		if folderID == "" && spaceID == "" {
			output.PrintError("VALIDATION_ERROR", "--folder or --space is required")
			return &exitError{code: 1}
		}
		if templateID == "" {
			output.PrintError("VALIDATION_ERROR", "--template-id is required")
			return &exitError{code: 1}
		}
		if name == "" {
			output.PrintError("VALIDATION_ERROR", "--name is required")
			return &exitError{code: 1}
		}
		req := &api.CreateFromTemplateRequest{Name: name}
		if optStr, _ := cmd.Flags().GetString("options"); optStr != "" {
			var v interface{}
			if err := json.Unmarshal([]byte(optStr), &v); err == nil {
				req.Options = v
			}
		}
		if folderID != "" {
			resp, err := client.CreateListFromFolderTemplate(ctx, folderID, templateID, req)
			if err != nil {
				return handleError(err)
			}
			output.JSON(resp)
			return nil
		}
		resp, err := client.CreateListFromSpaceTemplate(ctx, spaceID, templateID, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

func init() {
	templateListCmd.Flags().Int("page", 0, "Page number")

	templateCreateTaskCmd.Flags().String("list", "", "List ID (required)")
	templateCreateTaskCmd.Flags().String("template-id", "", "Template ID (required)")
	templateCreateTaskCmd.Flags().String("name", "", "Task name (required)")

	templateCreateFolderCmd.Flags().String("space", "", "Space ID (required)")
	templateCreateFolderCmd.Flags().String("template-id", "", "Template ID (required)")
	templateCreateFolderCmd.Flags().String("name", "", "Folder name (required)")
	templateCreateFolderCmd.Flags().String("options", "", "Template options (JSON)")

	templateCreateListCmd.Flags().String("folder", "", "Folder ID")
	templateCreateListCmd.Flags().String("space", "", "Space ID")
	templateCreateListCmd.Flags().String("template-id", "", "Template ID (required)")
	templateCreateListCmd.Flags().String("name", "", "List name (required)")
	templateCreateListCmd.Flags().String("options", "", "Template options (JSON)")

	templateCmd.AddCommand(templateListCmd)
	templateCmd.AddCommand(templateCreateTaskCmd)
	templateCmd.AddCommand(templateCreateFolderCmd)
	templateCmd.AddCommand(templateCreateListCmd)
	rootCmd.AddCommand(templateCmd)
}
