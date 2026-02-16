package cmd

import (
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
		resp, err := client.ListWorkspaces()
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

func init() {
	workspaceCmd.AddCommand(workspaceListCmd)
	rootCmd.AddCommand(workspaceCmd)
}
