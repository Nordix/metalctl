package client

import (
	"fmt"
	"path/filepath"
	"strings"

	//"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/kustomize/api/filesys"
	"sigs.k8s.io/kustomize/api/krusty"
	"sigs.k8s.io/kustomize/api/resid"
	"sigs.k8s.io/kustomize/api/resmap"
	"sigs.k8s.io/kustomize/api/resource"
	"sigs.k8s.io/kustomize/api/types"
	"sigs.k8s.io/yaml"
)

type GenOptions struct {
	sourcePath string
	outputPath string
}
type Unstructured struct {
	// Object is a JSON compatible map with string, float, int, bool, []interface{}, or
	// map[string]interface{}
	// children.
	Object map[string]interface{}
}

func (c *metalctlClient) Generate(options GenerateOptions) error {
	fSys := filesys.MakeFsOnDisk()
	k := krusty.MakeKustomizer(fSys, makeOptions())
	m, err := k.Run(options.SourcePath)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	c.ResMap = m

	manifests, err := c.GetAllManifests()
	fmt.Println("List of manifests:")
	for _, manifest := range manifests {
		fmt.Println(manifest.GetName())
	}

	writeResources(options.OutputPath, fSys, c.ResMap)
	return nil
}

// GetAllManifests returns all Manifests converted from resource map to list
func (c *metalctlClient) GetAllManifests() ([]Manifest, error) {
	manifestSet := []Manifest{}
	for _, res := range c.ResMap.Resources() {
		manifest, err := NewManifest(res)
		if err != nil {
			return manifestSet, err
		}
		manifestSet = append(manifestSet, manifest)
	}
	return manifestSet, nil
}

// Get ResourcesGvk  get Manifests with group,versio,kind filter
func (c *metalctlClient) GetResourcesGvk(group, version, kind string) ([]Manifest, error) {
	g := resid.Gvk{Group: group, Version: version, Kind: kind}
	selector := types.Selector{Gvk: g}
	return c.Select(selector)
}

// Kustomize Select function used for GVK filtering and then collect found ones in the list
func (c *metalctlClient) Select(selector types.Selector) ([]Manifest, error) {
	resources, err := c.ResMap.Select(selector)
	if err != nil {
		return []Manifest{}, err
	}
	manifestSet := []Manifest{}
	for _, res := range resources {
		var manifest Manifest
		manifest, err = NewManifest(res)
		if err != nil {
			return manifestSet, err
		}
		manifestSet = append(manifestSet, manifest)
	}
	return manifestSet, err
}

func makeOptions() *krusty.Options {
	opts := &krusty.Options{
		DoLegacyResourceSort: true,
		LoadRestrictions:     0,
		DoPrune:              false,
	}
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
