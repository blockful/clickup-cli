package cmd

import (
	"context"

	"github.com/blockful/clickup-cli/internal/api"
	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage workspace users",
}

var userInviteCmd = &cobra.Command{
	Use:   "invite",
	Short: "Invite a user to workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		email, _ := cmd.Flags().GetString("email")
		if email == "" {
			output.PrintError("VALIDATION_ERROR", "--email is required")
			return &exitError{code: 1}
		}
		req := &api.InviteUserRequest{Email: email}
		if cmd.Flags().Changed("admin") {
			v, _ := cmd.Flags().GetBool("admin")
			req.Admin = api.BoolPtr(v)
		}
		if cmd.Flags().Changed("custom-role-id") {
			v, _ := cmd.Flags().GetInt("custom-role-id")
			req.CustomRoleID = api.IntPtr(v)
		}
		if cmd.Flags().Changed("member-groups") {
			v, _ := cmd.Flags().GetIntSlice("member-groups")
			req.MemberGroups = v
		}
		resp, err := client.InviteUser(ctx, wid, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var userGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a user",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		resp, err := client.GetTeamUser(ctx, wid, id)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var userUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Edit a user",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		req := &api.EditUserRequest{}
		if cmd.Flags().Changed("username") {
			req.Username, _ = cmd.Flags().GetString("username")
		}
		if cmd.Flags().Changed("admin") {
			v, _ := cmd.Flags().GetBool("admin")
			req.Admin = api.BoolPtr(v)
		}
		if cmd.Flags().Changed("custom-role-id") {
			v, _ := cmd.Flags().GetInt("custom-role-id")
			req.CustomRoleID = api.IntPtr(v)
		}
		resp, err := client.EditUser(ctx, wid, id, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var userRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a user from workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		if err := client.RemoveUser(ctx, wid, id); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(userInviteCmd, userGetCmd, userUpdateCmd, userRemoveCmd)

	userInviteCmd.Flags().String("email", "", "User email (required)")
	userInviteCmd.Flags().Bool("admin", false, "Make user admin")
	userInviteCmd.Flags().Int("custom-role-id", 0, "Custom role ID")
	userInviteCmd.Flags().IntSlice("member-groups", nil, "Member group IDs")

	userGetCmd.Flags().String("id", "", "User ID (required)")

	userUpdateCmd.Flags().String("id", "", "User ID (required)")
	userUpdateCmd.Flags().String("username", "", "New username")
	userUpdateCmd.Flags().Bool("admin", false, "Admin status")
	userUpdateCmd.Flags().Int("custom-role-id", 0, "Custom role ID")

	userRemoveCmd.Flags().String("id", "", "User ID (required)")
}
