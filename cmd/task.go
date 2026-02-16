package cmd

import (
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
		opts.Watchers, _ = cmd.Flags().GetStringSlice("watchers")
		opts.Page, _ = cmd.Flags().GetInt("page")
		opts.OrderBy, _ = cmd.Flags().GetString("order-by")
		opts.Reverse, _ = cmd.Flags().GetBool("reverse")
		opts.Subtasks, _ = cmd.Flags().GetBool("subtasks")
		opts.IncludeClosed, _ = cmd.Flags().GetBool("include-closed")
		opts.Archived, _ = cmd.Flags().GetBool("archived")
		opts.IncludeMarkdown, _ = cmd.Flags().GetBool("include-markdown")
		opts.IncludeTiml, _ = cmd.Flags().GetBool("include-timl")
		opts.DueDateGt, _ = cmd.Flags().GetInt64("due-date-gt")
		opts.DueDateLt, _ = cmd.Flags().GetInt64("due-date-lt")
		opts.DateCreatedGt, _ = cmd.Flags().GetInt64("date-created-gt")
		opts.DateCreatedLt, _ = cmd.Flags().GetInt64("date-created-lt")
		opts.DateUpdatedGt, _ = cmd.Flags().GetInt64("date-updated-gt")
		opts.DateUpdatedLt, _ = cmd.Flags().GetInt64("date-updated-lt")
		opts.DateDoneGt, _ = cmd.Flags().GetInt64("date-done-gt")
		opts.DateDoneLt, _ = cmd.Flags().GetInt64("date-done-lt")
		opts.CustomFields, _ = cmd.Flags().GetString("custom-fields")
		opts.CustomItems, _ = cmd.Flags().GetIntSlice("custom-items")

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
		getOpts := api.GetTaskOptions{}
		getOpts.CustomTaskIDs, _ = cmd.Flags().GetBool("custom-task-ids")
		getOpts.TeamID, _ = cmd.Flags().GetString("team-id")
		getOpts.IncludeSubtasks, _ = cmd.Flags().GetBool("include-subtasks")
		getOpts.IncludeMarkdown, _ = cmd.Flags().GetBool("include-markdown")

		resp, err := client.GetTask(id, getOpts)
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
		req.MarkdownDescription, _ = cmd.Flags().GetString("markdown-description")
		req.Status, _ = cmd.Flags().GetString("status")
		req.Parent, _ = cmd.Flags().GetString("parent")
		req.LinksTo, _ = cmd.Flags().GetString("links-to")

		tags, _ := cmd.Flags().GetStringSlice("tag")
		if len(tags) > 0 {
			req.Tags = tags
		}

		assignees, _ := cmd.Flags().GetIntSlice("assignee")
		if len(assignees) > 0 {
			req.Assignees = assignees
		}

		priority, _ := cmd.Flags().GetInt("priority")
		if priority > 0 {
			req.Priority = &priority
		}

		if cmd.Flags().Changed("due-date") {
			v, _ := cmd.Flags().GetInt64("due-date")
			req.DueDate = api.Int64Ptr(v)
		}
		if cmd.Flags().Changed("due-date-time") {
			v, _ := cmd.Flags().GetBool("due-date-time")
			req.DueDateTime = api.BoolPtr(v)
		}
		if cmd.Flags().Changed("start-date") {
			v, _ := cmd.Flags().GetInt64("start-date")
			req.StartDate = api.Int64Ptr(v)
		}
		if cmd.Flags().Changed("start-date-time") {
			v, _ := cmd.Flags().GetBool("start-date-time")
			req.StartDateTime = api.BoolPtr(v)
		}
		if cmd.Flags().Changed("time-estimate") {
			v, _ := cmd.Flags().GetInt64("time-estimate")
			req.TimeEstimate = api.Int64Ptr(v)
		}
		req.NotifyAll, _ = cmd.Flags().GetBool("notify-all")

		cfStr, _ := cmd.Flags().GetString("custom-fields")
		if cfStr != "" {
			fields, err := api.ParseCustomFields(cfStr)
			if err != nil {
				output.PrintError("VALIDATION_ERROR", err.Error())
				return &exitError{code: 1}
			}
			req.CustomFields = fields
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
		if cmd.Flags().Changed("due-date") {
			v, _ := cmd.Flags().GetInt64("due-date")
			req.DueDate = api.Int64Ptr(v)
		}
		if cmd.Flags().Changed("due-date-time") {
			v, _ := cmd.Flags().GetBool("due-date-time")
			req.DueDateTime = api.BoolPtr(v)
		}
		if cmd.Flags().Changed("start-date") {
			v, _ := cmd.Flags().GetInt64("start-date")
			req.StartDate = api.Int64Ptr(v)
		}
		if cmd.Flags().Changed("start-date-time") {
			v, _ := cmd.Flags().GetBool("start-date-time")
			req.StartDateTime = api.BoolPtr(v)
		}
		if cmd.Flags().Changed("time-estimate") {
			v, _ := cmd.Flags().GetInt64("time-estimate")
			req.TimeEstimate = api.Int64Ptr(v)
		}
		if cmd.Flags().Changed("archived") {
			v, _ := cmd.Flags().GetBool("archived")
			req.Archived = api.BoolPtr(v)
		}
		if cmd.Flags().Changed("parent") {
			v, _ := cmd.Flags().GetString("parent")
			req.Parent = api.StringPtr(v)
		}

		addAssignees, _ := cmd.Flags().GetIntSlice("assignees-add")
		remAssignees, _ := cmd.Flags().GetIntSlice("assignees-rem")
		if len(addAssignees) > 0 || len(remAssignees) > 0 {
			req.Assignees = &api.UpdateTaskAssignees{Add: addAssignees, Rem: remAssignees}
		}

		updateOpts := api.UpdateTaskOptions{}
		updateOpts.CustomTaskIDs, _ = cmd.Flags().GetBool("custom-task-ids")
		updateOpts.TeamID, _ = cmd.Flags().GetString("team-id")

		resp, err := client.UpdateTask(id, req, updateOpts)
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

var taskSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search tasks across a workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		teamID := getWorkspaceID(cmd)

		opts := &api.SearchTasksOptions{}
		opts.Statuses, _ = cmd.Flags().GetStringSlice("status")
		opts.Assignees, _ = cmd.Flags().GetStringSlice("assignee")
		opts.Tags, _ = cmd.Flags().GetStringSlice("tag")
		opts.Page, _ = cmd.Flags().GetInt("page")
		opts.OrderBy, _ = cmd.Flags().GetString("order-by")
		opts.Reverse, _ = cmd.Flags().GetBool("reverse")
		opts.Subtasks, _ = cmd.Flags().GetBool("subtasks")
		opts.IncludeClosed, _ = cmd.Flags().GetBool("include-closed")
		opts.IncludeMarkdown, _ = cmd.Flags().GetBool("include-markdown")
		opts.DueDateGt, _ = cmd.Flags().GetInt64("due-date-gt")
		opts.DueDateLt, _ = cmd.Flags().GetInt64("due-date-lt")
		opts.DateCreatedGt, _ = cmd.Flags().GetInt64("date-created-gt")
		opts.DateCreatedLt, _ = cmd.Flags().GetInt64("date-created-lt")
		opts.DateUpdatedGt, _ = cmd.Flags().GetInt64("date-updated-gt")
		opts.DateUpdatedLt, _ = cmd.Flags().GetInt64("date-updated-lt")
		opts.DateDoneGt, _ = cmd.Flags().GetInt64("date-done-gt")
		opts.DateDoneLt, _ = cmd.Flags().GetInt64("date-done-lt")
		opts.CustomFields, _ = cmd.Flags().GetString("custom-fields")
		opts.CustomItems, _ = cmd.Flags().GetIntSlice("custom-items")
		opts.ListIDs, _ = cmd.Flags().GetStringSlice("list-ids")
		opts.ProjectIDs, _ = cmd.Flags().GetStringSlice("project-ids")
		opts.SpaceIDs, _ = cmd.Flags().GetStringSlice("space-ids")
		opts.FolderIDs, _ = cmd.Flags().GetStringSlice("folder-ids")

		resp, err := client.SearchTasks(teamID, opts)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

func init() {
	// task list
	taskListCmd.Flags().String("list", "", "List ID")
	taskListCmd.Flags().StringSlice("status", nil, "Filter by status")
	taskListCmd.Flags().StringSlice("assignee", nil, "Filter by assignee")
	taskListCmd.Flags().StringSlice("tag", nil, "Filter by tag")
	taskListCmd.Flags().StringSlice("watchers", nil, "Filter by watchers")
	taskListCmd.Flags().Int("page", 0, "Page number")
	taskListCmd.Flags().String("order-by", "", "Order by field")
	taskListCmd.Flags().Bool("reverse", false, "Reverse order")
	taskListCmd.Flags().Bool("subtasks", false, "Include subtasks")
	taskListCmd.Flags().Bool("include-closed", false, "Include closed tasks")
	taskListCmd.Flags().Bool("archived", false, "Filter archived tasks")
	taskListCmd.Flags().Bool("include-markdown", false, "Include markdown description")
	taskListCmd.Flags().Bool("include-timl", false, "Include tasks in multiple lists")
	taskListCmd.Flags().Int64("due-date-gt", 0, "Due date greater than (Unix ms)")
	taskListCmd.Flags().Int64("due-date-lt", 0, "Due date less than (Unix ms)")
	taskListCmd.Flags().Int64("date-created-gt", 0, "Date created greater than (Unix ms)")
	taskListCmd.Flags().Int64("date-created-lt", 0, "Date created less than (Unix ms)")
	taskListCmd.Flags().Int64("date-updated-gt", 0, "Date updated greater than (Unix ms)")
	taskListCmd.Flags().Int64("date-updated-lt", 0, "Date updated less than (Unix ms)")
	taskListCmd.Flags().Int64("date-done-gt", 0, "Date done greater than (Unix ms)")
	taskListCmd.Flags().Int64("date-done-lt", 0, "Date done less than (Unix ms)")
	taskListCmd.Flags().String("custom-fields", "", "Custom fields filter (JSON array)")
	taskListCmd.Flags().IntSlice("custom-items", nil, "Filter by task type (custom item IDs)")

	// task get
	taskGetCmd.Flags().String("id", "", "Task ID")
	taskGetCmd.Flags().Bool("custom-task-ids", false, "Use custom task IDs")
	taskGetCmd.Flags().String("team-id", "", "Team ID (required when custom-task-ids=true)")
	taskGetCmd.Flags().Bool("include-subtasks", false, "Include subtasks")
	taskGetCmd.Flags().Bool("include-markdown", false, "Include markdown description")

	// task create
	taskCreateCmd.Flags().String("list", "", "List ID")
	taskCreateCmd.Flags().String("name", "", "Task name")
	taskCreateCmd.Flags().String("description", "", "Task description")
	taskCreateCmd.Flags().String("markdown-description", "", "Task description in markdown")
	taskCreateCmd.Flags().IntSlice("assignee", nil, "Assignee user IDs")
	taskCreateCmd.Flags().String("status", "", "Task status")
	taskCreateCmd.Flags().Int("priority", 0, "Priority (1=urgent, 2=high, 3=normal, 4=low)")
	taskCreateCmd.Flags().StringSlice("tag", nil, "Task tags")
	taskCreateCmd.Flags().Int64("due-date", 0, "Due date (Unix ms)")
	taskCreateCmd.Flags().Bool("due-date-time", false, "Due date includes time")
	taskCreateCmd.Flags().Int64("start-date", 0, "Start date (Unix ms)")
	taskCreateCmd.Flags().Bool("start-date-time", false, "Start date includes time")
	taskCreateCmd.Flags().Int64("time-estimate", 0, "Time estimate (ms)")
	taskCreateCmd.Flags().Bool("notify-all", false, "Notify all assignees")
	taskCreateCmd.Flags().String("parent", "", "Parent task ID (for subtasks)")
	taskCreateCmd.Flags().String("links-to", "", "Task ID to link to")
	taskCreateCmd.Flags().String("custom-fields", "", "Custom fields (JSON array)")

	// task update
	taskUpdateCmd.Flags().String("id", "", "Task ID")
	taskUpdateCmd.Flags().String("name", "", "Task name")
	taskUpdateCmd.Flags().String("description", "", "Task description")
	taskUpdateCmd.Flags().String("status", "", "Task status")
	taskUpdateCmd.Flags().Int("priority", 0, "Priority (1=urgent, 2=high, 3=normal, 4=low)")
	taskUpdateCmd.Flags().IntSlice("assignees-add", nil, "Assignee IDs to add")
	taskUpdateCmd.Flags().IntSlice("assignees-rem", nil, "Assignee IDs to remove")
	taskUpdateCmd.Flags().Int64("due-date", 0, "Due date (Unix ms)")
	taskUpdateCmd.Flags().Bool("due-date-time", false, "Due date includes time")
	taskUpdateCmd.Flags().Int64("start-date", 0, "Start date (Unix ms)")
	taskUpdateCmd.Flags().Bool("start-date-time", false, "Start date includes time")
	taskUpdateCmd.Flags().Int64("time-estimate", 0, "Time estimate (ms)")
	taskUpdateCmd.Flags().Bool("archived", false, "Archive task")
	taskUpdateCmd.Flags().String("parent", "", "Parent task ID")
	taskUpdateCmd.Flags().Bool("custom-task-ids", false, "Use custom task IDs")
	taskUpdateCmd.Flags().String("team-id", "", "Team ID (required when custom-task-ids=true)")

	// task delete
	taskDeleteCmd.Flags().String("id", "", "Task ID")

	// task search
	taskSearchCmd.Flags().String("workspace", "", "Workspace/Team ID")
	taskSearchCmd.Flags().StringSlice("status", nil, "Filter by status")
	taskSearchCmd.Flags().StringSlice("assignee", nil, "Filter by assignee")
	taskSearchCmd.Flags().StringSlice("tag", nil, "Filter by tag")
	taskSearchCmd.Flags().Int("page", 0, "Page number")
	taskSearchCmd.Flags().String("order-by", "", "Order by field")
	taskSearchCmd.Flags().Bool("reverse", false, "Reverse order")
	taskSearchCmd.Flags().Bool("subtasks", false, "Include subtasks")
	taskSearchCmd.Flags().Bool("include-closed", false, "Include closed tasks")
	taskSearchCmd.Flags().Bool("include-markdown", false, "Include markdown description")
	taskSearchCmd.Flags().Int64("due-date-gt", 0, "Due date greater than (Unix ms)")
	taskSearchCmd.Flags().Int64("due-date-lt", 0, "Due date less than (Unix ms)")
	taskSearchCmd.Flags().Int64("date-created-gt", 0, "Date created greater than (Unix ms)")
	taskSearchCmd.Flags().Int64("date-created-lt", 0, "Date created less than (Unix ms)")
	taskSearchCmd.Flags().Int64("date-updated-gt", 0, "Date updated greater than (Unix ms)")
	taskSearchCmd.Flags().Int64("date-updated-lt", 0, "Date updated less than (Unix ms)")
	taskSearchCmd.Flags().Int64("date-done-gt", 0, "Date done greater than (Unix ms)")
	taskSearchCmd.Flags().Int64("date-done-lt", 0, "Date done less than (Unix ms)")
	taskSearchCmd.Flags().String("custom-fields", "", "Custom fields filter (JSON array)")
	taskSearchCmd.Flags().IntSlice("custom-items", nil, "Filter by task type")
	taskSearchCmd.Flags().StringSlice("list-ids", nil, "Filter by list IDs")
	taskSearchCmd.Flags().StringSlice("project-ids", nil, "Filter by project/folder IDs")
	taskSearchCmd.Flags().StringSlice("space-ids", nil, "Filter by space IDs")
	taskSearchCmd.Flags().StringSlice("folder-ids", nil, "Filter by folder IDs")

	taskCmd.AddCommand(taskListCmd)
	taskCmd.AddCommand(taskGetCmd)
	taskCmd.AddCommand(taskCreateCmd)
	taskCmd.AddCommand(taskUpdateCmd)
	taskCmd.AddCommand(taskDeleteCmd)
	taskCmd.AddCommand(taskSearchCmd)
	rootCmd.AddCommand(taskCmd)
}
