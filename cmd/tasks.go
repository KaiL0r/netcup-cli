package cmd

import (
	"github.com/spf13/cobra"

	"github.com/KaiL0r/netcup-cli/api"
)

// ======================
// Commands
// ======================

var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Manage tasks",
}

var tasksListCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		limit, _ := cmd.Flags().GetInt("limit")
		offset, _ := cmd.Flags().GetInt("offset")
		q, _ := cmd.Flags().GetString("q")
		serverid, _ := cmd.Flags().GetInt("serverid")
		state, _ := cmd.Flags().GetString("state")

		params := api.ListTasksParams{
			Limit:    limit,
			Offset:   offset,
			Q:        q,
			ServerID: serverid,
			State:    state,
		}

		tasks, err := apiClient.ListTasks(params)
		if err != nil {
			return err
		}

		return printJSON(tasks)
	},
}

var tasksGetCmd = &cobra.Command{
	Use:   "get <uuid>",
	Short: "Get a task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		uuid := args[0]

		task, err := apiClient.GetTask(uuid)
		if err != nil {
			return err
		}

		return printJSON(task)
	},
}

var tasksCancelCmd = &cobra.Command{
	Use:   "cancel <uuid>",
	Short: "Cancel a task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		uuid := args[0]

		err := apiClient.CancelTask(uuid)
		if err != nil {
			return err
		}

		return nil
	},
}

// ======================
// INIT
// ======================

func init() {
	tasksListCmd.Flags().Int("limit", 10, "Limit")
	tasksListCmd.Flags().Int("offset", 0, "Offset")
	tasksListCmd.Flags().String("q", "", "Search query")
	tasksListCmd.Flags().Int("serverid", 0, "Filter by server ID")
	tasksListCmd.Flags().String("state", "", "Task state")

	tasksCmd.AddCommand(tasksListCmd)
	tasksCmd.AddCommand(tasksGetCmd)
	tasksCmd.AddCommand(tasksCancelCmd)

	rootCmd.AddCommand(tasksCmd)
}
