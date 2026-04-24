package cmd

import (
	"fmt"

	"github.com/KaiL0r/netcup-cli/api"
	"github.com/spf13/cobra"
)

// ======================
// USERS command
// ======================

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Manage users",
}

var usersGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get one user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}

		u, err := apiClient.GetUser(id)
		if err != nil {
			return err
		}

		return printJSON(u)
	},
}

var usersUpdateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Update a user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}

		var u api.UserSave

		if cmd.Flags().Changed("language") {
			if v, _ := cmd.Flags().GetString("language"); v != "" {
				language := api.UserLanguage(v)
				u.Language = &language
			}
		}
		if cmd.Flags().Changed("timezone") {
			if v, _ := cmd.Flags().GetString("timezone"); v != "" {
				u.TimeZone = &v
			}
		}
		if cmd.Flags().Changed("api-ip-login-restrictions") {
			if v, _ := cmd.Flags().GetString("api-ip-login-restrictions"); v != "" {
				u.ApiIpLoginRestrictions = &v
			}
		}
		if cmd.Flags().Changed("password") {
			if v, _ := cmd.Flags().GetString("password"); v != "" {
				u.Password = &v
			}
		}
		if cmd.Flags().Changed("old-password") {
			if v, _ := cmd.Flags().GetString("old-password"); v != "" {
				u.OldPassword = &v
			}
		}
		if cmd.Flags().Changed("soap-webservice-password") {
			if v, _ := cmd.Flags().GetString("soap-webservice-password"); v != "" {
				u.SoapWebservicePassword = &v
			}
		}

		if cmd.Flags().Changed("show-nickname") {
			v, _ := cmd.Flags().GetBool("show-nickname")
			u.ShowNickname = &v
		}
		if cmd.Flags().Changed("passwordless-mode") {
			v, _ := cmd.Flags().GetBool("passwordless-mode")
			u.PasswordlessMode = &v
		}
		if cmd.Flags().Changed("secure-mode") {
			v, _ := cmd.Flags().GetBool("secure-mode")
			u.SecureMode = &v
		}
		if cmd.Flags().Changed("secure-mode-app-access") {
			v, _ := cmd.Flags().GetBool("secure-mode-app-access")
			u.SecureModeAppAccess = &v
		}

		updated, err := apiClient.UpdateUser(id, u)
		if err != nil {
			return err
		}

		return printJSON(updated)
	},
}

var usersGetLogsCmd = &cobra.Command{
	Use:   "logs <id>",
	Short: "Get logs for a user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}
		limit, _ := cmd.Flags().GetInt("limit")
		offset, _ := cmd.Flags().GetInt("offset")

		u, err := apiClient.GetUserLogs(id, limit, offset)
		if err != nil {
			return err
		}

		return printJSON(u)
	},
}

var usersApiKeysCmd = &cobra.Command{
	Use:   "ssh-keys",
	Short: "Manage user's SSH keys",
}

var usersApiKeysListCmd = &cobra.Command{
	Use:   "list <userId>",
	Short: "List SSH keys for a user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}

		keys, err := apiClient.ListUserSSHKeys(id)
		if err != nil {
			return err
		}

		return printJSON(keys)
	},
}

var usersApiKeysCreateCmd = &cobra.Command{
	Use:   "create <userId> <name> <key>",
	Short: "Create API key for a user",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}
		name := args[1]
		key := args[2]

		sshKey, err := apiClient.CreateUserApiKey(id, name, key)
		if err != nil {
			return err
		}

		return printJSON(sshKey)
	},
}

var usersApiKeysDeleteCmd = &cobra.Command{
	Use:   "delete <userId> <keyId>",
	Short: "Delete an API key for a user",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}
		keyId, err := parseID(args[1])
		if err != nil {
			return err
		}

		if err := apiClient.DeleteUserApiKey(id, keyId); err != nil {
			return err
		}

		fmt.Printf("apikey %d for user %d deleted\n", keyId, id)
		return nil
	},
}

func init() {
	usersUpdateCmd.Flags().String("language", "", "Language (en|de)")
	usersUpdateCmd.Flags().String("timezone", "", "Timezone")
	usersUpdateCmd.Flags().String("api-ip-login-restrictions", "", "API IP login restrictions")
	usersUpdateCmd.Flags().String("password", "", "Set new password")
	usersUpdateCmd.Flags().String("old-password", "", "Current password (required when changing password)")
	usersUpdateCmd.Flags().String("soap-webservice-password", "", "SOAP webservice password")
	usersUpdateCmd.Flags().Bool("show-nickname", false, "Show nickname")
	usersUpdateCmd.Flags().Bool("passwordless-mode", false, "Passwordless mode")
	usersUpdateCmd.Flags().Bool("secure-mode", false, "Secure mode")
	usersUpdateCmd.Flags().Bool("secure-mode-app-access", false, "Secure mode app access")

	usersGetLogsCmd.Flags().Int("limit", 10, "Limit")
	usersGetLogsCmd.Flags().Int("offset", 0, "Offset")

	usersCmd.AddCommand(usersGetCmd)
	usersCmd.AddCommand(usersUpdateCmd)
	usersCmd.AddCommand(usersGetLogsCmd)

	usersApiKeysCmd.AddCommand(usersApiKeysListCmd, usersApiKeysCreateCmd, usersApiKeysDeleteCmd)
	usersCmd.AddCommand(usersApiKeysCmd)

	rootCmd.AddCommand(usersCmd)
}
