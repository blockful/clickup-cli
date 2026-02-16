package cmd

import (
	"github.com/blockful/clickup-cli/internal/api"
	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
)

var checklistCmd = &cobra.Command{
	Use:   "checklist",
	Short: "Manage checklists",
}

var checklistItemCmd = &cobra.Command{
	Use:   "checklist-item",
	Short: "Manage checklist items",
}

var checklistCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a checklist on a task",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		taskID, _ := cmd.Flags().GetString("task")
		name, _ := cmd.Flags().GetString("name")

		if taskID == "" || name == "" {
			output.PrintError("VALIDATION_ERROR", "--task and --name are required")
			return &exitError{code: 1}
		}

		resp, err := client.CreateChecklist(taskID, &api.CreateChecklistRequest{Name: name})
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var checklistUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a checklist",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		id, _ := cmd.Flags().GetString("id")
		name, _ := cmd.Flags().GetString("name")
		position, _ := cmd.Flags().GetInt("position")

		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}

		req := &api.EditChecklistRequest{}
		if name != "" {
			req.Name = name
		}
		if cmd.Flags().Changed("position") {
			req.Position = &position
		}

		if err := client.EditChecklist(id, req); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

var checklistDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a checklist",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}

		if err := client.DeleteChecklist(id); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

var checklistItemCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a checklist item",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		checklistID, _ := cmd.Flags().GetString("checklist")
		name, _ := cmd.Flags().GetString("name")
		assignee, _ := cmd.Flags().GetInt("assignee")

		if checklistID == "" {
			output.PrintError("VALIDATION_ERROR", "--checklist is required")
			return &exitError{code: 1}
		}

		req := &api.CreateChecklistItemRequest{Name: name}
		if cmd.Flags().Changed("assignee") {
			req.Assignee = &assignee
		}

		resp, err := client.CreateChecklistItem(checklistID, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var checklistItemUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a checklist item",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		checklistID, _ := cmd.Flags().GetString("checklist")
		itemID, _ := cmd.Flags().GetString("id")
		name, _ := cmd.Flags().GetString("name")
		resolved, _ := cmd.Flags().GetBool("resolved")
		assignee, _ := cmd.Flags().GetString("assignee")
		parent, _ := cmd.Flags().GetString("parent")

		if checklistID == "" || itemID == "" {
			output.PrintError("VALIDATION_ERROR", "--checklist and --id are required")
			return &exitError{code: 1}
		}

		req := &api.EditChecklistItemRequest{}
		if name != "" {
			req.Name = name
		}
		if cmd.Flags().Changed("resolved") {
			req.Resolved = &resolved
		}
		if cmd.Flags().Changed("assignee") {
			req.Assignee = &assignee
		}
		if cmd.Flags().Changed("parent") {
			req.Parent = &parent
		}

		resp, err := client.EditChecklistItem(checklistID, itemID, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var checklistItemDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a checklist item",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		checklistID, _ := cmd.Flags().GetString("checklist")
		itemID, _ := cmd.Flags().GetString("id")

		if checklistID == "" || itemID == "" {
			output.PrintError("VALIDATION_ERROR", "--checklist and --id are required")
			return &exitError{code: 1}
		}

		if err := client.DeleteChecklistItem(checklistID, itemID); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

func init() {
	rootCmd.AddCommand(checklistCmd)
	rootCmd.AddCommand(checklistItemCmd)

	checklistCmd.AddCommand(checklistCreateCmd, checklistUpdateCmd, checklistDeleteCmd)
	checklistItemCmd.AddCommand(checklistItemCreateCmd, checklistItemUpdateCmd, checklistItemDeleteCmd)

	checklistCreateCmd.Flags().String("task", "", "Task ID (required)")
	checklistCreateCmd.Flags().String("name", "", "Checklist name (required)")

	checklistUpdateCmd.Flags().String("id", "", "Checklist ID (required)")
	checklistUpdateCmd.Flags().String("name", "", "New name")
	checklistUpdateCmd.Flags().Int("position", 0, "Position")

	checklistDeleteCmd.Flags().String("id", "", "Checklist ID (required)")

	checklistItemCreateCmd.Flags().String("checklist", "", "Checklist ID (required)")
	checklistItemCreateCmd.Flags().String("name", "", "Item name")
	checklistItemCreateCmd.Flags().Int("assignee", 0, "Assignee user ID")

	checklistItemUpdateCmd.Flags().String("checklist", "", "Checklist ID (required)")
	checklistItemUpdateCmd.Flags().String("id", "", "Checklist item ID (required)")
	checklistItemUpdateCmd.Flags().String("name", "", "New name")
	checklistItemUpdateCmd.Flags().Bool("resolved", false, "Resolved status")
	checklistItemUpdateCmd.Flags().String("assignee", "", "Assignee user ID or null")
	checklistItemUpdateCmd.Flags().String("parent", "", "Parent checklist item ID or null")

	checklistItemDeleteCmd.Flags().String("checklist", "", "Checklist ID (required)")
	checklistItemDeleteCmd.Flags().String("id", "", "Checklist item ID (required)")
}
