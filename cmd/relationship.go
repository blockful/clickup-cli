package cmd

import (
	"context"

	"github.com/blockful/clickup-cli/internal/api"
	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
)

var dependencyCmd = &cobra.Command{
	Use:   "dependency",
	Short: "Manage task dependencies",
}

var dependencyAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a dependency to a task",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		taskID, _ := cmd.Flags().GetString("task")
		dependsOn, _ := cmd.Flags().GetString("depends-on")
		dependencyOf, _ := cmd.Flags().GetString("dependency-of")
		depType, _ := cmd.Flags().GetString("type")

		if taskID == "" {
			output.PrintError("VALIDATION_ERROR", "--task is required")
			return &exitError{code: 1}
		}
		if dependsOn == "" && dependencyOf == "" {
			output.PrintError("VALIDATION_ERROR", "--depends-on or --dependency-of is required")
			return &exitError{code: 1}
		}

		req := &api.AddDependencyRequest{
			DependsOn:    dependsOn,
			DependencyOf: dependencyOf,
			Type:         depType,
		}
		resp, err := client.AddDependency(ctx, taskID, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var dependencyRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a dependency from a task",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		taskID, _ := cmd.Flags().GetString("task")
		dependsOn, _ := cmd.Flags().GetString("depends-on")
		dependencyOf, _ := cmd.Flags().GetString("dependency-of")

		if taskID == "" {
			output.PrintError("VALIDATION_ERROR", "--task is required")
			return &exitError{code: 1}
		}
		if dependsOn == "" && dependencyOf == "" {
			output.PrintError("VALIDATION_ERROR", "--depends-on or --dependency-of is required")
			return &exitError{code: 1}
		}

		if err := client.DeleteDependency(ctx, taskID, dependsOn, dependencyOf); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Manage task links",
}

var linkAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a link between tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		taskID, _ := cmd.Flags().GetString("task")
		linksTo, _ := cmd.Flags().GetString("links-to")

		if taskID == "" || linksTo == "" {
			output.PrintError("VALIDATION_ERROR", "--task and --links-to are required")
			return &exitError{code: 1}
		}

		resp, err := client.AddTaskLink(ctx, taskID, linksTo)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var linkRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a link between tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		taskID, _ := cmd.Flags().GetString("task")
		linksTo, _ := cmd.Flags().GetString("links-to")

		if taskID == "" || linksTo == "" {
			output.PrintError("VALIDATION_ERROR", "--task and --links-to are required")
			return &exitError{code: 1}
		}

		if err := client.DeleteTaskLink(ctx, taskID, linksTo); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

func init() {
	taskCmd.AddCommand(dependencyCmd)
	dependencyCmd.AddCommand(dependencyAddCmd, dependencyRemoveCmd)

	dependencyAddCmd.Flags().String("task", "", "Task ID (required)")
	dependencyAddCmd.Flags().String("depends-on", "", "Task ID this task depends on")
	dependencyAddCmd.Flags().String("dependency-of", "", "Task ID that depends on this task")
	dependencyAddCmd.Flags().String("type", "", "Dependency type")

	dependencyRemoveCmd.Flags().String("task", "", "Task ID (required)")
	dependencyRemoveCmd.Flags().String("depends-on", "", "Task ID this task depends on")
	dependencyRemoveCmd.Flags().String("dependency-of", "", "Task ID that depends on this task")

	taskCmd.AddCommand(linkCmd)
	linkCmd.AddCommand(linkAddCmd, linkRemoveCmd)

	linkAddCmd.Flags().String("task", "", "Task ID (required)")
	linkAddCmd.Flags().String("links-to", "", "Task ID to link to (required)")

	linkRemoveCmd.Flags().String("task", "", "Task ID (required)")
	linkRemoveCmd.Flags().String("links-to", "", "Task ID to unlink (required)")
}
