package cmd

import (
	io_pkg "io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"sigs.k8s.io/kustomize/api/filesys"
	"sigs.k8s.io/kustomize/api/konfig"
	"sigs.k8s.io/kustomize/api/krusty"
	"sigs.k8s.io/kustomize/api/resmap"
	"sigs.k8s.io/kustomize/api/resource"
	//"sigs.k8s.io/kustomize/inter"
	"sigs.k8s.io/yaml"
)

type Options struct {
	kustomizationPath string
	outputPath        string
	outOrder          reorderOutput
}

// NewOptions creates a Options object
func NewManifestOptions(p, o string) *Options {
	return &Options{
		kustomizationPath: p,
		outputPath:        o,
	}
}

var o Options
var generateManifestCmd = &cobra.Command{
	Use:   "manifest",
	Short: "Generate manifests",
	Long: LongDesc(`
			TODO`),
	Example: Examples(`
			TODO`),
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := o.Validate(args)
		if err != nil {
			return err
		}
		stdOut := os.Stdout
		return o.RunBuild(stdOut)
	},
}

func (o *Options) RunBuild(out io_pkg.Writer) error {
	fSys := filesys.MakeFsOnDisk()

	k := krusty.MakeKustomizer(fSys, o.makeOptions())
	m, err := k.Run(o.kustomizationPath)
	if err != nil {
		return err
	}
	return o.emitResources(out, fSys, m)
}

func init() {
	generateManifestCmd.Flags().StringVarP(&gtp.sourcePath, "sourcePath", "s", "", "TODO")

	genCmd.AddCommand(generateManifestCmd)
}

// Validate validates build command.
func (o *Options) Validate(args []string) (err error) {
	if len(args) > 2 {
		return errors.New(
			"specify one path to " +
				konfig.DefaultKustomizationFileName())
	}
	if len(args) == 0 {
		o.kustomizationPath = filesys.SelfDir
	} else {
		o.kustomizationPath = args[0]
	}
	err = validateFlagLoadRestrictor()
	if err != nil {
		return err
	}
	o.outOrder, err = validateFlagReorderOutput()
	return
}

func (o *Options) makeOptions() *krusty.Options {
	opts := &krusty.Options{
		DoLegacyResourceSort: o.outOrder == legacy,
		LoadRestrictions:     getFlagLoadRestrictorValue(),
		DoPrune:              false,
	}
	if isFlagEnablePluginsSet() {
		c, err := konfig.EnabledPluginConfig()
		if err != nil {
			log.Fatal(err)
		}
		opts.PluginConfig = c
	} else {
		opts.PluginConfig = konfig.DisabledPluginConfig()
	}
	return opts
}

func (o *Options) emitResources(
	out io_pkg.Writer, fSys filesys.FileSystem, m resmap.ResMap) error {
	if o.outputPath != "" && fSys.IsDir(o.outputPath) {
		return writeIndividualFiles(fSys, o.outputPath, m)
	}
	res, err := m.AsYaml()
	if err != nil {
		return err
	}
	if o.outputPath != "" {
		return fSys.WriteFile(o.outputPath, res)
	}
	_, err = out.Write(res)
	return err
}

func writeIndividualFiles(
	fSys filesys.FileSystem, folderPath string, m resmap.ResMap) error {
	byNamespace := m.GroupedByCurrentNamespace()
	for namespace, resList := range byNamespace {
		for _, res := range resList {
			fName := fileName(res)
			if len(byNamespace) > 1 {
				fName = strings.ToLower(namespace) + "_" + fName
			}
			err := writeFile(fSys, folderPath, fName, res)
			if err != nil {
				return err
			}
		}
	}
	for _, res := range m.NonNamespaceable() {
		err := writeFile(fSys, folderPath, fileName(res), res)
		if err != nil {
			return err
		}
	}
	return nil
}

func fileName(res *resource.Resource) string {
	return strings.ToLower(res.GetGvk().String()) +
		"_" + strings.ToLower(res.GetName()) + ".yaml"
}

func writeFile(
	fSys filesys.FileSystem, path, fName string, res *resource.Resource) error {
	out, err := yaml.Marshal(res.Map())
	if err != nil {
		return err
	}
	return fSys.WriteFile(filepath.Join(path, fName), out)
}

//func runGenerateTemplate(name string) error {
//	fmt.Printf("%s\n", name)
//	return nil
//}
