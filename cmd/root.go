package cmd

import (
	"fmt"
	"os"

	"github.com/KaiL0r/netcup-cli/internal/api"
	"github.com/spf13/cobra"
)

// ======================
// CLIENT INJECTION
// ======================

var newClient = api.MustClient
var apiClient *api.Client

var rootCmd = &cobra.Command{
	Use:   "netcup-cli",
	Short: "CLI for Netcup SCP API",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// initialize shared client for all subcommands
		if apiClient == nil {
			apiClient = newClient()
			if apiClient == nil {
				return fmt.Errorf("failed to create API client")
			}

			// check token initially
			apiClient.Auth.GetAccessToken()
		}
		return nil
	},
}

func RootCmd() *cobra.Command {
	return rootCmd
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
