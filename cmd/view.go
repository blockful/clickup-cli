package cmd

import (
	"context"
	"encoding/json"

	"github.com/blockful/clickup-cli/internal/api"
	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
)

func parseJSONFlag(cmd *cobra.Command, name string) interface{} {
	s, _ := cmd.Flags().GetString(name)
	if s == "" {
		return nil
	}
	var v interface{}
	if err := json.Unmarshal([]byte(s), &v); err != nil {
		return nil
	}
	return v
}

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Manage views",
}

var viewListCmd = &cobra.Command{
	Use:   "list",
	Short: "List views at workspace/space/folder/list level",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		spaceID, _ := cmd.Flags().GetString("space")
		folderID, _ := cmd.Flags().GetString("folder")
		listID, _ := cmd.Flags().GetString("list")

		var resp *api.ViewsResponse
		var err error

		switch {
		case listID != "":
			resp, err = client.GetListViews(ctx, listID)
		case folderID != "":
			resp, err = client.GetFolderViews(ctx, folderID)
		case spaceID != "":
			resp, err = client.GetSpaceViews(ctx, spaceID)
		default:
			wid := getWorkspaceID(cmd)
			resp, err = client.GetTeamViews(ctx, wid)
		}

		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var viewGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a view by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		resp, err := client.GetView(ctx, id)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var viewCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a view",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		name, _ := cmd.Flags().GetString("name")
		viewType, _ := cmd.Flags().GetString("type")
		spaceID, _ := cmd.Flags().GetString("space")
		folderID, _ := cmd.Flags().GetString("folder")
		listID, _ := cmd.Flags().GetString("list")

		if name == "" || viewType == "" {
			output.PrintError("VALIDATION_ERROR", "--name and --type are required")
			return &exitError{code: 1}
		}

		req := &api.CreateViewRequest{
			Name:        name,
			Type:        viewType,
			Filters:     parseJSONFlag(cmd, "filters"),
			Sorting:     parseJSONFlag(cmd, "sorting"),
			Grouping:    parseJSONFlag(cmd, "grouping"),
			Columns:     parseJSONFlag(cmd, "columns"),
			TeamSidebar: parseJSONFlag(cmd, "team-sidebar"),
			Settings:    parseJSONFlag(cmd, "settings"),
			Divide:      parseJSONFlag(cmd, "divide"),
		}
		var resp *api.ViewResponse
		var err error

		switch {
		case listID != "":
			resp, err = client.CreateListView(ctx, listID, req)
		case folderID != "":
			resp, err = client.CreateFolderView(ctx, folderID, req)
		case spaceID != "":
			resp, err = client.CreateSpaceView(ctx, spaceID, req)
		default:
			wid := getWorkspaceID(cmd)
			resp, err = client.CreateTeamView(ctx, wid, req)
		}

		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var viewUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a view",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		id, _ := cmd.Flags().GetString("id")
		name, _ := cmd.Flags().GetString("name")

		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}

		viewType, _ := cmd.Flags().GetString("type")
		req := &api.UpdateViewRequest{
			Name:        name,
			Type:        viewType,
			Filters:     parseJSONFlag(cmd, "filters"),
			Sorting:     parseJSONFlag(cmd, "sorting"),
			Grouping:    parseJSONFlag(cmd, "grouping"),
			Columns:     parseJSONFlag(cmd, "columns"),
			TeamSidebar: parseJSONFlag(cmd, "team-sidebar"),
			Settings:    parseJSONFlag(cmd, "settings"),
			Divide:      parseJSONFlag(cmd, "divide"),
		}
		resp, err := client.UpdateView(ctx, id, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var viewDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a view",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		if err := client.DeleteView(ctx, id); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

var viewTasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Get tasks in a view",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		id, _ := cmd.Flags().GetString("id")
		page, _ := cmd.Flags().GetInt("page")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		resp, err := client.GetViewTasks(ctx, id, page)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
	viewCmd.AddCommand(viewListCmd, viewGetCmd, viewCreateCmd, viewUpdateCmd, viewDeleteCmd, viewTasksCmd)

	viewListCmd.Flags().String("space", "", "Space ID")
	viewListCmd.Flags().String("folder", "", "Folder ID")
	viewListCmd.Flags().String("list", "", "List ID")

	viewGetCmd.Flags().String("id", "", "View ID (required)")

	viewCreateCmd.Flags().String("name", "", "View name (required)")
	viewCreateCmd.Flags().String("type", "", "View type: list, board, calendar, etc. (required)")
	viewCreateCmd.Flags().String("space", "", "Space ID")
	viewCreateCmd.Flags().String("folder", "", "Folder ID")
	viewCreateCmd.Flags().String("list", "", "List ID")
	viewCreateCmd.Flags().String("filters", "", "Filters config (JSON)")
	viewCreateCmd.Flags().String("sorting", "", "Sorting config (JSON)")
	viewCreateCmd.Flags().String("grouping", "", "Grouping config (JSON)")
	viewCreateCmd.Flags().String("columns", "", "Columns config (JSON)")
	viewCreateCmd.Flags().String("team-sidebar", "", "Team sidebar config (JSON)")
	viewCreateCmd.Flags().String("settings", "", "Settings config (JSON)")
	viewCreateCmd.Flags().String("divide", "", "Divide config (JSON)")

	viewUpdateCmd.Flags().String("id", "", "View ID (required)")
	viewUpdateCmd.Flags().String("name", "", "New name")
	viewUpdateCmd.Flags().String("type", "", "View type")
	viewUpdateCmd.Flags().String("filters", "", "Filters config (JSON)")
	viewUpdateCmd.Flags().String("sorting", "", "Sorting config (JSON)")
	viewUpdateCmd.Flags().String("grouping", "", "Grouping config (JSON)")
	viewUpdateCmd.Flags().String("columns", "", "Columns config (JSON)")
	viewUpdateCmd.Flags().String("team-sidebar", "", "Team sidebar config (JSON)")
	viewUpdateCmd.Flags().String("settings", "", "Settings config (JSON)")
	viewUpdateCmd.Flags().String("divide", "", "Divide config (JSON)")

	viewDeleteCmd.Flags().String("id", "", "View ID (required)")

	viewTasksCmd.Flags().String("id", "", "View ID (required)")
	viewTasksCmd.Flags().Int("page", 0, "Page number")
}
