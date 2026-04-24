package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/KaiL0r/netcup-cli/api"
)

// ======================
// Root command
// ======================

var serversCmd = &cobra.Command{
	Use:   "servers",
	Short: "Manage servers",
}

// ======================
// LIST
// ======================

var serversListCmd = &cobra.Command{
	Use:   "list",
	Short: "List servers",
	RunE: func(cmd *cobra.Command, args []string) error {

		limit, _ := cmd.Flags().GetInt("limit")
		offset, _ := cmd.Flags().GetInt("offset")
		name, _ := cmd.Flags().GetString("name")
		ip, _ := cmd.Flags().GetString("ip")
		query, _ := cmd.Flags().GetString("query")

		servers, err := apiClient.ListServers(limit, offset, name, ip, query)
		if err != nil {
			return err
		}

		return printJSON(servers)
	},
}

// ======================
// UPDATE root
// ======================

var serversUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update server attributes",
}

// ======================
// UPDATE: STATE
// ======================

var serversUpdateStateCmd = &cobra.Command{
	Use:   "state <id> <state>",
	Short: "Set server state (ON | OFF | SUSPENDED)",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}
		state := api.SetServerState(strings.ToUpper(args[1]))

		stateOptionStr, _ := cmd.Flags().GetString("state-option")
		stateOption := api.SetServerStateOption(strings.ToUpper(stateOptionStr))

		task, err := apiClient.UpdateServerState(id, state, stateOption)

		if err != nil {
			return err
		}
		return printJSON(task)
	},
}

// ======================
// UPDATE: AUTOSTART
// ======================

var serversUpdateAutostartCmd = &cobra.Command{
	Use:   "autostart <id> <bool>",
	Short: "Set server autostart",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}
		value, err := parseBoolArg(args[1])
		if err != nil {
			return err
		}

		task, err := apiClient.UpdateServerAutostart(id, value)

		if err != nil {
			return err
		}
		return printJSON(task)
	},
}

// ======================
// UPDATE: BOOTORDER
// ======================

var serversUpdateBootorderCmd = &cobra.Command{
	Use:   "bootorder <id> <HDD|CDROM|NETWORK> ...",
	Short: "Set boot order",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}

		var bootorder []api.BootOrder
		for _, v := range args[1:] {
			bootorder = append(bootorder, api.BootOrder(strings.ToUpper(v)))
		}

		task, err := apiClient.UpdateServerBootorder(id, bootorder)

		if err != nil {
			return err
		}
		return printJSON(task)
	},
}

// ======================
// UPDATE: OS OPTIMIZATION
// ======================

var serversUpdateOSCmd = &cobra.Command{
	Use:   "os <id> <type>",
	Short: "Set OS optimization (LINUX | WINDOWS | BSD | LINUX_LEGACY | UNKNOWN)",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}
		osOptimization := api.OsOptimization(strings.ToUpper(args[1]))

		task, err := apiClient.UpdateServerOsoptimization(id, osOptimization)
		if err != nil {
			return err
		}
		return printJSON(task)
	},
}

// ======================
// UPDATE: CPU TOPOLOGY
// ======================

var serversUpdateCPUCmd = &cobra.Command{
	Use:   "cpu <id> <sockets> <cores>",
	Short: "Set CPU topology",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}
		sockets, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("invalid sockets %q: %w", args[1], err)
		}
		cores, err := strconv.Atoi(args[2])
		if err != nil {
			return fmt.Errorf("invalid cores %q: %w", args[2], err)
		}

		task, err := apiClient.UpdateServerCpuTopology(id, sockets, cores)
		if err != nil {
			return err
		}
		return printJSON(task)
	},
}

// ======================
// UPDATE: UEFI
// ======================

var serversUpdateUEFICmd = &cobra.Command{
	Use:   "uefi <id> <bool>",
	Short: "Enable/disable UEFI",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}
		val, err := parseBoolArg(args[1])
		if err != nil {
			return err
		}

		task, err := apiClient.UpdateServerUefi(id, val)
		if err != nil {
			return err
		}
		return printJSON(task)
	},
}

// ======================
// UPDATE: HOSTNAME
// ======================

var serversUpdateHostnameCmd = &cobra.Command{
	Use:   "hostname <id> <hostname>",
	Short: "Set hostname",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}

		task, err := apiClient.UpdateServerHostname(id, args[1])
		if err != nil {
			return err
		}
		return printJSON(task)
	},
}

// ======================
// UPDATE: NICKNAME
// ======================

var serversUpdateNicknameCmd = &cobra.Command{
	Use:   "nickname <id> <nickname>",
	Short: "Set nickname",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}

		_, err = apiClient.UpdateServerNickname(id, args[1])
		if err != nil {
			return err
		}
		// return printJSON(task)
		return nil
	},
}

// ======================
// UPDATE: KEYBOARD LAYOUT
// ======================

var serversUpdateKeyboardLayoutCmd = &cobra.Command{
	Use:   "keyboard-layout <id> <layout>",
	Short: "Set keyboard layout (en-us, de, ...) wrong layout will result in 'en-us'",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}
		task, err := apiClient.UpdateServerKeyboardLayout(id, args[1])
		if err != nil {
			return err
		}
		return printJSON(task)
	},
}

// ======================
// UPDATE: ROOT PASSWORD
// ======================

