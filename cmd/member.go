package cmd

import (
	"context"

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
		wid := getWorkspaceID(cmd)
		resp, err := client.GetGroups(ctx, wid)
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

		req := &api.CreateGroupRequest{Name: name, Handle: handle}
		resp, err := client.CreateGroup(ctx, wid, req)
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
		req := &api.GuestPermissionRequest{PermissionLevel: permLevel}
		resp, err := client.AddGuestToTask(ctx, taskID, guestID, req)
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
		if err := client.RemoveGuestFromTask(ctx, taskID, guestID); err != nil {
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
		req := &api.GuestPermissionRequest{PermissionLevel: permLevel}
		resp, err := client.AddGuestToList(ctx, listID, guestID, req)
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
		if err := client.RemoveGuestFromList(ctx, listID, guestID); err != nil {
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
		req := &api.GuestPermissionRequest{PermissionLevel: permLevel}
		resp, err := client.AddGuestToFolder(ctx, folderID, guestID, req)
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
		if err := client.RemoveGuestFromFolder(ctx, folderID, guestID); err != nil {
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

	groupCmd.AddCommand(groupListCmd, groupCreateCmd, groupDeleteCmd)
	groupCreateCmd.Flags().String("name", "", "Group name (required)")
	groupCreateCmd.Flags().String("handle", "", "Group handle")
	groupDeleteCmd.Flags().String("id", "", "Group ID (required)")

	guestCmd.AddCommand(guestInviteCmd, guestGetCmd, guestRemoveCmd)
	guestCmd.AddCommand(guestAddToTaskCmd, guestRemoveFromTaskCmd)
	guestCmd.AddCommand(guestAddToListCmd, guestRemoveFromListCmd)
	guestCmd.AddCommand(guestAddToFolderCmd, guestRemoveFromFolderCmd)
	guestInviteCmd.Flags().String("email", "", "Guest email (required)")
	guestGetCmd.Flags().String("id", "", "Guest ID (required)")
	guestRemoveCmd.Flags().String("id", "", "Guest ID (required)")

	guestAddToTaskCmd.Flags().String("task", "", "Task ID (required)")
	guestAddToTaskCmd.Flags().Int("guest-id", 0, "Guest ID (required)")
	guestAddToTaskCmd.Flags().String("permission-level", "read", "Permission level")
	guestRemoveFromTaskCmd.Flags().String("task", "", "Task ID (required)")
	guestRemoveFromTaskCmd.Flags().Int("guest-id", 0, "Guest ID (required)")

	guestAddToListCmd.Flags().String("list", "", "List ID (required)")
	guestAddToListCmd.Flags().Int("guest-id", 0, "Guest ID (required)")
	guestAddToListCmd.Flags().String("permission-level", "read", "Permission level")
	guestRemoveFromListCmd.Flags().String("list", "", "List ID (required)")
	guestRemoveFromListCmd.Flags().Int("guest-id", 0, "Guest ID (required)")

	guestAddToFolderCmd.Flags().String("folder", "", "Folder ID (required)")
	guestAddToFolderCmd.Flags().Int("guest-id", 0, "Guest ID (required)")
	guestAddToFolderCmd.Flags().String("permission-level", "read", "Permission level")
	guestRemoveFromFolderCmd.Flags().String("folder", "", "Folder ID (required)")
	guestRemoveFromFolderCmd.Flags().Int("guest-id", 0, "Guest ID (required)")
}
