package cmd

import (
	"strconv"

	"github.com/blockful/clickup-cli/internal/api"
	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
)

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage tasks",
}

var taskListCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks in a list",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		listID, _ := cmd.Flags().GetString("list")
		if listID == "" {
			output.PrintError("VALIDATION_ERROR", "--list is required")
			return &exitError{code: 1}
		}

		opts := &api.ListTasksOptions{}
		opts.Statuses, _ = cmd.Flags().GetStringSlice("status")
		opts.Assignees, _ = cmd.Flags().GetStringSlice("assignee")
		opts.Tags, _ = cmd.Flags().GetStringSlice("tag")
		opts.Page, _ = cmd.Flags().GetInt("page")
		opts.OrderBy, _ = cmd.Flags().GetString("order-by")
		opts.Reverse, _ = cmd.Flags().GetBool("reverse")
		opts.Subtasks, _ = cmd.Flags().GetBool("subtasks")
		opts.IncludeClosed, _ = cmd.Flags().GetBool("include-closed")

		resp, err := client.ListTasks(listID, opts)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var taskGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a task by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		resp, err := client.GetTask(id)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var taskCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new task",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		listID, _ := cmd.Flags().GetString("list")
		if listID == "" {
			output.PrintError("VALIDATION_ERROR", "--list is required")
			return &exitError{code: 1}
		}
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			output.PrintError("VALIDATION_ERROR", "--name is required")
			return &exitError{code: 1}
		}

		req := &api.CreateTaskRequest{Name: name}
		req.Description, _ = cmd.Flags().GetString("description")
		req.Status, _ = cmd.Flags().GetString("status")

		tags, _ := cmd.Flags().GetStringSlice("tag")
		if len(tags) > 0 {
			req.Tags = tags
		}

		assigneeStr, _ := cmd.Flags().GetString("assignee")
		if assigneeStr != "" {
			if id, err := strconv.Atoi(assigneeStr); err == nil {
				req.Assignees = []int{id}
			}
		}

		priority, _ := cmd.Flags().GetInt("priority")
		if priority > 0 {
			req.Priority = &priority
		}

		resp, err := client.CreateTask(listID, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var taskUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}

		req := &api.UpdateTaskRequest{}
		if cmd.Flags().Changed("name") {
			v, _ := cmd.Flags().GetString("name")
			req.Name = api.StringPtr(v)
		}
		if cmd.Flags().Changed("description") {
			v, _ := cmd.Flags().GetString("description")
			req.Description = api.StringPtr(v)
		}
		if cmd.Flags().Changed("status") {
			v, _ := cmd.Flags().GetString("status")
			req.Status = api.StringPtr(v)
		}
		if cmd.Flags().Changed("priority") {
			v, _ := cmd.Flags().GetInt("priority")
			req.Priority = api.IntPtr(v)
		}

		resp, err := client.UpdateTask(id, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var taskDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		if err := client.DeleteTask(id); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"message": "task deleted", "id": id})
		return nil
	},
}

func init() {
	taskListCmd.Flags().String("list", "", "List ID")
	taskListCmd.Flags().StringSlice("status", nil, "Filter by status")
	taskListCmd.Flags().StringSlice("assignee", nil, "Filter by assignee")
	taskListCmd.Flags().StringSlice("tag", nil, "Filter by tag")
	taskListCmd.Flags().Int("page", 0, "Page number")
	taskListCmd.Flags().String("order-by", "", "Order by field")
	taskListCmd.Flags().Bool("reverse", false, "Reverse order")
	taskListCmd.Flags().Bool("subtasks", false, "Include subtasks")
	taskListCmd.Flags().Bool("include-closed", false, "Include closed tasks")

	taskGetCmd.Flags().String("id", "", "Task ID")

	taskCreateCmd.Flags().String("list", "", "List ID")
	taskCreateCmd.Flags().String("name", "", "Task name")
	taskCreateCmd.Flags().String("description", "", "Task description")
	taskCreateCmd.Flags().String("assignee", "", "Assignee user ID")
	taskCreateCmd.Flags().String("status", "", "Task status")
	taskCreateCmd.Flags().Int("priority", 0, "Priority (1=urgent, 2=high, 3=normal, 4=low)")
	taskCreateCmd.Flags().StringSlice("tag", nil, "Task tags")

	taskUpdateCmd.Flags().String("id", "", "Task ID")
	taskUpdateCmd.Flags().String("name", "", "Task name")
	taskUpdateCmd.Flags().String("description", "", "Task description")
	taskUpdateCmd.Flags().String("status", "", "Task status")
	taskUpdateCmd.Flags().Int("priority", 0, "Priority (1=urgent, 2=high, 3=normal, 4=low)")

	taskDeleteCmd.Flags().String("id", "", "Task ID")

	taskCmd.AddCommand(taskListCmd)
	taskCmd.AddCommand(taskGetCmd)
	taskCmd.AddCommand(taskCreateCmd)
	taskCmd.AddCommand(taskUpdateCmd)
	taskCmd.AddCommand(taskDeleteCmd)
	rootCmd.AddCommand(taskCmd)
}