var serversUpdatePasswordCmd = &cobra.Command{
	Use:   "root-password <id> <password>",
	Short: "Set root password",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}
		task, err := apiClient.UpdateServerRootPassword(id, args[1])
		if err != nil {
			return err
		}
		return printJSON(task)
	},
}

// ======================
// GET
// ======================

var serversGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get one server",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}

		servers, err := apiClient.GetServer(id)
		if err != nil {
			return err
		}

		return printJSON(servers)
	},
}

// ======================
// GET GPU DRIVER
// ======================

var serversGetGpuDriverCmd = &cobra.Command{
	Use:   "gpu-drivers <id>",
	Short: "Generate presigned download URL for GPU driver if available",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}

		servers, err := apiClient.GetServerGpuDriver(id)
		if err != nil {
			return err
		}

		return printJSON(servers)
	},
}

// ======================
// GET GUEST AGENT
// ======================

var serversGetGuestAgentCmd = &cobra.Command{
	Use:   "guest-agent <id>",
	Short: "Get guest agent data for server",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}

		servers, err := apiClient.GetServerGuestAgent(id)
		if err != nil {
			return err
		}

		return printJSON(servers)
	},
}

// ======================
// GET LOGS
// ======================

var serversLogsCmd = &cobra.Command{
	Use:   "logs <serverId>",
	Short: "Get server logs",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}

		limit, _ := cmd.Flags().GetInt("limit")
		offset, _ := cmd.Flags().GetInt("offset")

		logs, err := apiClient.ListServerLogs(id, limit, offset)
		if err != nil {
			return err
		}

		return printJSON(logs)
	},
}

// ======================
// RESCUE SYSTEM
// ======================

var serversRescueCmd = &cobra.Command{
	Use:   "rescue-system",
	Short: "Manage server rescue system",
}

var serversRescueGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get rescue system status for server",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}

		status, err := apiClient.GetServerRescueSystem(id)
		if err != nil {
			return err
		}

		return printJSON(status)
	},
}

var serversRescueActivateCmd = &cobra.Command{
	Use:   "activate <id>",
	Short: "Activate rescue system for server",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}

		task, err := apiClient.ActivateServerRescueSystem(id)
		if err != nil {
			return err
		}
		return printJSON(task)
	},
}

var serversRescueDeactivateCmd = &cobra.Command{
	Use:   "deactivate <id>",
	Short: "Deactivate rescue system for server",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}

		task, err := apiClient.DeactivateServerRescueSystem(id)
		if err != nil {
			return err
		}
		return printJSON(task)
	},
}

// ======================
// STORAGE OPTIMIZE
// ======================

var serversStorageOptimizeCmd = &cobra.Command{
	Use:   "storage-optimize <id>",
	Short: "Optimize storage of a server",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}
		disks, err := cmd.Flags().GetStringArray("disks")
		if err != nil {
			return err
		}
		startAfter, err := cmd.Flags().GetBool("start-after")
		if err != nil {
			return err
		}

		task, err := apiClient.OptimizeServerStorage(id, disks, startAfter)
		if err != nil {
			return err
		}
		return printJSON(task)
	},
}

// ======================
// INIT
// ======================

func init() {
	// add "list" flags
	serversListCmd.Flags().Int("limit", 10, "Limit")
	serversListCmd.Flags().Int("offset", 0, "Offset")
	serversListCmd.Flags().String("name", "", "Filter by name")
	serversListCmd.Flags().String("ip", "", "Filter by IP")
	serversListCmd.Flags().String("query", "", "Search query")

	// add "update state" flag
	serversUpdateStateCmd.Flags().String("state-option", "", "Valid values for state ON: POWERCYCLE, RESET. Valid values for state OFF: POWEROFF")

	serversLogsCmd.Flags().Int("limit", 10, "Limit")
	serversLogsCmd.Flags().Int("offset", 0, "Offset")

	// update commands
	serversUpdateCmd.AddCommand(
		serversUpdateStateCmd,
		serversUpdateAutostartCmd,
		serversUpdateBootorderCmd,
		serversUpdateOSCmd,
		serversUpdateCPUCmd,
		serversUpdateUEFICmd,
		serversUpdateHostnameCmd,
		serversUpdateNicknameCmd,
		serversUpdateKeyboardLayoutCmd,
		serversUpdatePasswordCmd,
		serversStorageOptimizeCmd,
	)

	// storage-optimize flags
	serversStorageOptimizeCmd.Flags().StringArray("disks", []string{}, "List of disk identifiers to optimize (can be provided multiple times)")
	serversStorageOptimizeCmd.Flags().Bool("start-after", true, "Start server after optimization")

	serversCmd.AddCommand(serversListCmd)
	serversCmd.AddCommand(serversUpdateCmd)
	serversCmd.AddCommand(serversGetCmd)
	serversCmd.AddCommand(serversGetGpuDriverCmd)
	serversCmd.AddCommand(serversGetGuestAgentCmd)
	serversCmd.AddCommand(serversLogsCmd)

	// rescue-system subcommands
	serversRescueCmd.AddCommand(serversRescueGetCmd, serversRescueActivateCmd, serversRescueDeactivateCmd)
	serversCmd.AddCommand(serversRescueCmd)

	rootCmd.AddCommand(serversCmd)
}
