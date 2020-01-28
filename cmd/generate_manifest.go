package cmd

import (
	"github.com/spf13/cobra"
)

type Options struct {
	sourcePath string
	outputPath string
}

// NewDeployTargetClusterCmd represents the deployTargetCluster command
func NewGenerateManifestCmd() *cobra.Command {
	o := &Options{}
	generateManifestCmd := &cobra.Command{
		Use:   "generate-manifest",
		Short: "Generate manifests",
		Long: LongDesc(`
				TODO`),
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.RunBuild()
		},
	}
	o.initFlags(generateManifestCmd)
	return generateManifestCmd
}

func (o *Options) initFlags(cmd *cobra.Command) {
	flags := cmd.Flags()
	flags.StringVar(&o.sourcePath, "sourcePath", "", "TODO")
	cmd.MarkFlagRequired("sourcePath")
	flags.StringVar(&o.outputPath, "outputPath", "", "TODO")
	cmd.MarkFlagRequired("outputPath")
}

func (o *Options) RunBuild() error {
	return nil
}
