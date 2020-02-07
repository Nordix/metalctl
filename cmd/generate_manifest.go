package cmd

import (
	"fmt"
	"metalctl/pkg/client"
	"metalctl/pkg/utils"

	"github.com/spf13/cobra"
)

// NewDeployTargetClusterCmd represents the deployTargetCluster command
func NewGenerateManifestCmd() *cobra.Command {
	o := &client.GenerateOptions{}
	generateManifestCmd := &cobra.Command{
		Use:   "generate-manifest",
		Short: "Generate manifests",
		Long: utils.LongDesc(`
				TODO`),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGenerate(o)
		},
	}
	initFlags(o, generateManifestCmd)
	return generateManifestCmd
}

func initFlags(o *client.GenerateOptions, cmd *cobra.Command) {
	flags := cmd.Flags()
	flags.StringVar(&o.SourcePath, "sourcePath", "", "TODO")
	cmd.MarkFlagRequired("sourcePath")
	flags.StringVar(&o.OutputPath, "outputPath", "", "TODO")
	cmd.MarkFlagRequired("outputPath")
}

func runGenerate(o *client.GenerateOptions) error {
	c, err := client.New(cfgFile)
	if err != nil {
		return err
	}

	fmt.Println("Generating manifests...")

	c.Generate(client.GenerateOptions{
		SourcePath: o.SourcePath,
		OutputPath: o.OutputPath,
	})
	if err != nil {
		return err
	}
	return nil
}
