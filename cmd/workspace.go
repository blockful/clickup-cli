package cmd

import (
	"context"

	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
)

var workspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "Manage workspaces",
}

var workspaceListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all workspaces",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		resp, err := client.ListWorkspaces(ctx)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var workspaceSeatsCmd = &cobra.Command{
	Use:   "seats",
	Short: "Get workspace seats",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		resp, err := client.GetWorkspaceSeats(ctx, wid)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var workspacePlanCmd = &cobra.Command{
	Use:   "plan",
	Short: "Get workspace plan",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		resp, err := client.GetWorkspacePlan(ctx, wid)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

func init() {
	workspaceCmd.AddCommand(workspaceListCmd, workspaceSeatsCmd, workspacePlanCmd)
	rootCmd.AddCommand(workspaceCmd)
}
