package cmd

import (
	"context"

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
		ctx := context.Background()
		spaceID, _ := cmd.Flags().GetString("space")
		if spaceID == "" {
			output.PrintError("VALIDATION_ERROR", "--space is required")
			return &exitError{code: 1}
		}
		resp, err := client.ListFolders(ctx, spaceID)
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
		ctx := context.Background()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		resp, err := client.GetFolder(ctx, id)
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
		ctx := context.Background()
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
		resp, err := client.CreateFolder(ctx, spaceID, &api.CreateFolderRequest{Name: name})
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var folderUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a folder",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			output.PrintError("VALIDATION_ERROR", "--name is required")
			return &exitError{code: 1}
		}
		resp, err := client.UpdateFolder(ctx, id, &api.UpdateFolderRequest{Name: name})
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var folderDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a folder",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		if err := client.DeleteFolder(ctx, id); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"message": "folder deleted", "id": id})
		return nil
	},
}

func init() {
	folderListCmd.Flags().String("space", "", "Space ID")
	folderGetCmd.Flags().String("id", "", "Folder ID")
	folderCreateCmd.Flags().String("space", "", "Space ID")
	folderCreateCmd.Flags().String("name", "", "Folder name")
	folderUpdateCmd.Flags().String("id", "", "Folder ID")
	folderUpdateCmd.Flags().String("name", "", "Folder name")
	folderDeleteCmd.Flags().String("id", "", "Folder ID")

	folderCmd.AddCommand(folderListCmd)
	folderCmd.AddCommand(folderGetCmd)
	folderCmd.AddCommand(folderCreateCmd)
	folderCmd.AddCommand(folderUpdateCmd)
	folderCmd.AddCommand(folderDeleteCmd)
	rootCmd.AddCommand(folderCmd)
}
