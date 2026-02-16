package cmd

import (
	"github.com/blockful/clickup-cli/internal/api"
	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Manage lists",
}

var listListCmd = &cobra.Command{
	Use:   "list",
	Short: "List lists in a folder",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		folderID, _ := cmd.Flags().GetString("folder")
		if folderID == "" {
			output.PrintError("VALIDATION_ERROR", "--folder is required")
			return &exitError{code: 1}
		}
		resp, err := client.ListLists(folderID)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var listGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a list by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		resp, err := client.GetList(id)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var listCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new list",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		folderID, _ := cmd.Flags().GetString("folder")
		if folderID == "" {
			output.PrintError("VALIDATION_ERROR", "--folder is required")
			return &exitError{code: 1}
		}
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			output.PrintError("VALIDATION_ERROR", "--name is required")
			return &exitError{code: 1}
		}
		resp, err := client.CreateList(folderID, &api.CreateListRequest{Name: name})
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

func init() {
	listListCmd.Flags().String("folder", "", "Folder ID")
	listGetCmd.Flags().String("id", "", "List ID")
	listCreateCmd.Flags().String("folder", "", "Folder ID")
	listCreateCmd.Flags().String("name", "", "List name")

	listCmd.AddCommand(listListCmd)
	listCmd.AddCommand(listGetCmd)
	listCmd.AddCommand(listCreateCmd)
	rootCmd.AddCommand(listCmd)
}
