package cmd

import (
	"github.com/blockful/clickup-cli/internal/api"
	"github.com/blockful/clickup-cli/internal/config"
	"github.com/blockful/clickup-cli/internal/output"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "clickup",
	Short: "ClickUp CLI — optimized for AI agents",
	Long:  "A command-line interface for ClickUp, designed for AI agent workflows. All output is JSON by default.",
	SilenceUsage:  true,
	SilenceErrors: true,
}

// clientFactory can be overridden in tests to inject a mock client.
var clientFactory func() api.ClientInterface

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(config.Init)

	rootCmd.PersistentFlags().String("token", "", "ClickUp API token (overrides config)")
	rootCmd.PersistentFlags().String("workspace", "", "Default workspace ID (overrides config)")
	rootCmd.PersistentFlags().String("format", "json", "Output format: json or text")
	rootCmd.PersistentFlags().Bool("verbose", false, "Enable verbose output")

	_ = viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	_ = viper.BindPFlag("workspace", rootCmd.PersistentFlags().Lookup("workspace"))
}

func getClient() api.ClientInterface {
	if clientFactory != nil {
		return clientFactory()
	}
	token := config.GetToken()
	if token == "" {
		output.PrintErrorAndExit("AUTH_REQUIRED", "No API token configured. Run 'clickup auth login' first.", 2)
	}
	return api.NewClient(token)
}

func getWorkspaceID(cmd *cobra.Command) string {
	id, _ := cmd.Flags().GetString("workspace")
	if id == "" {
		id = config.GetWorkspace()
	}
	if id == "" {
		output.PrintErrorAndExit("WORKSPACE_REQUIRED", "Workspace ID required. Use --workspace flag or set default with config.", 1)
	}
	return id
}

func handleError(err error) error {
	if clientErr, ok := err.(*api.ClientError); ok {
		output.PrintError(clientErr.Code, clientErr.Message)
		if clientErr.Code == "UNAUTHORIZED" {
			return &exitError{code: 2}
		}
	} else {
		output.PrintError("ERROR", err.Error())
	}
	return &exitError{code: 1}
}

type exitError struct {
	code int
}

func (e *exitError) Error() string {
	return ""
}
