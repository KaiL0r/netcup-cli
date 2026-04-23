package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authorize netcup-cli with API credentials (through OAuth)",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {

		if _, err := apiClient.Auth.GetAccessToken(); err != nil {
			if _, err := apiClient.Auth.DeviceFlow(); err != nil {
				fmt.Println(fmt.Errorf("Authentication failed: %v", err))
			}
		}

		fmt.Println("OK")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
