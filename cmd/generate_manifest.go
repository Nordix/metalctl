package cmd

import (
	"fmt"

	"metalctl/pkg/client"

	"github.com/spf13/cobra"
)

type generateOptions struct {
	sourcePath string
	outputPath string
}

// NewDeployTargetClusterCmd represents the deployTargetCluster command
func NewGenerateManifestCmd() *cobra.Command {
	o := &generateOptions{}
	generateManifestCmd := &cobra.Command{
		Use:   "generate-manifest",
		Short: "Generate manifests",
		Long: LongDesc(`
				TODO`),
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.runGenerate()
		},
	}
	o.initFlags(generateManifestCmd)
	return generateManifestCmd
}

func (o *generateOptions) initFlags(cmd *cobra.Command) {
	flags := cmd.Flags()
	flags.StringVar(&o.sourcePath, "sourcePath", "", "TODO")
	cmd.MarkFlagRequired("sourcePath")
	flags.StringVar(&o.outputPath, "outputPath", "", "TODO")
	cmd.MarkFlagRequired("outputPath")
}

func (o *generateOptions) runGenerate() error {
	c, err := client.New(cfgFile)
	if err != nil {
		return err
	}

	fmt.Println("Generating manifests...")

	c.Generate(client.GenerateOptions{
		SourcePath: o.sourcePath,
		OutputPath: o.outputPath,
	})
	if err != nil {
		return err
	}
	return nil
}
