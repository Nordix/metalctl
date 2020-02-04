package client

import (
	"metalctl/pkg/client"
)

type Client struct {
	options *client.GenerateOptions
}

var cfgFile string

// MakeClient returns an instance of Client
func MakeClient() *Client {
	return &Client{}
}

// Run performs a generate
func (c *Client) Generate(sourcePath string, outputPath string) error {
	metalclient, err := client.New(cfgFile)
	if err != nil {
		return err
	}
	return metalclient.Generate(client.GenerateOptions{
		SourcePath: sourcePath,
		OutputPath: outputPath,
	})
}
