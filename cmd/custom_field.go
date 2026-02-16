package cmd

import (
	"encoding/json"

	"github.com/blockful/clickup-cli/internal/api"
	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
)

var customFieldCmd = &cobra.Command{
	Use:   "custom-field",
	Short: "Manage custom fields",
}

var customFieldListCmd = &cobra.Command{
	Use:   "list",
	Short: "List custom fields for a list, folder, space, or workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		listID, _ := cmd.Flags().GetString("list")
		folderID, _ := cmd.Flags().GetString("folder")
		spaceID, _ := cmd.Flags().GetString("space")
		workspaceID, _ := cmd.Flags().GetString("workspace")

		var resp *api.CustomFieldsResponse
		var err error

		switch {
		case listID != "":
			resp, err = client.GetListCustomFields(listID)
		case folderID != "":
			resp, err = client.GetFolderCustomFields(folderID)
		case spaceID != "":
			resp, err = client.GetSpaceCustomFields(spaceID)
		case workspaceID != "":
			resp, err = client.GetWorkspaceCustomFields(workspaceID)
		default:
			output.PrintError("VALIDATION_ERROR", "one of --list, --folder, --space, or --workspace is required")
			return &exitError{code: 1}
		}

		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var customFieldSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a custom field value on a task",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		taskID, _ := cmd.Flags().GetString("task")
		fieldID, _ := cmd.Flags().GetString("field")
		value, _ := cmd.Flags().GetString("value")

		if taskID == "" || fieldID == "" || value == "" {
			output.PrintError("VALIDATION_ERROR", "--task, --field, and --value are required")
			return &exitError{code: 1}
		}

		// Try to parse value as JSON, fall back to string
		var v interface{}
		if err := json.Unmarshal([]byte(value), &v); err != nil {
			v = value
		}

		req := &api.SetCustomFieldRequest{Value: v}
		if err := client.SetCustomFieldValue(taskID, fieldID, req); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

var customFieldRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a custom field value from a task",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		taskID, _ := cmd.Flags().GetString("task")
		fieldID, _ := cmd.Flags().GetString("field")

		if taskID == "" || fieldID == "" {
			output.PrintError("VALIDATION_ERROR", "--task and --field are required")
			return &exitError{code: 1}
		}

		if err := client.RemoveCustomFieldValue(taskID, fieldID); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

func init() {
	rootCmd.AddCommand(customFieldCmd)
	customFieldCmd.AddCommand(customFieldListCmd)
	customFieldCmd.AddCommand(customFieldSetCmd)
	customFieldCmd.AddCommand(customFieldRemoveCmd)

	customFieldListCmd.Flags().String("list", "", "List ID")
	customFieldListCmd.Flags().String("folder", "", "Folder ID")
	customFieldListCmd.Flags().String("space", "", "Space ID")
	customFieldListCmd.Flags().String("workspace", "", "Workspace ID")

	customFieldSetCmd.Flags().String("task", "", "Task ID (required)")
	customFieldSetCmd.Flags().String("field", "", "Field ID (required)")
	customFieldSetCmd.Flags().String("value", "", "Field value as JSON or string (required)")

	customFieldRemoveCmd.Flags().String("task", "", "Task ID (required)")
	customFieldRemoveCmd.Flags().String("field", "", "Field ID (required)")
}
