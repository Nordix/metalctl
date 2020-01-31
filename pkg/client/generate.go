package client

import (
	"path/filepath"
	"strings"

	"sigs.k8s.io/kustomize/api/filesys"
	"sigs.k8s.io/kustomize/api/krusty"
	"sigs.k8s.io/kustomize/api/resmap"
	"sigs.k8s.io/kustomize/api/resource"
	"sigs.k8s.io/yaml"
)

type GenOptions struct {
	sourcePath string
	outputPath string
}

func (c *metalctlClient) Generate(options GenerateOptions) error {
	fSys := filesys.MakeFsOnDisk()

	k := krusty.MakeKustomizer(fSys, makeOptions())
	m, err := k.Run(options.SourcePath)
	if err != nil {
		return err
	}
	writeResources(options.OutputPath, fSys, m)
	return nil
}

func makeOptions() *krusty.Options {
	opts := &krusty.Options{
		DoLegacyResourceSort: true,
		LoadRestrictions:     0,
		DoPrune:              false,
	}
	// TODO analyse if we want to use plugin usability control logic
	/* 	if build.IsFlagEnablePluginsSet() {
		c, err := konfig.EnabledPluginConfig()
		if err != nil {
			log.Fatal(err)
		}
		opts.PluginConfig = c
	} else {
		opts.PluginConfig = konfig.DisabledPluginConfig()
	} */
	return opts
}

func writeResources(
	outputPath string, fSys filesys.FileSystem, m resmap.ResMap) error {
	if outputPath != "" && fSys.IsDir(outputPath) {
		return writeIndividualFiles(fSys, outputPath, m)
	}
	res, err := m.AsYaml()
	if err != nil {
		return err
	}
	fSys.WriteFile(outputPath, res)
	return nil

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
