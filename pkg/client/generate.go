package client

import (
	"fmt"
	"path/filepath"
	"strings"

	//"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/kustomize/api/filesys"
	"sigs.k8s.io/kustomize/api/krusty"
	"sigs.k8s.io/kustomize/api/resmap"
	"sigs.k8s.io/kustomize/api/resource"
	"sigs.k8s.io/yaml"
)

//TODO add type conversions between structured and unstructured objects
/* type Unstructured struct {
	// Object is a JSON compatible map with string, float, int, bool, []interface{}, or
	// map[string]interface{}
	// children.
	Object map[string]interface{}
} */

func (c *metalctlClient) Generate(options GenerateOptions) error {
	var fSys filesys.FileSystem
	if options.FileSystemType == "internal" {
		fmt.Println("Setting fs internal")
		fSys = filesys.MakeFsInMemory()
	} else {
		fSys = filesys.MakeFsOnDisk()
	}
	c.Init(fSys, options)
	fmt.Println("Calling Kustomizer")
	k := krusty.MakeKustomizer(fSys, makeOptions())
	fmt.Println("Calling Run")
	m, err := k.Run(options.GeneratePath + options.KustomizePath)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	c.ResMap = m
	fSys = filesys.MakeFsOnDisk()
	fmt.Println("Calling write", c.ResMap)
	writeResources(options.OutputPath, fSys, c.ResMap)
	return nil
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
	fmt.Println("Writing to", outputPath)
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
