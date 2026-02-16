package cmd

import (
	"context"

	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
)

var sharedCmd = &cobra.Command{
	Use:   "shared",
	Short: "Manage shared hierarchy",
}

var sharedListCmd = &cobra.Command{
	Use:   "list",
	Short: "List shared hierarchy",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		resp, err := client.GetSharedHierarchy(ctx, wid)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(sharedCmd)
	sharedCmd.AddCommand(sharedListCmd)
}
