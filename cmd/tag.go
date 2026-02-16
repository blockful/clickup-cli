package cmd

import (
	"context"

	"github.com/blockful/clickup-cli/internal/api"
	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Manage tags",
}

var tagListCmd = &cobra.Command{
	Use:   "list",
	Short: "List tags in a space",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		spaceID, _ := cmd.Flags().GetString("space")
		if spaceID == "" {
			output.PrintError("VALIDATION_ERROR", "--space is required")
			return &exitError{code: 1}
		}
		resp, err := client.GetSpaceTags(ctx, spaceID)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var tagCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a tag in a space",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		spaceID, _ := cmd.Flags().GetString("space")
		name, _ := cmd.Flags().GetString("name")
		fg, _ := cmd.Flags().GetString("fg")
		bg, _ := cmd.Flags().GetString("bg")

		if spaceID == "" || name == "" {
			output.PrintError("VALIDATION_ERROR", "--space and --name are required")
			return &exitError{code: 1}
		}

		req := &api.CreateTagRequest{Tag: api.Tag{Name: name, TagFg: fg, TagBg: bg}}
		if err := client.CreateSpaceTag(ctx, spaceID, req); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

var tagUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a tag in a space",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		spaceID, _ := cmd.Flags().GetString("space")
		tagName, _ := cmd.Flags().GetString("name")
		newName, _ := cmd.Flags().GetString("new-name")
		fg, _ := cmd.Flags().GetString("fg")
		bg, _ := cmd.Flags().GetString("bg")

		if spaceID == "" || tagName == "" {
			output.PrintError("VALIDATION_ERROR", "--space and --name are required")
			return &exitError{code: 1}
		}

		tag := api.Tag{TagFg: fg, TagBg: bg}
		if newName != "" {
			tag.Name = newName
		} else {
			tag.Name = tagName
		}
		req := &api.UpdateTagRequest{Tag: tag}
		if err := client.UpdateSpaceTag(ctx, spaceID, tagName, req); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

var tagDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a tag from a space",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		spaceID, _ := cmd.Flags().GetString("space")
		tagName, _ := cmd.Flags().GetString("name")

		if spaceID == "" || tagName == "" {
			output.PrintError("VALIDATION_ERROR", "--space and --name are required")
			return &exitError{code: 1}
		}

		if err := client.DeleteSpaceTag(ctx, spaceID, tagName); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

var tagAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a tag to a task",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		taskID, _ := cmd.Flags().GetString("task")
		tagName, _ := cmd.Flags().GetString("name")

		if taskID == "" || tagName == "" {
			output.PrintError("VALIDATION_ERROR", "--task and --name are required")
			return &exitError{code: 1}
		}

		if err := client.AddTagToTask(ctx, taskID, tagName); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

var tagRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a tag from a task",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		taskID, _ := cmd.Flags().GetString("task")
		tagName, _ := cmd.Flags().GetString("name")

		if taskID == "" || tagName == "" {
			output.PrintError("VALIDATION_ERROR", "--task and --name are required")
			return &exitError{code: 1}
		}

		if err := client.RemoveTagFromTask(ctx, taskID, tagName); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)
	tagCmd.AddCommand(tagListCmd, tagCreateCmd, tagUpdateCmd, tagDeleteCmd, tagAddCmd, tagRemoveCmd)

	tagListCmd.Flags().String("space", "", "Space ID (required)")

	tagCreateCmd.Flags().String("space", "", "Space ID (required)")
	tagCreateCmd.Flags().String("name", "", "Tag name (required)")
	tagCreateCmd.Flags().String("fg", "#000000", "Foreground color")
	tagCreateCmd.Flags().String("bg", "#000000", "Background color")

	tagUpdateCmd.Flags().String("space", "", "Space ID (required)")
	tagUpdateCmd.Flags().String("name", "", "Current tag name (required)")
	tagUpdateCmd.Flags().String("new-name", "", "New tag name")
	tagUpdateCmd.Flags().String("fg", "", "Foreground color")
	tagUpdateCmd.Flags().String("bg", "", "Background color")

	tagDeleteCmd.Flags().String("space", "", "Space ID (required)")
	tagDeleteCmd.Flags().String("name", "", "Tag name (required)")

	tagAddCmd.Flags().String("task", "", "Task ID (required)")
	tagAddCmd.Flags().String("name", "", "Tag name (required)")

	tagRemoveCmd.Flags().String("task", "", "Task ID (required)")
	tagRemoveCmd.Flags().String("name", "", "Tag name (required)")
}
