package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func SetVersionInfo(v, c, d string) {
	version = v
	commit = c
	date = d
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		info := map[string]string{
			"version": version,
			"commit":  commit,
			"date":    date,
		}
		out, _ := json.Marshal(info)
		fmt.Println(string(out))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
