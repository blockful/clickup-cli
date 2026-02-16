package cmd

import (
	"context"
	"encoding/json"

	"github.com/blockful/clickup-cli/internal/api"
	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
)

var memberCmd = &cobra.Command{
	Use:   "member",
	Short: "Manage members",
}

var memberListCmd = &cobra.Command{
	Use:   "list",
	Short: "List members of a list or task",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		listID, _ := cmd.Flags().GetString("list")
		taskID, _ := cmd.Flags().GetString("task")

		var resp *api.MembersResponse
		var err error

		switch {
		case taskID != "":
			resp, err = client.GetTaskMembers(ctx, taskID)
		case listID != "":
			resp, err = client.GetListMembers(ctx, listID)
		default:
			output.PrintError("VALIDATION_ERROR", "--list or --task is required")
			return &exitError{code: 1}
		}
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Manage user groups",
}

var groupListCmd = &cobra.Command{
	Use:   "list",
	Short: "List groups",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		teamID, _ := cmd.Flags().GetString("team-id")
		if teamID == "" {
			teamID = getWorkspaceID(cmd)
		}
		groupIDs, _ := cmd.Flags().GetStringSlice("group-ids")
		resp, err := client.GetGroups(ctx, teamID, groupIDs)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var groupCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a group",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		name, _ := cmd.Flags().GetString("name")
		handle, _ := cmd.Flags().GetString("handle")

		if name == "" {
			output.PrintError("VALIDATION_ERROR", "--name is required")
			return &exitError{code: 1}
		}

		members, _ := cmd.Flags().GetIntSlice("members")
		req := &api.CreateGroupRequest{Name: name, Handle: handle, Members: members}
		resp, err := client.CreateGroup(ctx, wid, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var groupUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a group",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		req := &api.UpdateGroupRequest{}
		if cmd.Flags().Changed("name") {
			req.Name, _ = cmd.Flags().GetString("name")
		}
		if cmd.Flags().Changed("handle") {
			req.Handle, _ = cmd.Flags().GetString("handle")
		}
		if cmd.Flags().Changed("members") {
			membersJSON, _ := cmd.Flags().GetString("members")
			var m struct {
				Add []int `json:"add"`
				Rem []int `json:"rem"`
			}
			if err := json.Unmarshal([]byte(membersJSON), &m); err != nil {
				output.PrintError("VALIDATION_ERROR", "invalid --members JSON: "+err.Error())
				return &exitError{code: 1}
			}
			req.Members = &struct {
				Add []int `json:"add,omitempty"`
				Rem []int `json:"rem,omitempty"`
			}{Add: m.Add, Rem: m.Rem}
		}
		resp, err := client.UpdateGroup(ctx, id, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var groupDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a group",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		if err := client.DeleteGroup(ctx, id); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

var guestCmd = &cobra.Command{
	Use:   "guest",
	Short: "Manage guests",
}

var guestInviteCmd = &cobra.Command{
	Use:   "invite",
	Short: "Invite a guest to workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		email, _ := cmd.Flags().GetString("email")
		if email == "" {
			output.PrintError("VALIDATION_ERROR", "--email is required")
			return &exitError{code: 1}
		}
		req := &api.InviteGuestRequest{Email: email}
		if cmd.Flags().Changed("can-edit-tags") {
			v, _ := cmd.Flags().GetBool("can-edit-tags"); req.CanEditTags = &v
		}
		if cmd.Flags().Changed("can-see-time-spent") {
			v, _ := cmd.Flags().GetBool("can-see-time-spent"); req.CanSeeTimeSpent = &v
		}
		if cmd.Flags().Changed("can-see-time-estimated") {
			v, _ := cmd.Flags().GetBool("can-see-time-estimated"); req.CanSeeTimeEstimated = &v
		}
		if cmd.Flags().Changed("can-create-views") {
			v, _ := cmd.Flags().GetBool("can-create-views"); req.CanCreateViews = &v
		}
		if cmd.Flags().Changed("can-see-points-estimated") {
			v, _ := cmd.Flags().GetBool("can-see-points-estimated"); req.CanSeePointsEstimated = &v
		}
		if cmd.Flags().Changed("custom-role-id") {
			v, _ := cmd.Flags().GetInt("custom-role-id"); req.CustomRoleID = &v
		}
		if err := client.InviteGuest(ctx, wid, req); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

var guestGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a guest",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		resp, err := client.GetGuest(ctx, wid, id)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var guestRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a guest from workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		if err := client.RemoveGuest(ctx, wid, id); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

var guestEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a guest",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		req := &api.EditGuestRequest{}
		if cmd.Flags().Changed("can-edit-tags") {
			v, _ := cmd.Flags().GetBool("can-edit-tags"); req.CanEditTags = &v
		}
		if cmd.Flags().Changed("can-see-time-spent") {
			v, _ := cmd.Flags().GetBool("can-see-time-spent"); req.CanSeeTimeSpent = &v
		}
		if cmd.Flags().Changed("can-see-time-estimated") {
			v, _ := cmd.Flags().GetBool("can-see-time-estimated"); req.CanSeeTimeEstimated = &v
		}
		if cmd.Flags().Changed("can-create-views") {
			v, _ := cmd.Flags().GetBool("can-create-views"); req.CanCreateViews = &v
		}
		if cmd.Flags().Changed("can-see-points-estimated") {
			v, _ := cmd.Flags().GetBool("can-see-points-estimated"); req.CanSeePointsEstimated = &v
		}
		if cmd.Flags().Changed("custom-role-id") {
			v, _ := cmd.Flags().GetInt("custom-role-id"); req.CustomRoleID = &v
		}
		resp, err := client.EditGuest(ctx, wid, id, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var guestAddToTaskCmd = &cobra.Command{
	Use:   "add-to-task",
	Short: "Add a guest to a task",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		taskID, _ := cmd.Flags().GetString("task")
		guestID, _ := cmd.Flags().GetInt("guest-id")
		permLevel, _ := cmd.Flags().GetString("permission-level")
		if taskID == "" {
			output.PrintError("VALIDATION_ERROR", "--task is required")
			return &exitError{code: 1}
		}
		if guestID == 0 {
			output.PrintError("VALIDATION_ERROR", "--guest-id is required")
			return &exitError{code: 1}
		}
		includeShared, _ := cmd.Flags().GetBool("include-shared")
		req := &api.GuestPermissionRequest{PermissionLevel: permLevel}
		resp, err := client.AddGuestToTask(ctx, taskID, guestID, req, includeShared, getTaskScopedOpts(cmd))
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var guestRemoveFromTaskCmd = &cobra.Command{
	Use:   "remove-from-task",
	Short: "Remove a guest from a task",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		taskID, _ := cmd.Flags().GetString("task")
		guestID, _ := cmd.Flags().GetInt("guest-id")
		if taskID == "" {
			output.PrintError("VALIDATION_ERROR", "--task is required")
			return &exitError{code: 1}
		}
		if guestID == 0 {
			output.PrintError("VALIDATION_ERROR", "--guest-id is required")
			return &exitError{code: 1}
		}
		includeShared, _ := cmd.Flags().GetBool("include-shared")
		if err := client.RemoveGuestFromTask(ctx, taskID, guestID, includeShared, getTaskScopedOpts(cmd)); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

var guestAddToListCmd = &cobra.Command{
	Use:   "add-to-list",
	Short: "Add a guest to a list",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		listID, _ := cmd.Flags().GetString("list")
		guestID, _ := cmd.Flags().GetInt("guest-id")
		permLevel, _ := cmd.Flags().GetString("permission-level")
		if listID == "" {
			output.PrintError("VALIDATION_ERROR", "--list is required")
			return &exitError{code: 1}
		}
		if guestID == 0 {
			output.PrintError("VALIDATION_ERROR", "--guest-id is required")
			return &exitError{code: 1}
		}
		includeShared, _ := cmd.Flags().GetBool("include-shared")
		req := &api.GuestPermissionRequest{PermissionLevel: permLevel}
		resp, err := client.AddGuestToList(ctx, listID, guestID, req, includeShared)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var guestRemoveFromListCmd = &cobra.Command{
	Use:   "remove-from-list",
	Short: "Remove a guest from a list",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		listID, _ := cmd.Flags().GetString("list")
		guestID, _ := cmd.Flags().GetInt("guest-id")
		if listID == "" {
			output.PrintError("VALIDATION_ERROR", "--list is required")
			return &exitError{code: 1}
		}
		if guestID == 0 {
			output.PrintError("VALIDATION_ERROR", "--guest-id is required")
			return &exitError{code: 1}
		}
		includeShared, _ := cmd.Flags().GetBool("include-shared")
		if err := client.RemoveGuestFromList(ctx, listID, guestID, includeShared); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

var guestAddToFolderCmd = &cobra.Command{
	Use:   "add-to-folder",
	Short: "Add a guest to a folder",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		folderID, _ := cmd.Flags().GetString("folder")
		guestID, _ := cmd.Flags().GetInt("guest-id")
		permLevel, _ := cmd.Flags().GetString("permission-level")
		if folderID == "" {
			output.PrintError("VALIDATION_ERROR", "--folder is required")
			return &exitError{code: 1}
		}
		if guestID == 0 {
			output.PrintError("VALIDATION_ERROR", "--guest-id is required")
			return &exitError{code: 1}
		}
		includeShared, _ := cmd.Flags().GetBool("include-shared")
		req := &api.GuestPermissionRequest{PermissionLevel: permLevel}
		resp, err := client.AddGuestToFolder(ctx, folderID, guestID, req, includeShared)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var guestRemoveFromFolderCmd = &cobra.Command{
	Use:   "remove-from-folder",
	Short: "Remove a guest from a folder",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		folderID, _ := cmd.Flags().GetString("folder")
		guestID, _ := cmd.Flags().GetInt("guest-id")
		if folderID == "" {
			output.PrintError("VALIDATION_ERROR", "--folder is required")
			return &exitError{code: 1}
		}
		if guestID == 0 {
			output.PrintError("VALIDATION_ERROR", "--guest-id is required")
			return &exitError{code: 1}
		}
		includeShared, _ := cmd.Flags().GetBool("include-shared")
		if err := client.RemoveGuestFromFolder(ctx, folderID, guestID, includeShared); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

func init() {
	rootCmd.AddCommand(memberCmd, groupCmd, guestCmd)

	memberCmd.AddCommand(memberListCmd)
	memberListCmd.Flags().String("list", "", "List ID")
	memberListCmd.Flags().String("task", "", "Task ID")

	groupCmd.AddCommand(groupListCmd, groupCreateCmd, groupUpdateCmd, groupDeleteCmd)
	groupListCmd.Flags().String("team-id", "", "Team ID (overrides workspace)")
	groupListCmd.Flags().StringSlice("group-ids", nil, "Filter by group IDs")
	groupCreateCmd.Flags().String("name", "", "Group name (required)")
	groupCreateCmd.Flags().String("handle", "", "Group handle")
	groupCreateCmd.Flags().IntSlice("members", nil, "Member user IDs")
	groupUpdateCmd.Flags().String("id", "", "Group ID (required)")
	groupUpdateCmd.Flags().String("name", "", "New group name")
	groupUpdateCmd.Flags().String("handle", "", "New group handle")
	groupUpdateCmd.Flags().String("members", "", "Members JSON: {\"add\":[id,...],\"rem\":[id,...]}")
	groupDeleteCmd.Flags().String("id", "", "Group ID (required)")

	guestCmd.AddCommand(guestInviteCmd, guestGetCmd, guestEditCmd, guestRemoveCmd)
	guestCmd.AddCommand(guestAddToTaskCmd, guestRemoveFromTaskCmd)
	guestCmd.AddCommand(guestAddToListCmd, guestRemoveFromListCmd)
	guestCmd.AddCommand(guestAddToFolderCmd, guestRemoveFromFolderCmd)
	guestInviteCmd.Flags().String("email", "", "Guest email (required)")
	guestInviteCmd.Flags().Bool("can-edit-tags", false, "Can edit tags")
	guestInviteCmd.Flags().Bool("can-see-time-spent", false, "Can see time spent")
	guestInviteCmd.Flags().Bool("can-see-time-estimated", false, "Can see time estimated")
	guestInviteCmd.Flags().Bool("can-create-views", false, "Can create views")
	guestInviteCmd.Flags().Bool("can-see-points-estimated", false, "Can see points estimated")
	guestInviteCmd.Flags().Int("custom-role-id", 0, "Custom role ID")

	guestEditCmd.Flags().String("id", "", "Guest ID (required)")
	guestEditCmd.Flags().Bool("can-edit-tags", false, "Can edit tags")
	guestEditCmd.Flags().Bool("can-see-time-spent", false, "Can see time spent")
	guestEditCmd.Flags().Bool("can-see-time-estimated", false, "Can see time estimated")
	guestEditCmd.Flags().Bool("can-create-views", false, "Can create views")
	guestEditCmd.Flags().Bool("can-see-points-estimated", false, "Can see points estimated")
	guestEditCmd.Flags().Int("custom-role-id", 0, "Custom role ID")
	guestGetCmd.Flags().String("id", "", "Guest ID (required)")
	guestRemoveCmd.Flags().String("id", "", "Guest ID (required)")

	guestAddToTaskCmd.Flags().String("task", "", "Task ID (required)")
	guestAddToTaskCmd.Flags().Int("guest-id", 0, "Guest ID (required)")
	guestAddToTaskCmd.Flags().String("permission-level", "read", "Permission level")
	guestAddToTaskCmd.Flags().Bool("include-shared", false, "Include shared items in response")
	addTaskScopedFlags(guestAddToTaskCmd)
	guestRemoveFromTaskCmd.Flags().String("task", "", "Task ID (required)")
	guestRemoveFromTaskCmd.Flags().Int("guest-id", 0, "Guest ID (required)")
	guestRemoveFromTaskCmd.Flags().Bool("include-shared", false, "Include shared items in response")
	addTaskScopedFlags(guestRemoveFromTaskCmd)

	guestAddToListCmd.Flags().String("list", "", "List ID (required)")
	guestAddToListCmd.Flags().Int("guest-id", 0, "Guest ID (required)")
	guestAddToListCmd.Flags().String("permission-level", "read", "Permission level")
	guestAddToListCmd.Flags().Bool("include-shared", false, "Include shared items in response")
	guestRemoveFromListCmd.Flags().String("list", "", "List ID (required)")
	guestRemoveFromListCmd.Flags().Int("guest-id", 0, "Guest ID (required)")
	guestRemoveFromListCmd.Flags().Bool("include-shared", false, "Include shared items in response")

	guestAddToFolderCmd.Flags().String("folder", "", "Folder ID (required)")
	guestAddToFolderCmd.Flags().Int("guest-id", 0, "Guest ID (required)")
	guestAddToFolderCmd.Flags().String("permission-level", "read", "Permission level")
	guestAddToFolderCmd.Flags().Bool("include-shared", false, "Include shared items in response")
	guestRemoveFromFolderCmd.Flags().String("folder", "", "Folder ID (required)")
	guestRemoveFromFolderCmd.Flags().Int("guest-id", 0, "Guest ID (required)")
	guestRemoveFromFolderCmd.Flags().Bool("include-shared", false, "Include shared items in response")
}
