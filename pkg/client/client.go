package client

import (
	"sigs.k8s.io/kustomize/api/filesys"
	"sigs.k8s.io/kustomize/api/resmap"
)

type GenerateOptions struct {
	SourcePath     string
	OutputPath     string
	Basepath       string
	GeneratePath   string
	FileSystemType string
	KustomizePath  string
}

type Client interface {
	// Generate generates templates into set location
	Init(fSys filesys.FileSystem, options GenerateOptions) error
	Generate(options GenerateOptions) error
}

// metalctlClient implements Client.
type metalctlClient struct {
	ResMap resmap.ResMap
}

// Ensure metalctlClient implements Client.
var _ Client = &metalctlClient{}

// New returns a metalctlClient.
func New(path string) (Client, error) {
	return newMetalctlClient(path)
}

func newMetalctlClient(path string) (*metalctlClient, error) {
	return &metalctlClient{}, nil
}
