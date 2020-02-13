package client

import (
	"fmt"
	"io"
	_ "metalctl/pkg/database"
	"os"

	"github.com/phogolabs/parcello"
	"sigs.k8s.io/kustomize/api/filesys"
)

func (c *metalctlClient) Init(fSys filesys.FileSystem, options GenerateOptions) error {
	err := parcello.Manager.Walk(options.Basepath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			panic(err)
		}
		if !info.IsDir() {
			fmt.Println("Opening path:", path)
			file, err := parcello.Open(path)
			if err != nil {
				fmt.Println("Failed to open path:", path)
				panic(err)
			}
			dest, err := fSys.Create(options.GeneratePath + path)
			if err != nil {
				fmt.Println("Failed to create path:", options.GeneratePath+path)
				panic(err)
			}
			if _, err = io.Copy(dest, file); err != nil {
				fmt.Println(err)
				panic(err)
			}
		} else {
			fmt.Println("Creating PATH", options.GeneratePath+path)
			err = fSys.MkdirAll(options.GeneratePath + path)
			if err != nil {
				fmt.Println("Directory creation failed")
				panic(err)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
