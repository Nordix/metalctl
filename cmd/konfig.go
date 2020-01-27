package cmd

import (
	"github.com/spf13/cobra"
)

var konfigCmd = &cobra.Command{
	Use:   "konfig",
	Short: "Display Cluster API provider configuration and templates for creating workload clusters",
	Long:  `Display Cluster API provider configuration and templates for creating workload clusters`,
}

func init() {
	RootCmd.AddCommand(konfigCmd)
}
