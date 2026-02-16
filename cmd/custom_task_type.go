package cmd

import (
	"context"

	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
)

var customTaskTypeCmd = &cobra.Command{
	Use:   "custom-task-type",
	Short: "Manage custom task types",
}

var customTaskTypeListCmd = &cobra.Command{
	Use:   "list",
	Short: "List custom task types",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		resp, err := client.GetCustomTaskTypes(ctx, wid)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(customTaskTypeCmd)
	customTaskTypeCmd.AddCommand(customTaskTypeListCmd)
}
