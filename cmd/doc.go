package cmd

import (
	"github.com/blockful/clickup-cli/internal/api"
	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
)

var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "Manage docs (v3 API)",
}

var docListCmd = &cobra.Command{
	Use:   "list",
	Short: "List/search docs in a workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		wid := getWorkspaceID(cmd)

		resp, err := client.SearchDocs(wid)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var docGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a doc by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		wid := getWorkspaceID(cmd)
		docID, _ := cmd.Flags().GetString("id")
		if docID == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}

		resp, err := client.GetDoc(wid, docID)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var docCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a doc",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		wid := getWorkspaceID(cmd)
		name, _ := cmd.Flags().GetString("name")
		visibility, _ := cmd.Flags().GetString("visibility")
		parentID, _ := cmd.Flags().GetString("parent-id")
		parentType, _ := cmd.Flags().GetInt("parent-type")

		if name == "" {
			output.PrintError("VALIDATION_ERROR", "--name is required")
			return &exitError{code: 1}
		}

		req := &api.CreateDocRequest{Name: name, Visibility: visibility}
		if parentID != "" {
			req.Parent = &api.DocParent{ID: parentID, Type: parentType}
		}

		resp, err := client.CreateDoc(wid, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var docPageListCmd = &cobra.Command{
	Use:   "page-list",
	Short: "List pages in a doc",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		wid := getWorkspaceID(cmd)
		docID, _ := cmd.Flags().GetString("doc")
		if docID == "" {
			output.PrintError("VALIDATION_ERROR", "--doc is required")
			return &exitError{code: 1}
		}

		resp, err := client.GetDocPageListing(wid, docID)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var docPageGetCmd = &cobra.Command{
	Use:   "page-get",
	Short: "Get a page from a doc",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		wid := getWorkspaceID(cmd)
		docID, _ := cmd.Flags().GetString("doc")
		pageID, _ := cmd.Flags().GetString("page")
		if docID == "" || pageID == "" {
			output.PrintError("VALIDATION_ERROR", "--doc and --page are required")
			return &exitError{code: 1}
		}

		resp, err := client.GetPage(wid, docID, pageID)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var docPageCreateCmd = &cobra.Command{
	Use:   "page-create",
	Short: "Create a page in a doc",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		wid := getWorkspaceID(cmd)
		docID, _ := cmd.Flags().GetString("doc")
		name, _ := cmd.Flags().GetString("name")
		content, _ := cmd.Flags().GetString("content")
		contentHtml, _ := cmd.Flags().GetString("content-html")
		parentPageID, _ := cmd.Flags().GetString("parent-page")

		if docID == "" || name == "" {
			output.PrintError("VALIDATION_ERROR", "--doc and --name are required")
			return &exitError{code: 1}
		}

		req := &api.CreatePageRequest{
			Name:         name,
			Content:      content,
			ContentHtml:  contentHtml,
			ParentPageID: parentPageID,
		}

		resp, err := client.CreatePage(wid, docID, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var docPageUpdateCmd = &cobra.Command{
	Use:   "page-update",
	Short: "Update a page in a doc",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		wid := getWorkspaceID(cmd)
		docID, _ := cmd.Flags().GetString("doc")
		pageID, _ := cmd.Flags().GetString("page")
		name, _ := cmd.Flags().GetString("name")
		content, _ := cmd.Flags().GetString("content")
		contentHtml, _ := cmd.Flags().GetString("content-html")

		if docID == "" || pageID == "" {
			output.PrintError("VALIDATION_ERROR", "--doc and --page are required")
			return &exitError{code: 1}
		}

		req := &api.EditPageRequest{
			Name:        name,
			Content:     content,
			ContentHtml: contentHtml,
		}

		resp, err := client.EditPage(wid, docID, pageID, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(docCmd)
	docCmd.AddCommand(docListCmd, docGetCmd, docCreateCmd, docPageListCmd, docPageGetCmd, docPageCreateCmd, docPageUpdateCmd)

	docGetCmd.Flags().String("id", "", "Doc ID (required)")

	docCreateCmd.Flags().String("name", "", "Doc name (required)")
	docCreateCmd.Flags().String("visibility", "", "Visibility")
	docCreateCmd.Flags().String("parent-id", "", "Parent ID")
	docCreateCmd.Flags().Int("parent-type", 0, "Parent type")

	docPageListCmd.Flags().String("doc", "", "Doc ID (required)")

	docPageGetCmd.Flags().String("doc", "", "Doc ID (required)")
	docPageGetCmd.Flags().String("page", "", "Page ID (required)")

	docPageCreateCmd.Flags().String("doc", "", "Doc ID (required)")
	docPageCreateCmd.Flags().String("name", "", "Page name (required)")
	docPageCreateCmd.Flags().String("content", "", "Page content (markdown)")
	docPageCreateCmd.Flags().String("content-html", "", "Page content (HTML)")
	docPageCreateCmd.Flags().String("parent-page", "", "Parent page ID")

	docPageUpdateCmd.Flags().String("doc", "", "Doc ID (required)")
	docPageUpdateCmd.Flags().String("page", "", "Page ID (required)")
	docPageUpdateCmd.Flags().String("name", "", "New name")
	docPageUpdateCmd.Flags().String("content", "", "New content (markdown)")
	docPageUpdateCmd.Flags().String("content-html", "", "New content (HTML)")
}
