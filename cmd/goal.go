package cmd

import (
	"context"
	"github.com/blockful/clickup-cli/internal/api"
	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
)

var goalCmd = &cobra.Command{
	Use:   "goal",
	Short: "Manage goals",
}

var goalListCmd = &cobra.Command{
	Use:   "list",
	Short: "List goals",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		includeCompleted, _ := cmd.Flags().GetBool("include-completed")
		resp, err := client.GetGoals(ctx, wid, includeCompleted)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var goalGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a goal",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		resp, err := client.GetGoal(ctx, id)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var goalCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a goal",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		name, _ := cmd.Flags().GetString("name")
		dueDate, _ := cmd.Flags().GetInt64("due-date")
		description, _ := cmd.Flags().GetString("description")
		color, _ := cmd.Flags().GetString("color")
		multipleOwners, _ := cmd.Flags().GetBool("multiple-owners")

		if name == "" {
			output.PrintError("VALIDATION_ERROR", "--name is required")
			return &exitError{code: 1}
		}

		req := &api.CreateGoalRequest{
			Name:           name,
			DueDate:        dueDate,
			Description:    description,
			Color:          color,
			MultipleOwners: multipleOwners,
		}

		resp, err := client.CreateGoal(ctx, wid, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var goalUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a goal",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}

		req := &api.UpdateGoalRequest{}
		if cmd.Flags().Changed("name") {
			req.Name, _ = cmd.Flags().GetString("name")
		}
		if cmd.Flags().Changed("description") {
			req.Description, _ = cmd.Flags().GetString("description")
		}
		if cmd.Flags().Changed("color") {
			req.Color, _ = cmd.Flags().GetString("color")
		}

		resp, err := client.UpdateGoal(ctx, id, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var goalDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a goal",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		if err := client.DeleteGoal(ctx, id); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

func init() {
	rootCmd.AddCommand(goalCmd)
	goalCmd.AddCommand(goalListCmd, goalGetCmd, goalCreateCmd, goalUpdateCmd, goalDeleteCmd)

	goalListCmd.Flags().Bool("include-completed", false, "Include completed goals")

	goalGetCmd.Flags().String("id", "", "Goal ID (required)")

	goalCreateCmd.Flags().String("name", "", "Goal name (required)")
	goalCreateCmd.Flags().Int64("due-date", 0, "Due date (unix ms)")
	goalCreateCmd.Flags().String("description", "", "Description")
	goalCreateCmd.Flags().String("color", "", "Color hex")
	goalCreateCmd.Flags().Bool("multiple-owners", false, "Allow multiple owners")

	goalUpdateCmd.Flags().String("id", "", "Goal ID (required)")
	goalUpdateCmd.Flags().String("name", "", "New name")
	goalUpdateCmd.Flags().String("description", "", "New description")
	goalUpdateCmd.Flags().String("color", "", "New color")

	goalDeleteCmd.Flags().String("id", "", "Goal ID (required)")
}
