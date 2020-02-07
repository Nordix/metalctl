/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"io"
	"github.com/spf13/cobra"
	"metalctl/pkg/utils"
)

var cfgFile string

// NewMetalCTLCommand creates a root `metalctl` command with the default commands attached
func NewMetalCTLCommand(out io.Writer) (*cobra.Command, error) {
	rootCmd, err := NewRootCmd(out)
	return AddDefaultMetalCTLCommands(rootCmd), err
}

func NewRootCmd(out io.Writer) (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:          "metalctl",
		SilenceUsage: true,
		Short:        "metalctl controls a management cluster and move to target cluster",
		Long: utils.LongDesc(`
			Get started with Cluster API using metalctl for initializing a management cluster by installing
			CAPI+CAPBM+BMO  providers, and then use clusterctl for creating yaml templates for your workload clusters
			ande deploy them. After that pivot from management cluster to targer cluster`),
	}
	return rootCmd, nil
}

// AddDefaultMetalCTLCommands is a convenience function for adding all of the
// default commands to metalctl
func AddDefaultMetalCTLCommands(cmd *cobra.Command) *cobra.Command {
	cmd.AddCommand(NewGenerateManifestCmd())
	return cmd
}

