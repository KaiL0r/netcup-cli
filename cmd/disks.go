package cmd

import (
	"github.com/KaiL0r/netcup-cli/internal/api"
	"github.com/spf13/cobra"
)

// ======================
// Commands
// ======================

var disksCmd = &cobra.Command{
	Use:   "disks",
	Short: "Manage disks of a server",
}

var disksListCmd = &cobra.Command{
	Use:   "list <serverId>",
	Short: "List disks of a server",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		serverId, err := parseID(args[0])
		if err != nil {
			return err
		}

		disks, err := apiClient.ListDisks(serverId)
		if err != nil {
			return err
		}

		return printJSON(disks)
	},
}

var disksGetCmd = &cobra.Command{
	Use:   "get <serverId> <diskName>",
	Short: "Get a disk of a server",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		serverId, err := parseID(args[0])
		if err != nil {
			return err
		}

		disk, err := apiClient.GetDisk(serverId, args[1])
		if err != nil {
			return err
		}

		return printJSON(disk)
	},
}

var disksListSupportedDriversCmd = &cobra.Command{
	Use:   "list-drivers <serverId>",
	Short: "List available storage drivers of a server",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		serverId, err := parseID(args[0])
		if err != nil {
			return err
		}

		disks, err := apiClient.ListDiskSupportedDrivers(serverId)
		if err != nil {
			return err
		}

		return printJSON(disks)
	},
}

var disksSetDriversCmd = &cobra.Command{
	Use:   "set-driver <serverId> <driver>",
	Short: "Sets storage driver of a server",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		serverId, err := parseID(args[0])
		if err != nil {
			return err
		}

		disks, err := apiClient.UpdateDiskDriver(serverId, api.StorageDriver(args[1]))
		if err != nil {
			return err
		}

		return printJSON(disks)
	},
}

var disksFormatCmd = &cobra.Command{
	Use:   "format <serverId> <diskName>",
	Short: "Formats disk of a server (WARNING: This will delete all data on the disk!)",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		serverId, err := parseID(args[0])
		if err != nil {
			return err
		}

		task, err := apiClient.FormatDisk(serverId, args[1])
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
	disksCmd.AddCommand(disksListCmd)
	disksCmd.AddCommand(disksGetCmd)
	disksCmd.AddCommand(disksListSupportedDriversCmd)
	disksCmd.AddCommand(disksSetDriversCmd)
	disksCmd.AddCommand(disksFormatCmd)

	serversCmd.AddCommand(disksCmd)
}
