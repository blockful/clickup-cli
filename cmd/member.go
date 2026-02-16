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
	guestInviteCmd.Flags().String("email", "", "Guest email (required)")
	guestGetCmd.Flags().String("id", "", "Guest ID (required)")
	guestRemoveCmd.Flags().String("id", "", "Guest ID (required)")
}
