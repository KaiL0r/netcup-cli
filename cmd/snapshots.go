package cmd

import (
	"github.com/spf13/cobra"

	"github.com/KaiL0r/netcup-cli/api"
)

var snapshotsCmd = &cobra.Command{
	Use:   "snapshots",
	Short: "Manage server snapshots",
}

var snapshotsListCmd = &cobra.Command{
	Use:   "list <serverId>",
	Short: "List snapshots for a server",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}

		snaps, err := apiClient.ListServerSnapshots(id)
		if err != nil {
			return err
		}
		return printJSON(snaps)
	},
}

var snapshotsGetCmd = &cobra.Command{
	Use:   "get <serverId> <name>",
	Short: "Get a snapshot",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}
		snap, err := apiClient.GetServerSnapshot(id, args[1])
		if err != nil {
			return err
		}
		return printJSON(snap)
	},
}

var snapshotsCreateCmd = &cobra.Command{
	Use:   "create <serverId> <name>",
	Short: "Create a snapshot for a server",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}

		create := api.SnapshotCreate{Name: args[1]}

		if cmd.Flags().Changed("description") {
			if v, _ := cmd.Flags().GetString("description"); v != "" {
				create.Description = &v
			}
		}
		if cmd.Flags().Changed("diskname") {
			if v, _ := cmd.Flags().GetString("diskname"); v != "" {
				create.Diskname = &v
			}
		}
		if cmd.Flags().Changed("online") {
			if v, err := cmd.Flags().GetBool("online"); err == nil {
				create.OnlineSnapshot = &v
			}
		}

		task, err := apiClient.CreateServerSnapshot(id, create)
		if err != nil {
			return err
		}
		return printJSON(task)
	},
}

var snapshotsDeleteCmd = &cobra.Command{
	Use:   "delete <serverId> <name>",
	Short: "Delete a snapshot",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}
		task, err := apiClient.DeleteServerSnapshot(id, args[1])
		if err != nil {
			return err
		}
		return printJSON(task)
	},
}

var snapshotsExportCmd = &cobra.Command{
	Use:   "export <serverId> <name>",
	Short: "Export a snapshot",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}
		task, err := apiClient.ExportServerSnapshot(id, args[1])
		if err != nil {
			return err
		}
		return printJSON(task)
	},
}

var snapshotsRevertCmd = &cobra.Command{
	Use:   "revert <serverId> <name>",
	Short: "Revert a snapshot",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseID(args[0])
		if err != nil {
			return err
		}
		task, err := apiClient.RevertServerSnapshot(id, args[1])
		if err != nil {
			return err
		}
		return printJSON(task)
	},
}

func init() {
	snapshotsCreateCmd.Flags().String("description", "", "")
	snapshotsCreateCmd.Flags().String("diskname", "", "The name of the disk to snapshot, must be set if attribute onlineSnapshot is not set")
	snapshotsCreateCmd.Flags().Bool("online", true, "Whether to create an online snapshot")

	snapshotsCmd.AddCommand(snapshotsListCmd, snapshotsGetCmd, snapshotsCreateCmd, snapshotsDeleteCmd, snapshotsExportCmd, snapshotsRevertCmd)
	snapshotsListCmd.Flags().Int("limit", 10, "Limit")
	snapshotsListCmd.Flags().Int("offset", 0, "Offset")
	serversCmd.AddCommand(snapshotsCmd)
}
