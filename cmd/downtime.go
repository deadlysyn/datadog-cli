package cmd

import "github.com/spf13/cobra"

var downtimeCmd = &cobra.Command{
	Use:   "downtime",
	Short: "List & modify downtime",
}

func init() {
	RootCmd.AddCommand(downtimeCmd)
}
