package cmd

import (
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
		wsID := getWorkspaceID(cmd)
		resp, err := client.ListSpaces(wsID)
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
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		resp, err := client.GetSpace(id)
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
		wsID := getWorkspaceID(cmd)
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			output.PrintError("VALIDATION_ERROR", "--name is required")
			return &exitError{code: 1}
		}
		resp, err := client.CreateSpace(wsID, &api.CreateSpaceRequest{Name: name})
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

func init() {
	spaceListCmd.Flags().String("workspace", "", "Workspace ID")
	spaceGetCmd.Flags().String("id", "", "Space ID")
	spaceCreateCmd.Flags().String("workspace", "", "Workspace ID")
	spaceCreateCmd.Flags().String("name", "", "Space name")

	spaceCmd.AddCommand(spaceListCmd)
	spaceCmd.AddCommand(spaceGetCmd)
	spaceCmd.AddCommand(spaceCreateCmd)
	rootCmd.AddCommand(spaceCmd)
}
