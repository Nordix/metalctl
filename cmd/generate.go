package cmd

import (
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "generate",
	Short: "REPLACE THIS",
	Long:  `REPLACE THIS`,
}

func init() {
	RootCmd.AddCommand(genCmd)
}
