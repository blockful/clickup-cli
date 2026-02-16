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

		owners, _ := cmd.Flags().GetIntSlice("owners")
		req := &api.CreateGoalRequest{
			Name:           name,
			DueDate:        dueDate,
			Description:    description,
			Color:          color,
			MultipleOwners: multipleOwners,
			Owners:         owners,
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
		if cmd.Flags().Changed("rem-owners") {
			req.RemOwners, _ = cmd.Flags().GetIntSlice("rem-owners")
		}
		if cmd.Flags().Changed("add-owners") {
			req.AddOwners, _ = cmd.Flags().GetIntSlice("add-owners")
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

var keyResultCmd = &cobra.Command{
	Use:   "key-result",
	Short: "Manage key results",
}

var keyResultCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a key result for a goal",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		goalID, _ := cmd.Flags().GetString("goal-id")
		if goalID == "" {
			output.PrintError("VALIDATION_ERROR", "--goal-id is required")
			return &exitError{code: 1}
		}
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			output.PrintError("VALIDATION_ERROR", "--name is required")
			return &exitError{code: 1}
		}
		typ, _ := cmd.Flags().GetString("type")
		owners, _ := cmd.Flags().GetIntSlice("owners")
		stepsStart, _ := cmd.Flags().GetInt("steps-start")
		stepsEnd, _ := cmd.Flags().GetInt("steps-end")
		unit, _ := cmd.Flags().GetString("unit")
		taskIDs, _ := cmd.Flags().GetStringSlice("task-ids")
		listIDs, _ := cmd.Flags().GetStringSlice("list-ids")

		req := &api.CreateKeyResultRequest{
			Name:       name,
			Type:       typ,
			Owners:     owners,
			StepsStart: stepsStart,
			StepsEnd:   stepsEnd,
			Unit:       unit,
			TaskIDs:    taskIDs,
			ListIDs:    listIDs,
		}
		resp, err := client.CreateKeyResult(ctx, goalID, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var keyResultUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a key result",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		req := &api.UpdateKeyResultRequest{}
		if cmd.Flags().Changed("steps-current") {
			v, _ := cmd.Flags().GetInt("steps-current")
			req.StepsCurrent = &v
		}
		if cmd.Flags().Changed("note") {
			req.Note, _ = cmd.Flags().GetString("note")
		}
		resp, err := client.UpdateKeyResult(ctx, id, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var keyResultDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a key result",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		if err := client.DeleteKeyResult(ctx, id); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

func init() {
	rootCmd.AddCommand(goalCmd)
	goalCmd.AddCommand(goalListCmd, goalGetCmd, goalCreateCmd, goalUpdateCmd, goalDeleteCmd, keyResultCmd)
	keyResultCmd.AddCommand(keyResultCreateCmd, keyResultUpdateCmd, keyResultDeleteCmd)

	keyResultCreateCmd.Flags().String("goal-id", "", "Goal ID (required)")
	keyResultCreateCmd.Flags().String("name", "", "Key result name (required)")
	keyResultCreateCmd.Flags().String("type", "number", "Type (number, percentage, automatic, boolean)")
	keyResultCreateCmd.Flags().IntSlice("owners", nil, "Owner user IDs")
	keyResultCreateCmd.Flags().Int("steps-start", 0, "Starting value")
	keyResultCreateCmd.Flags().Int("steps-end", 0, "Target value")
	keyResultCreateCmd.Flags().String("unit", "", "Unit label")
	keyResultCreateCmd.Flags().StringSlice("task-ids", nil, "Task IDs (for automatic type)")
	keyResultCreateCmd.Flags().StringSlice("list-ids", nil, "List IDs (for automatic type)")

	keyResultUpdateCmd.Flags().String("id", "", "Key Result ID (required)")
	keyResultUpdateCmd.Flags().Int("steps-current", 0, "Current steps value")
	keyResultUpdateCmd.Flags().String("note", "", "Note")

	keyResultDeleteCmd.Flags().String("id", "", "Key Result ID (required)")

	goalListCmd.Flags().Bool("include-completed", false, "Include completed goals")

	goalGetCmd.Flags().String("id", "", "Goal ID (required)")

	goalCreateCmd.Flags().String("name", "", "Goal name (required)")
	goalCreateCmd.Flags().Int64("due-date", 0, "Due date (unix ms)")
	goalCreateCmd.Flags().String("description", "", "Description")
	goalCreateCmd.Flags().String("color", "", "Color hex")
	goalCreateCmd.Flags().Bool("multiple-owners", false, "Allow multiple owners")
	goalCreateCmd.Flags().IntSlice("owners", nil, "Owner user IDs")

	goalUpdateCmd.Flags().String("id", "", "Goal ID (required)")
	goalUpdateCmd.Flags().String("name", "", "New name")
	goalUpdateCmd.Flags().String("description", "", "New description")
	goalUpdateCmd.Flags().String("color", "", "New color")
	goalUpdateCmd.Flags().IntSlice("rem-owners", nil, "Owner IDs to remove")
	goalUpdateCmd.Flags().IntSlice("add-owners", nil, "Owner IDs to add")

	goalDeleteCmd.Flags().String("id", "", "Goal ID (required)")
}
