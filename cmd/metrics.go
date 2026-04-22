package cmd

import (
	"github.com/spf13/cobra"
)

// ======================
// Commands
// ======================

var metricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Get metrics of a server",
}

var metricsCpuCmd = &cobra.Command{
	Use:   "cpu <serverId>",
	Short: "Get CPU metrics of a server",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		serverId, err := parseID(args[0])
		if err != nil {
			return err
		}
		hours, _ := cmd.Flags().GetInt("hours")

		disks, err := apiClient.MetricsCpu(serverId, hours)
		if err != nil {
			return err
		}

		return printJSON(disks)
	},
}

var metricsDiskCmd = &cobra.Command{
	Use:   "disk <serverId>",
	Short: "Get disk metrics of a server",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		serverId, err := parseID(args[0])
		if err != nil {
			return err
		}
		hours, _ := cmd.Flags().GetInt("hours")

		disks, err := apiClient.MetricsDisk(serverId, hours)
		if err != nil {
			return err
		}

		return printJSON(disks)
	},
}

var metricsNetworkCmd = &cobra.Command{
	Use:   "network <serverId>",
	Short: "Get network metrics of a server",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		serverId, err := parseID(args[0])
		if err != nil {
			return err
		}
		hours, _ := cmd.Flags().GetInt("hours")

		disks, err := apiClient.MetricsNetwork(serverId, hours)
		if err != nil {
			return err
		}

		return printJSON(disks)
	},
}

var metricsNetworkPacketsCmd = &cobra.Command{
	Use:   "network-packet <serverId>",
	Short: "Get network packet metrics of a server",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		serverId, err := parseID(args[0])
		if err != nil {
			return err
		}
		hours, _ := cmd.Flags().GetInt("hours")

		disks, err := apiClient.MetricsNetworkPackets(serverId, hours)
		if err != nil {
			return err
		}

		return printJSON(disks)
	},
}

// ======================
// INIT
// ======================

func init() {
	metricsCpuCmd.Flags().String("hours", "", "Number of hours to retrieve metrics for")
	metricsCmd.AddCommand(metricsCpuCmd)

	metricsDiskCmd.Flags().String("hours", "", "Number of hours to retrieve metrics for")
	metricsCmd.AddCommand(metricsDiskCmd)

	metricsNetworkCmd.Flags().String("hours", "", "Number of hours to retrieve metrics for")
	metricsCmd.AddCommand(metricsNetworkCmd)

	metricsNetworkPacketsCmd.Flags().String("hours", "", "Number of hours to retrieve metrics for")
	metricsCmd.AddCommand(metricsNetworkPacketsCmd)

	serversCmd.AddCommand(metricsCmd)
}
