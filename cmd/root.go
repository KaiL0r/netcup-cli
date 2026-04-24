package cmd

import (
	"os"

	"github.com/KaiL0r/netcup-cli/api"
	"github.com/spf13/cobra"
)

var apiClient *api.Client

var rootCmd = &cobra.Command{
	Use:   "netcup-cli",
	Short: "CLI for Netcup SCP API",
}

func RootCmd() *cobra.Command {
	return rootCmd
}

func Execute(c *api.Client) {
	apiClient = c

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
