package client

type GenerateOptions struct {
	SourcePath string
	OutputPath string
}

type Client interface {
	// Generate generates templates into set location
	Generate(options GenerateOptions) error
}

// metalctlClient implements Client.
type metalctlClient struct {
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
