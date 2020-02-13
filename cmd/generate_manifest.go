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
	//TODO: define defaults when complete architecture is clearer
	flags.StringVar(&o.SourcePath, "sourcePath", "source", "TODO")
	flags.StringVar(&o.OutputPath, "outputPath", "/tmp/manifests", "TODO")
	flags.StringVar(&o.Basepath, "basePath", "/baseresourcesexamples", "TODO")
	flags.StringVar(&o.GeneratePath, "generatePath", "/tmp/templates", "TODO")
	flags.StringVar(&o.KustomizePath, "kustomizePath", "/baseresourcesexamples/rbac", "TODO")
	flags.StringVar(&o.FileSystemType, "filesystemtype", "internal", "TODO")
}

func runGenerate(o *client.GenerateOptions) error {
	c, err := client.New(cfgFile)
	if err != nil {
		return err
	}

	fmt.Println("Generating manifests...")

	err = c.Generate(client.GenerateOptions{
		SourcePath:     o.SourcePath,
		OutputPath:     o.OutputPath,
		Basepath:       o.Basepath,
		GeneratePath:   o.GeneratePath,
		KustomizePath:  o.KustomizePath,
		FileSystemType: o.FileSystemType,
	})
	if err != nil {
		return err
	}
	return nil
}
