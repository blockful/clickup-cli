package cmd

import (
	"context"
	"encoding/json"

	"github.com/blockful/clickup-cli/internal/api"
	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
)

var spaceCmd = &cobra.Command{
	Use:   "space",
	Short: "Manage spaces",
}

var spaceListCmd = &cobra.Command{
	Use:   "list",
	Short: "List spaces in a workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wsID := getWorkspaceID(cmd)
		resp, err := client.ListSpaces(ctx, wsID)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var spaceGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a space by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		resp, err := client.GetSpace(ctx, id)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var spaceCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new space",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wsID := getWorkspaceID(cmd)
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			output.PrintError("VALIDATION_ERROR", "--name is required")
			return &exitError{code: 1}
		}
		req := &api.CreateSpaceRequest{Name: name}
		req.MultipleAssignees, _ = cmd.Flags().GetBool("multiple-assignees")
		featuresStr, _ := cmd.Flags().GetString("features")
		if featuresStr != "" {
			var features map[string]interface{}
			if err := json.Unmarshal([]byte(featuresStr), &features); err != nil {
				output.PrintError("VALIDATION_ERROR", "invalid --features JSON: "+err.Error())
				return &exitError{code: 1}
			}
			req.Features = features
		}
		resp, err := client.CreateSpace(ctx, wsID, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var spaceUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a space",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		req := &api.UpdateSpaceRequest{}
		if cmd.Flags().Changed("name") {
			req.Name, _ = cmd.Flags().GetString("name")
		}
		if cmd.Flags().Changed("multiple-assignees") {
			v, _ := cmd.Flags().GetBool("multiple-assignees")
			req.MultipleAssignees = api.BoolPtr(v)
		}
		if cmd.Flags().Changed("private") {
			v, _ := cmd.Flags().GetBool("private")
			req.Private = api.BoolPtr(v)
		}
		if cmd.Flags().Changed("admin-can-manage") {
			v, _ := cmd.Flags().GetBool("admin-can-manage")
			req.AdminCanManage = api.BoolPtr(v)
		}
		featuresStr, _ := cmd.Flags().GetString("features")
		if featuresStr != "" {
			var features map[string]interface{}
			if err := json.Unmarshal([]byte(featuresStr), &features); err != nil {
				output.PrintError("VALIDATION_ERROR", "invalid --features JSON: "+err.Error())
				return &exitError{code: 1}
			}
			req.Features = features
		}
		resp, err := client.UpdateSpace(ctx, id, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var spaceDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a space",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		if err := client.DeleteSpace(ctx, id); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"message": "space deleted", "id": id})
		return nil
	},
}

func init() {
	spaceListCmd.Flags().String("workspace", "", "Workspace ID")

	spaceGetCmd.Flags().String("id", "", "Space ID")

	spaceCreateCmd.Flags().String("workspace", "", "Workspace ID")
	spaceCreateCmd.Flags().String("name", "", "Space name")
	spaceCreateCmd.Flags().Bool("multiple-assignees", false, "Enable multiple assignees")
	spaceCreateCmd.Flags().String("features", "", "Space features (JSON object)")

	spaceUpdateCmd.Flags().String("id", "", "Space ID")
	spaceUpdateCmd.Flags().String("name", "", "Space name")
	spaceUpdateCmd.Flags().Bool("multiple-assignees", false, "Enable multiple assignees")
	spaceUpdateCmd.Flags().Bool("private", false, "Make space private")
	spaceUpdateCmd.Flags().Bool("admin-can-manage", false, "Allow admins to manage")
	spaceUpdateCmd.Flags().String("features", "", "Space features (JSON object)")

	spaceDeleteCmd.Flags().String("id", "", "Space ID")

	spaceCmd.AddCommand(spaceListCmd)
	spaceCmd.AddCommand(spaceGetCmd)
	spaceCmd.AddCommand(spaceCreateCmd)
	spaceCmd.AddCommand(spaceUpdateCmd)
	spaceCmd.AddCommand(spaceDeleteCmd)
	rootCmd.AddCommand(spaceCmd)
}
