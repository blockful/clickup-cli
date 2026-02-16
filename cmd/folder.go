package cmd

import (
	"github.com/blockful/clickup-cli/internal/api"
	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
)

var folderCmd = &cobra.Command{
	Use:   "folder",
	Short: "Manage folders",
}

var folderListCmd = &cobra.Command{
	Use:   "list",
	Short: "List folders in a space",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		spaceID, _ := cmd.Flags().GetString("space")
		if spaceID == "" {
			output.PrintError("VALIDATION_ERROR", "--space is required")
			return &exitError{code: 1}
		}
		resp, err := client.ListFolders(spaceID)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var folderGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a folder by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		resp, err := client.GetFolder(id)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var folderCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new folder",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		spaceID, _ := cmd.Flags().GetString("space")
		if spaceID == "" {
			output.PrintError("VALIDATION_ERROR", "--space is required")
			return &exitError{code: 1}
		}
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			output.PrintError("VALIDATION_ERROR", "--name is required")
			return &exitError{code: 1}
		}
		resp, err := client.CreateFolder(spaceID, &api.CreateFolderRequest{Name: name})
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

func init() {
	folderListCmd.Flags().String("space", "", "Space ID")
	folderGetCmd.Flags().String("id", "", "Folder ID")
	folderCreateCmd.Flags().String("space", "", "Space ID")
	folderCreateCmd.Flags().String("name", "", "Folder name")

	folderCmd.AddCommand(folderListCmd)
	folderCmd.AddCommand(folderGetCmd)
	folderCmd.AddCommand(folderCreateCmd)
	rootCmd.AddCommand(folderCmd)
}
