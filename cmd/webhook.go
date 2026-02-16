package cmd

import (
	"context"
	"strconv"

	"github.com/blockful/clickup-cli/internal/api"
	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
)

var webhookCmd = &cobra.Command{
	Use:   "webhook",
	Short: "Manage webhooks",
}

var webhookListCmd = &cobra.Command{
	Use:   "list",
	Short: "List webhooks",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		resp, err := client.GetWebhooks(ctx, wid)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var webhookCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a webhook",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		wid := getWorkspaceID(cmd)
		endpoint, _ := cmd.Flags().GetString("endpoint")
		events, _ := cmd.Flags().GetStringSlice("events")

		if endpoint == "" || len(events) == 0 {
			output.PrintError("VALIDATION_ERROR", "--endpoint and --events are required")
			return &exitError{code: 1}
		}

		req := &api.CreateWebhookRequest{Endpoint: endpoint, Events: events}
		if s, _ := cmd.Flags().GetString("space-id"); s != "" {
			if v, err := strconv.Atoi(s); err == nil {
				req.SpaceID = &v
			}
		}
		if s, _ := cmd.Flags().GetString("folder-id"); s != "" {
			if v, err := strconv.Atoi(s); err == nil {
				req.FolderID = &v
			}
		}
		if s, _ := cmd.Flags().GetString("list-id"); s != "" {
			if v, err := strconv.Atoi(s); err == nil {
				req.ListID = &v
			}
		}
		if s, _ := cmd.Flags().GetString("task-id"); s != "" {
			req.TaskID = &s
		}
		resp, err := client.CreateWebhook(ctx, wid, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var webhookUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a webhook",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		id, _ := cmd.Flags().GetString("id")
		endpoint, _ := cmd.Flags().GetString("endpoint")
		events, _ := cmd.Flags().GetString("events")
		status, _ := cmd.Flags().GetString("status")

		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}

		req := &api.UpdateWebhookRequest{Endpoint: endpoint, Events: events, Status: status}
		resp, err := client.UpdateWebhook(ctx, id, req)
		if err != nil {
			return handleError(err)
		}
		output.JSON(resp)
		return nil
	},
}

var webhookDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a webhook",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			output.PrintError("VALIDATION_ERROR", "--id is required")
			return &exitError{code: 1}
		}
		if err := client.DeleteWebhook(ctx, id); err != nil {
			return handleError(err)
		}
		output.JSON(map[string]string{"status": "ok"})
		return nil
	},
}

func init() {
	rootCmd.AddCommand(webhookCmd)
	webhookCmd.AddCommand(webhookListCmd, webhookCreateCmd, webhookUpdateCmd, webhookDeleteCmd)

	webhookCreateCmd.Flags().String("endpoint", "", "Webhook URL (required)")
	webhookCreateCmd.Flags().StringSlice("events", nil, "Events to subscribe to (required)")
	webhookCreateCmd.Flags().String("space-id", "", "Space ID to scope webhook")
	webhookCreateCmd.Flags().String("folder-id", "", "Folder ID to scope webhook")
	webhookCreateCmd.Flags().String("list-id", "", "List ID to scope webhook")
	webhookCreateCmd.Flags().String("task-id", "", "Task ID to scope webhook")

	webhookUpdateCmd.Flags().String("id", "", "Webhook ID (required)")
	webhookUpdateCmd.Flags().String("endpoint", "", "Webhook URL")
	webhookUpdateCmd.Flags().String("events", "", "Events (use * for all)")
	webhookUpdateCmd.Flags().String("status", "", "Status (active/inactive)")

	webhookDeleteCmd.Flags().String("id", "", "Webhook ID (required)")
}
