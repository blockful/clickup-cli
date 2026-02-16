package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/blockful/clickup-cli/internal/api"
	"github.com/blockful/clickup-cli/internal/config"
	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage authentication",
}

var authLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate with ClickUp API token",
	RunE: func(cmd *cobra.Command, args []string) error {
		token, _ := cmd.Flags().GetString("token")

		if token == "" {
			fmt.Print("Enter your ClickUp API token: ")
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				output.PrintError("INPUT_ERROR", "failed to read token")
				return &exitError{code: 1}
			}
			token = strings.TrimSpace(input)
		}

		if token == "" {
			output.PrintError("VALIDATION_ERROR", "token cannot be empty")
			return &exitError{code: 1}
		}

		// Validate token by calling /v2/user
		client := api.NewClient(token)
		ctx := context.Background()
		user, err := client.GetUser(ctx)
		if err != nil {
			return handleError(err)
		}

		// Save token
		if err := config.SetToken(token); err != nil {
			output.PrintError("CONFIG_ERROR", fmt.Sprintf("failed to save token: %v", err))
			return &exitError{code: 1}
		}

		output.JSON(map[string]interface{}{
			"message": "authenticated successfully",
			"user":    user.User,
			"config":  config.ConfigFilePath(),
		})
		return nil
	},
}

var authWhoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Show current authenticated user",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()
		user, err := client.GetUser(ctx)
		if err != nil {
			return handleError(err)
		}
		output.JSON(user.User)
		return nil
	},
}

func init() {
	authLoginCmd.Flags().String("token", "", "API token (if not provided, will prompt)")
	authCmd.AddCommand(authLoginCmd)
	authCmd.AddCommand(authWhoamiCmd)
	rootCmd.AddCommand(authCmd)
}
