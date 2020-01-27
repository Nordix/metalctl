package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type generateTemplateOptions struct {
	sourcePath string
	//flavor                 string
	//bootstrapProvider      string
	//infrastructureProvider string

	//targetNamespace   string
	//kubernetesVersion string
	//controlplaneCount int
	//workerCount       int
}

var gtp = &generateTemplateOptions{}

var generateTemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "Generate templates for management and target cluster",
	Long: LongDesc(`
		TODO`),

	Example: Examples(`
		TODO`),

	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runGenerateTemplate(args[0])
	},
}

func init() {
	generateTemplateCmd.Flags().StringVarP(&gtp.sourcePath, "sourcePath", "s", "", "TODO")

	genCmd.AddCommand(generateTemplateCmd)
}

func runGenerateTemplate(name string) error {
	fmt.Printf("%s\n", name)
	return nil
}
